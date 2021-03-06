// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetOpenAnswersFromUserTestHandlerFunc turns a function with the right signature into a get open answers from user test handler
type GetOpenAnswersFromUserTestHandlerFunc func(GetOpenAnswersFromUserTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOpenAnswersFromUserTestHandlerFunc) Handle(params GetOpenAnswersFromUserTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetOpenAnswersFromUserTestHandler interface for that can handle valid get open answers from user test params
type GetOpenAnswersFromUserTestHandler interface {
	Handle(GetOpenAnswersFromUserTestParams, *models.User) middleware.Responder
}

// NewGetOpenAnswersFromUserTest creates a new http.Handler for the get open answers from user test operation
func NewGetOpenAnswersFromUserTest(ctx *middleware.Context, handler GetOpenAnswersFromUserTestHandler) *GetOpenAnswersFromUserTest {
	return &GetOpenAnswersFromUserTest{Context: ctx, Handler: handler}
}

/* GetOpenAnswersFromUserTest swagger:route GET /users/{username}/solvableTests/{testid}/openAnswers user getOpenAnswersFromUserTest

Returns all open answers that the user is answering to a test. It should be only one or zero

Returns all open answers that the user is answering to a test. It should be only one or zero

*/
type GetOpenAnswersFromUserTest struct {
	Context *middleware.Context
	Handler GetOpenAnswersFromUserTestHandler
}

func (o *GetOpenAnswersFromUserTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetOpenAnswersFromUserTestParams()
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
