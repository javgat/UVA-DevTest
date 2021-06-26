// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeletePruebaParams creates a new DeletePruebaParams object
//
// There are no default values defined in the spec.
func NewDeletePruebaParams() DeletePruebaParams {

	return DeletePruebaParams{}
}

// DeletePruebaParams contains all the bound params for the delete prueba operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeletePrueba
type DeletePruebaParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id of the prueba to delete
	  Required: true
	  In: path
	*/
	Pruebaid int64
	/*Id of the question to delete its prueba
	  Required: true
	  In: path
	*/
	Questionid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeletePruebaParams() beforehand.
func (o *DeletePruebaParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPruebaid, rhkPruebaid, _ := route.Params.GetOK("pruebaid")
	if err := o.bindPruebaid(rPruebaid, rhkPruebaid, route.Formats); err != nil {
		res = append(res, err)
	}

	rQuestionid, rhkQuestionid, _ := route.Params.GetOK("questionid")
	if err := o.bindQuestionid(rQuestionid, rhkQuestionid, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPruebaid binds and validates parameter Pruebaid from path.
func (o *DeletePruebaParams) bindPruebaid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pruebaid", "path", "int64", raw)
	}
	o.Pruebaid = value

	return nil
}

// bindQuestionid binds and validates parameter Questionid from path.
func (o *DeletePruebaParams) bindQuestionid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
