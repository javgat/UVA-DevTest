// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTeamOKCode is the HTTP code returned for type GetTeamOK
const GetTeamOKCode int = 200

/*GetTeamOK team found

swagger:response getTeamOK
*/
type GetTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.Team `json:"body,omitempty"`
}

// NewGetTeamOK creates GetTeamOK with default headers values
func NewGetTeamOK() *GetTeamOK {

	return &GetTeamOK{}
}

// WithPayload adds the payload to the get team o k response
func (o *GetTeamOK) WithPayload(payload *models.Team) *GetTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team o k response
func (o *GetTeamOK) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeamBadRequestCode is the HTTP code returned for type GetTeamBadRequest
const GetTeamBadRequestCode int = 400

/*GetTeamBadRequest Incorrect Request, or invalida data

swagger:response getTeamBadRequest
*/
type GetTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTeamBadRequest creates GetTeamBadRequest with default headers values
func NewGetTeamBadRequest() *GetTeamBadRequest {

	return &GetTeamBadRequest{}
}

// WithPayload adds the payload to the get team bad request response
func (o *GetTeamBadRequest) WithPayload(payload *models.Error) *GetTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team bad request response
func (o *GetTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeamForbiddenCode is the HTTP code returned for type GetTeamForbidden
const GetTeamForbiddenCode int = 403

/*GetTeamForbidden Not authorized to this content

swagger:response getTeamForbidden
*/
type GetTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTeamForbidden creates GetTeamForbidden with default headers values
func NewGetTeamForbidden() *GetTeamForbidden {

	return &GetTeamForbidden{}
}

// WithPayload adds the payload to the get team forbidden response
func (o *GetTeamForbidden) WithPayload(payload *models.Error) *GetTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team forbidden response
func (o *GetTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeamGoneCode is the HTTP code returned for type GetTeamGone
const GetTeamGoneCode int = 410

/*GetTeamGone That user (password and name) does not exist

swagger:response getTeamGone
*/
type GetTeamGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTeamGone creates GetTeamGone with default headers values
func NewGetTeamGone() *GetTeamGone {

	return &GetTeamGone{}
}

// WithPayload adds the payload to the get team gone response
func (o *GetTeamGone) WithPayload(payload *models.Error) *GetTeamGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team gone response
func (o *GetTeamGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeamInternalServerErrorCode is the HTTP code returned for type GetTeamInternalServerError
const GetTeamInternalServerErrorCode int = 500

/*GetTeamInternalServerError Internal error

swagger:response getTeamInternalServerError
*/
type GetTeamInternalServerError struct {
}

// NewGetTeamInternalServerError creates GetTeamInternalServerError with default headers values
func NewGetTeamInternalServerError() *GetTeamInternalServerError {

	return &GetTeamInternalServerError{}
}

// WriteResponse to the client
func (o *GetTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
