// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginUser login user
//
// swagger:model LoginUser
type LoginUser struct {

	// loginid
	// Example: carlosg72 || carlos@mail.com
	// Required: true
	// Max Length: 100
	Loginid *string `json:"loginid"`

	// pass
	// Example: password
	// Required: true
	// Pattern: ^.{6,}$
	// Format: password
	Pass *strfmt.Password `json:"pass"`
}

// Validate validates this login user
func (m *LoginUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoginid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginUser) validateLoginid(formats strfmt.Registry) error {

	if err := validate.Required("loginid", "body", m.Loginid); err != nil {
		return err
	}

	if err := validate.MaxLength("loginid", "body", *m.Loginid, 100); err != nil {
		return err
	}

	return nil
}

func (m *LoginUser) validatePass(formats strfmt.Registry) error {

	if err := validate.Required("pass", "body", m.Pass); err != nil {
		return err
	}

	if err := validate.Pattern("pass", "body", m.Pass.String(), `^.{6,}$`); err != nil {
		return err
	}

	if err := validate.FormatOf("pass", "body", "password", m.Pass.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login user based on context it is used
func (m *LoginUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginUser) UnmarshalBinary(b []byte) error {
	var res LoginUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
