// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetEditQuestionsOfUserOKCode is the HTTP code returned for type GetEditQuestionsOfUserOK
const GetEditQuestionsOfUserOKCode int = 200

/*GetEditQuestionsOfUserOK questions found

swagger:response getEditQuestionsOfUserOK
*/
type GetEditQuestionsOfUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Question `json:"body,omitempty"`
}

// NewGetEditQuestionsOfUserOK creates GetEditQuestionsOfUserOK with default headers values
func NewGetEditQuestionsOfUserOK() *GetEditQuestionsOfUserOK {

	return &GetEditQuestionsOfUserOK{}
}

// WithPayload adds the payload to the get edit questions of user o k response
func (o *GetEditQuestionsOfUserOK) WithPayload(payload []*models.Question) *GetEditQuestionsOfUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit questions of user o k response
func (o *GetEditQuestionsOfUserOK) SetPayload(payload []*models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditQuestionsOfUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetEditQuestionsOfUserBadRequestCode is the HTTP code returned for type GetEditQuestionsOfUserBadRequest
const GetEditQuestionsOfUserBadRequestCode int = 400

/*GetEditQuestionsOfUserBadRequest Incorrect Request, or invalida data

swagger:response getEditQuestionsOfUserBadRequest
*/
type GetEditQuestionsOfUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEditQuestionsOfUserBadRequest creates GetEditQuestionsOfUserBadRequest with default headers values
func NewGetEditQuestionsOfUserBadRequest() *GetEditQuestionsOfUserBadRequest {

	return &GetEditQuestionsOfUserBadRequest{}
}

// WithPayload adds the payload to the get edit questions of user bad request response
func (o *GetEditQuestionsOfUserBadRequest) WithPayload(payload *models.Error) *GetEditQuestionsOfUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit questions of user bad request response
func (o *GetEditQuestionsOfUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditQuestionsOfUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEditQuestionsOfUserForbiddenCode is the HTTP code returned for type GetEditQuestionsOfUserForbidden
const GetEditQuestionsOfUserForbiddenCode int = 403

/*GetEditQuestionsOfUserForbidden Not authorized to this content

swagger:response getEditQuestionsOfUserForbidden
*/
type GetEditQuestionsOfUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEditQuestionsOfUserForbidden creates GetEditQuestionsOfUserForbidden with default headers values
func NewGetEditQuestionsOfUserForbidden() *GetEditQuestionsOfUserForbidden {

	return &GetEditQuestionsOfUserForbidden{}
}

// WithPayload adds the payload to the get edit questions of user forbidden response
func (o *GetEditQuestionsOfUserForbidden) WithPayload(payload *models.Error) *GetEditQuestionsOfUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit questions of user forbidden response
func (o *GetEditQuestionsOfUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditQuestionsOfUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEditQuestionsOfUserGoneCode is the HTTP code returned for type GetEditQuestionsOfUserGone
const GetEditQuestionsOfUserGoneCode int = 410

/*GetEditQuestionsOfUserGone That user (password and name) does not exist

swagger:response getEditQuestionsOfUserGone
*/
type GetEditQuestionsOfUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEditQuestionsOfUserGone creates GetEditQuestionsOfUserGone with default headers values
func NewGetEditQuestionsOfUserGone() *GetEditQuestionsOfUserGone {

	return &GetEditQuestionsOfUserGone{}
}

// WithPayload adds the payload to the get edit questions of user gone response
func (o *GetEditQuestionsOfUserGone) WithPayload(payload *models.Error) *GetEditQuestionsOfUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get edit questions of user gone response
func (o *GetEditQuestionsOfUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEditQuestionsOfUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEditQuestionsOfUserInternalServerErrorCode is the HTTP code returned for type GetEditQuestionsOfUserInternalServerError
const GetEditQuestionsOfUserInternalServerErrorCode int = 500

/*GetEditQuestionsOfUserInternalServerError Internal error

swagger:response getEditQuestionsOfUserInternalServerError
*/
type GetEditQuestionsOfUserInternalServerError struct {
}

// NewGetEditQuestionsOfUserInternalServerError creates GetEditQuestionsOfUserInternalServerError with default headers values
func NewGetEditQuestionsOfUserInternalServerError() *GetEditQuestionsOfUserInternalServerError {

	return &GetEditQuestionsOfUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetEditQuestionsOfUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
