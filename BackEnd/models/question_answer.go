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

// QuestionAnswer question answer
//
// swagger:model QuestionAnswer
type QuestionAnswer struct {

	// corregida
	// Example: true
	// Required: true
	Corregida *bool `json:"corregida"`

	// id pregunta
	// Example: 1
	// Required: true
	IDPregunta *int64 `json:"idPregunta"`

	// id respuesta
	// Example: 1
	// Required: true
	IDRespuesta *int64 `json:"idRespuesta"`

	// indices opciones
	IndicesOpciones []int64 `json:"indicesOpciones"`

	// puntuacion
	// Example: 1
	// Required: true
	Puntuacion *int64 `json:"puntuacion"`

	// respuesta
	// Example: Javadoc
	Respuesta string `json:"respuesta,omitempty"`
}

// Validate validates this question answer
func (m *QuestionAnswer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorregida(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDPregunta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDRespuesta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePuntuacion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuestionAnswer) validateCorregida(formats strfmt.Registry) error {

	if err := validate.Required("corregida", "body", m.Corregida); err != nil {
		return err
	}

	return nil
}

func (m *QuestionAnswer) validateIDPregunta(formats strfmt.Registry) error {

	if err := validate.Required("idPregunta", "body", m.IDPregunta); err != nil {
		return err
	}

	return nil
}

func (m *QuestionAnswer) validateIDRespuesta(formats strfmt.Registry) error {

	if err := validate.Required("idRespuesta", "body", m.IDRespuesta); err != nil {
		return err
	}

	return nil
}

func (m *QuestionAnswer) validatePuntuacion(formats strfmt.Registry) error {

	if err := validate.Required("puntuacion", "body", m.Puntuacion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this question answer based on context it is used
func (m *QuestionAnswer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuestionAnswer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuestionAnswer) UnmarshalBinary(b []byte) error {
	var res QuestionAnswer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
