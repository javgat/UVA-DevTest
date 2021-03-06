// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetQuestionFromTeamOKCode is the HTTP code returned for type GetQuestionFromTeamOK
const GetQuestionFromTeamOKCode int = 200

/*GetQuestionFromTeamOK question found

swagger:response getQuestionFromTeamOK
*/
type GetQuestionFromTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.Question `json:"body,omitempty"`
}

// NewGetQuestionFromTeamOK creates GetQuestionFromTeamOK with default headers values
func NewGetQuestionFromTeamOK() *GetQuestionFromTeamOK {

	return &GetQuestionFromTeamOK{}
}

// WithPayload adds the payload to the get question from team o k response
func (o *GetQuestionFromTeamOK) WithPayload(payload *models.Question) *GetQuestionFromTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from team o k response
func (o *GetQuestionFromTeamOK) SetPayload(payload *models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromTeamBadRequestCode is the HTTP code returned for type GetQuestionFromTeamBadRequest
const GetQuestionFromTeamBadRequestCode int = 400

/*GetQuestionFromTeamBadRequest Incorrect Request, or invalida data

swagger:response getQuestionFromTeamBadRequest
*/
type GetQuestionFromTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionFromTeamBadRequest creates GetQuestionFromTeamBadRequest with default headers values
func NewGetQuestionFromTeamBadRequest() *GetQuestionFromTeamBadRequest {

	return &GetQuestionFromTeamBadRequest{}
}

// WithPayload adds the payload to the get question from team bad request response
func (o *GetQuestionFromTeamBadRequest) WithPayload(payload *models.Error) *GetQuestionFromTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from team bad request response
func (o *GetQuestionFromTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromTeamForbiddenCode is the HTTP code returned for type GetQuestionFromTeamForbidden
const GetQuestionFromTeamForbiddenCode int = 403

/*GetQuestionFromTeamForbidden Not authorized to this content

swagger:response getQuestionFromTeamForbidden
*/
type GetQuestionFromTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionFromTeamForbidden creates GetQuestionFromTeamForbidden with default headers values
func NewGetQuestionFromTeamForbidden() *GetQuestionFromTeamForbidden {

	return &GetQuestionFromTeamForbidden{}
}

// WithPayload adds the payload to the get question from team forbidden response
func (o *GetQuestionFromTeamForbidden) WithPayload(payload *models.Error) *GetQuestionFromTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from team forbidden response
func (o *GetQuestionFromTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromTeamGoneCode is the HTTP code returned for type GetQuestionFromTeamGone
const GetQuestionFromTeamGoneCode int = 410

/*GetQuestionFromTeamGone That user (password and name) does not exist

swagger:response getQuestionFromTeamGone
*/
type GetQuestionFromTeamGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionFromTeamGone creates GetQuestionFromTeamGone with default headers values
func NewGetQuestionFromTeamGone() *GetQuestionFromTeamGone {

	return &GetQuestionFromTeamGone{}
}

// WithPayload adds the payload to the get question from team gone response
func (o *GetQuestionFromTeamGone) WithPayload(payload *models.Error) *GetQuestionFromTeamGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from team gone response
func (o *GetQuestionFromTeamGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromTeamGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromTeamInternalServerErrorCode is the HTTP code returned for type GetQuestionFromTeamInternalServerError
const GetQuestionFromTeamInternalServerErrorCode int = 500

/*GetQuestionFromTeamInternalServerError Internal error

swagger:response getQuestionFromTeamInternalServerError
*/
type GetQuestionFromTeamInternalServerError struct {
}

// NewGetQuestionFromTeamInternalServerError creates GetQuestionFromTeamInternalServerError with default headers values
func NewGetQuestionFromTeamInternalServerError() *GetQuestionFromTeamInternalServerError {

	return &GetQuestionFromTeamInternalServerError{}
}

// WriteResponse to the client
func (o *GetQuestionFromTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
