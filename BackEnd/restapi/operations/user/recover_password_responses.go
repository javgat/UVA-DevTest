// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// RecoverPasswordOKCode is the HTTP code returned for type RecoverPasswordOK
const RecoverPasswordOKCode int = 200

/*RecoverPasswordOK Resource password modified correctly

swagger:response recoverPasswordOK
*/
type RecoverPasswordOK struct {
}

// NewRecoverPasswordOK creates RecoverPasswordOK with default headers values
func NewRecoverPasswordOK() *RecoverPasswordOK {

	return &RecoverPasswordOK{}
}

// WriteResponse to the client
func (o *RecoverPasswordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RecoverPasswordBadRequestCode is the HTTP code returned for type RecoverPasswordBadRequest
const RecoverPasswordBadRequestCode int = 400

/*RecoverPasswordBadRequest Incorrect Request, or invalida data

swagger:response recoverPasswordBadRequest
*/
type RecoverPasswordBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRecoverPasswordBadRequest creates RecoverPasswordBadRequest with default headers values
func NewRecoverPasswordBadRequest() *RecoverPasswordBadRequest {

	return &RecoverPasswordBadRequest{}
}

// WithPayload adds the payload to the recover password bad request response
func (o *RecoverPasswordBadRequest) WithPayload(payload *models.Error) *RecoverPasswordBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recover password bad request response
func (o *RecoverPasswordBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RecoverPasswordBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RecoverPasswordForbiddenCode is the HTTP code returned for type RecoverPasswordForbidden
const RecoverPasswordForbiddenCode int = 403

/*RecoverPasswordForbidden Not authorized to this content

swagger:response recoverPasswordForbidden
*/
type RecoverPasswordForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRecoverPasswordForbidden creates RecoverPasswordForbidden with default headers values
func NewRecoverPasswordForbidden() *RecoverPasswordForbidden {

	return &RecoverPasswordForbidden{}
}

// WithPayload adds the payload to the recover password forbidden response
func (o *RecoverPasswordForbidden) WithPayload(payload *models.Error) *RecoverPasswordForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recover password forbidden response
func (o *RecoverPasswordForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RecoverPasswordForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RecoverPasswordGoneCode is the HTTP code returned for type RecoverPasswordGone
const RecoverPasswordGoneCode int = 410

/*RecoverPasswordGone That resource does not exist

swagger:response recoverPasswordGone
*/
type RecoverPasswordGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRecoverPasswordGone creates RecoverPasswordGone with default headers values
func NewRecoverPasswordGone() *RecoverPasswordGone {

	return &RecoverPasswordGone{}
}

// WithPayload adds the payload to the recover password gone response
func (o *RecoverPasswordGone) WithPayload(payload *models.Error) *RecoverPasswordGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recover password gone response
func (o *RecoverPasswordGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RecoverPasswordGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RecoverPasswordInternalServerErrorCode is the HTTP code returned for type RecoverPasswordInternalServerError
const RecoverPasswordInternalServerErrorCode int = 500

/*RecoverPasswordInternalServerError Internal error

swagger:response recoverPasswordInternalServerError
*/
type RecoverPasswordInternalServerError struct {
}

// NewRecoverPasswordInternalServerError creates RecoverPasswordInternalServerError with default headers values
func NewRecoverPasswordInternalServerError() *RecoverPasswordInternalServerError {

	return &RecoverPasswordInternalServerError{}
}

// WriteResponse to the client
func (o *RecoverPasswordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
