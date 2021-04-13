// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// StartAnswerHandlerFunc turns a function with the right signature into a start answer handler
type StartAnswerHandlerFunc func(StartAnswerParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn StartAnswerHandlerFunc) Handle(params StartAnswerParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// StartAnswerHandler interface for that can handle valid start answer params
type StartAnswerHandler interface {
	Handle(StartAnswerParams, *models.User) middleware.Responder
}

// NewStartAnswer creates a new http.Handler for the start answer operation
func NewStartAnswer(ctx *middleware.Context, handler StartAnswerHandler) *StartAnswer {
	return &StartAnswer{Context: ctx, Handler: handler}
}

/* StartAnswer swagger:route POST /users/{username}/publishedTests/{testid}/answers user answer startAnswer

Starts a new answer

Starts a new answer

*/
type StartAnswer struct {
	Context *middleware.Context
	Handler StartAnswerHandler
}

func (o *StartAnswer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStartAnswerParams()
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
