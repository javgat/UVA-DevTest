// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// RemoveTeamToTestHandlerFunc turns a function with the right signature into a remove team to test handler
type RemoveTeamToTestHandlerFunc func(RemoveTeamToTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveTeamToTestHandlerFunc) Handle(params RemoveTeamToTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// RemoveTeamToTestHandler interface for that can handle valid remove team to test params
type RemoveTeamToTestHandler interface {
	Handle(RemoveTeamToTestParams, *models.User) middleware.Responder
}

// NewRemoveTeamToTest creates a new http.Handler for the remove team to test operation
func NewRemoveTeamToTest(ctx *middleware.Context, handler RemoveTeamToTestHandler) *RemoveTeamToTest {
	return &RemoveTeamToTest{Context: ctx, Handler: handler}
}

/* RemoveTeamToTest swagger:route DELETE /tests/{testid}/teams/{teamname} test removeTeamToTest

Removes a team to administer a test

Removes a team to administer a test

*/
type RemoveTeamToTest struct {
	Context *middleware.Context
	Handler RemoveTeamToTestHandler
}

func (o *RemoveTeamToTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRemoveTeamToTestParams()
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
