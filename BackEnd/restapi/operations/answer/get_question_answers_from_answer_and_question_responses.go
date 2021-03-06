// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetQuestionAnswersFromAnswerAndQuestionOKCode is the HTTP code returned for type GetQuestionAnswersFromAnswerAndQuestionOK
const GetQuestionAnswersFromAnswerAndQuestionOKCode int = 200

/*GetQuestionAnswersFromAnswerAndQuestionOK QuestionAnswers found

swagger:response getQuestionAnswersFromAnswerAndQuestionOK
*/
type GetQuestionAnswersFromAnswerAndQuestionOK struct {

	/*
	  In: Body
	*/
	Payload []*models.QuestionAnswer `json:"body,omitempty"`
}

// NewGetQuestionAnswersFromAnswerAndQuestionOK creates GetQuestionAnswersFromAnswerAndQuestionOK with default headers values
func NewGetQuestionAnswersFromAnswerAndQuestionOK() *GetQuestionAnswersFromAnswerAndQuestionOK {

	return &GetQuestionAnswersFromAnswerAndQuestionOK{}
}

// WithPayload adds the payload to the get question answers from answer and question o k response
func (o *GetQuestionAnswersFromAnswerAndQuestionOK) WithPayload(payload []*models.QuestionAnswer) *GetQuestionAnswersFromAnswerAndQuestionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question answers from answer and question o k response
func (o *GetQuestionAnswersFromAnswerAndQuestionOK) SetPayload(payload []*models.QuestionAnswer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionAnswersFromAnswerAndQuestionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.QuestionAnswer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetQuestionAnswersFromAnswerAndQuestionBadRequestCode is the HTTP code returned for type GetQuestionAnswersFromAnswerAndQuestionBadRequest
const GetQuestionAnswersFromAnswerAndQuestionBadRequestCode int = 400

/*GetQuestionAnswersFromAnswerAndQuestionBadRequest Incorrect Request, or invalida data

swagger:response getQuestionAnswersFromAnswerAndQuestionBadRequest
*/
type GetQuestionAnswersFromAnswerAndQuestionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionAnswersFromAnswerAndQuestionBadRequest creates GetQuestionAnswersFromAnswerAndQuestionBadRequest with default headers values
func NewGetQuestionAnswersFromAnswerAndQuestionBadRequest() *GetQuestionAnswersFromAnswerAndQuestionBadRequest {

	return &GetQuestionAnswersFromAnswerAndQuestionBadRequest{}
}

// WithPayload adds the payload to the get question answers from answer and question bad request response
func (o *GetQuestionAnswersFromAnswerAndQuestionBadRequest) WithPayload(payload *models.Error) *GetQuestionAnswersFromAnswerAndQuestionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question answers from answer and question bad request response
func (o *GetQuestionAnswersFromAnswerAndQuestionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionAnswersFromAnswerAndQuestionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionAnswersFromAnswerAndQuestionForbiddenCode is the HTTP code returned for type GetQuestionAnswersFromAnswerAndQuestionForbidden
const GetQuestionAnswersFromAnswerAndQuestionForbiddenCode int = 403

/*GetQuestionAnswersFromAnswerAndQuestionForbidden Not authorized to this content

swagger:response getQuestionAnswersFromAnswerAndQuestionForbidden
*/
type GetQuestionAnswersFromAnswerAndQuestionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionAnswersFromAnswerAndQuestionForbidden creates GetQuestionAnswersFromAnswerAndQuestionForbidden with default headers values
func NewGetQuestionAnswersFromAnswerAndQuestionForbidden() *GetQuestionAnswersFromAnswerAndQuestionForbidden {

	return &GetQuestionAnswersFromAnswerAndQuestionForbidden{}
}

// WithPayload adds the payload to the get question answers from answer and question forbidden response
func (o *GetQuestionAnswersFromAnswerAndQuestionForbidden) WithPayload(payload *models.Error) *GetQuestionAnswersFromAnswerAndQuestionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question answers from answer and question forbidden response
func (o *GetQuestionAnswersFromAnswerAndQuestionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionAnswersFromAnswerAndQuestionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionAnswersFromAnswerAndQuestionGoneCode is the HTTP code returned for type GetQuestionAnswersFromAnswerAndQuestionGone
const GetQuestionAnswersFromAnswerAndQuestionGoneCode int = 410

/*GetQuestionAnswersFromAnswerAndQuestionGone That resource does not exist

swagger:response getQuestionAnswersFromAnswerAndQuestionGone
*/
type GetQuestionAnswersFromAnswerAndQuestionGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionAnswersFromAnswerAndQuestionGone creates GetQuestionAnswersFromAnswerAndQuestionGone with default headers values
func NewGetQuestionAnswersFromAnswerAndQuestionGone() *GetQuestionAnswersFromAnswerAndQuestionGone {

	return &GetQuestionAnswersFromAnswerAndQuestionGone{}
}

// WithPayload adds the payload to the get question answers from answer and question gone response
func (o *GetQuestionAnswersFromAnswerAndQuestionGone) WithPayload(payload *models.Error) *GetQuestionAnswersFromAnswerAndQuestionGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question answers from answer and question gone response
func (o *GetQuestionAnswersFromAnswerAndQuestionGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionAnswersFromAnswerAndQuestionGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionAnswersFromAnswerAndQuestionInternalServerErrorCode is the HTTP code returned for type GetQuestionAnswersFromAnswerAndQuestionInternalServerError
const GetQuestionAnswersFromAnswerAndQuestionInternalServerErrorCode int = 500

/*GetQuestionAnswersFromAnswerAndQuestionInternalServerError Internal error

swagger:response getQuestionAnswersFromAnswerAndQuestionInternalServerError
*/
type GetQuestionAnswersFromAnswerAndQuestionInternalServerError struct {
}

// NewGetQuestionAnswersFromAnswerAndQuestionInternalServerError creates GetQuestionAnswersFromAnswerAndQuestionInternalServerError with default headers values
func NewGetQuestionAnswersFromAnswerAndQuestionInternalServerError() *GetQuestionAnswersFromAnswerAndQuestionInternalServerError {

	return &GetQuestionAnswersFromAnswerAndQuestionInternalServerError{}
}

// WriteResponse to the client
func (o *GetQuestionAnswersFromAnswerAndQuestionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
