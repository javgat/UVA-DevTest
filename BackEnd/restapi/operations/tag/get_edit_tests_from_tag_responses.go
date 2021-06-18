// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetEditTestsFromTagOKCode is the HTTP code returned for type GetEditTestsFromTagOK
const GetEditTestsFromTagOKCode int = 200

/*GetEditTestsFromTagOK tests found

swagger:response getEditTestsFromTagOK
*/
type GetEditTestsFromTagOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetEditTestsFromTagOK creates GetEditTestsFromTagOK with default headers values
func NewGetEditTestsFromTagOK() *GetEditTestsFromTagOK {

	return &GetEditTestsFromTagOK{}
}

// WithPayload adds the payload to the get edit tests from tag o k response
func (o *GetEditTestsFromTagOK) WithPayload(payload []*models.Test) *GetEditTestsFromTagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit tests from tag o k response
func (o *GetEditTestsFromTagOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditTestsFromTagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetEditTestsFromTagBadRequestCode is the HTTP code returned for type GetEditTestsFromTagBadRequest
const GetEditTestsFromTagBadRequestCode int = 400

/*GetEditTestsFromTagBadRequest Incorrect Request, or invalida data

swagger:response getEditTestsFromTagBadRequest
*/
type GetEditTestsFromTagBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEditTestsFromTagBadRequest creates GetEditTestsFromTagBadRequest with default headers values
func NewGetEditTestsFromTagBadRequest() *GetEditTestsFromTagBadRequest {

	return &GetEditTestsFromTagBadRequest{}
}

// WithPayload adds the payload to the get edit tests from tag bad request response
func (o *GetEditTestsFromTagBadRequest) WithPayload(payload *models.Error) *GetEditTestsFromTagBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit tests from tag bad request response
func (o *GetEditTestsFromTagBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditTestsFromTagBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEditTestsFromTagForbiddenCode is the HTTP code returned for type GetEditTestsFromTagForbidden
const GetEditTestsFromTagForbiddenCode int = 403

/*GetEditTestsFromTagForbidden Not authorized to this content

swagger:response getEditTestsFromTagForbidden
*/
type GetEditTestsFromTagForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEditTestsFromTagForbidden creates GetEditTestsFromTagForbidden with default headers values
func NewGetEditTestsFromTagForbidden() *GetEditTestsFromTagForbidden {

	return &GetEditTestsFromTagForbidden{}
}

// WithPayload adds the payload to the get edit tests from tag forbidden response
func (o *GetEditTestsFromTagForbidden) WithPayload(payload *models.Error) *GetEditTestsFromTagForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit tests from tag forbidden response
func (o *GetEditTestsFromTagForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditTestsFromTagForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEditTestsFromTagGoneCode is the HTTP code returned for type GetEditTestsFromTagGone
const GetEditTestsFromTagGoneCode int = 410

/*GetEditTestsFromTagGone That user (password and name) does not exist

swagger:response getEditTestsFromTagGone
*/
type GetEditTestsFromTagGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEditTestsFromTagGone creates GetEditTestsFromTagGone with default headers values
func NewGetEditTestsFromTagGone() *GetEditTestsFromTagGone {

	return &GetEditTestsFromTagGone{}
}

// WithPayload adds the payload to the get edit tests from tag gone response
func (o *GetEditTestsFromTagGone) WithPayload(payload *models.Error) *GetEditTestsFromTagGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit tests from tag gone response
func (o *GetEditTestsFromTagGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditTestsFromTagGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEditTestsFromTagInternalServerErrorCode is the HTTP code returned for type GetEditTestsFromTagInternalServerError
const GetEditTestsFromTagInternalServerErrorCode int = 500

/*GetEditTestsFromTagInternalServerError Internal error

swagger:response getEditTestsFromTagInternalServerError
*/
type GetEditTestsFromTagInternalServerError struct {
}

// NewGetEditTestsFromTagInternalServerError creates GetEditTestsFromTagInternalServerError with default headers values
func NewGetEditTestsFromTagInternalServerError() *GetEditTestsFromTagInternalServerError {

	return &GetEditTestsFromTagInternalServerError{}
}

// WriteResponse to the client
func (o *GetEditTestsFromTagInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
