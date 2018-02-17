// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHistoriqueParams creates a new HistoriqueParams object
// with the default values initialized.
func NewHistoriqueParams() HistoriqueParams {
	var ()
	return HistoriqueParams{}
}

// HistoriqueParams contains all the bound params for the historique operation
// typically these are obtained from a http.Request
//
// swagger:parameters historique
type HistoriqueParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Objetid string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *HistoriqueParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rObjetid, rhkObjetid, _ := route.Params.GetOK("objetid")
	if err := o.bindObjetid(rObjetid, rhkObjetid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HistoriqueParams) bindObjetid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Objetid = raw

	return nil
}