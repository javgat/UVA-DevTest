// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetTeamsOfUserParams creates a new GetTeamsOfUserParams object
//
// There are no default values defined in the spec.
func NewGetTeamsOfUserParams() GetTeamsOfUserParams {

	return GetTeamsOfUserParams{}
}

// GetTeamsOfUserParams contains all the bound params for the get teams of user operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTeamsOfUser
type GetTeamsOfUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	LikeStartTeamname *string
	/*
	  In: query
	*/
	LikeTeamname *string
	/*max number of elements to be returned
	  Minimum: 0
	  In: query
	*/
	Limit *int64
	/*first elements to be skipped at being returned
	  Minimum: 0
	  In: query
	*/
	Offset *int64
	/*Username of the user to get their teams
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTeamsOfUserParams() beforehand.
func (o *GetTeamsOfUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLikeStartTeamname, qhkLikeStartTeamname, _ := qs.GetOK("likeStartTeamname")
	if err := o.bindLikeStartTeamname(qLikeStartTeamname, qhkLikeStartTeamname, route.Formats); err != nil {
		res = append(res, err)
	}

	qLikeTeamname, qhkLikeTeamname, _ := qs.GetOK("likeTeamname")
	if err := o.bindLikeTeamname(qLikeTeamname, qhkLikeTeamname, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLikeStartTeamname binds and validates parameter LikeStartTeamname from query.
func (o *GetTeamsOfUserParams) bindLikeStartTeamname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LikeStartTeamname = &raw

	return nil
}

// bindLikeTeamname binds and validates parameter LikeTeamname from query.
func (o *GetTeamsOfUserParams) bindLikeTeamname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LikeTeamname = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetTeamsOfUserParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetTeamsOfUserParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", *o.Limit, 0, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetTeamsOfUserParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetTeamsOfUserParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", *o.Offset, 0, false); err != nil {
		return err
	}

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *GetTeamsOfUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
