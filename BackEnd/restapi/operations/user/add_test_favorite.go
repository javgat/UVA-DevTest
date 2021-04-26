// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// AddTestFavoriteHandlerFunc turns a function with the right signature into a add test favorite handler
type AddTestFavoriteHandlerFunc func(AddTestFavoriteParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTestFavoriteHandlerFunc) Handle(params AddTestFavoriteParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AddTestFavoriteHandler interface for that can handle valid add test favorite params
type AddTestFavoriteHandler interface {
	Handle(AddTestFavoriteParams, *models.User) middleware.Responder
}

// NewAddTestFavorite creates a new http.Handler for the add test favorite operation
func NewAddTestFavorite(ctx *middleware.Context, handler AddTestFavoriteHandler) *AddTestFavorite {
	return &AddTestFavorite{Context: ctx, Handler: handler}
}

/* AddTestFavorite swagger:route PUT /users/{username}/favoriteTests/{testid} user addTestFavorite

Marks test as favorite for the user

Marks test as favorite for the user

*/
type AddTestFavorite struct {
	Context *middleware.Context
	Handler AddTestFavoriteHandler
}

func (o *AddTestFavorite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddTestFavoriteParams()
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