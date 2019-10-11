// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeviceListHandlerFunc turns a function with the right signature into a device list handler
type DeviceListHandlerFunc func(DeviceListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeviceListHandlerFunc) Handle(params DeviceListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeviceListHandler interface for that can handle valid device list params
type DeviceListHandler interface {
	Handle(DeviceListParams, interface{}) middleware.Responder
}

// NewDeviceList creates a new http.Handler for the device list operation
func NewDeviceList(ctx *middleware.Context, handler DeviceListHandler) *DeviceList {
	return &DeviceList{Context: ctx, Handler: handler}
}

/*DeviceList swagger:route GET /devices/ wireguard deviceList

get wireguard devices

*/
type DeviceList struct {
	Context *middleware.Context
	Handler DeviceListHandler
}

func (o *DeviceList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeviceListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}