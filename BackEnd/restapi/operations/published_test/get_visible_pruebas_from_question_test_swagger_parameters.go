// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetVisiblePruebasFromQuestionTestParams creates a new GetVisiblePruebasFromQuestionTestParams object
//
// There are no default values defined in the spec.
func NewGetVisiblePruebasFromQuestionTestParams() GetVisiblePruebasFromQuestionTestParams {

	return GetVisiblePruebasFromQuestionTestParams{}
}

// GetVisiblePruebasFromQuestionTestParams contains all the bound params for the get visible pruebas from question test operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetVisiblePruebasFromQuestionTest
type GetVisiblePruebasFromQuestionTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of the question to find its pruebas visibles
	  Required: true
	  In: path
	*/
	Questionid int64
	/*Id of the test
	  Required: true
	  In: path
	*/
	Testid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVisiblePruebasFromQuestionTestParams() beforehand.
func (o *GetVisiblePruebasFromQuestionTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rQuestionid, rhkQuestionid, _ := route.Params.GetOK("questionid")
	if err := o.bindQuestionid(rQuestionid, rhkQuestionid, route.Formats); err != nil {
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

// bindQuestionid binds and validates parameter Questionid from path.
func (o *GetVisiblePruebasFromQuestionTestParams) bindQuestionid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("questionid", "path", "int64", raw)
	}
	o.Questionid = value

	return nil
}

// bindTestid binds and validates parameter Testid from path.
func (o *GetVisiblePruebasFromQuestionTestParams) bindTestid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
