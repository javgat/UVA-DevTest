// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetEditQuestionsOfUserHandlerFunc turns a function with the right signature into a get edit questions of user handler
type GetEditQuestionsOfUserHandlerFunc func(GetEditQuestionsOfUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEditQuestionsOfUserHandlerFunc) Handle(params GetEditQuestionsOfUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetEditQuestionsOfUserHandler interface for that can handle valid get edit questions of user params
type GetEditQuestionsOfUserHandler interface {
	Handle(GetEditQuestionsOfUserParams, *models.User) middleware.Responder
}

// NewGetEditQuestionsOfUser creates a new http.Handler for the get edit questions of user operation
func NewGetEditQuestionsOfUser(ctx *middleware.Context, handler GetEditQuestionsOfUserHandler) *GetEditQuestionsOfUser {
	return &GetEditQuestionsOfUser{Context: ctx, Handler: handler}
}

/* GetEditQuestionsOfUser swagger:route GET /users/{username}/editQuestions user getEditQuestionsOfUser

Returns all non-published questions owned by the user

Returns all non-published questions owned by the user

*/
type GetEditQuestionsOfUser struct {
	Context *middleware.Context
	Handler GetEditQuestionsOfUserHandler
}

func (o *GetEditQuestionsOfUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetEditQuestionsOfUserParams()
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
