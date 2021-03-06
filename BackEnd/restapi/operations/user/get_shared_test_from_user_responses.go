// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetSharedTestFromUserOKCode is the HTTP code returned for type GetSharedTestFromUserOK
const GetSharedTestFromUserOKCode int = 200

/*GetSharedTestFromUserOK test found

swagger:response getSharedTestFromUserOK
*/
type GetSharedTestFromUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Test `json:"body,omitempty"`
}

// NewGetSharedTestFromUserOK creates GetSharedTestFromUserOK with default headers values
func NewGetSharedTestFromUserOK() *GetSharedTestFromUserOK {

	return &GetSharedTestFromUserOK{}
}

// WithPayload adds the payload to the get shared test from user o k response
func (o *GetSharedTestFromUserOK) WithPayload(payload *models.Test) *GetSharedTestFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared test from user o k response
func (o *GetSharedTestFromUserOK) SetPayload(payload *models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedTestFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedTestFromUserBadRequestCode is the HTTP code returned for type GetSharedTestFromUserBadRequest
const GetSharedTestFromUserBadRequestCode int = 400

/*GetSharedTestFromUserBadRequest Incorrect Request, or invalida data

swagger:response getSharedTestFromUserBadRequest
*/
type GetSharedTestFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSharedTestFromUserBadRequest creates GetSharedTestFromUserBadRequest with default headers values
func NewGetSharedTestFromUserBadRequest() *GetSharedTestFromUserBadRequest {

	return &GetSharedTestFromUserBadRequest{}
}

// WithPayload adds the payload to the get shared test from user bad request response
func (o *GetSharedTestFromUserBadRequest) WithPayload(payload *models.Error) *GetSharedTestFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared test from user bad request response
func (o *GetSharedTestFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedTestFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedTestFromUserForbiddenCode is the HTTP code returned for type GetSharedTestFromUserForbidden
const GetSharedTestFromUserForbiddenCode int = 403

/*GetSharedTestFromUserForbidden Not authorized to this content

swagger:response getSharedTestFromUserForbidden
*/
type GetSharedTestFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSharedTestFromUserForbidden creates GetSharedTestFromUserForbidden with default headers values
func NewGetSharedTestFromUserForbidden() *GetSharedTestFromUserForbidden {

	return &GetSharedTestFromUserForbidden{}
}

// WithPayload adds the payload to the get shared test from user forbidden response
func (o *GetSharedTestFromUserForbidden) WithPayload(payload *models.Error) *GetSharedTestFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared test from user forbidden response
func (o *GetSharedTestFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedTestFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedTestFromUserGoneCode is the HTTP code returned for type GetSharedTestFromUserGone
const GetSharedTestFromUserGoneCode int = 410

/*GetSharedTestFromUserGone That user (password and name) does not exist

swagger:response getSharedTestFromUserGone
*/
type GetSharedTestFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSharedTestFromUserGone creates GetSharedTestFromUserGone with default headers values
func NewGetSharedTestFromUserGone() *GetSharedTestFromUserGone {

	return &GetSharedTestFromUserGone{}
}

// WithPayload adds the payload to the get shared test from user gone response
func (o *GetSharedTestFromUserGone) WithPayload(payload *models.Error) *GetSharedTestFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared test from user gone response
func (o *GetSharedTestFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedTestFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedTestFromUserInternalServerErrorCode is the HTTP code returned for type GetSharedTestFromUserInternalServerError
const GetSharedTestFromUserInternalServerErrorCode int = 500

/*GetSharedTestFromUserInternalServerError Internal error

swagger:response getSharedTestFromUserInternalServerError
*/
type GetSharedTestFromUserInternalServerError struct {
}

// NewGetSharedTestFromUserInternalServerError creates GetSharedTestFromUserInternalServerError with default headers values
func NewGetSharedTestFromUserInternalServerError() *GetSharedTestFromUserInternalServerError {

	return &GetSharedTestFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetSharedTestFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
