// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// QrcodeHandlerFunc turns a function with the right signature into a qrcode handler
type QrcodeHandlerFunc func(QrcodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn QrcodeHandlerFunc) Handle(params QrcodeParams) middleware.Responder {
	return fn(params)
}

// QrcodeHandler interface for that can handle valid qrcode params
type QrcodeHandler interface {
	Handle(QrcodeParams) middleware.Responder
}

// NewQrcode creates a new http.Handler for the qrcode operation
func NewQrcode(ctx *middleware.Context, handler QrcodeHandler) *Qrcode {
	return &Qrcode{Context: ctx, Handler: handler}
}

/*Qrcode swagger:route POST /qrcode qrcode

Créé un QRcode

Créé un QRcode avec les metadata associés

*/
type Qrcode struct {
	Context *middleware.Context
	Handler QrcodeHandler
}

func (o *Qrcode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewQrcodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}