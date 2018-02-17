// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AssociationHandlerFunc turns a function with the right signature into a association handler
type AssociationHandlerFunc func(AssociationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociationHandlerFunc) Handle(params AssociationParams) middleware.Responder {
	return fn(params)
}

// AssociationHandler interface for that can handle valid association params
type AssociationHandler interface {
	Handle(AssociationParams) middleware.Responder
}

// NewAssociation creates a new http.Handler for the association operation
func NewAssociation(ctx *middleware.Context, handler AssociationHandler) *Association {
	return &Association{Context: ctx, Handler: handler}
}

/*Association swagger:route POST /association/{c_id} association

Met des objets dans un objet

Met des objets unitaire dans objets conteneurs

*/
type Association struct {
	Context *middleware.Context
	Handler AssociationHandler
}

func (o *Association) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}