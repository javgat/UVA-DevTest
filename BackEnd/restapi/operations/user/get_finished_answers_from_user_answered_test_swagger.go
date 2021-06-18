// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetFinishedAnswersFromUserAnsweredTestHandlerFunc turns a function with the right signature into a get finished answers from user answered test handler
type GetFinishedAnswersFromUserAnsweredTestHandlerFunc func(GetFinishedAnswersFromUserAnsweredTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFinishedAnswersFromUserAnsweredTestHandlerFunc) Handle(params GetFinishedAnswersFromUserAnsweredTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetFinishedAnswersFromUserAnsweredTestHandler interface for that can handle valid get finished answers from user answered test params
type GetFinishedAnswersFromUserAnsweredTestHandler interface {
	Handle(GetFinishedAnswersFromUserAnsweredTestParams, *models.User) middleware.Responder
}

// NewGetFinishedAnswersFromUserAnsweredTest creates a new http.Handler for the get finished answers from user answered test operation
func NewGetFinishedAnswersFromUserAnsweredTest(ctx *middleware.Context, handler GetFinishedAnswersFromUserAnsweredTestHandler) *GetFinishedAnswersFromUserAnsweredTest {
	return &GetFinishedAnswersFromUserAnsweredTest{Context: ctx, Handler: handler}
}

/* GetFinishedAnswersFromUserAnsweredTest swagger:route GET /users/{username}/answeredTests/{testid}/finishedAnswers user getFinishedAnswersFromUserAnsweredTest

Returns all answers that the user has answered to a test and are finished

Returns all answers that the user has answered to a test and are finished

*/
type GetFinishedAnswersFromUserAnsweredTest struct {
	Context *middleware.Context
	Handler GetFinishedAnswersFromUserAnsweredTestHandler
}

func (o *GetFinishedAnswersFromUserAnsweredTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFinishedAnswersFromUserAnsweredTestParams()
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
