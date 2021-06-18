// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetSharedQuestionFromUserHandlerFunc turns a function with the right signature into a get shared question from user handler
type GetSharedQuestionFromUserHandlerFunc func(GetSharedQuestionFromUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSharedQuestionFromUserHandlerFunc) Handle(params GetSharedQuestionFromUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetSharedQuestionFromUserHandler interface for that can handle valid get shared question from user params
type GetSharedQuestionFromUserHandler interface {
	Handle(GetSharedQuestionFromUserParams, *models.User) middleware.Responder
}

// NewGetSharedQuestionFromUser creates a new http.Handler for the get shared question from user operation
func NewGetSharedQuestionFromUser(ctx *middleware.Context, handler GetSharedQuestionFromUserHandler) *GetSharedQuestionFromUser {
	return &GetSharedQuestionFromUser{Context: ctx, Handler: handler}
}

/* GetSharedQuestionFromUser swagger:route GET /users/{username}/sharedQuestions/{questionid} user getSharedQuestionFromUser

Returns a question shared to a user

Returns a question shared to a user

*/
type GetSharedQuestionFromUser struct {
	Context *middleware.Context
	Handler GetSharedQuestionFromUserHandler
}

func (o *GetSharedQuestionFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSharedQuestionFromUserParams()
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
