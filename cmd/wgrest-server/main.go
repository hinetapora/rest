package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/hinetapora/wgrest/handlers"
	"github.com/hinetapora/wgrest/storage"
	"github.com/hinetapora/wgrest/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"golang.org/x/crypto/acme/autocert"
)

var (
	appVersion string // Populated during build time
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "conf",
			Value:   "/etc/wgrest/wgrest.conf",
			Usage:   "wgrest config file path",
			EnvVars: []string{"WGREST_CONF"},
		},
		&cli.BoolFlag{
			Name:  "version",
			Value: false,
			Usage: "Print version and exit",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "listen",
			Value:   "127.0.0.1:8000",
			Usage:   "Listen address",
			EnvVars: []string{"WGREST_LISTEN"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "data-dir",
			Value:   "/var/lib/wgrest",
			Usage:   "Data dir",
			EnvVars: []string{"WGREST_DATA_DIR"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "static-auth-token",
			Value:   "",
			Usage:   "It is used for bearer token authorization",
			EnvVars: []string{"WGREST_STATIC_AUTH_TOKEN"},
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:    "tls-domain",
			Value:   cli.NewStringSlice(),
			Usage:   "TLS Domains",
			EnvVars: []string{"WGREST_TLS_DOMAIN"},
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:    "demo",
			Value:   false,
			Usage:   "Demo mode",
			EnvVars: []string{"WGREST_DEMO"},
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:    "device-allowed-ips",
			Value:   cli.NewStringSlice("0.0.0.0/0", "::0/0"),
			Usage:   "Default device allowed ips. You can overwrite it through api",
			EnvVars: []string{"WGREST_DEVICE_ALLOWED_IPS"},
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:    "device-dns-servers",
			Value:   cli.NewStringSlice("8.8.8.8", "1.1.1.1", "2001:4860:4860::8888", "2606:4700:4700::1111"),
			Usage:   "Default device DNS servers. You can overwrite it through api",
			EnvVars: []string{"WGREST_DEVICE_DNS_SERVERS"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "device-host",
			Value:   "",
			Usage:   "Default device host. You can overwrite it through api",
			EnvVars: []string{"WGREST_DEVICE_HOST"},
		}),
	}

	app := &cli.App{
		Name:   "wgrest",
		Usage:  "wgrest - rest api for wireguard",
		Flags:  flags,
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewTomlSourceFromFlagFunc("conf")),
		Action: func(c *cli.Context) error {
			if c.Bool("version") {
				fmt.Printf("wgrest version: %s\n", appVersion)
				return nil
			}

			e := echo.New()
			e.HideBanner = true

			// Basic middleware
			e.Use(middleware.Logger())
			e.Use(middleware.Recover())
			e.Pre(middleware.Rewrite(map[string]string{
				"^/devices":   "/",
				"^/devices/*": "/",
			}))

			e.GET("/version", getVersionHandler)

			dataDir := c.String("data-dir")
			e.File("/", path.Join(dataDir, "public", "index.html"))
			e.Static("/", path.Join(dataDir, "public"))

			cacheDir := path.Join(dataDir, ".cache")
			tlsDomains := c.StringSlice("tls-domain")
			if len(tlsDomains) > 0 {
				e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(tlsDomains...)
				e.AutoTLSManager.Cache = autocert.DirCache(cacheDir)
			}

			v1 := e.Group("/v1")
			v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				Skipper:          middleware.DefaultSkipper,
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
				AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Language", "Link", "Authorization"},
				AllowCredentials: true,
			}))

			staticAuthToken := c.String("static-auth-token")
			if staticAuthToken != "" {
				v1.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
					return key == staticAuthToken, nil
				}))
			}

			wgStorage, err := storage.NewFileStorage(dataDir)
			if err != nil {
				return err
			}

			defaultDeviceHost := c.String("device-host")
			if defaultDeviceHost == "" {
				defaultDeviceHost, err = utils.GetExternalIP()
				if err != nil {
					log.Printf("failed to identify external ip: %s", err.Error())
				}
			}

			defaultDeviceOptions := storage.StoreDeviceOptions{
				AllowedIPs: c.StringSlice("device-allowed-ips"),
				DNSServers: c.StringSlice("device-dns-servers"),
				Host:       defaultDeviceHost,
			}

			wc, err := handlers.NewWireGuardContainer(handlers.WireGuardContainerOptions{
				Storage:              wgStorage,
				DefaultDeviceOptions: defaultDeviceOptions,
			})
			if err != nil {
				return err
			}

			// CreateDevice - Create new device
			v1.POST("/devices/", wc.CreateDevice)

			// CreateDevicePeer - Create new device peer
			v1.POST("/devices/:name/peers/", wc.CreateDevicePeer)

			// DeleteDevice - Delete Device
			v1.DELETE("/devices/:name/", wc.DeleteDevice)

			// DeleteDevicePeer - Delete device's peer
			v1.DELETE("/devices/:name/peers/:urlSafePubKey/", wc.DeleteDevicePeer)

			// GetDevice - Get device info
			v1.GET("/devices/:name/", wc.GetDevice)

			// GetDevicePeer - Get device peer info
			v1.GET("/devices/:name/peers/:urlSafePubKey/", wc.GetDevicePeer)

			// ListDevicePeers - Peers list
			v1.GET("/devices/:name/peers/", wc.ListDevicePeers)

			// ListDevices - Devices list
			v1.GET("/devices/", wc.ListDevices)

			// UpdateDevice - Update device
			v1.PATCH("/devices/:name/", wc.UpdateDevice)

			// UpdateDevicePeer - Update device's peer
			v1.PATCH("/devices/:name/peers/:urlSafePubKey/", wc.UpdateDevicePeer)

			// GetDevicePeerQuickConfig - Get device peer quick config
			v1.GET("/devices/:name/peers/:urlSafePubKey/quick.conf", wc.GetDevicePeerQuickConfig)

			// GetDevicePeerQuickConfigQRCodePNG - Get device peer quick config QR code
			v1.GET("/devices/:name/peers/:urlSafePubKey/quick.conf.png", wc.GetDevicePeerQuickConfigQRCodePNG)

			// GetDeviceOptions - Get device options
			v1.GET("/devices/:name/options/", wc.GetDeviceOptions)

			// UpdateDeviceOptions - Update device's options
			v1.PATCH("/devices/:name/options/", wc.UpdateDeviceOptions)

			listen := c.String("listen")
			// Start server
			if len(tlsDomains) > 0 {
				return e.StartAutoTLS(listen)
			} else {
				return e.Start(listen)
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getVersionHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, appVersion)
}
