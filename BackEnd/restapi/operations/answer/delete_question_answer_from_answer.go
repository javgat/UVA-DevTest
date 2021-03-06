// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// DeleteQuestionAnswerFromAnswerHandlerFunc turns a function with the right signature into a delete question answer from answer handler
type DeleteQuestionAnswerFromAnswerHandlerFunc func(DeleteQuestionAnswerFromAnswerParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteQuestionAnswerFromAnswerHandlerFunc) Handle(params DeleteQuestionAnswerFromAnswerParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteQuestionAnswerFromAnswerHandler interface for that can handle valid delete question answer from answer params
type DeleteQuestionAnswerFromAnswerHandler interface {
	Handle(DeleteQuestionAnswerFromAnswerParams, *models.User) middleware.Responder
}

// NewDeleteQuestionAnswerFromAnswer creates a new http.Handler for the delete question answer from answer operation
func NewDeleteQuestionAnswerFromAnswer(ctx *middleware.Context, handler DeleteQuestionAnswerFromAnswerHandler) *DeleteQuestionAnswerFromAnswer {
	return &DeleteQuestionAnswerFromAnswer{Context: ctx, Handler: handler}
}

/* DeleteQuestionAnswerFromAnswer swagger:route DELETE /answers/{answerid}/qanswers/{questionid} answer deleteQuestionAnswerFromAnswer

Deletes an answer's questionAnswer. Only if answer is open

Deletes an answers's questionAnswer. Only if answer is open

*/
type DeleteQuestionAnswerFromAnswer struct {
	Context *middleware.Context
	Handler DeleteQuestionAnswerFromAnswerHandler
}

func (o *DeleteQuestionAnswerFromAnswer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteQuestionAnswerFromAnswerParams()
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
