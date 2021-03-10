// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetUsersFromTeamHandlerFunc turns a function with the right signature into a get users from team handler
type GetUsersFromTeamHandlerFunc func(GetUsersFromTeamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersFromTeamHandlerFunc) Handle(params GetUsersFromTeamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetUsersFromTeamHandler interface for that can handle valid get users from team params
type GetUsersFromTeamHandler interface {
	Handle(GetUsersFromTeamParams, *models.User) middleware.Responder
}

// NewGetUsersFromTeam creates a new http.Handler for the get users from team operation
func NewGetUsersFromTeam(ctx *middleware.Context, handler GetUsersFromTeamHandler) *GetUsersFromTeam {
	return &GetUsersFromTeam{Context: ctx, Handler: handler}
}

/* GetUsersFromTeam swagger:route GET /teams/{teamname}/users user team getUsersFromTeam

Returns all users that are members of a team

Returns all users that are members of a team

*/
type GetUsersFromTeam struct {
	Context *middleware.Context
	Handler GetUsersFromTeamHandler
}

func (o *GetUsersFromTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersFromTeamParams()
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
