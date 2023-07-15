BUILDDIR ?= dist
OSS ?= linux darwin freebsd windows
ARCHS ?= amd64 arm64
VERSION ?= $(shell git describe --tags `git rev-list -1 HEAD`)

build: $(BUILDDIR)/wgrest

clean:
	rm -rf "$(BUILDDIR)"

install: build

define wgrest
$(BUILDDIR)/wgrest-$(1)-$(2): export CGO_ENABLED := 0
$(BUILDDIR)/wgrest-$(1)-$(2): export GOOS := $(1)
$(BUILDDIR)/wgrest-$(1)-$(2): export GOARCH := $(2)
$(BUILDDIR)/wgrest-$(1)-$(2):
	go build \
	-ldflags="-s -w -X main.appVersion=$(VERSION)" \
	-trimpath -v -o "$(BUILDDIR)/wgrest-$(1)-$(2)" \
	cmd/wgrest-server/main.go
endef
$(foreach OS,$(OSS),$(foreach ARCH,$(ARCHS),$(eval $(call wgrest,$(OS),$(ARCH)))))

$(BUILDDIR)/wgrest: $(foreach OS,$(OSS),$(foreach ARCH,$(ARCHS),$(BUILDDIR)/wgrest-$(OS)-$(ARCH)))
	@mkdir -vp "$(BUILDDIR)"

package: $(BUILDDIR)/wgrest
	mkdir -p pkg/debian/DEBIAN
	mkdir -p pkg/debian/usr/bin
	cp $(BUILDDIR)/wgrest-linux-amd64 pkg/debian/usr/bin/wgrest
	cp debian/control pkg/debian/DEBIAN
	dpkg-deb --build pkg/debian

package-darwin: $(BUILDDIR)/wgrest-darwin-amd64
	mkdir -p pkg/darwin
	cp $(BUILDDIR)/wgrest-darwin-amd64 pkg/darwin/wgrest

package-windows: $(BUILDDIR)/wgrest-windows-amd64
	mkdir -p pkg/windows
	cp $(BUILDDIR)/wgrest-windows-amd64 pkg/windows/wgrest.exe

go-echo-server:
	openapi-generator generate -g go-echo-server \
		-i openapi-spec.yaml \
		-o . \
		--git-host github.com \
		--git-user-id hinetapora \
		--git-repo-id wgrest

typescript-axios-client:
	swagger-codegen generate -l typescript-axios \
		--additional-properties modelPropertyNaming=original \
		-i openapi-spec.yaml \
		-o clients/typeascript-axios

.PHONY: clean build install package package-darwin package-windows
