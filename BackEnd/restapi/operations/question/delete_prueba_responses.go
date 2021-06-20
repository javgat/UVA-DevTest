// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// DeletePruebaOKCode is the HTTP code returned for type DeletePruebaOK
const DeletePruebaOKCode int = 200

/*DeletePruebaOK prueba deleted

swagger:response deletePruebaOK
*/
type DeletePruebaOK struct {
}

// NewDeletePruebaOK creates DeletePruebaOK with default headers values
func NewDeletePruebaOK() *DeletePruebaOK {

	return &DeletePruebaOK{}
}

// WriteResponse to the client
func (o *DeletePruebaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeletePruebaBadRequestCode is the HTTP code returned for type DeletePruebaBadRequest
const DeletePruebaBadRequestCode int = 400

/*DeletePruebaBadRequest Incorrect Request, or invalida data

swagger:response deletePruebaBadRequest
*/
type DeletePruebaBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePruebaBadRequest creates DeletePruebaBadRequest with default headers values
func NewDeletePruebaBadRequest() *DeletePruebaBadRequest {

	return &DeletePruebaBadRequest{}
}

// WithPayload adds the payload to the delete prueba bad request response
func (o *DeletePruebaBadRequest) WithPayload(payload *models.Error) *DeletePruebaBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete prueba bad request response
func (o *DeletePruebaBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePruebaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePruebaForbiddenCode is the HTTP code returned for type DeletePruebaForbidden
const DeletePruebaForbiddenCode int = 403

/*DeletePruebaForbidden Not authorized to this content

swagger:response deletePruebaForbidden
*/
type DeletePruebaForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePruebaForbidden creates DeletePruebaForbidden with default headers values
func NewDeletePruebaForbidden() *DeletePruebaForbidden {

	return &DeletePruebaForbidden{}
}

// WithPayload adds the payload to the delete prueba forbidden response
func (o *DeletePruebaForbidden) WithPayload(payload *models.Error) *DeletePruebaForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete prueba forbidden response
func (o *DeletePruebaForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePruebaForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePruebaGoneCode is the HTTP code returned for type DeletePruebaGone
const DeletePruebaGoneCode int = 410

/*DeletePruebaGone That resource does not exist

swagger:response deletePruebaGone
*/
type DeletePruebaGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePruebaGone creates DeletePruebaGone with default headers values
func NewDeletePruebaGone() *DeletePruebaGone {

	return &DeletePruebaGone{}
}

// WithPayload adds the payload to the delete prueba gone response
func (o *DeletePruebaGone) WithPayload(payload *models.Error) *DeletePruebaGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete prueba gone response
func (o *DeletePruebaGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePruebaGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePruebaInternalServerErrorCode is the HTTP code returned for type DeletePruebaInternalServerError
const DeletePruebaInternalServerErrorCode int = 500

/*DeletePruebaInternalServerError Internal error

swagger:response deletePruebaInternalServerError
*/
type DeletePruebaInternalServerError struct {
}

// NewDeletePruebaInternalServerError creates DeletePruebaInternalServerError with default headers values
func NewDeletePruebaInternalServerError() *DeletePruebaInternalServerError {

	return &DeletePruebaInternalServerError{}
}

// WriteResponse to the client
func (o *DeletePruebaInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}