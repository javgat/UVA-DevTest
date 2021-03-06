// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublishedTestOKCode is the HTTP code returned for type GetPublishedTestOK
const GetPublishedTestOKCode int = 200

/*GetPublishedTestOK PublishedTest found

swagger:response getPublishedTestOK
*/
type GetPublishedTestOK struct {

	/*
	  In: Body
	*/
	Payload *models.Test `json:"body,omitempty"`
}

// NewGetPublishedTestOK creates GetPublishedTestOK with default headers values
func NewGetPublishedTestOK() *GetPublishedTestOK {

	return &GetPublishedTestOK{}
}

// WithPayload adds the payload to the get published test o k response
func (o *GetPublishedTestOK) WithPayload(payload *models.Test) *GetPublishedTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published test o k response
func (o *GetPublishedTestOK) SetPayload(payload *models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestBadRequestCode is the HTTP code returned for type GetPublishedTestBadRequest
const GetPublishedTestBadRequestCode int = 400

/*GetPublishedTestBadRequest Incorrect Request, or invalida data

swagger:response getPublishedTestBadRequest
*/
type GetPublishedTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestBadRequest creates GetPublishedTestBadRequest with default headers values
func NewGetPublishedTestBadRequest() *GetPublishedTestBadRequest {

	return &GetPublishedTestBadRequest{}
}

// WithPayload adds the payload to the get published test bad request response
func (o *GetPublishedTestBadRequest) WithPayload(payload *models.Error) *GetPublishedTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published test bad request response
func (o *GetPublishedTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestForbiddenCode is the HTTP code returned for type GetPublishedTestForbidden
const GetPublishedTestForbiddenCode int = 403

/*GetPublishedTestForbidden Not authorized to this content

swagger:response getPublishedTestForbidden
*/
type GetPublishedTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestForbidden creates GetPublishedTestForbidden with default headers values
func NewGetPublishedTestForbidden() *GetPublishedTestForbidden {

	return &GetPublishedTestForbidden{}
}

// WithPayload adds the payload to the get published test forbidden response
func (o *GetPublishedTestForbidden) WithPayload(payload *models.Error) *GetPublishedTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published test forbidden response
func (o *GetPublishedTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestGoneCode is the HTTP code returned for type GetPublishedTestGone
const GetPublishedTestGoneCode int = 410

/*GetPublishedTestGone That user (password and name) does not exist

swagger:response getPublishedTestGone
*/
type GetPublishedTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestGone creates GetPublishedTestGone with default headers values
func NewGetPublishedTestGone() *GetPublishedTestGone {

	return &GetPublishedTestGone{}
}

// WithPayload adds the payload to the get published test gone response
func (o *GetPublishedTestGone) WithPayload(payload *models.Error) *GetPublishedTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published test gone response
func (o *GetPublishedTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestInternalServerErrorCode is the HTTP code returned for type GetPublishedTestInternalServerError
const GetPublishedTestInternalServerErrorCode int = 500

/*GetPublishedTestInternalServerError Internal error

swagger:response getPublishedTestInternalServerError
*/
type GetPublishedTestInternalServerError struct {
}

// NewGetPublishedTestInternalServerError creates GetPublishedTestInternalServerError with default headers values
func NewGetPublishedTestInternalServerError() *GetPublishedTestInternalServerError {

	return &GetPublishedTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublishedTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
