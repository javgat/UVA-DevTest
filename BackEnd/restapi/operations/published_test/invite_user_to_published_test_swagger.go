// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// InviteUserToPublishedTestHandlerFunc turns a function with the right signature into a invite user to published test handler
type InviteUserToPublishedTestHandlerFunc func(InviteUserToPublishedTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn InviteUserToPublishedTestHandlerFunc) Handle(params InviteUserToPublishedTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// InviteUserToPublishedTestHandler interface for that can handle valid invite user to published test params
type InviteUserToPublishedTestHandler interface {
	Handle(InviteUserToPublishedTestParams, *models.User) middleware.Responder
}

// NewInviteUserToPublishedTest creates a new http.Handler for the invite user to published test operation
func NewInviteUserToPublishedTest(ctx *middleware.Context, handler InviteUserToPublishedTestHandler) *InviteUserToPublishedTest {
	return &InviteUserToPublishedTest{Context: ctx, Handler: handler}
}

/* InviteUserToPublishedTest swagger:route PUT /publishedTests/{testid}/users/{username} publishedTest inviteUserToPublishedTest

Invites a user to do a test

Invites a user to do a test

*/
type InviteUserToPublishedTest struct {
	Context *middleware.Context
	Handler InviteUserToPublishedTestHandler
}

func (o *InviteUserToPublishedTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewInviteUserToPublishedTestParams()
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
