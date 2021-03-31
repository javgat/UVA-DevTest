// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTestsFromUserOKCode is the HTTP code returned for type GetTestsFromUserOK
const GetTestsFromUserOKCode int = 200

/*GetTestsFromUserOK tests found

swagger:response getTestsFromUserOK
*/
type GetTestsFromUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetTestsFromUserOK creates GetTestsFromUserOK with default headers values
func NewGetTestsFromUserOK() *GetTestsFromUserOK {

	return &GetTestsFromUserOK{}
}

// WithPayload adds the payload to the get tests from user o k response
func (o *GetTestsFromUserOK) WithPayload(payload []*models.Test) *GetTestsFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from user o k response
func (o *GetTestsFromUserOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTestsFromUserBadRequestCode is the HTTP code returned for type GetTestsFromUserBadRequest
const GetTestsFromUserBadRequestCode int = 400

/*GetTestsFromUserBadRequest Incorrect Request, or invalida data

swagger:response getTestsFromUserBadRequest
*/
type GetTestsFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsFromUserBadRequest creates GetTestsFromUserBadRequest with default headers values
func NewGetTestsFromUserBadRequest() *GetTestsFromUserBadRequest {

	return &GetTestsFromUserBadRequest{}
}

// WithPayload adds the payload to the get tests from user bad request response
func (o *GetTestsFromUserBadRequest) WithPayload(payload *models.Error) *GetTestsFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from user bad request response
func (o *GetTestsFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsFromUserForbiddenCode is the HTTP code returned for type GetTestsFromUserForbidden
const GetTestsFromUserForbiddenCode int = 403

/*GetTestsFromUserForbidden Not authorized to this content

swagger:response getTestsFromUserForbidden
*/
type GetTestsFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsFromUserForbidden creates GetTestsFromUserForbidden with default headers values
func NewGetTestsFromUserForbidden() *GetTestsFromUserForbidden {

	return &GetTestsFromUserForbidden{}
}

// WithPayload adds the payload to the get tests from user forbidden response
func (o *GetTestsFromUserForbidden) WithPayload(payload *models.Error) *GetTestsFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from user forbidden response
func (o *GetTestsFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsFromUserGoneCode is the HTTP code returned for type GetTestsFromUserGone
const GetTestsFromUserGoneCode int = 410

/*GetTestsFromUserGone That user (password and name) does not exist

swagger:response getTestsFromUserGone
*/
type GetTestsFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsFromUserGone creates GetTestsFromUserGone with default headers values
func NewGetTestsFromUserGone() *GetTestsFromUserGone {

	return &GetTestsFromUserGone{}
}

// WithPayload adds the payload to the get tests from user gone response
func (o *GetTestsFromUserGone) WithPayload(payload *models.Error) *GetTestsFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from user gone response
func (o *GetTestsFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsFromUserInternalServerErrorCode is the HTTP code returned for type GetTestsFromUserInternalServerError
const GetTestsFromUserInternalServerErrorCode int = 500

/*GetTestsFromUserInternalServerError Internal error

swagger:response getTestsFromUserInternalServerError
*/
type GetTestsFromUserInternalServerError struct {
}

// NewGetTestsFromUserInternalServerError creates GetTestsFromUserInternalServerError with default headers values
func NewGetTestsFromUserInternalServerError() *GetTestsFromUserInternalServerError {

	return &GetTestsFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetTestsFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}