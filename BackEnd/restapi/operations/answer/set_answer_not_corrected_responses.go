// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// SetAnswerNotCorrectedOKCode is the HTTP code returned for type SetAnswerNotCorrectedOK
const SetAnswerNotCorrectedOKCode int = 200

/*SetAnswerNotCorrectedOK Answer updated

swagger:response setAnswerNotCorrectedOK
*/
type SetAnswerNotCorrectedOK struct {
}

// NewSetAnswerNotCorrectedOK creates SetAnswerNotCorrectedOK with default headers values
func NewSetAnswerNotCorrectedOK() *SetAnswerNotCorrectedOK {

	return &SetAnswerNotCorrectedOK{}
}

// WriteResponse to the client
func (o *SetAnswerNotCorrectedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SetAnswerNotCorrectedForbiddenCode is the HTTP code returned for type SetAnswerNotCorrectedForbidden
const SetAnswerNotCorrectedForbiddenCode int = 403

/*SetAnswerNotCorrectedForbidden Not authorized to this content

swagger:response setAnswerNotCorrectedForbidden
*/
type SetAnswerNotCorrectedForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAnswerNotCorrectedForbidden creates SetAnswerNotCorrectedForbidden with default headers values
func NewSetAnswerNotCorrectedForbidden() *SetAnswerNotCorrectedForbidden {

	return &SetAnswerNotCorrectedForbidden{}
}

// WithPayload adds the payload to the set answer not corrected forbidden response
func (o *SetAnswerNotCorrectedForbidden) WithPayload(payload *models.Error) *SetAnswerNotCorrectedForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set answer not corrected forbidden response
func (o *SetAnswerNotCorrectedForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAnswerNotCorrectedForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetAnswerNotCorrectedGoneCode is the HTTP code returned for type SetAnswerNotCorrectedGone
const SetAnswerNotCorrectedGoneCode int = 410

/*SetAnswerNotCorrectedGone That resource does not exist

swagger:response setAnswerNotCorrectedGone
*/
type SetAnswerNotCorrectedGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetAnswerNotCorrectedGone creates SetAnswerNotCorrectedGone with default headers values
func NewSetAnswerNotCorrectedGone() *SetAnswerNotCorrectedGone {

	return &SetAnswerNotCorrectedGone{}
}

// WithPayload adds the payload to the set answer not corrected gone response
func (o *SetAnswerNotCorrectedGone) WithPayload(payload *models.Error) *SetAnswerNotCorrectedGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set answer not corrected gone response
func (o *SetAnswerNotCorrectedGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetAnswerNotCorrectedGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetAnswerNotCorrectedInternalServerErrorCode is the HTTP code returned for type SetAnswerNotCorrectedInternalServerError
const SetAnswerNotCorrectedInternalServerErrorCode int = 500

/*SetAnswerNotCorrectedInternalServerError Internal error

swagger:response setAnswerNotCorrectedInternalServerError
*/
type SetAnswerNotCorrectedInternalServerError struct {
}

// NewSetAnswerNotCorrectedInternalServerError creates SetAnswerNotCorrectedInternalServerError with default headers values
func NewSetAnswerNotCorrectedInternalServerError() *SetAnswerNotCorrectedInternalServerError {

	return &SetAnswerNotCorrectedInternalServerError{}
}

// WriteResponse to the client
func (o *SetAnswerNotCorrectedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
