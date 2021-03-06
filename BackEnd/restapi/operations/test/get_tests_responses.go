// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTestsOKCode is the HTTP code returned for type GetTestsOK
const GetTestsOKCode int = 200

/*GetTestsOK tests found

swagger:response getTestsOK
*/
type GetTestsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetTestsOK creates GetTestsOK with default headers values
func NewGetTestsOK() *GetTestsOK {

	return &GetTestsOK{}
}

// WithPayload adds the payload to the get tests o k response
func (o *GetTestsOK) WithPayload(payload []*models.Test) *GetTestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests o k response
func (o *GetTestsOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Test, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTestsBadRequestCode is the HTTP code returned for type GetTestsBadRequest
const GetTestsBadRequestCode int = 400

/*GetTestsBadRequest Incorrect Request, or invalida data

swagger:response getTestsBadRequest
*/
type GetTestsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsBadRequest creates GetTestsBadRequest with default headers values
func NewGetTestsBadRequest() *GetTestsBadRequest {

	return &GetTestsBadRequest{}
}

// WithPayload adds the payload to the get tests bad request response
func (o *GetTestsBadRequest) WithPayload(payload *models.Error) *GetTestsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests bad request response
func (o *GetTestsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsForbiddenCode is the HTTP code returned for type GetTestsForbidden
const GetTestsForbiddenCode int = 403

/*GetTestsForbidden Not authorized to this content

swagger:response getTestsForbidden
*/
type GetTestsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsForbidden creates GetTestsForbidden with default headers values
func NewGetTestsForbidden() *GetTestsForbidden {

	return &GetTestsForbidden{}
}

// WithPayload adds the payload to the get tests forbidden response
func (o *GetTestsForbidden) WithPayload(payload *models.Error) *GetTestsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests forbidden response
func (o *GetTestsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsInternalServerErrorCode is the HTTP code returned for type GetTestsInternalServerError
const GetTestsInternalServerErrorCode int = 500

/*GetTestsInternalServerError Internal error

swagger:response getTestsInternalServerError
*/
type GetTestsInternalServerError struct {
}

// NewGetTestsInternalServerError creates GetTestsInternalServerError with default headers values
func NewGetTestsInternalServerError() *GetTestsInternalServerError {

	return &GetTestsInternalServerError{}
}

// WriteResponse to the client
func (o *GetTestsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
