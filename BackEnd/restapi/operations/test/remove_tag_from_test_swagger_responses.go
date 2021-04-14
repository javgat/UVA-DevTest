// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// RemoveTagFromTestOKCode is the HTTP code returned for type RemoveTagFromTestOK
const RemoveTagFromTestOKCode int = 200

/*RemoveTagFromTestOK tag removed

swagger:response removeTagFromTestOK
*/
type RemoveTagFromTestOK struct {
}

// NewRemoveTagFromTestOK creates RemoveTagFromTestOK with default headers values
func NewRemoveTagFromTestOK() *RemoveTagFromTestOK {

	return &RemoveTagFromTestOK{}
}

// WriteResponse to the client
func (o *RemoveTagFromTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RemoveTagFromTestBadRequestCode is the HTTP code returned for type RemoveTagFromTestBadRequest
const RemoveTagFromTestBadRequestCode int = 400

/*RemoveTagFromTestBadRequest Incorrect Request, or invalida data

swagger:response removeTagFromTestBadRequest
*/
type RemoveTagFromTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTagFromTestBadRequest creates RemoveTagFromTestBadRequest with default headers values
func NewRemoveTagFromTestBadRequest() *RemoveTagFromTestBadRequest {

	return &RemoveTagFromTestBadRequest{}
}

// WithPayload adds the payload to the remove tag from test bad request response
func (o *RemoveTagFromTestBadRequest) WithPayload(payload *models.Error) *RemoveTagFromTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove tag from test bad request response
func (o *RemoveTagFromTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTagFromTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTagFromTestForbiddenCode is the HTTP code returned for type RemoveTagFromTestForbidden
const RemoveTagFromTestForbiddenCode int = 403

/*RemoveTagFromTestForbidden Not authorized to this content

swagger:response removeTagFromTestForbidden
*/
type RemoveTagFromTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTagFromTestForbidden creates RemoveTagFromTestForbidden with default headers values
func NewRemoveTagFromTestForbidden() *RemoveTagFromTestForbidden {

	return &RemoveTagFromTestForbidden{}
}

// WithPayload adds the payload to the remove tag from test forbidden response
func (o *RemoveTagFromTestForbidden) WithPayload(payload *models.Error) *RemoveTagFromTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove tag from test forbidden response
func (o *RemoveTagFromTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTagFromTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTagFromTestGoneCode is the HTTP code returned for type RemoveTagFromTestGone
const RemoveTagFromTestGoneCode int = 410

/*RemoveTagFromTestGone That user (password and name) does not exist

swagger:response removeTagFromTestGone
*/
type RemoveTagFromTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTagFromTestGone creates RemoveTagFromTestGone with default headers values
func NewRemoveTagFromTestGone() *RemoveTagFromTestGone {

	return &RemoveTagFromTestGone{}
}

// WithPayload adds the payload to the remove tag from test gone response
func (o *RemoveTagFromTestGone) WithPayload(payload *models.Error) *RemoveTagFromTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove tag from test gone response
func (o *RemoveTagFromTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTagFromTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTagFromTestInternalServerErrorCode is the HTTP code returned for type RemoveTagFromTestInternalServerError
const RemoveTagFromTestInternalServerErrorCode int = 500

/*RemoveTagFromTestInternalServerError Internal error

swagger:response removeTagFromTestInternalServerError
*/
type RemoveTagFromTestInternalServerError struct {
}

// NewRemoveTagFromTestInternalServerError creates RemoveTagFromTestInternalServerError with default headers values
func NewRemoveTagFromTestInternalServerError() *RemoveTagFromTestInternalServerError {

	return &RemoveTagFromTestInternalServerError{}
}

// WriteResponse to the client
func (o *RemoveTagFromTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
