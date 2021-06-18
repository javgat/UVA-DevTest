// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// PutRoleHandlerFunc turns a function with the right signature into a put role handler
type PutRoleHandlerFunc func(PutRoleParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PutRoleHandlerFunc) Handle(params PutRoleParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PutRoleHandler interface for that can handle valid put role params
type PutRoleHandler interface {
	Handle(PutRoleParams, *models.User) middleware.Responder
}

// NewPutRole creates a new http.Handler for the put role operation
func NewPutRole(ctx *middleware.Context, handler PutRoleHandler) *PutRole {
	return &PutRole{Context: ctx, Handler: handler}
}

/* PutRole swagger:route PUT /users/{username}/role user putRole

Modifies the role of the user <username>

Modifies the role of the user <username>

*/
type PutRole struct {
	Context *middleware.Context
	Handler PutRoleHandler
}

func (o *PutRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutRoleParams()
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
