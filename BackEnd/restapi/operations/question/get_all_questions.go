// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetAllQuestionsHandlerFunc turns a function with the right signature into a get all questions handler
type GetAllQuestionsHandlerFunc func(GetAllQuestionsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllQuestionsHandlerFunc) Handle(params GetAllQuestionsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetAllQuestionsHandler interface for that can handle valid get all questions params
type GetAllQuestionsHandler interface {
	Handle(GetAllQuestionsParams, *models.User) middleware.Responder
}

// NewGetAllQuestions creates a new http.Handler for the get all questions operation
func NewGetAllQuestions(ctx *middleware.Context, handler GetAllQuestionsHandler) *GetAllQuestions {
	return &GetAllQuestions{Context: ctx, Handler: handler}
}

/* GetAllQuestions swagger:route GET /allQuestions question getAllQuestions

Returns all questions. Admin operation.

Returns all questions. Admin operation.

*/
type GetAllQuestions struct {
	Context *middleware.Context
	Handler GetAllQuestionsHandler
}

func (o *GetAllQuestions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllQuestionsParams()
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
