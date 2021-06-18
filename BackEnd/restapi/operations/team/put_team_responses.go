// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PutTeamOKCode is the HTTP code returned for type PutTeamOK
const PutTeamOKCode int = 200

/*PutTeamOK team updated

swagger:response putTeamOK
*/
type PutTeamOK struct {
}

// NewPutTeamOK creates PutTeamOK with default headers values
func NewPutTeamOK() *PutTeamOK {

	return &PutTeamOK{}
}

// WriteResponse to the client
func (o *PutTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutTeamBadRequestCode is the HTTP code returned for type PutTeamBadRequest
const PutTeamBadRequestCode int = 400

/*PutTeamBadRequest Incorrect Request, or invalida data

swagger:response putTeamBadRequest
*/
type PutTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTeamBadRequest creates PutTeamBadRequest with default headers values
func NewPutTeamBadRequest() *PutTeamBadRequest {

	return &PutTeamBadRequest{}
}

// WithPayload adds the payload to the put team bad request response
func (o *PutTeamBadRequest) WithPayload(payload *models.Error) *PutTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put team bad request response
func (o *PutTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTeamForbiddenCode is the HTTP code returned for type PutTeamForbidden
const PutTeamForbiddenCode int = 403

/*PutTeamForbidden Not authorized to this content

swagger:response putTeamForbidden
*/
type PutTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTeamForbidden creates PutTeamForbidden with default headers values
func NewPutTeamForbidden() *PutTeamForbidden {

	return &PutTeamForbidden{}
}

// WithPayload adds the payload to the put team forbidden response
func (o *PutTeamForbidden) WithPayload(payload *models.Error) *PutTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put team forbidden response
func (o *PutTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTeamConflictCode is the HTTP code returned for type PutTeamConflict
const PutTeamConflictCode int = 409

/*PutTeamConflict A user with same username/email already exists

swagger:response putTeamConflict
*/
type PutTeamConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTeamConflict creates PutTeamConflict with default headers values
func NewPutTeamConflict() *PutTeamConflict {

	return &PutTeamConflict{}
}

// WithPayload adds the payload to the put team conflict response
func (o *PutTeamConflict) WithPayload(payload *models.Error) *PutTeamConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put team conflict response
func (o *PutTeamConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTeamConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTeamGoneCode is the HTTP code returned for type PutTeamGone
const PutTeamGoneCode int = 410

/*PutTeamGone That user (password and name) does not exist

swagger:response putTeamGone
*/
type PutTeamGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTeamGone creates PutTeamGone with default headers values
func NewPutTeamGone() *PutTeamGone {

	return &PutTeamGone{}
}

// WithPayload adds the payload to the put team gone response
func (o *PutTeamGone) WithPayload(payload *models.Error) *PutTeamGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put team gone response
func (o *PutTeamGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTeamGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTeamInternalServerErrorCode is the HTTP code returned for type PutTeamInternalServerError
const PutTeamInternalServerErrorCode int = 500

/*PutTeamInternalServerError Internal error

swagger:response putTeamInternalServerError
*/
type PutTeamInternalServerError struct {
}

// NewPutTeamInternalServerError creates PutTeamInternalServerError with default headers values
func NewPutTeamInternalServerError() *PutTeamInternalServerError {

	return &PutTeamInternalServerError{}
}

// WriteResponse to the client
func (o *PutTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
