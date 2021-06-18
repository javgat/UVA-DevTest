// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// CopyQuestionHandlerFunc turns a function with the right signature into a copy question handler
type CopyQuestionHandlerFunc func(CopyQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CopyQuestionHandlerFunc) Handle(params CopyQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CopyQuestionHandler interface for that can handle valid copy question params
type CopyQuestionHandler interface {
	Handle(CopyQuestionParams, *models.User) middleware.Responder
}

// NewCopyQuestion creates a new http.Handler for the copy question operation
func NewCopyQuestion(ctx *middleware.Context, handler CopyQuestionHandler) *CopyQuestion {
	return &CopyQuestion{Context: ctx, Handler: handler}
}

/* CopyQuestion swagger:route POST /users/{username}/questions/{questionid}/copiedQuestions user question copyQuestion

Creates a question copied from another

Creates a question copied from another

*/
type CopyQuestion struct {
	Context *middleware.Context
	Handler CopyQuestionHandler
}

func (o *CopyQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCopyQuestionParams()
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
