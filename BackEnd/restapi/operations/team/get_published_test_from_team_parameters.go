// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetPublishedTestFromTeamParams creates a new GetPublishedTestFromTeamParams object
//
// There are no default values defined in the spec.
func NewGetPublishedTestFromTeamParams() GetPublishedTestFromTeamParams {

	return GetPublishedTestFromTeamParams{}
}

// GetPublishedTestFromTeamParams contains all the bound params for the get published test from team operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetPublishedTestFromTeam
type GetPublishedTestFromTeamParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Teamname of the team to get its publishedTest
	  Required: true
	  In: path
	*/
	Teamname string
	/*Id of the publishedTest to find
	  Required: true
	  In: path
	*/
	Testid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPublishedTestFromTeamParams() beforehand.
func (o *GetPublishedTestFromTeamParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTeamname, rhkTeamname, _ := route.Params.GetOK("teamname")
	if err := o.bindTeamname(rTeamname, rhkTeamname, route.Formats); err != nil {
		res = append(res, err)
	}

	rTestid, rhkTestid, _ := route.Params.GetOK("testid")
	if err := o.bindTestid(rTestid, rhkTestid, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTeamname binds and validates parameter Teamname from path.
func (o *GetPublishedTestFromTeamParams) bindTeamname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Teamname = raw

	return nil
}

// bindTestid binds and validates parameter Testid from path.
func (o *GetPublishedTestFromTeamParams) bindTestid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("testid", "path", "int64", raw)
	}
	o.Testid = value

	return nil
}
