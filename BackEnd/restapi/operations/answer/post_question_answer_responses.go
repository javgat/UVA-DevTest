// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PostQuestionAnswerCreatedCode is the HTTP code returned for type PostQuestionAnswerCreated
const PostQuestionAnswerCreatedCode int = 201

/*PostQuestionAnswerCreated QuestionAnswer created

swagger:response postQuestionAnswerCreated
*/
type PostQuestionAnswerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.QuestionAnswer `json:"body,omitempty"`
}

// NewPostQuestionAnswerCreated creates PostQuestionAnswerCreated with default headers values
func NewPostQuestionAnswerCreated() *PostQuestionAnswerCreated {

	return &PostQuestionAnswerCreated{}
}

// WithPayload adds the payload to the post question answer created response
func (o *PostQuestionAnswerCreated) WithPayload(payload *models.QuestionAnswer) *PostQuestionAnswerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post question answer created response
func (o *PostQuestionAnswerCreated) SetPayload(payload *models.QuestionAnswer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostQuestionAnswerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostQuestionAnswerBadRequestCode is the HTTP code returned for type PostQuestionAnswerBadRequest
const PostQuestionAnswerBadRequestCode int = 400

/*PostQuestionAnswerBadRequest Incorrect Request, or invalida data

swagger:response postQuestionAnswerBadRequest
*/
type PostQuestionAnswerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostQuestionAnswerBadRequest creates PostQuestionAnswerBadRequest with default headers values
func NewPostQuestionAnswerBadRequest() *PostQuestionAnswerBadRequest {

	return &PostQuestionAnswerBadRequest{}
}

// WithPayload adds the payload to the post question answer bad request response
func (o *PostQuestionAnswerBadRequest) WithPayload(payload *models.Error) *PostQuestionAnswerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post question answer bad request response
func (o *PostQuestionAnswerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostQuestionAnswerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostQuestionAnswerForbiddenCode is the HTTP code returned for type PostQuestionAnswerForbidden
const PostQuestionAnswerForbiddenCode int = 403

/*PostQuestionAnswerForbidden Not authorized to this content

swagger:response postQuestionAnswerForbidden
*/
type PostQuestionAnswerForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostQuestionAnswerForbidden creates PostQuestionAnswerForbidden with default headers values
func NewPostQuestionAnswerForbidden() *PostQuestionAnswerForbidden {

	return &PostQuestionAnswerForbidden{}
}

// WithPayload adds the payload to the post question answer forbidden response
func (o *PostQuestionAnswerForbidden) WithPayload(payload *models.Error) *PostQuestionAnswerForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post question answer forbidden response
func (o *PostQuestionAnswerForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostQuestionAnswerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostQuestionAnswerGoneCode is the HTTP code returned for type PostQuestionAnswerGone
const PostQuestionAnswerGoneCode int = 410

/*PostQuestionAnswerGone That user (password and name) does not exist

swagger:response postQuestionAnswerGone
*/
type PostQuestionAnswerGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostQuestionAnswerGone creates PostQuestionAnswerGone with default headers values
func NewPostQuestionAnswerGone() *PostQuestionAnswerGone {

	return &PostQuestionAnswerGone{}
}

// WithPayload adds the payload to the post question answer gone response
func (o *PostQuestionAnswerGone) WithPayload(payload *models.Error) *PostQuestionAnswerGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post question answer gone response
func (o *PostQuestionAnswerGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostQuestionAnswerGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostQuestionAnswerInternalServerErrorCode is the HTTP code returned for type PostQuestionAnswerInternalServerError
const PostQuestionAnswerInternalServerErrorCode int = 500

/*PostQuestionAnswerInternalServerError Internal error

swagger:response postQuestionAnswerInternalServerError
*/
type PostQuestionAnswerInternalServerError struct {
}

// NewPostQuestionAnswerInternalServerError creates PostQuestionAnswerInternalServerError with default headers values
func NewPostQuestionAnswerInternalServerError() *PostQuestionAnswerInternalServerError {

	return &PostQuestionAnswerInternalServerError{}
}

// WriteResponse to the client
func (o *PostQuestionAnswerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
