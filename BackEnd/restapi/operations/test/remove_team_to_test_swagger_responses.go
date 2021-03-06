// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// RemoveTeamToTestOKCode is the HTTP code returned for type RemoveTeamToTestOK
const RemoveTeamToTestOKCode int = 200

/*RemoveTeamToTestOK team removed

swagger:response removeTeamToTestOK
*/
type RemoveTeamToTestOK struct {
}

// NewRemoveTeamToTestOK creates RemoveTeamToTestOK with default headers values
func NewRemoveTeamToTestOK() *RemoveTeamToTestOK {

	return &RemoveTeamToTestOK{}
}

// WriteResponse to the client
func (o *RemoveTeamToTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RemoveTeamToTestBadRequestCode is the HTTP code returned for type RemoveTeamToTestBadRequest
const RemoveTeamToTestBadRequestCode int = 400

/*RemoveTeamToTestBadRequest Incorrect Request, or invalida data

swagger:response removeTeamToTestBadRequest
*/
type RemoveTeamToTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTeamToTestBadRequest creates RemoveTeamToTestBadRequest with default headers values
func NewRemoveTeamToTestBadRequest() *RemoveTeamToTestBadRequest {

	return &RemoveTeamToTestBadRequest{}
}

// WithPayload adds the payload to the remove team to test bad request response
func (o *RemoveTeamToTestBadRequest) WithPayload(payload *models.Error) *RemoveTeamToTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove team to test bad request response
func (o *RemoveTeamToTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTeamToTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTeamToTestForbiddenCode is the HTTP code returned for type RemoveTeamToTestForbidden
const RemoveTeamToTestForbiddenCode int = 403

/*RemoveTeamToTestForbidden Not authorized to this content

swagger:response removeTeamToTestForbidden
*/
type RemoveTeamToTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTeamToTestForbidden creates RemoveTeamToTestForbidden with default headers values
func NewRemoveTeamToTestForbidden() *RemoveTeamToTestForbidden {

	return &RemoveTeamToTestForbidden{}
}

// WithPayload adds the payload to the remove team to test forbidden response
func (o *RemoveTeamToTestForbidden) WithPayload(payload *models.Error) *RemoveTeamToTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove team to test forbidden response
func (o *RemoveTeamToTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTeamToTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTeamToTestGoneCode is the HTTP code returned for type RemoveTeamToTestGone
const RemoveTeamToTestGoneCode int = 410

/*RemoveTeamToTestGone That user (password and name) does not exist

swagger:response removeTeamToTestGone
*/
type RemoveTeamToTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTeamToTestGone creates RemoveTeamToTestGone with default headers values
func NewRemoveTeamToTestGone() *RemoveTeamToTestGone {

	return &RemoveTeamToTestGone{}
}

// WithPayload adds the payload to the remove team to test gone response
func (o *RemoveTeamToTestGone) WithPayload(payload *models.Error) *RemoveTeamToTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove team to test gone response
func (o *RemoveTeamToTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTeamToTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTeamToTestInternalServerErrorCode is the HTTP code returned for type RemoveTeamToTestInternalServerError
const RemoveTeamToTestInternalServerErrorCode int = 500

/*RemoveTeamToTestInternalServerError Internal error

swagger:response removeTeamToTestInternalServerError
*/
type RemoveTeamToTestInternalServerError struct {
}

// NewRemoveTeamToTestInternalServerError creates RemoveTeamToTestInternalServerError with default headers values
func NewRemoveTeamToTestInternalServerError() *RemoveTeamToTestInternalServerError {

	return &RemoveTeamToTestInternalServerError{}
}

// WriteResponse to the client
func (o *RemoveTeamToTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
