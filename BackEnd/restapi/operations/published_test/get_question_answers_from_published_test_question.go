// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetQuestionAnswersFromPublishedTestQuestionHandlerFunc turns a function with the right signature into a get question answers from published test question handler
type GetQuestionAnswersFromPublishedTestQuestionHandlerFunc func(GetQuestionAnswersFromPublishedTestQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQuestionAnswersFromPublishedTestQuestionHandlerFunc) Handle(params GetQuestionAnswersFromPublishedTestQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetQuestionAnswersFromPublishedTestQuestionHandler interface for that can handle valid get question answers from published test question params
type GetQuestionAnswersFromPublishedTestQuestionHandler interface {
	Handle(GetQuestionAnswersFromPublishedTestQuestionParams, *models.User) middleware.Responder
}

// NewGetQuestionAnswersFromPublishedTestQuestion creates a new http.Handler for the get question answers from published test question operation
func NewGetQuestionAnswersFromPublishedTestQuestion(ctx *middleware.Context, handler GetQuestionAnswersFromPublishedTestQuestionHandler) *GetQuestionAnswersFromPublishedTestQuestion {
	return &GetQuestionAnswersFromPublishedTestQuestion{Context: ctx, Handler: handler}
}

/* GetQuestionAnswersFromPublishedTestQuestion swagger:route GET /publishedTests/{testid}/questions/{questionid}/qanswers publishedTest getQuestionAnswersFromPublishedTestQuestion

Returns all questions answers to a question of a published test

Returns all questions answers to a question of a published test

*/
type GetQuestionAnswersFromPublishedTestQuestion struct {
	Context *middleware.Context
	Handler GetQuestionAnswersFromPublishedTestQuestionHandler
}

func (o *GetQuestionAnswersFromPublishedTestQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetQuestionAnswersFromPublishedTestQuestionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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