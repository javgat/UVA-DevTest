// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAddTagToTestParams creates a new AddTagToTestParams object
//
// There are no default values defined in the spec.
func NewAddTagToTestParams() AddTagToTestParams {

	return AddTagToTestParams{}
}

// AddTagToTestParams contains all the bound params for the add tag to test operation
// typically these are obtained from a http.Request
//
// swagger:parameters AddTagToTest
type AddTagToTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Tag to add
	  Required: true
	  In: path
	*/
	Tag string
	/*Id of the test to add a tag
	  Required: true
	  In: path
	*/
	Testid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddTagToTestParams() beforehand.
func (o *AddTagToTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTag, rhkTag, _ := route.Params.GetOK("tag")
	if err := o.bindTag(rTag, rhkTag, route.Formats); err != nil {
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

// bindTag binds and validates parameter Tag from path.
func (o *AddTagToTestParams) bindTag(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Tag = raw

	return nil
}

// bindTestid binds and validates parameter Testid from path.
func (o *AddTagToTestParams) bindTestid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
