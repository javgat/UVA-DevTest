// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetAllEditQuestionsHandlerFunc turns a function with the right signature into a get all edit questions handler
type GetAllEditQuestionsHandlerFunc func(GetAllEditQuestionsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllEditQuestionsHandlerFunc) Handle(params GetAllEditQuestionsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetAllEditQuestionsHandler interface for that can handle valid get all edit questions params
type GetAllEditQuestionsHandler interface {
	Handle(GetAllEditQuestionsParams, *models.User) middleware.Responder
}

// NewGetAllEditQuestions creates a new http.Handler for the get all edit questions operation
func NewGetAllEditQuestions(ctx *middleware.Context, handler GetAllEditQuestionsHandler) *GetAllEditQuestions {
	return &GetAllEditQuestions{Context: ctx, Handler: handler}
}

/* GetAllEditQuestions swagger:route GET /allEditQuestions question getAllEditQuestions

Returns all non-published questions

Returns all non-published questions

*/
type GetAllEditQuestions struct {
	Context *middleware.Context
	Handler GetAllEditQuestionsHandler
}

func (o *GetAllEditQuestions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllEditQuestionsParams()
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
