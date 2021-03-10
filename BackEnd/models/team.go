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

// Team team
//
// swagger:model Team
type Team struct {

	// description
	// Example: DevTest Team
	Description string `json:"description,omitempty"`

	// teamname
	// Example: devtestTeam
	// Required: true
	// Pattern: ^[^@ \t\r\n]+$
	Teamname *string `json:"teamname"`
}

// Validate validates this team
func (m *Team) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Team) validateTeamname(formats strfmt.Registry) error {

	if err := validate.Required("teamname", "body", m.Teamname); err != nil {
		return err
	}

	if err := validate.Pattern("teamname", "body", *m.Teamname, `^[^@ \t\r\n]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this team based on context it is used
func (m *Team) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Team) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Team) UnmarshalBinary(b []byte) error {
	var res Team
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
