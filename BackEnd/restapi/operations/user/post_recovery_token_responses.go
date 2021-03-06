// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PostRecoveryTokenCreatedCode is the HTTP code returned for type PostRecoveryTokenCreated
const PostRecoveryTokenCreatedCode int = 201

/*PostRecoveryTokenCreated Password Recovery Token Created

swagger:response postRecoveryTokenCreated
*/
type PostRecoveryTokenCreated struct {
}

// NewPostRecoveryTokenCreated creates PostRecoveryTokenCreated with default headers values
func NewPostRecoveryTokenCreated() *PostRecoveryTokenCreated {

	return &PostRecoveryTokenCreated{}
}

// WriteResponse to the client
func (o *PostRecoveryTokenCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostRecoveryTokenBadRequestCode is the HTTP code returned for type PostRecoveryTokenBadRequest
const PostRecoveryTokenBadRequestCode int = 400

/*PostRecoveryTokenBadRequest Incorrect Request, or invalida data

swagger:response postRecoveryTokenBadRequest
*/
type PostRecoveryTokenBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRecoveryTokenBadRequest creates PostRecoveryTokenBadRequest with default headers values
func NewPostRecoveryTokenBadRequest() *PostRecoveryTokenBadRequest {

	return &PostRecoveryTokenBadRequest{}
}

// WithPayload adds the payload to the post recovery token bad request response
func (o *PostRecoveryTokenBadRequest) WithPayload(payload *models.Error) *PostRecoveryTokenBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post recovery token bad request response
func (o *PostRecoveryTokenBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRecoveryTokenBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRecoveryTokenForbiddenCode is the HTTP code returned for type PostRecoveryTokenForbidden
const PostRecoveryTokenForbiddenCode int = 403

/*PostRecoveryTokenForbidden Not authorized to this content

swagger:response postRecoveryTokenForbidden
*/
type PostRecoveryTokenForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRecoveryTokenForbidden creates PostRecoveryTokenForbidden with default headers values
func NewPostRecoveryTokenForbidden() *PostRecoveryTokenForbidden {

	return &PostRecoveryTokenForbidden{}
}

// WithPayload adds the payload to the post recovery token forbidden response
func (o *PostRecoveryTokenForbidden) WithPayload(payload *models.Error) *PostRecoveryTokenForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post recovery token forbidden response
func (o *PostRecoveryTokenForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRecoveryTokenForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRecoveryTokenGoneCode is the HTTP code returned for type PostRecoveryTokenGone
const PostRecoveryTokenGoneCode int = 410

/*PostRecoveryTokenGone That resource does not exist

swagger:response postRecoveryTokenGone
*/
type PostRecoveryTokenGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRecoveryTokenGone creates PostRecoveryTokenGone with default headers values
func NewPostRecoveryTokenGone() *PostRecoveryTokenGone {

	return &PostRecoveryTokenGone{}
}

// WithPayload adds the payload to the post recovery token gone response
func (o *PostRecoveryTokenGone) WithPayload(payload *models.Error) *PostRecoveryTokenGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post recovery token gone response
func (o *PostRecoveryTokenGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRecoveryTokenGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRecoveryTokenInternalServerErrorCode is the HTTP code returned for type PostRecoveryTokenInternalServerError
const PostRecoveryTokenInternalServerErrorCode int = 500

/*PostRecoveryTokenInternalServerError Internal error

swagger:response postRecoveryTokenInternalServerError
*/
type PostRecoveryTokenInternalServerError struct {
}

// NewPostRecoveryTokenInternalServerError creates PostRecoveryTokenInternalServerError with default headers values
func NewPostRecoveryTokenInternalServerError() *PostRecoveryTokenInternalServerError {

	return &PostRecoveryTokenInternalServerError{}
}

// WriteResponse to the client
func (o *PostRecoveryTokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
