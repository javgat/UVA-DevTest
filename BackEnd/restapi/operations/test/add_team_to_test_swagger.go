// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// AddTeamToTestHandlerFunc turns a function with the right signature into a add team to test handler
type AddTeamToTestHandlerFunc func(AddTeamToTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTeamToTestHandlerFunc) Handle(params AddTeamToTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AddTeamToTestHandler interface for that can handle valid add team to test params
type AddTeamToTestHandler interface {
	Handle(AddTeamToTestParams, *models.User) middleware.Responder
}

// NewAddTeamToTest creates a new http.Handler for the add team to test operation
func NewAddTeamToTest(ctx *middleware.Context, handler AddTeamToTestHandler) *AddTeamToTest {
	return &AddTeamToTest{Context: ctx, Handler: handler}
}

/* AddTeamToTest swagger:route PUT /tests/{testid}/teams/{teamname} test addTeamToTest

Adds a team to administer a test

Adds a team to administer a test

*/
type AddTeamToTest struct {
	Context *middleware.Context
	Handler AddTeamToTestHandler
}

func (o *AddTeamToTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddTeamToTestParams()
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
