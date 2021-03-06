// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// SetAnswerNotVisibleHandlerFunc turns a function with the right signature into a set answer not visible handler
type SetAnswerNotVisibleHandlerFunc func(SetAnswerNotVisibleParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn SetAnswerNotVisibleHandlerFunc) Handle(params SetAnswerNotVisibleParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// SetAnswerNotVisibleHandler interface for that can handle valid set answer not visible params
type SetAnswerNotVisibleHandler interface {
	Handle(SetAnswerNotVisibleParams, *models.User) middleware.Responder
}

// NewSetAnswerNotVisible creates a new http.Handler for the set answer not visible operation
func NewSetAnswerNotVisible(ctx *middleware.Context, handler SetAnswerNotVisibleHandler) *SetAnswerNotVisible {
	return &SetAnswerNotVisible{Context: ctx, Handler: handler}
}

/* SetAnswerNotVisible swagger:route DELETE /answers/{answerid}/visible answer setAnswerNotVisible

Marks answer as not visible. Only for TestAdmins

Marks answer as not visible

*/
type SetAnswerNotVisible struct {
	Context *middleware.Context
	Handler SetAnswerNotVisibleHandler
}

func (o *SetAnswerNotVisible) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSetAnswerNotVisibleParams()
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
