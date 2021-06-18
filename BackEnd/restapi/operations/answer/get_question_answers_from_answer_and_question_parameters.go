// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetQuestionAnswersFromAnswerAndQuestionParams creates a new GetQuestionAnswersFromAnswerAndQuestionParams object
//
// There are no default values defined in the spec.
func NewGetQuestionAnswersFromAnswerAndQuestionParams() GetQuestionAnswersFromAnswerAndQuestionParams {

	return GetQuestionAnswersFromAnswerAndQuestionParams{}
}

// GetQuestionAnswersFromAnswerAndQuestionParams contains all the bound params for the get question answers from answer and question operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetQuestionAnswersFromAnswerAndQuestion
type GetQuestionAnswersFromAnswerAndQuestionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of the answer
	  Required: true
	  In: path
	*/
	Answerid int64
	/*Id of the question
	  Required: true
	  In: path
	*/
	Questionid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetQuestionAnswersFromAnswerAndQuestionParams() beforehand.
func (o *GetQuestionAnswersFromAnswerAndQuestionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAnswerid, rhkAnswerid, _ := route.Params.GetOK("answerid")
	if err := o.bindAnswerid(rAnswerid, rhkAnswerid, route.Formats); err != nil {
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

// bindAnswerid binds and validates parameter Answerid from path.
func (o *GetQuestionAnswersFromAnswerAndQuestionParams) bindAnswerid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("answerid", "path", "int64", raw)
	}
	o.Answerid = value

	return nil
}

// bindQuestionid binds and validates parameter Questionid from path.
func (o *GetQuestionAnswersFromAnswerAndQuestionParams) bindQuestionid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
