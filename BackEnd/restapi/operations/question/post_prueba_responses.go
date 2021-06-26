// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PostPruebaCreatedCode is the HTTP code returned for type PostPruebaCreated
const PostPruebaCreatedCode int = 201

/*PostPruebaCreated prueba created

swagger:response postPruebaCreated
*/
type PostPruebaCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Prueba `json:"body,omitempty"`
}

// NewPostPruebaCreated creates PostPruebaCreated with default headers values
func NewPostPruebaCreated() *PostPruebaCreated {

	return &PostPruebaCreated{}
}

// WithPayload adds the payload to the post prueba created response
func (o *PostPruebaCreated) WithPayload(payload *models.Prueba) *PostPruebaCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post prueba created response
func (o *PostPruebaCreated) SetPayload(payload *models.Prueba) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPruebaCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPruebaBadRequestCode is the HTTP code returned for type PostPruebaBadRequest
const PostPruebaBadRequestCode int = 400

/*PostPruebaBadRequest Incorrect Request, or invalida data

swagger:response postPruebaBadRequest
*/
type PostPruebaBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostPruebaBadRequest creates PostPruebaBadRequest with default headers values
func NewPostPruebaBadRequest() *PostPruebaBadRequest {

	return &PostPruebaBadRequest{}
}

// WithPayload adds the payload to the post prueba bad request response
func (o *PostPruebaBadRequest) WithPayload(payload *models.Error) *PostPruebaBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post prueba bad request response
func (o *PostPruebaBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPruebaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPruebaForbiddenCode is the HTTP code returned for type PostPruebaForbidden
const PostPruebaForbiddenCode int = 403

/*PostPruebaForbidden Not authorized to this content

swagger:response postPruebaForbidden
*/
type PostPruebaForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostPruebaForbidden creates PostPruebaForbidden with default headers values
func NewPostPruebaForbidden() *PostPruebaForbidden {

	return &PostPruebaForbidden{}
}

// WithPayload adds the payload to the post prueba forbidden response
func (o *PostPruebaForbidden) WithPayload(payload *models.Error) *PostPruebaForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post prueba forbidden response
func (o *PostPruebaForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPruebaForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPruebaGoneCode is the HTTP code returned for type PostPruebaGone
const PostPruebaGoneCode int = 410

/*PostPruebaGone That resource does not exist

swagger:response postPruebaGone
*/
type PostPruebaGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostPruebaGone creates PostPruebaGone with default headers values
func NewPostPruebaGone() *PostPruebaGone {

	return &PostPruebaGone{}
}

// WithPayload adds the payload to the post prueba gone response
func (o *PostPruebaGone) WithPayload(payload *models.Error) *PostPruebaGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post prueba gone response
func (o *PostPruebaGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPruebaGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPruebaInternalServerErrorCode is the HTTP code returned for type PostPruebaInternalServerError
const PostPruebaInternalServerErrorCode int = 500

/*PostPruebaInternalServerError Internal error

swagger:response postPruebaInternalServerError
*/
type PostPruebaInternalServerError struct {
}

// NewPostPruebaInternalServerError creates PostPruebaInternalServerError with default headers values
func NewPostPruebaInternalServerError() *PostPruebaInternalServerError {

	return &PostPruebaInternalServerError{}
}

// WriteResponse to the client
func (o *PostPruebaInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
