// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPendingTestsFromUserHandlerFunc turns a function with the right signature into a get pending tests from user handler
type GetPendingTestsFromUserHandlerFunc func(GetPendingTestsFromUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPendingTestsFromUserHandlerFunc) Handle(params GetPendingTestsFromUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPendingTestsFromUserHandler interface for that can handle valid get pending tests from user params
type GetPendingTestsFromUserHandler interface {
	Handle(GetPendingTestsFromUserParams, *models.User) middleware.Responder
}

// NewGetPendingTestsFromUser creates a new http.Handler for the get pending tests from user operation
func NewGetPendingTestsFromUser(ctx *middleware.Context, handler GetPendingTestsFromUserHandler) *GetPendingTestsFromUser {
	return &GetPendingTestsFromUser{Context: ctx, Handler: handler}
}

/* GetPendingTestsFromUser swagger:route GET /users/{username}/pendingTests user getPendingTestsFromUser

Returns all publishedTests that the user is invited to and hasn't answered yet

Returns all publishedTests that the user is invited to and hasn't answered yet

*/
type GetPendingTestsFromUser struct {
	Context *middleware.Context
	Handler GetPendingTestsFromUserHandler
}

func (o *GetPendingTestsFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPendingTestsFromUserParams()
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
