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

// Test test
//
// swagger:model Test
type Test struct {

	// acceso publico
	// Example: true
	// Required: true
	AccesoPublico *bool `json:"accesoPublico"`

	// description
	// Example: En este test se evaluaran los conocimientos respecto al lenguaje de programación Java
	// Required: true
	Description *string `json:"description"`

	// editable
	// Example: false
	// Required: true
	Editable *bool `json:"editable"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// max seconds
	// Example: 32600
	// Required: true
	MaxSeconds *int64 `json:"maxSeconds"`

	// title
	// Example: Test de introduccion a Java
	// Required: true
	Title *string `json:"title"`

	// username
	// Example: javgat
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this test
func (m *Test) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccesoPublico(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEditable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Test) validateAccesoPublico(formats strfmt.Registry) error {

	if err := validate.Required("accesoPublico", "body", m.AccesoPublico); err != nil {
		return err
	}

	return nil
}

func (m *Test) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Test) validateEditable(formats strfmt.Registry) error {

	if err := validate.Required("editable", "body", m.Editable); err != nil {
		return err
	}

	return nil
}

func (m *Test) validateMaxSeconds(formats strfmt.Registry) error {

	if err := validate.Required("maxSeconds", "body", m.MaxSeconds); err != nil {
		return err
	}

	return nil
}

func (m *Test) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Test) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this test based on context it is used
func (m *Test) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Test) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Test) UnmarshalBinary(b []byte) error {
	var res Test
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
