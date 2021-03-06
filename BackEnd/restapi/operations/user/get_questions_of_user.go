// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetQuestionsOfUserHandlerFunc turns a function with the right signature into a get questions of user handler
type GetQuestionsOfUserHandlerFunc func(GetQuestionsOfUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQuestionsOfUserHandlerFunc) Handle(params GetQuestionsOfUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetQuestionsOfUserHandler interface for that can handle valid get questions of user params
type GetQuestionsOfUserHandler interface {
	Handle(GetQuestionsOfUserParams, *models.User) middleware.Responder
}

// NewGetQuestionsOfUser creates a new http.Handler for the get questions of user operation
func NewGetQuestionsOfUser(ctx *middleware.Context, handler GetQuestionsOfUserHandler) *GetQuestionsOfUser {
	return &GetQuestionsOfUser{Context: ctx, Handler: handler}
}

/* GetQuestionsOfUser swagger:route GET /users/{username}/questions user getQuestionsOfUser

Returns all questions owned by the user

Returns all questions owned by the user

*/
type GetQuestionsOfUser struct {
	Context *middleware.Context
	Handler GetQuestionsOfUserHandler
}

func (o *GetQuestionsOfUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetQuestionsOfUserParams()
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
