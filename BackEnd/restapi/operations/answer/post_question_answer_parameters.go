// Code generated by go-swagger; DO NOT EDIT.

package answer

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

// NewPostQuestionAnswerParams creates a new PostQuestionAnswerParams object
//
// There are no default values defined in the spec.
func NewPostQuestionAnswerParams() PostQuestionAnswerParams {

	return PostQuestionAnswerParams{}
}

// PostQuestionAnswerParams contains all the bound params for the post question answer operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostQuestionAnswer
type PostQuestionAnswerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of the answer
	  Required: true
	  In: path
	*/
	Answerid int64
	/*QuestionAnswer to post
	  Required: true
	  In: body
	*/
	QuestionAnswer *models.QuestionAnswer
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostQuestionAnswerParams() beforehand.
func (o *PostQuestionAnswerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAnswerid, rhkAnswerid, _ := route.Params.GetOK("answerid")
	if err := o.bindAnswerid(rAnswerid, rhkAnswerid, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.QuestionAnswer
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("questionAnswer", "body", ""))
			} else {
				res = append(res, errors.NewParseError("questionAnswer", "body", "", err))
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
				o.QuestionAnswer = &body
			}
		}
	} else {
		res = append(res, errors.Required("questionAnswer", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAnswerid binds and validates parameter Answerid from path.
func (o *PostQuestionAnswerParams) bindAnswerid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
