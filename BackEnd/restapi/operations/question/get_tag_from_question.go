// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetTagFromQuestionHandlerFunc turns a function with the right signature into a get tag from question handler
type GetTagFromQuestionHandlerFunc func(GetTagFromQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTagFromQuestionHandlerFunc) Handle(params GetTagFromQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTagFromQuestionHandler interface for that can handle valid get tag from question params
type GetTagFromQuestionHandler interface {
	Handle(GetTagFromQuestionParams, *models.User) middleware.Responder
}

// NewGetTagFromQuestion creates a new http.Handler for the get tag from question operation
func NewGetTagFromQuestion(ctx *middleware.Context, handler GetTagFromQuestionHandler) *GetTagFromQuestion {
	return &GetTagFromQuestion{Context: ctx, Handler: handler}
}

/* GetTagFromQuestion swagger:route GET /questions/{questionid}/tags/{tag} question getTagFromQuestion

Returns a tag from a question.

Returns a tag from a question.

*/
type GetTagFromQuestion struct {
	Context *middleware.Context
	Handler GetTagFromQuestionHandler
}

func (o *GetTagFromQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTagFromQuestionParams()
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