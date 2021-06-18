// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// DeleteQuestionHandlerFunc turns a function with the right signature into a delete question handler
type DeleteQuestionHandlerFunc func(DeleteQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteQuestionHandlerFunc) Handle(params DeleteQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteQuestionHandler interface for that can handle valid delete question params
type DeleteQuestionHandler interface {
	Handle(DeleteQuestionParams, *models.User) middleware.Responder
}

// NewDeleteQuestion creates a new http.Handler for the delete question operation
func NewDeleteQuestion(ctx *middleware.Context, handler DeleteQuestionHandler) *DeleteQuestion {
	return &DeleteQuestion{Context: ctx, Handler: handler}
}

/* DeleteQuestion swagger:route DELETE /questions/{questionid} question deleteQuestion

Deletes a question

Deletes a question

*/
type DeleteQuestion struct {
	Context *middleware.Context
	Handler DeleteQuestionHandler
}

func (o *DeleteQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteQuestionParams()
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
