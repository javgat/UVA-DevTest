// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetTagFromTestHandlerFunc turns a function with the right signature into a get tag from test handler
type GetTagFromTestHandlerFunc func(GetTagFromTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTagFromTestHandlerFunc) Handle(params GetTagFromTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTagFromTestHandler interface for that can handle valid get tag from test params
type GetTagFromTestHandler interface {
	Handle(GetTagFromTestParams, *models.User) middleware.Responder
}

// NewGetTagFromTest creates a new http.Handler for the get tag from test operation
func NewGetTagFromTest(ctx *middleware.Context, handler GetTagFromTestHandler) *GetTagFromTest {
	return &GetTagFromTest{Context: ctx, Handler: handler}
}

/* GetTagFromTest swagger:route GET /tests/{testid}/tags/{tag} test getTagFromTest

Returns a tag from a test.

Returns a tag from a test.

*/
type GetTagFromTest struct {
	Context *middleware.Context
	Handler GetTagFromTestHandler
}

func (o *GetTagFromTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTagFromTestParams()
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
