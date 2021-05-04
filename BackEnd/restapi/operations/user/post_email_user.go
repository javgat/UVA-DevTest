// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// PostEmailUserHandlerFunc turns a function with the right signature into a post email user handler
type PostEmailUserHandlerFunc func(PostEmailUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PostEmailUserHandlerFunc) Handle(params PostEmailUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PostEmailUserHandler interface for that can handle valid post email user params
type PostEmailUserHandler interface {
	Handle(PostEmailUserParams, *models.User) middleware.Responder
}

// NewPostEmailUser creates a new http.Handler for the post email user operation
func NewPostEmailUser(ctx *middleware.Context, handler PostEmailUserHandler) *PostEmailUser {
	return &PostEmailUser{Context: ctx, Handler: handler}
}

/* PostEmailUser swagger:route POST /emailUsers user postEmailUser

Adds a user without specifying username or password

Adds a user without specifying username or password

*/
type PostEmailUser struct {
	Context *middleware.Context
	Handler PostEmailUserHandler
}

func (o *PostEmailUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostEmailUserParams()
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
