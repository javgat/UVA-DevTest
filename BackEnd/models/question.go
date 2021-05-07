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

// Question question
//
// swagger:model Question
type Question struct {

	// acceso publico no publicada
	// Example: false
	// Required: true
	AccesoPublicoNoPublicada *bool `json:"accesoPublicoNoPublicada"`

	// auto correct
	// Example: true
	// Required: true
	AutoCorrect *bool `json:"autoCorrect"`

	// editable
	// Example: false
	// Required: true
	Editable *bool `json:"editable"`

	// eleccion unica
	// Example: false
	EleccionUnica bool `json:"eleccionUnica,omitempty"`

	// estimated time
	// Example: 60
	// Required: true
	// Minimum: 0
	EstimatedTime *int64 `json:"estimatedTime"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// only present in GetQuestionsFromAnswer
	// Example: false
	IsRespondida bool `json:"isRespondida,omitempty"`

	// penalizacion
	// Required: true
	// Maximum: 100
	// Minimum: 0
	Penalizacion *int64 `json:"penalizacion"`

	// question
	// Example: ¿Cual es el lenguaje que tiene un nombre más largo de todos?
	// Required: true
	Question *string `json:"question"`

	// solucion
	// Example: Javadoc
	Solucion string `json:"solucion,omitempty"`

	// tipo pregunta
	// Required: true
	// Enum: [opciones string codigo]
	TipoPregunta *string `json:"tipoPregunta"`

	// title
	// Example: Paralelismo en C
	// Required: true
	Title *string `json:"title"`

	// username
	// Example: javgat
	// Required: true
	Username *string `json:"username"`

	// valor final
	// Minimum: 0
	ValorFinal *int64 `json:"valorFinal,omitempty"`
}

// Validate validates this question
func (m *Question) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccesoPublicoNoPublicada(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoCorrect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEditable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePenalizacion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTipoPregunta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
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

func (m *Question) validateAccesoPublicoNoPublicada(formats strfmt.Registry) error {

	if err := validate.Required("accesoPublicoNoPublicada", "body", m.AccesoPublicoNoPublicada); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateAutoCorrect(formats strfmt.Registry) error {

	if err := validate.Required("autoCorrect", "body", m.AutoCorrect); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateEditable(formats strfmt.Registry) error {

	if err := validate.Required("editable", "body", m.Editable); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateEstimatedTime(formats strfmt.Registry) error {

	if err := validate.Required("estimatedTime", "body", m.EstimatedTime); err != nil {
		return err
	}

	if err := validate.MinimumInt("estimatedTime", "body", *m.EstimatedTime, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Question) validatePenalizacion(formats strfmt.Registry) error {

	if err := validate.Required("penalizacion", "body", m.Penalizacion); err != nil {
		return err
	}

	if err := validate.MinimumInt("penalizacion", "body", *m.Penalizacion, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("penalizacion", "body", *m.Penalizacion, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateQuestion(formats strfmt.Registry) error {

	if err := validate.Required("question", "body", m.Question); err != nil {
		return err
	}

	return nil
}

var questionTypeTipoPreguntaPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["opciones","string","codigo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		questionTypeTipoPreguntaPropEnum = append(questionTypeTipoPreguntaPropEnum, v)
	}
}

const (

	// QuestionTipoPreguntaOpciones captures enum value "opciones"
	QuestionTipoPreguntaOpciones string = "opciones"

	// QuestionTipoPreguntaString captures enum value "string"
	QuestionTipoPreguntaString string = "string"

	// QuestionTipoPreguntaCodigo captures enum value "codigo"
	QuestionTipoPreguntaCodigo string = "codigo"
)

// prop value enum
func (m *Question) validateTipoPreguntaEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, questionTypeTipoPreguntaPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Question) validateTipoPregunta(formats strfmt.Registry) error {

	if err := validate.Required("tipoPregunta", "body", m.TipoPregunta); err != nil {
		return err
	}

	// value enum
	if err := m.validateTipoPreguntaEnum("tipoPregunta", "body", *m.TipoPregunta); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *Question) validateValorFinal(formats strfmt.Registry) error {
	if swag.IsZero(m.ValorFinal) { // not required
		return nil
	}

	if err := validate.MinimumInt("valorFinal", "body", *m.ValorFinal, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this question based on context it is used
func (m *Question) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Question) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Question) UnmarshalBinary(b []byte) error {
	var res Question
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
