// Code generated by go-swagger; DO NOT EDIT.

package wireguard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPeerGetParams creates a new PeerGetParams object
// no default values defined in spec.
func NewPeerGetParams() PeerGetParams {

	return PeerGetParams{}
}

// PeerGetParams contains all the bound params for the peer get operation
// typically these are obtained from a http.Request
//
// swagger:parameters PeerGet
type PeerGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Dev string
	/*
	  Required: true
	  In: path
	*/
	PeerID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPeerGetParams() beforehand.
func (o *PeerGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDev, rhkDev, _ := route.Params.GetOK("dev")
	if err := o.bindDev(rDev, rhkDev, route.Formats); err != nil {
		res = append(res, err)
	}

	rPeerID, rhkPeerID, _ := route.Params.GetOK("peer_id")
	if err := o.bindPeerID(rPeerID, rhkPeerID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDev binds and validates parameter Dev from path.
func (o *PeerGetParams) bindDev(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Dev = raw

	return nil
}

// bindPeerID binds and validates parameter PeerID from path.
func (o *PeerGetParams) bindPeerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PeerID = raw

	return nil
}
