// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPublishedTestsHandlerFunc turns a function with the right signature into a get published tests handler
type GetPublishedTestsHandlerFunc func(GetPublishedTestsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublishedTestsHandlerFunc) Handle(params GetPublishedTestsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPublishedTestsHandler interface for that can handle valid get published tests params
type GetPublishedTestsHandler interface {
	Handle(GetPublishedTestsParams, *models.User) middleware.Responder
}

// NewGetPublishedTests creates a new http.Handler for the get published tests operation
func NewGetPublishedTests(ctx *middleware.Context, handler GetPublishedTestsHandler) *GetPublishedTests {
	return &GetPublishedTests{Context: ctx, Handler: handler}
}

/* GetPublishedTests swagger:route GET /publishedTests publishedTest getPublishedTests

Returns all publishedTests

Returns all publishedTests

*/
type GetPublishedTests struct {
	Context *middleware.Context
	Handler GetPublishedTestsHandler
}

func (o *GetPublishedTests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPublishedTestsParams()
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
