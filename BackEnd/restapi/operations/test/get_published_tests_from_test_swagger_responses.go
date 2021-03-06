// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublishedTestsFromTestOKCode is the HTTP code returned for type GetPublishedTestsFromTestOK
const GetPublishedTestsFromTestOKCode int = 200

/*GetPublishedTestsFromTestOK Tests Found

swagger:response getPublishedTestsFromTestOK
*/
type GetPublishedTestsFromTestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTestOK creates GetPublishedTestsFromTestOK with default headers values
func NewGetPublishedTestsFromTestOK() *GetPublishedTestsFromTestOK {

	return &GetPublishedTestsFromTestOK{}
}

// WithPayload adds the payload to the get published tests from test o k response
func (o *GetPublishedTestsFromTestOK) WithPayload(payload []*models.Test) *GetPublishedTestsFromTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from test o k response
func (o *GetPublishedTestsFromTestOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPublishedTestsFromTestBadRequestCode is the HTTP code returned for type GetPublishedTestsFromTestBadRequest
const GetPublishedTestsFromTestBadRequestCode int = 400

/*GetPublishedTestsFromTestBadRequest Incorrect Request, or invalida data

swagger:response getPublishedTestsFromTestBadRequest
*/
type GetPublishedTestsFromTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTestBadRequest creates GetPublishedTestsFromTestBadRequest with default headers values
func NewGetPublishedTestsFromTestBadRequest() *GetPublishedTestsFromTestBadRequest {

	return &GetPublishedTestsFromTestBadRequest{}
}

// WithPayload adds the payload to the get published tests from test bad request response
func (o *GetPublishedTestsFromTestBadRequest) WithPayload(payload *models.Error) *GetPublishedTestsFromTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from test bad request response
func (o *GetPublishedTestsFromTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsFromTestForbiddenCode is the HTTP code returned for type GetPublishedTestsFromTestForbidden
const GetPublishedTestsFromTestForbiddenCode int = 403

/*GetPublishedTestsFromTestForbidden Not authorized to this content

swagger:response getPublishedTestsFromTestForbidden
*/
type GetPublishedTestsFromTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTestForbidden creates GetPublishedTestsFromTestForbidden with default headers values
func NewGetPublishedTestsFromTestForbidden() *GetPublishedTestsFromTestForbidden {

	return &GetPublishedTestsFromTestForbidden{}
}

// WithPayload adds the payload to the get published tests from test forbidden response
func (o *GetPublishedTestsFromTestForbidden) WithPayload(payload *models.Error) *GetPublishedTestsFromTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from test forbidden response
func (o *GetPublishedTestsFromTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsFromTestGoneCode is the HTTP code returned for type GetPublishedTestsFromTestGone
const GetPublishedTestsFromTestGoneCode int = 410

/*GetPublishedTestsFromTestGone That resource does not exist

swagger:response getPublishedTestsFromTestGone
*/
type GetPublishedTestsFromTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTestGone creates GetPublishedTestsFromTestGone with default headers values
func NewGetPublishedTestsFromTestGone() *GetPublishedTestsFromTestGone {

	return &GetPublishedTestsFromTestGone{}
}

// WithPayload adds the payload to the get published tests from test gone response
func (o *GetPublishedTestsFromTestGone) WithPayload(payload *models.Error) *GetPublishedTestsFromTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from test gone response
func (o *GetPublishedTestsFromTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsFromTestInternalServerErrorCode is the HTTP code returned for type GetPublishedTestsFromTestInternalServerError
const GetPublishedTestsFromTestInternalServerErrorCode int = 500

/*GetPublishedTestsFromTestInternalServerError Internal error

swagger:response getPublishedTestsFromTestInternalServerError
*/
type GetPublishedTestsFromTestInternalServerError struct {
}

// NewGetPublishedTestsFromTestInternalServerError creates GetPublishedTestsFromTestInternalServerError with default headers values
func NewGetPublishedTestsFromTestInternalServerError() *GetPublishedTestsFromTestInternalServerError {

	return &GetPublishedTestsFromTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
