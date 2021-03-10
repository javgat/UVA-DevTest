// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// PutPasswordHandlerFunc turns a function with the right signature into a put password handler
type PutPasswordHandlerFunc func(PutPasswordParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPasswordHandlerFunc) Handle(params PutPasswordParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PutPasswordHandler interface for that can handle valid put password params
type PutPasswordHandler interface {
	Handle(PutPasswordParams, *models.User) middleware.Responder
}

// NewPutPassword creates a new http.Handler for the put password operation
func NewPutPassword(ctx *middleware.Context, handler PutPasswordHandler) *PutPassword {
	return &PutPassword{Context: ctx, Handler: handler}
}

/* PutPassword swagger:route PUT /passwords/{username} user auth putPassword

Modifies the password of the user <username>

Modifies the password of the user <username>

*/
type PutPassword struct {
	Context *middleware.Context
	Handler PutPasswordHandler
}

func (o *PutPassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutPasswordParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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
