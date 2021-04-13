// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSharedQuestionFromUserParams creates a new GetSharedQuestionFromUserParams object
//
// There are no default values defined in the spec.
func NewGetSharedQuestionFromUserParams() GetSharedQuestionFromUserParams {

	return GetSharedQuestionFromUserParams{}
}

// GetSharedQuestionFromUserParams contains all the bound params for the get shared question from user operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSharedQuestionFromUser
type GetSharedQuestionFromUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id of the question to find
	  Required: true
	  In: path
	*/
	Questionid int64
	/*Username of the user who is shared the question
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSharedQuestionFromUserParams() beforehand.
func (o *GetSharedQuestionFromUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rQuestionid, rhkQuestionid, _ := route.Params.GetOK("questionid")
	if err := o.bindQuestionid(rQuestionid, rhkQuestionid, route.Formats); err != nil {
		res = append(res, err)
	}

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindQuestionid binds and validates parameter Questionid from path.
func (o *GetSharedQuestionFromUserParams) bindQuestionid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindUsername binds and validates parameter Username from path.
func (o *GetSharedQuestionFromUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
