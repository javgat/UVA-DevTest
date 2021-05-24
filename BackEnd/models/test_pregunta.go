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

// TestPregunta test pregunta
//
// swagger:model TestPregunta
type TestPregunta struct {

	// posicion
	// Example: 5
	// Required: true
	Posicion *int64 `json:"posicion"`

	// valor final
	// Example: 1
	// Required: true
	// Minimum: 0
	ValorFinal *int64 `json:"valorFinal"`
}

// Validate validates this test pregunta
func (m *TestPregunta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosicion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValorFinal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestPregunta) validatePosicion(formats strfmt.Registry) error {

	if err := validate.Required("posicion", "body", m.Posicion); err != nil {
		return err
	}

	return nil
}

func (m *TestPregunta) validateValorFinal(formats strfmt.Registry) error {

	if err := validate.Required("valorFinal", "body", m.ValorFinal); err != nil {
		return err
	}

	if err := validate.MinimumInt("valorFinal", "body", *m.ValorFinal, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this test pregunta based on context it is used
func (m *TestPregunta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestPregunta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestPregunta) UnmarshalBinary(b []byte) error {
	var res TestPregunta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
