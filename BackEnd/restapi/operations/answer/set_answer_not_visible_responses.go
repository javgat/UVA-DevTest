// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// SetAnswerNotVisibleOKCode is the HTTP code returned for type SetAnswerNotVisibleOK
const SetAnswerNotVisibleOKCode int = 200

/*SetAnswerNotVisibleOK Answer updated

swagger:response setAnswerNotVisibleOK
*/
type SetAnswerNotVisibleOK struct {
}

// NewSetAnswerNotVisibleOK creates SetAnswerNotVisibleOK with default headers values
func NewSetAnswerNotVisibleOK() *SetAnswerNotVisibleOK {

	return &SetAnswerNotVisibleOK{}
}

// WriteResponse to the client
func (o *SetAnswerNotVisibleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SetAnswerNotVisibleForbiddenCode is the HTTP code returned for type SetAnswerNotVisibleForbidden
const SetAnswerNotVisibleForbiddenCode int = 403

/*SetAnswerNotVisibleForbidden Not authorized to this content

swagger:response setAnswerNotVisibleForbidden
*/
type SetAnswerNotVisibleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAnswerNotVisibleForbidden creates SetAnswerNotVisibleForbidden with default headers values
func NewSetAnswerNotVisibleForbidden() *SetAnswerNotVisibleForbidden {

	return &SetAnswerNotVisibleForbidden{}
}

// WithPayload adds the payload to the set answer not visible forbidden response
func (o *SetAnswerNotVisibleForbidden) WithPayload(payload *models.Error) *SetAnswerNotVisibleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set answer not visible forbidden response
func (o *SetAnswerNotVisibleForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAnswerNotVisibleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetAnswerNotVisibleGoneCode is the HTTP code returned for type SetAnswerNotVisibleGone
const SetAnswerNotVisibleGoneCode int = 410

/*SetAnswerNotVisibleGone That resource does not exist

swagger:response setAnswerNotVisibleGone
*/
type SetAnswerNotVisibleGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAnswerNotVisibleGone creates SetAnswerNotVisibleGone with default headers values
func NewSetAnswerNotVisibleGone() *SetAnswerNotVisibleGone {

	return &SetAnswerNotVisibleGone{}
}

// WithPayload adds the payload to the set answer not visible gone response
func (o *SetAnswerNotVisibleGone) WithPayload(payload *models.Error) *SetAnswerNotVisibleGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set answer not visible gone response
func (o *SetAnswerNotVisibleGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAnswerNotVisibleGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetAnswerNotVisibleInternalServerErrorCode is the HTTP code returned for type SetAnswerNotVisibleInternalServerError
const SetAnswerNotVisibleInternalServerErrorCode int = 500

/*SetAnswerNotVisibleInternalServerError Internal error

swagger:response setAnswerNotVisibleInternalServerError
*/
type SetAnswerNotVisibleInternalServerError struct {
}

// NewSetAnswerNotVisibleInternalServerError creates SetAnswerNotVisibleInternalServerError with default headers values
func NewSetAnswerNotVisibleInternalServerError() *SetAnswerNotVisibleInternalServerError {

	return &SetAnswerNotVisibleInternalServerError{}
}

// WriteResponse to the client
func (o *SetAnswerNotVisibleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
