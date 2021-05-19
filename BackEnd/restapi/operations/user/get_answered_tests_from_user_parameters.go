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

// NewGetAnsweredTestsFromUserParams creates a new GetAnsweredTestsFromUserParams object
//
// There are no default values defined in the spec.
func NewGetAnsweredTestsFromUserParams() GetAnsweredTestsFromUserParams {

	return GetAnsweredTestsFromUserParams{}
}

// GetAnsweredTestsFromUserParams contains all the bound params for the get answered tests from user operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAnsweredTestsFromUser
type GetAnsweredTestsFromUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	LikeTitle *string
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
	/*Indicates which element is first returned. In case of tie it unties with newdate first
	  In: query
	*/
	Orderby *string
	/*
	  In: query
	  Collection Format: pipes
	*/
	Tags [][]string
	/*Username of the user who has answered the publishedTests
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAnsweredTestsFromUserParams() beforehand.
func (o *GetAnsweredTestsFromUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLikeTitle, qhkLikeTitle, _ := qs.GetOK("likeTitle")
	if err := o.bindLikeTitle(qLikeTitle, qhkLikeTitle, route.Formats); err != nil {
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

	qOrderby, qhkOrderby, _ := qs.GetOK("orderby")
	if err := o.bindOrderby(qOrderby, qhkOrderby, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
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

// bindLikeTitle binds and validates parameter LikeTitle from query.
func (o *GetAnsweredTestsFromUserParams) bindLikeTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LikeTitle = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetAnsweredTestsFromUserParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetAnsweredTestsFromUserParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", *o.Limit, 0, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetAnsweredTestsFromUserParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetAnsweredTestsFromUserParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", *o.Offset, 0, false); err != nil {
		return err
	}

	return nil
}

// bindOrderby binds and validates parameter Orderby from query.
func (o *GetAnsweredTestsFromUserParams) bindOrderby(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Orderby = &raw

	if err := o.validateOrderby(formats); err != nil {
		return err
	}

	return nil
}

// validateOrderby carries on validations for parameter Orderby
func (o *GetAnsweredTestsFromUserParams) validateOrderby(formats strfmt.Registry) error {

	if err := validate.EnumCase("orderby", "query", *o.Orderby, []interface{}{"newDate", "oldDate", "moreFav", "lessFav", "moreTime", "lessTime"}, true); err != nil {
		return err
	}

	return nil
}

// bindTags binds and validates array parameter Tags from query.
//
// Arrays are parsed according to CollectionFormat: "pipes" (defaults to "csv" when empty).
func (o *GetAnsweredTestsFromUserParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvTags string
	if len(rawData) > 0 {
		qvTags = rawData[len(rawData)-1]
	}

	// CollectionFormat: pipes
	tagsIC := swag.SplitByFormat(qvTags, "pipes")
	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR [][]string
	for _, tagsIV := range tagsIC {
		// items.CollectionFormat: csv
		tagsIIC := swag.SplitByFormat(tagsIV, "csv")
		if len(tagsIIC) > 0 {

			var tagsIIR []string
			for _, tagsIIV := range tagsIIC {
				tagsII := tagsIIV

				tagsIIR = append(tagsIIR, tagsII)
			}

			tagsIR = append(tagsIR, tagsIIR)
		}
	}

	o.Tags = tagsIR

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *GetAnsweredTestsFromUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
