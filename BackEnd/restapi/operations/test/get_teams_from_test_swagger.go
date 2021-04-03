// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetTeamsFromTestHandlerFunc turns a function with the right signature into a get teams from test handler
type GetTeamsFromTestHandlerFunc func(GetTeamsFromTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeamsFromTestHandlerFunc) Handle(params GetTeamsFromTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTeamsFromTestHandler interface for that can handle valid get teams from test params
type GetTeamsFromTestHandler interface {
	Handle(GetTeamsFromTestParams, *models.User) middleware.Responder
}

// NewGetTeamsFromTest creates a new http.Handler for the get teams from test operation
func NewGetTeamsFromTest(ctx *middleware.Context, handler GetTeamsFromTestHandler) *GetTeamsFromTest {
	return &GetTeamsFromTest{Context: ctx, Handler: handler}
}

/* GetTeamsFromTest swagger:route GET /tests/{testid}/teams test getTeamsFromTest

Returns all teams from a test.

Returns all teams from a test.

*/
type GetTeamsFromTest struct {
	Context *middleware.Context
	Handler GetTeamsFromTestHandler
}

func (o *GetTeamsFromTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTeamsFromTestParams()
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