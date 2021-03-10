// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"uva-devtest/models"
)

// NewLoginParams creates a new LoginParams object
//
// There are no default values defined in the spec.
func NewLoginParams() LoginParams {

	return LoginParams{}
}

// LoginParams contains all the bound params for the login operation
// typically these are obtained from a http.Request
//
// swagger:parameters Login
type LoginParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*User who is trying to generate a token
	  Required: true
	  In: body
	*/
	LoginUser *models.LoginUser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLoginParams() beforehand.
func (o *LoginParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.LoginUser
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("loginUser", "body", ""))
			} else {
				res = append(res, errors.NewParseError("loginUser", "body", "", err))
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
				o.LoginUser = &body
			}
		}
	} else {
		res = append(res, errors.Required("loginUser", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
