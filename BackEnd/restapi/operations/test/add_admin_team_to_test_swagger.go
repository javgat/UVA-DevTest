// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// AddAdminTeamToTestHandlerFunc turns a function with the right signature into a add admin team to test handler
type AddAdminTeamToTestHandlerFunc func(AddAdminTeamToTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AddAdminTeamToTestHandlerFunc) Handle(params AddAdminTeamToTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AddAdminTeamToTestHandler interface for that can handle valid add admin team to test params
type AddAdminTeamToTestHandler interface {
	Handle(AddAdminTeamToTestParams, *models.User) middleware.Responder
}

// NewAddAdminTeamToTest creates a new http.Handler for the add admin team to test operation
func NewAddAdminTeamToTest(ctx *middleware.Context, handler AddAdminTeamToTestHandler) *AddAdminTeamToTest {
	return &AddAdminTeamToTest{Context: ctx, Handler: handler}
}

/* AddAdminTeamToTest swagger:route PUT /tests/{testid}/adminTeams/{teamname} test addAdminTeamToTest

Adds a team to administer a test

Adds a team to administer a test

*/
type AddAdminTeamToTest struct {
	Context *middleware.Context
	Handler AddAdminTeamToTestHandler
}

func (o *AddAdminTeamToTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddAdminTeamToTestParams()
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
