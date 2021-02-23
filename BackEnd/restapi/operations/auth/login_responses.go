// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// LoginOKCode is the HTTP code returned for type LoginOK
const LoginOKCode int = 200

/*LoginOK Successful authentication

swagger:response loginOK
*/
type LoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.JWTJSON `json:"body,omitempty"`
}

// NewLoginOK creates LoginOK with default headers values
func NewLoginOK() *LoginOK {

	return &LoginOK{}
}

// WithPayload adds the payload to the login o k response
func (o *LoginOK) WithPayload(payload *models.JWTJSON) *LoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login o k response
func (o *LoginOK) SetPayload(payload *models.JWTJSON) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginBadRequestCode is the HTTP code returned for type LoginBadRequest
const LoginBadRequestCode int = 400

/*LoginBadRequest Incorrect Request, or invalida data

swagger:response loginBadRequest
*/
type LoginBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoginBadRequest creates LoginBadRequest with default headers values
func NewLoginBadRequest() *LoginBadRequest {

	return &LoginBadRequest{}
}

// WithPayload adds the payload to the login bad request response
func (o *LoginBadRequest) WithPayload(payload *models.Error) *LoginBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login bad request response
func (o *LoginBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginGoneCode is the HTTP code returned for type LoginGone
const LoginGoneCode int = 410

/*LoginGone That user (password and name) does not exist

swagger:response loginGone
*/
type LoginGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoginGone creates LoginGone with default headers values
func NewLoginGone() *LoginGone {

	return &LoginGone{}
}

// WithPayload adds the payload to the login gone response
func (o *LoginGone) WithPayload(payload *models.Error) *LoginGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login gone response
func (o *LoginGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginInternalServerErrorCode is the HTTP code returned for type LoginInternalServerError
const LoginInternalServerErrorCode int = 500

/*LoginInternalServerError Internal error

swagger:response loginInternalServerError
*/
type LoginInternalServerError struct {
}

// NewLoginInternalServerError creates LoginInternalServerError with default headers values
func NewLoginInternalServerError() *LoginInternalServerError {

	return &LoginInternalServerError{}
}

// WriteResponse to the client
func (o *LoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}