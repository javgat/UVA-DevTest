// Code generated by go-swagger; DO NOT EDIT.

package tiporol

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PutTipoRolOKCode is the HTTP code returned for type PutTipoRolOK
const PutTipoRolOKCode int = 200

/*PutTipoRolOK TipoRol modified

swagger:response putTipoRolOK
*/
type PutTipoRolOK struct {
}

// NewPutTipoRolOK creates PutTipoRolOK with default headers values
func NewPutTipoRolOK() *PutTipoRolOK {

	return &PutTipoRolOK{}
}

// WriteResponse to the client
func (o *PutTipoRolOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutTipoRolBadRequestCode is the HTTP code returned for type PutTipoRolBadRequest
const PutTipoRolBadRequestCode int = 400

/*PutTipoRolBadRequest Incorrect Request, or invalida data

swagger:response putTipoRolBadRequest
*/
type PutTipoRolBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTipoRolBadRequest creates PutTipoRolBadRequest with default headers values
func NewPutTipoRolBadRequest() *PutTipoRolBadRequest {

	return &PutTipoRolBadRequest{}
}

// WithPayload adds the payload to the put tipo rol bad request response
func (o *PutTipoRolBadRequest) WithPayload(payload *models.Error) *PutTipoRolBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tipo rol bad request response
func (o *PutTipoRolBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTipoRolBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTipoRolForbiddenCode is the HTTP code returned for type PutTipoRolForbidden
const PutTipoRolForbiddenCode int = 403

/*PutTipoRolForbidden Not authorized to this content

swagger:response putTipoRolForbidden
*/
type PutTipoRolForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTipoRolForbidden creates PutTipoRolForbidden with default headers values
func NewPutTipoRolForbidden() *PutTipoRolForbidden {

	return &PutTipoRolForbidden{}
}

// WithPayload adds the payload to the put tipo rol forbidden response
func (o *PutTipoRolForbidden) WithPayload(payload *models.Error) *PutTipoRolForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tipo rol forbidden response
func (o *PutTipoRolForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTipoRolForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTipoRolConflictCode is the HTTP code returned for type PutTipoRolConflict
const PutTipoRolConflictCode int = 409

/*PutTipoRolConflict A user with same username/email already exists

swagger:response putTipoRolConflict
*/
type PutTipoRolConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTipoRolConflict creates PutTipoRolConflict with default headers values
func NewPutTipoRolConflict() *PutTipoRolConflict {

	return &PutTipoRolConflict{}
}

// WithPayload adds the payload to the put tipo rol conflict response
func (o *PutTipoRolConflict) WithPayload(payload *models.Error) *PutTipoRolConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tipo rol conflict response
func (o *PutTipoRolConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTipoRolConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTipoRolGoneCode is the HTTP code returned for type PutTipoRolGone
const PutTipoRolGoneCode int = 410

/*PutTipoRolGone That resource does not exist

swagger:response putTipoRolGone
*/
type PutTipoRolGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTipoRolGone creates PutTipoRolGone with default headers values
func NewPutTipoRolGone() *PutTipoRolGone {

	return &PutTipoRolGone{}
}

// WithPayload adds the payload to the put tipo rol gone response
func (o *PutTipoRolGone) WithPayload(payload *models.Error) *PutTipoRolGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tipo rol gone response
func (o *PutTipoRolGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTipoRolGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTipoRolInternalServerErrorCode is the HTTP code returned for type PutTipoRolInternalServerError
const PutTipoRolInternalServerErrorCode int = 500

/*PutTipoRolInternalServerError Internal error

swagger:response putTipoRolInternalServerError
*/
type PutTipoRolInternalServerError struct {
}

// NewPutTipoRolInternalServerError creates PutTipoRolInternalServerError with default headers values
func NewPutTipoRolInternalServerError() *PutTipoRolInternalServerError {

	return &PutTipoRolInternalServerError{}
}

// WriteResponse to the client
func (o *PutTipoRolInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
