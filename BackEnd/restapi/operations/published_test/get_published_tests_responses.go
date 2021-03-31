// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublishedTestsOKCode is the HTTP code returned for type GetPublishedTestsOK
const GetPublishedTestsOKCode int = 200

/*GetPublishedTestsOK PublishedTests found

swagger:response getPublishedTestsOK
*/
type GetPublishedTestsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetPublishedTestsOK creates GetPublishedTestsOK with default headers values
func NewGetPublishedTestsOK() *GetPublishedTestsOK {

	return &GetPublishedTestsOK{}
}

// WithPayload adds the payload to the get published tests o k response
func (o *GetPublishedTestsOK) WithPayload(payload []*models.Test) *GetPublishedTestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests o k response
func (o *GetPublishedTestsOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPublishedTestsBadRequestCode is the HTTP code returned for type GetPublishedTestsBadRequest
const GetPublishedTestsBadRequestCode int = 400

/*GetPublishedTestsBadRequest Incorrect Request, or invalida data

swagger:response getPublishedTestsBadRequest
*/
type GetPublishedTestsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsBadRequest creates GetPublishedTestsBadRequest with default headers values
func NewGetPublishedTestsBadRequest() *GetPublishedTestsBadRequest {

	return &GetPublishedTestsBadRequest{}
}

// WithPayload adds the payload to the get published tests bad request response
func (o *GetPublishedTestsBadRequest) WithPayload(payload *models.Error) *GetPublishedTestsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests bad request response
func (o *GetPublishedTestsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsForbiddenCode is the HTTP code returned for type GetPublishedTestsForbidden
const GetPublishedTestsForbiddenCode int = 403

/*GetPublishedTestsForbidden Not authorized to this content

swagger:response getPublishedTestsForbidden
*/
type GetPublishedTestsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsForbidden creates GetPublishedTestsForbidden with default headers values
func NewGetPublishedTestsForbidden() *GetPublishedTestsForbidden {

	return &GetPublishedTestsForbidden{}
}

// WithPayload adds the payload to the get published tests forbidden response
func (o *GetPublishedTestsForbidden) WithPayload(payload *models.Error) *GetPublishedTestsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests forbidden response
func (o *GetPublishedTestsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsInternalServerErrorCode is the HTTP code returned for type GetPublishedTestsInternalServerError
const GetPublishedTestsInternalServerErrorCode int = 500

/*GetPublishedTestsInternalServerError Internal error

swagger:response getPublishedTestsInternalServerError
*/
type GetPublishedTestsInternalServerError struct {
}

// NewGetPublishedTestsInternalServerError creates GetPublishedTestsInternalServerError with default headers values
func NewGetPublishedTestsInternalServerError() *GetPublishedTestsInternalServerError {

	return &GetPublishedTestsInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublishedTestsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}