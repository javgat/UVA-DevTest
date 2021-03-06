// Code generated by go-swagger; DO NOT EDIT.

package tiporol

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTipoRolHandlerFunc turns a function with the right signature into a get tipo rol handler
type GetTipoRolHandlerFunc func(GetTipoRolParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTipoRolHandlerFunc) Handle(params GetTipoRolParams) middleware.Responder {
	return fn(params)
}

// GetTipoRolHandler interface for that can handle valid get tipo rol params
type GetTipoRolHandler interface {
	Handle(GetTipoRolParams) middleware.Responder
}

// NewGetTipoRol creates a new http.Handler for the get tipo rol operation
func NewGetTipoRol(ctx *middleware.Context, handler GetTipoRolHandler) *GetTipoRol {
	return &GetTipoRol{Context: ctx, Handler: handler}
}

/* GetTipoRol swagger:route GET /tipoRoles/{rolNombre} tiporol getTipoRol

Returns a TipoRol

Returns a TipoRol

*/
type GetTipoRol struct {
	Context *middleware.Context
	Handler GetTipoRolHandler
}

func (o *GetTipoRol) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTipoRolParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
