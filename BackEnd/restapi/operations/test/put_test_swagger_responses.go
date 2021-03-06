// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PutTestOKCode is the HTTP code returned for type PutTestOK
const PutTestOKCode int = 200

/*PutTestOK Test updated

swagger:response putTestOK
*/
type PutTestOK struct {
}

// NewPutTestOK creates PutTestOK with default headers values
func NewPutTestOK() *PutTestOK {

	return &PutTestOK{}
}

// WriteResponse to the client
func (o *PutTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutTestBadRequestCode is the HTTP code returned for type PutTestBadRequest
const PutTestBadRequestCode int = 400

/*PutTestBadRequest Incorrect Request, or invalida data

swagger:response putTestBadRequest
*/
type PutTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTestBadRequest creates PutTestBadRequest with default headers values
func NewPutTestBadRequest() *PutTestBadRequest {

	return &PutTestBadRequest{}
}

// WithPayload adds the payload to the put test bad request response
func (o *PutTestBadRequest) WithPayload(payload *models.Error) *PutTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put test bad request response
func (o *PutTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTestForbiddenCode is the HTTP code returned for type PutTestForbidden
const PutTestForbiddenCode int = 403

/*PutTestForbidden Not authorized to this content

swagger:response putTestForbidden
*/
type PutTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTestForbidden creates PutTestForbidden with default headers values
func NewPutTestForbidden() *PutTestForbidden {

	return &PutTestForbidden{}
}

// WithPayload adds the payload to the put test forbidden response
func (o *PutTestForbidden) WithPayload(payload *models.Error) *PutTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put test forbidden response
func (o *PutTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTestInternalServerErrorCode is the HTTP code returned for type PutTestInternalServerError
const PutTestInternalServerErrorCode int = 500

/*PutTestInternalServerError Internal error

swagger:response putTestInternalServerError
*/
type PutTestInternalServerError struct {
}

// NewPutTestInternalServerError creates PutTestInternalServerError with default headers values
func NewPutTestInternalServerError() *PutTestInternalServerError {

	return &PutTestInternalServerError{}
}

// WriteResponse to the client
func (o *PutTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
