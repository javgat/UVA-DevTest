// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// RemoveAdminTeamToTestOKCode is the HTTP code returned for type RemoveAdminTeamToTestOK
const RemoveAdminTeamToTestOKCode int = 200

/*RemoveAdminTeamToTestOK team removed

swagger:response removeAdminTeamToTestOK
*/
type RemoveAdminTeamToTestOK struct {
}

// NewRemoveAdminTeamToTestOK creates RemoveAdminTeamToTestOK with default headers values
func NewRemoveAdminTeamToTestOK() *RemoveAdminTeamToTestOK {

	return &RemoveAdminTeamToTestOK{}
}

// WriteResponse to the client
func (o *RemoveAdminTeamToTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RemoveAdminTeamToTestBadRequestCode is the HTTP code returned for type RemoveAdminTeamToTestBadRequest
const RemoveAdminTeamToTestBadRequestCode int = 400

/*RemoveAdminTeamToTestBadRequest Incorrect Request, or invalida data

swagger:response removeAdminTeamToTestBadRequest
*/
type RemoveAdminTeamToTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveAdminTeamToTestBadRequest creates RemoveAdminTeamToTestBadRequest with default headers values
func NewRemoveAdminTeamToTestBadRequest() *RemoveAdminTeamToTestBadRequest {

	return &RemoveAdminTeamToTestBadRequest{}
}

// WithPayload adds the payload to the remove admin team to test bad request response
func (o *RemoveAdminTeamToTestBadRequest) WithPayload(payload *models.Error) *RemoveAdminTeamToTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove admin team to test bad request response
func (o *RemoveAdminTeamToTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAdminTeamToTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveAdminTeamToTestForbiddenCode is the HTTP code returned for type RemoveAdminTeamToTestForbidden
const RemoveAdminTeamToTestForbiddenCode int = 403

/*RemoveAdminTeamToTestForbidden Not authorized to this content

swagger:response removeAdminTeamToTestForbidden
*/
type RemoveAdminTeamToTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveAdminTeamToTestForbidden creates RemoveAdminTeamToTestForbidden with default headers values
func NewRemoveAdminTeamToTestForbidden() *RemoveAdminTeamToTestForbidden {

	return &RemoveAdminTeamToTestForbidden{}
}

// WithPayload adds the payload to the remove admin team to test forbidden response
func (o *RemoveAdminTeamToTestForbidden) WithPayload(payload *models.Error) *RemoveAdminTeamToTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove admin team to test forbidden response
func (o *RemoveAdminTeamToTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAdminTeamToTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveAdminTeamToTestGoneCode is the HTTP code returned for type RemoveAdminTeamToTestGone
const RemoveAdminTeamToTestGoneCode int = 410

/*RemoveAdminTeamToTestGone That resource does not exist

swagger:response removeAdminTeamToTestGone
*/
type RemoveAdminTeamToTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveAdminTeamToTestGone creates RemoveAdminTeamToTestGone with default headers values
func NewRemoveAdminTeamToTestGone() *RemoveAdminTeamToTestGone {

	return &RemoveAdminTeamToTestGone{}
}

// WithPayload adds the payload to the remove admin team to test gone response
func (o *RemoveAdminTeamToTestGone) WithPayload(payload *models.Error) *RemoveAdminTeamToTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove admin team to test gone response
func (o *RemoveAdminTeamToTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAdminTeamToTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveAdminTeamToTestInternalServerErrorCode is the HTTP code returned for type RemoveAdminTeamToTestInternalServerError
const RemoveAdminTeamToTestInternalServerErrorCode int = 500

/*RemoveAdminTeamToTestInternalServerError Internal error

swagger:response removeAdminTeamToTestInternalServerError
*/
type RemoveAdminTeamToTestInternalServerError struct {
}

// NewRemoveAdminTeamToTestInternalServerError creates RemoveAdminTeamToTestInternalServerError with default headers values
func NewRemoveAdminTeamToTestInternalServerError() *RemoveAdminTeamToTestInternalServerError {

	return &RemoveAdminTeamToTestInternalServerError{}
}

// WriteResponse to the client
func (o *RemoveAdminTeamToTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
