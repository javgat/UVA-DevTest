// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetMemberOKCode is the HTTP code returned for type GetMemberOK
const GetMemberOKCode int = 200

/*GetMemberOK user found

swagger:response getMemberOK
*/
type GetMemberOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetMemberOK creates GetMemberOK with default headers values
func NewGetMemberOK() *GetMemberOK {

	return &GetMemberOK{}
}

// WithPayload adds the payload to the get member o k response
func (o *GetMemberOK) WithPayload(payload *models.User) *GetMemberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get member o k response
func (o *GetMemberOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMemberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMemberBadRequestCode is the HTTP code returned for type GetMemberBadRequest
const GetMemberBadRequestCode int = 400

/*GetMemberBadRequest Incorrect Request, or invalida data

swagger:response getMemberBadRequest
*/
type GetMemberBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMemberBadRequest creates GetMemberBadRequest with default headers values
func NewGetMemberBadRequest() *GetMemberBadRequest {

	return &GetMemberBadRequest{}
}

// WithPayload adds the payload to the get member bad request response
func (o *GetMemberBadRequest) WithPayload(payload *models.Error) *GetMemberBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get member bad request response
func (o *GetMemberBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMemberBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMemberForbiddenCode is the HTTP code returned for type GetMemberForbidden
const GetMemberForbiddenCode int = 403

/*GetMemberForbidden Not authorized to this content

swagger:response getMemberForbidden
*/
type GetMemberForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMemberForbidden creates GetMemberForbidden with default headers values
func NewGetMemberForbidden() *GetMemberForbidden {

	return &GetMemberForbidden{}
}

// WithPayload adds the payload to the get member forbidden response
func (o *GetMemberForbidden) WithPayload(payload *models.Error) *GetMemberForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get member forbidden response
func (o *GetMemberForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMemberForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMemberGoneCode is the HTTP code returned for type GetMemberGone
const GetMemberGoneCode int = 410

/*GetMemberGone That user (password and name) does not exist

swagger:response getMemberGone
*/
type GetMemberGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMemberGone creates GetMemberGone with default headers values
func NewGetMemberGone() *GetMemberGone {

	return &GetMemberGone{}
}

// WithPayload adds the payload to the get member gone response
func (o *GetMemberGone) WithPayload(payload *models.Error) *GetMemberGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get member gone response
func (o *GetMemberGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMemberGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMemberInternalServerErrorCode is the HTTP code returned for type GetMemberInternalServerError
const GetMemberInternalServerErrorCode int = 500

/*GetMemberInternalServerError Internal error

swagger:response getMemberInternalServerError
*/
type GetMemberInternalServerError struct {
}

// NewGetMemberInternalServerError creates GetMemberInternalServerError with default headers values
func NewGetMemberInternalServerError() *GetMemberInternalServerError {

	return &GetMemberInternalServerError{}
}

// WriteResponse to the client
func (o *GetMemberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
