// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"uva-devtest/models"
)

// NewInviteTeamToPublishedTestParams creates a new InviteTeamToPublishedTestParams object
//
// There are no default values defined in the spec.
func NewInviteTeamToPublishedTestParams() InviteTeamToPublishedTestParams {

	return InviteTeamToPublishedTestParams{}
}

// InviteTeamToPublishedTestParams contains all the bound params for the invite team to published test operation
// typically these are obtained from a http.Request
//
// swagger:parameters InviteTeamToPublishedTest
type InviteTeamToPublishedTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Message sent to the users as a notification
	  In: body
	*/
	Message *models.Message
	/*Teamname of the team to invite to test
	  Required: true
	  In: path
	*/
	Teamname string
	/*Id of the test to find
	  Required: true
	  In: path
	*/
	Testid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewInviteTeamToPublishedTestParams() beforehand.
func (o *InviteTeamToPublishedTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Message
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("message", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Message = &body
			}
		}
	}

	rTeamname, rhkTeamname, _ := route.Params.GetOK("teamname")
	if err := o.bindTeamname(rTeamname, rhkTeamname, route.Formats); err != nil {
		res = append(res, err)
	}

	rTestid, rhkTestid, _ := route.Params.GetOK("testid")
	if err := o.bindTestid(rTestid, rhkTestid, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTeamname binds and validates parameter Teamname from path.
func (o *InviteTeamToPublishedTestParams) bindTeamname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Teamname = raw

	return nil
}

// bindTestid binds and validates parameter Testid from path.
func (o *InviteTeamToPublishedTestParams) bindTestid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("testid", "path", "int64", raw)
	}
	o.Testid = value

	return nil
}
