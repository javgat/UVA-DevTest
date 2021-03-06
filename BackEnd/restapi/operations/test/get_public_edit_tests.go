// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPublicEditTestsHandlerFunc turns a function with the right signature into a get public edit tests handler
type GetPublicEditTestsHandlerFunc func(GetPublicEditTestsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublicEditTestsHandlerFunc) Handle(params GetPublicEditTestsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPublicEditTestsHandler interface for that can handle valid get public edit tests params
type GetPublicEditTestsHandler interface {
	Handle(GetPublicEditTestsParams, *models.User) middleware.Responder
}

// NewGetPublicEditTests creates a new http.Handler for the get public edit tests operation
func NewGetPublicEditTests(ctx *middleware.Context, handler GetPublicEditTestsHandler) *GetPublicEditTests {
	return &GetPublicEditTests{Context: ctx, Handler: handler}
}

/* GetPublicEditTests swagger:route GET /publicEditTests test getPublicEditTests

Returns all public non-published tests

Returns all public non-published tests

*/
type GetPublicEditTests struct {
	Context *middleware.Context
	Handler GetPublicEditTestsHandler
}

func (o *GetPublicEditTests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPublicEditTestsParams()
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
