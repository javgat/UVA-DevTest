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

// PasswordRecovery password recovery
//
// swagger:model PasswordRecovery
type PasswordRecovery struct {

	// mailtoken
	// Required: true
	Mailtoken *string `json:"mailtoken"`

	// newpass
	// Example: password
	// Required: true
	// Pattern: ^.{6,}$
	// Format: password
	Newpass *strfmt.Password `json:"newpass"`
}

// Validate validates this password recovery
func (m *PasswordRecovery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMailtoken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewpass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PasswordRecovery) validateMailtoken(formats strfmt.Registry) error {

	if err := validate.Required("mailtoken", "body", m.Mailtoken); err != nil {
		return err
	}

	return nil
}

func (m *PasswordRecovery) validateNewpass(formats strfmt.Registry) error {

	if err := validate.Required("newpass", "body", m.Newpass); err != nil {
		return err
	}

	if err := validate.Pattern("newpass", "body", m.Newpass.String(), `^.{6,}$`); err != nil {
		return err
	}

	if err := validate.FormatOf("newpass", "body", "password", m.Newpass.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this password recovery based on context it is used
func (m *PasswordRecovery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PasswordRecovery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordRecovery) UnmarshalBinary(b []byte) error {
	var res PasswordRecovery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
