// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetUserHandlerFunc turns a function with the right signature into a get user handler
type GetUserHandlerFunc func(GetUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserHandlerFunc) Handle(params GetUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetUserHandler interface for that can handle valid get user params
type GetUserHandler interface {
	Handle(GetUserParams, *models.User) middleware.Responder
}

// NewGetUser creates a new http.Handler for the get user operation
func NewGetUser(ctx *middleware.Context, handler GetUserHandler) *GetUser {
	return &GetUser{Context: ctx, Handler: handler}
}

/* GetUser swagger:route GET /users/{username} user getUser

Finds a user by its username

Finds a user by its username

*/
type GetUser struct {
	Context *middleware.Context
	Handler GetUserHandler
}

func (o *GetUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserParams()
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
