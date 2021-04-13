// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetInvitedTestFromUserOKCode is the HTTP code returned for type GetInvitedTestFromUserOK
const GetInvitedTestFromUserOKCode int = 200

/*GetInvitedTestFromUserOK publishedTest found

swagger:response getInvitedTestFromUserOK
*/
type GetInvitedTestFromUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Test `json:"body,omitempty"`
}

// NewGetInvitedTestFromUserOK creates GetInvitedTestFromUserOK with default headers values
func NewGetInvitedTestFromUserOK() *GetInvitedTestFromUserOK {

	return &GetInvitedTestFromUserOK{}
}

// WithPayload adds the payload to the get invited test from user o k response
func (o *GetInvitedTestFromUserOK) WithPayload(payload *models.Test) *GetInvitedTestFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get invited test from user o k response
func (o *GetInvitedTestFromUserOK) SetPayload(payload *models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInvitedTestFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInvitedTestFromUserBadRequestCode is the HTTP code returned for type GetInvitedTestFromUserBadRequest
const GetInvitedTestFromUserBadRequestCode int = 400

/*GetInvitedTestFromUserBadRequest Incorrect Request, or invalida data

swagger:response getInvitedTestFromUserBadRequest
*/
type GetInvitedTestFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInvitedTestFromUserBadRequest creates GetInvitedTestFromUserBadRequest with default headers values
func NewGetInvitedTestFromUserBadRequest() *GetInvitedTestFromUserBadRequest {

	return &GetInvitedTestFromUserBadRequest{}
}

// WithPayload adds the payload to the get invited test from user bad request response
func (o *GetInvitedTestFromUserBadRequest) WithPayload(payload *models.Error) *GetInvitedTestFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get invited test from user bad request response
func (o *GetInvitedTestFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInvitedTestFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInvitedTestFromUserForbiddenCode is the HTTP code returned for type GetInvitedTestFromUserForbidden
const GetInvitedTestFromUserForbiddenCode int = 403

/*GetInvitedTestFromUserForbidden Not authorized to this content

swagger:response getInvitedTestFromUserForbidden
*/
type GetInvitedTestFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInvitedTestFromUserForbidden creates GetInvitedTestFromUserForbidden with default headers values
func NewGetInvitedTestFromUserForbidden() *GetInvitedTestFromUserForbidden {

	return &GetInvitedTestFromUserForbidden{}
}

// WithPayload adds the payload to the get invited test from user forbidden response
func (o *GetInvitedTestFromUserForbidden) WithPayload(payload *models.Error) *GetInvitedTestFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get invited test from user forbidden response
func (o *GetInvitedTestFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInvitedTestFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInvitedTestFromUserGoneCode is the HTTP code returned for type GetInvitedTestFromUserGone
const GetInvitedTestFromUserGoneCode int = 410

/*GetInvitedTestFromUserGone That user (password and name) does not exist

swagger:response getInvitedTestFromUserGone
*/
type GetInvitedTestFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInvitedTestFromUserGone creates GetInvitedTestFromUserGone with default headers values
func NewGetInvitedTestFromUserGone() *GetInvitedTestFromUserGone {

	return &GetInvitedTestFromUserGone{}
}

// WithPayload adds the payload to the get invited test from user gone response
func (o *GetInvitedTestFromUserGone) WithPayload(payload *models.Error) *GetInvitedTestFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get invited test from user gone response
func (o *GetInvitedTestFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInvitedTestFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInvitedTestFromUserInternalServerErrorCode is the HTTP code returned for type GetInvitedTestFromUserInternalServerError
const GetInvitedTestFromUserInternalServerErrorCode int = 500

/*GetInvitedTestFromUserInternalServerError Internal error

swagger:response getInvitedTestFromUserInternalServerError
*/
type GetInvitedTestFromUserInternalServerError struct {
}

// NewGetInvitedTestFromUserInternalServerError creates GetInvitedTestFromUserInternalServerError with default headers values
func NewGetInvitedTestFromUserInternalServerError() *GetInvitedTestFromUserInternalServerError {

	return &GetInvitedTestFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetInvitedTestFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}