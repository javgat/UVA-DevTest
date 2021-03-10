// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetUserTeamRoleOKCode is the HTTP code returned for type GetUserTeamRoleOK
const GetUserTeamRoleOKCode int = 200

/*GetUserTeamRoleOK Role found

swagger:response getUserTeamRoleOK
*/
type GetUserTeamRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.TeamRole `json:"body,omitempty"`
}

// NewGetUserTeamRoleOK creates GetUserTeamRoleOK with default headers values
func NewGetUserTeamRoleOK() *GetUserTeamRoleOK {

	return &GetUserTeamRoleOK{}
}

// WithPayload adds the payload to the get user team role o k response
func (o *GetUserTeamRoleOK) WithPayload(payload *models.TeamRole) *GetUserTeamRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user team role o k response
func (o *GetUserTeamRoleOK) SetPayload(payload *models.TeamRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserTeamRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserTeamRoleBadRequestCode is the HTTP code returned for type GetUserTeamRoleBadRequest
const GetUserTeamRoleBadRequestCode int = 400

/*GetUserTeamRoleBadRequest Incorrect Request, or invalida data

swagger:response getUserTeamRoleBadRequest
*/
type GetUserTeamRoleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserTeamRoleBadRequest creates GetUserTeamRoleBadRequest with default headers values
func NewGetUserTeamRoleBadRequest() *GetUserTeamRoleBadRequest {

	return &GetUserTeamRoleBadRequest{}
}

// WithPayload adds the payload to the get user team role bad request response
func (o *GetUserTeamRoleBadRequest) WithPayload(payload *models.Error) *GetUserTeamRoleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user team role bad request response
func (o *GetUserTeamRoleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserTeamRoleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserTeamRoleForbiddenCode is the HTTP code returned for type GetUserTeamRoleForbidden
const GetUserTeamRoleForbiddenCode int = 403

/*GetUserTeamRoleForbidden Not authorized to this content

swagger:response getUserTeamRoleForbidden
*/
type GetUserTeamRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserTeamRoleForbidden creates GetUserTeamRoleForbidden with default headers values
func NewGetUserTeamRoleForbidden() *GetUserTeamRoleForbidden {

	return &GetUserTeamRoleForbidden{}
}

// WithPayload adds the payload to the get user team role forbidden response
func (o *GetUserTeamRoleForbidden) WithPayload(payload *models.Error) *GetUserTeamRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user team role forbidden response
func (o *GetUserTeamRoleForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserTeamRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserTeamRoleGoneCode is the HTTP code returned for type GetUserTeamRoleGone
const GetUserTeamRoleGoneCode int = 410

/*GetUserTeamRoleGone That user (password and name) does not exist

swagger:response getUserTeamRoleGone
*/
type GetUserTeamRoleGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserTeamRoleGone creates GetUserTeamRoleGone with default headers values
func NewGetUserTeamRoleGone() *GetUserTeamRoleGone {

	return &GetUserTeamRoleGone{}
}

// WithPayload adds the payload to the get user team role gone response
func (o *GetUserTeamRoleGone) WithPayload(payload *models.Error) *GetUserTeamRoleGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user team role gone response
func (o *GetUserTeamRoleGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserTeamRoleGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserTeamRoleInternalServerErrorCode is the HTTP code returned for type GetUserTeamRoleInternalServerError
const GetUserTeamRoleInternalServerErrorCode int = 500

/*GetUserTeamRoleInternalServerError Internal error

swagger:response getUserTeamRoleInternalServerError
*/
type GetUserTeamRoleInternalServerError struct {
}

// NewGetUserTeamRoleInternalServerError creates GetUserTeamRoleInternalServerError with default headers values
func NewGetUserTeamRoleInternalServerError() *GetUserTeamRoleInternalServerError {

	return &GetUserTeamRoleInternalServerError{}
}

// WriteResponse to the client
func (o *GetUserTeamRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
