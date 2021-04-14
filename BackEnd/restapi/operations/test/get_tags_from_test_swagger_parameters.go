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

// NewGetTagsFromTestParams creates a new GetTagsFromTestParams object
//
// There are no default values defined in the spec.
func NewGetTagsFromTestParams() GetTagsFromTestParams {

	return GetTagsFromTestParams{}
}

// GetTagsFromTestParams contains all the bound params for the get tags from test operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTagsFromTest
type GetTagsFromTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of the tset to find its tags
	  Required: true
	  In: path
	*/
	Testid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTagsFromTestParams() beforehand.
func (o *GetTagsFromTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

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
func (o *GetTagsFromTestParams) bindTestid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
