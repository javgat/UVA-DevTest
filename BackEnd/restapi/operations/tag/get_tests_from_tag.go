// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetTestsFromTagHandlerFunc turns a function with the right signature into a get tests from tag handler
type GetTestsFromTagHandlerFunc func(GetTestsFromTagParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTestsFromTagHandlerFunc) Handle(params GetTestsFromTagParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTestsFromTagHandler interface for that can handle valid get tests from tag params
type GetTestsFromTagHandler interface {
	Handle(GetTestsFromTagParams, *models.User) middleware.Responder
}

// NewGetTestsFromTag creates a new http.Handler for the get tests from tag operation
func NewGetTestsFromTag(ctx *middleware.Context, handler GetTestsFromTagHandler) *GetTestsFromTag {
	return &GetTestsFromTag{Context: ctx, Handler: handler}
}

/* GetTestsFromTag swagger:route GET /tags/{tag}/tests tag getTestsFromTag

Returns all tests from a tag.

Returns all tests from a tag.

*/
type GetTestsFromTag struct {
	Context *middleware.Context
	Handler GetTestsFromTagHandler
}

func (o *GetTestsFromTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTestsFromTagParams()
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
