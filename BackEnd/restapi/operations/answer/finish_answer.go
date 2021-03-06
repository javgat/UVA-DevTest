// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// FinishAnswerHandlerFunc turns a function with the right signature into a finish answer handler
type FinishAnswerHandlerFunc func(FinishAnswerParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn FinishAnswerHandlerFunc) Handle(params FinishAnswerParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// FinishAnswerHandler interface for that can handle valid finish answer params
type FinishAnswerHandler interface {
	Handle(FinishAnswerParams, *models.User) middleware.Responder
}

// NewFinishAnswer creates a new http.Handler for the finish answer operation
func NewFinishAnswer(ctx *middleware.Context, handler FinishAnswerHandler) *FinishAnswer {
	return &FinishAnswer{Context: ctx, Handler: handler}
}

/* FinishAnswer swagger:route PUT /answers/{answerid} answer finishAnswer

Finishes an answer

Finishes an answers

*/
type FinishAnswer struct {
	Context *middleware.Context
	Handler FinishAnswerHandler
}

func (o *FinishAnswer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFinishAnswerParams()
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
