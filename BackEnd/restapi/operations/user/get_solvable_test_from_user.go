// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetSolvableTestFromUserHandlerFunc turns a function with the right signature into a get solvable test from user handler
type GetSolvableTestFromUserHandlerFunc func(GetSolvableTestFromUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSolvableTestFromUserHandlerFunc) Handle(params GetSolvableTestFromUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetSolvableTestFromUserHandler interface for that can handle valid get solvable test from user params
type GetSolvableTestFromUserHandler interface {
	Handle(GetSolvableTestFromUserParams, *models.User) middleware.Responder
}

// NewGetSolvableTestFromUser creates a new http.Handler for the get solvable test from user operation
func NewGetSolvableTestFromUser(ctx *middleware.Context, handler GetSolvableTestFromUserHandler) *GetSolvableTestFromUser {
	return &GetSolvableTestFromUser{Context: ctx, Handler: handler}
}

/* GetSolvableTestFromUser swagger:route GET /users/{username}/solvableTests/{testid} user getSolvableTestFromUser

Returns a publishedTest that the user can answer, including public ones and team ones. The Test DTO will contain the number of answers the user has for the test

Returns a publishedTest that the user can answer, including public ones and team ones.  The Test DTO will contain the number of answers the user has for the test

*/
type GetSolvableTestFromUser struct {
	Context *middleware.Context
	Handler GetSolvableTestFromUserHandler
}

func (o *GetSolvableTestFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSolvableTestFromUserParams()
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
