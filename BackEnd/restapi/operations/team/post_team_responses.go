// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PostTeamCreatedCode is the HTTP code returned for type PostTeamCreated
const PostTeamCreatedCode int = 201

/*PostTeamCreated team created

swagger:response postTeamCreated
*/
type PostTeamCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Team `json:"body,omitempty"`
}

// NewPostTeamCreated creates PostTeamCreated with default headers values
func NewPostTeamCreated() *PostTeamCreated {

	return &PostTeamCreated{}
}

// WithPayload adds the payload to the post team created response
func (o *PostTeamCreated) WithPayload(payload *models.Team) *PostTeamCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post team created response
func (o *PostTeamCreated) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTeamCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTeamBadRequestCode is the HTTP code returned for type PostTeamBadRequest
const PostTeamBadRequestCode int = 400

/*PostTeamBadRequest Incorrect Request, or invalida data

swagger:response postTeamBadRequest
*/
type PostTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTeamBadRequest creates PostTeamBadRequest with default headers values
func NewPostTeamBadRequest() *PostTeamBadRequest {

	return &PostTeamBadRequest{}
}

// WithPayload adds the payload to the post team bad request response
func (o *PostTeamBadRequest) WithPayload(payload *models.Error) *PostTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post team bad request response
func (o *PostTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTeamForbiddenCode is the HTTP code returned for type PostTeamForbidden
const PostTeamForbiddenCode int = 403

/*PostTeamForbidden Not authorized to this content

swagger:response postTeamForbidden
*/
type PostTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTeamForbidden creates PostTeamForbidden with default headers values
func NewPostTeamForbidden() *PostTeamForbidden {

	return &PostTeamForbidden{}
}

// WithPayload adds the payload to the post team forbidden response
func (o *PostTeamForbidden) WithPayload(payload *models.Error) *PostTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post team forbidden response
func (o *PostTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTeamConflictCode is the HTTP code returned for type PostTeamConflict
const PostTeamConflictCode int = 409

/*PostTeamConflict A user with same username/email already exists

swagger:response postTeamConflict
*/
type PostTeamConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTeamConflict creates PostTeamConflict with default headers values
func NewPostTeamConflict() *PostTeamConflict {

	return &PostTeamConflict{}
}

// WithPayload adds the payload to the post team conflict response
func (o *PostTeamConflict) WithPayload(payload *models.Error) *PostTeamConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post team conflict response
func (o *PostTeamConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTeamConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTeamInternalServerErrorCode is the HTTP code returned for type PostTeamInternalServerError
const PostTeamInternalServerErrorCode int = 500

/*PostTeamInternalServerError Internal error

swagger:response postTeamInternalServerError
*/
type PostTeamInternalServerError struct {
}

// NewPostTeamInternalServerError creates PostTeamInternalServerError with default headers values
func NewPostTeamInternalServerError() *PostTeamInternalServerError {

	return &PostTeamInternalServerError{}
}

// WriteResponse to the client
func (o *PostTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
