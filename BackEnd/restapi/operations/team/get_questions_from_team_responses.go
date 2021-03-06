// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetQuestionsFromTeamOKCode is the HTTP code returned for type GetQuestionsFromTeamOK
const GetQuestionsFromTeamOKCode int = 200

/*GetQuestionsFromTeamOK questions found

swagger:response getQuestionsFromTeamOK
*/
type GetQuestionsFromTeamOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Question `json:"body,omitempty"`
}

// NewGetQuestionsFromTeamOK creates GetQuestionsFromTeamOK with default headers values
func NewGetQuestionsFromTeamOK() *GetQuestionsFromTeamOK {

	return &GetQuestionsFromTeamOK{}
}

// WithPayload adds the payload to the get questions from team o k response
func (o *GetQuestionsFromTeamOK) WithPayload(payload []*models.Question) *GetQuestionsFromTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get questions from team o k response
func (o *GetQuestionsFromTeamOK) SetPayload(payload []*models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionsFromTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Question, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetQuestionsFromTeamBadRequestCode is the HTTP code returned for type GetQuestionsFromTeamBadRequest
const GetQuestionsFromTeamBadRequestCode int = 400

/*GetQuestionsFromTeamBadRequest Incorrect Request, or invalida data

swagger:response getQuestionsFromTeamBadRequest
*/
type GetQuestionsFromTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionsFromTeamBadRequest creates GetQuestionsFromTeamBadRequest with default headers values
func NewGetQuestionsFromTeamBadRequest() *GetQuestionsFromTeamBadRequest {

	return &GetQuestionsFromTeamBadRequest{}
}

// WithPayload adds the payload to the get questions from team bad request response
func (o *GetQuestionsFromTeamBadRequest) WithPayload(payload *models.Error) *GetQuestionsFromTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get questions from team bad request response
func (o *GetQuestionsFromTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionsFromTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionsFromTeamForbiddenCode is the HTTP code returned for type GetQuestionsFromTeamForbidden
const GetQuestionsFromTeamForbiddenCode int = 403

/*GetQuestionsFromTeamForbidden Not authorized to this content

swagger:response getQuestionsFromTeamForbidden
*/
type GetQuestionsFromTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionsFromTeamForbidden creates GetQuestionsFromTeamForbidden with default headers values
func NewGetQuestionsFromTeamForbidden() *GetQuestionsFromTeamForbidden {

	return &GetQuestionsFromTeamForbidden{}
}

// WithPayload adds the payload to the get questions from team forbidden response
func (o *GetQuestionsFromTeamForbidden) WithPayload(payload *models.Error) *GetQuestionsFromTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get questions from team forbidden response
func (o *GetQuestionsFromTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionsFromTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionsFromTeamGoneCode is the HTTP code returned for type GetQuestionsFromTeamGone
const GetQuestionsFromTeamGoneCode int = 410

/*GetQuestionsFromTeamGone That user (password and name) does not exist

swagger:response getQuestionsFromTeamGone
*/
type GetQuestionsFromTeamGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionsFromTeamGone creates GetQuestionsFromTeamGone with default headers values
func NewGetQuestionsFromTeamGone() *GetQuestionsFromTeamGone {

	return &GetQuestionsFromTeamGone{}
}

// WithPayload adds the payload to the get questions from team gone response
func (o *GetQuestionsFromTeamGone) WithPayload(payload *models.Error) *GetQuestionsFromTeamGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get questions from team gone response
func (o *GetQuestionsFromTeamGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionsFromTeamGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionsFromTeamInternalServerErrorCode is the HTTP code returned for type GetQuestionsFromTeamInternalServerError
const GetQuestionsFromTeamInternalServerErrorCode int = 500

/*GetQuestionsFromTeamInternalServerError Internal error

swagger:response getQuestionsFromTeamInternalServerError
*/
type GetQuestionsFromTeamInternalServerError struct {
}

// NewGetQuestionsFromTeamInternalServerError creates GetQuestionsFromTeamInternalServerError with default headers values
func NewGetQuestionsFromTeamInternalServerError() *GetQuestionsFromTeamInternalServerError {

	return &GetQuestionsFromTeamInternalServerError{}
}

// WriteResponse to the client
func (o *GetQuestionsFromTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
