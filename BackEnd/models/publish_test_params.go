// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublishTestParams publish test params
//
// swagger:model PublishTestParams
type PublishTestParams struct {

	// True si no necesitaras invitacion para hacer el test
	// Required: true
	AccesoPublico *bool `json:"accesoPublico"`

	// auto correct
	// Required: true
	AutoCorrect *bool `json:"autoCorrect"`

	// max minutes
	// Example: 4
	// Required: true
	MaxMinutes *int64 `json:"maxMinutes"`

	// tiempo estricto
	// Required: true
	TiempoEstricto *bool `json:"tiempoEstricto"`

	// title
	// Example: Nuevo titulo
	// Required: true
	// Max Length: 100
	Title *string `json:"title"`

	// visibilidad
	// Required: true
	// Enum: [alEntregar alCorregir manual]
	Visibilidad *string `json:"visibilidad"`
}

// Validate validates this publish test params
func (m *PublishTestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccesoPublico(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoCorrect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxMinutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTiempoEstricto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibilidad(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublishTestParams) validateAccesoPublico(formats strfmt.Registry) error {

	if err := validate.Required("accesoPublico", "body", m.AccesoPublico); err != nil {
		return err
	}

	return nil
}

func (m *PublishTestParams) validateAutoCorrect(formats strfmt.Registry) error {

	if err := validate.Required("autoCorrect", "body", m.AutoCorrect); err != nil {
		return err
	}

	return nil
}

func (m *PublishTestParams) validateMaxMinutes(formats strfmt.Registry) error {

	if err := validate.Required("maxMinutes", "body", m.MaxMinutes); err != nil {
		return err
	}

	return nil
}

func (m *PublishTestParams) validateTiempoEstricto(formats strfmt.Registry) error {

	if err := validate.Required("tiempoEstricto", "body", m.TiempoEstricto); err != nil {
		return err
	}

	return nil
}

func (m *PublishTestParams) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 100); err != nil {
		return err
	}

	return nil
}

var publishTestParamsTypeVisibilidadPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alEntregar","alCorregir","manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		publishTestParamsTypeVisibilidadPropEnum = append(publishTestParamsTypeVisibilidadPropEnum, v)
	}
}

const (

	// PublishTestParamsVisibilidadAlEntregar captures enum value "alEntregar"
	PublishTestParamsVisibilidadAlEntregar string = "alEntregar"

	// PublishTestParamsVisibilidadAlCorregir captures enum value "alCorregir"
	PublishTestParamsVisibilidadAlCorregir string = "alCorregir"

	// PublishTestParamsVisibilidadManual captures enum value "manual"
	PublishTestParamsVisibilidadManual string = "manual"
)

// prop value enum
func (m *PublishTestParams) validateVisibilidadEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, publishTestParamsTypeVisibilidadPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PublishTestParams) validateVisibilidad(formats strfmt.Registry) error {

	if err := validate.Required("visibilidad", "body", m.Visibilidad); err != nil {
		return err
	}

	// value enum
	if err := m.validateVisibilidadEnum("visibilidad", "body", *m.Visibilidad); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this publish test params based on context it is used
func (m *PublishTestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublishTestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublishTestParams) UnmarshalBinary(b []byte) error {
	var res PublishTestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
