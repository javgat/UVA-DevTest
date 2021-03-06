// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetAdminsHandlerFunc turns a function with the right signature into a get admins handler
type GetAdminsHandlerFunc func(GetAdminsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAdminsHandlerFunc) Handle(params GetAdminsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetAdminsHandler interface for that can handle valid get admins params
type GetAdminsHandler interface {
	Handle(GetAdminsParams, *models.User) middleware.Responder
}

// NewGetAdmins creates a new http.Handler for the get admins operation
func NewGetAdmins(ctx *middleware.Context, handler GetAdminsHandler) *GetAdmins {
	return &GetAdmins{Context: ctx, Handler: handler}
}

/* GetAdmins swagger:route GET /teams/{teamname}/admins team getAdmins

Returns all users that are admins of a team

Returns all users that are admins of a team

*/
type GetAdmins struct {
	Context *middleware.Context
	Handler GetAdminsHandler
}

func (o *GetAdmins) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAdminsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
