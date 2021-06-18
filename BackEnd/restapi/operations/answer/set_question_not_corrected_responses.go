// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// SetQuestionNotCorrectedOKCode is the HTTP code returned for type SetQuestionNotCorrectedOK
const SetQuestionNotCorrectedOKCode int = 200

/*SetQuestionNotCorrectedOK Answer updated

swagger:response setQuestionNotCorrectedOK
*/
type SetQuestionNotCorrectedOK struct {
}

// NewSetQuestionNotCorrectedOK creates SetQuestionNotCorrectedOK with default headers values
func NewSetQuestionNotCorrectedOK() *SetQuestionNotCorrectedOK {

	return &SetQuestionNotCorrectedOK{}
}

// WriteResponse to the client
func (o *SetQuestionNotCorrectedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SetQuestionNotCorrectedForbiddenCode is the HTTP code returned for type SetQuestionNotCorrectedForbidden
const SetQuestionNotCorrectedForbiddenCode int = 403

/*SetQuestionNotCorrectedForbidden Not authorized to this content

swagger:response setQuestionNotCorrectedForbidden
*/
type SetQuestionNotCorrectedForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetQuestionNotCorrectedForbidden creates SetQuestionNotCorrectedForbidden with default headers values
func NewSetQuestionNotCorrectedForbidden() *SetQuestionNotCorrectedForbidden {

	return &SetQuestionNotCorrectedForbidden{}
}

// WithPayload adds the payload to the set question not corrected forbidden response
func (o *SetQuestionNotCorrectedForbidden) WithPayload(payload *models.Error) *SetQuestionNotCorrectedForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set question not corrected forbidden response
func (o *SetQuestionNotCorrectedForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetQuestionNotCorrectedForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetQuestionNotCorrectedGoneCode is the HTTP code returned for type SetQuestionNotCorrectedGone
const SetQuestionNotCorrectedGoneCode int = 410

/*SetQuestionNotCorrectedGone That resource does not exist

swagger:response setQuestionNotCorrectedGone
*/
type SetQuestionNotCorrectedGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetQuestionNotCorrectedGone creates SetQuestionNotCorrectedGone with default headers values
func NewSetQuestionNotCorrectedGone() *SetQuestionNotCorrectedGone {

	return &SetQuestionNotCorrectedGone{}
}

// WithPayload adds the payload to the set question not corrected gone response
func (o *SetQuestionNotCorrectedGone) WithPayload(payload *models.Error) *SetQuestionNotCorrectedGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set question not corrected gone response
func (o *SetQuestionNotCorrectedGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetQuestionNotCorrectedGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetQuestionNotCorrectedInternalServerErrorCode is the HTTP code returned for type SetQuestionNotCorrectedInternalServerError
const SetQuestionNotCorrectedInternalServerErrorCode int = 500

/*SetQuestionNotCorrectedInternalServerError Internal error

swagger:response setQuestionNotCorrectedInternalServerError
*/
type SetQuestionNotCorrectedInternalServerError struct {
}

// NewSetQuestionNotCorrectedInternalServerError creates SetQuestionNotCorrectedInternalServerError with default headers values
func NewSetQuestionNotCorrectedInternalServerError() *SetQuestionNotCorrectedInternalServerError {

	return &SetQuestionNotCorrectedInternalServerError{}
}

// WriteResponse to the client
func (o *SetQuestionNotCorrectedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
