// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"uva-devtest/models"
)

// NewPostPublishedTestParams creates a new PostPublishedTestParams object
//
// There are no default values defined in the spec.
func NewPostPublishedTestParams() PostPublishedTestParams {

	return PostPublishedTestParams{}
}

// PostPublishedTestParams contains all the bound params for the post published test operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostPublishedTest
type PostPublishedTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Some attributes of the new test
	  Required: true
	  In: body
	*/
	PublishTestParams *models.PublishTestParams
	/*Id of the test to publish
	  Required: true
	  In: path
	*/
	Testid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostPublishedTestParams() beforehand.
func (o *PostPublishedTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PublishTestParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("publishTestParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("publishTestParams", "body", "", err))
			}
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
				o.PublishTestParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("publishTestParams", "body", ""))
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

// bindTestid binds and validates parameter Testid from path.
func (o *PostPublishedTestParams) bindTestid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
