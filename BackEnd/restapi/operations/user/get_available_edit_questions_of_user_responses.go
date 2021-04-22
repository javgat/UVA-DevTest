// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetAvailableEditQuestionsOfUserOKCode is the HTTP code returned for type GetAvailableEditQuestionsOfUserOK
const GetAvailableEditQuestionsOfUserOKCode int = 200

/*GetAvailableEditQuestionsOfUserOK questions found

swagger:response getAvailableEditQuestionsOfUserOK
*/
type GetAvailableEditQuestionsOfUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Question `json:"body,omitempty"`
}

// NewGetAvailableEditQuestionsOfUserOK creates GetAvailableEditQuestionsOfUserOK with default headers values
func NewGetAvailableEditQuestionsOfUserOK() *GetAvailableEditQuestionsOfUserOK {

	return &GetAvailableEditQuestionsOfUserOK{}
}

// WithPayload adds the payload to the get available edit questions of user o k response
func (o *GetAvailableEditQuestionsOfUserOK) WithPayload(payload []*models.Question) *GetAvailableEditQuestionsOfUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available edit questions of user o k response
func (o *GetAvailableEditQuestionsOfUserOK) SetPayload(payload []*models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableEditQuestionsOfUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetAvailableEditQuestionsOfUserBadRequestCode is the HTTP code returned for type GetAvailableEditQuestionsOfUserBadRequest
const GetAvailableEditQuestionsOfUserBadRequestCode int = 400

/*GetAvailableEditQuestionsOfUserBadRequest Incorrect Request, or invalida data

swagger:response getAvailableEditQuestionsOfUserBadRequest
*/
type GetAvailableEditQuestionsOfUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableEditQuestionsOfUserBadRequest creates GetAvailableEditQuestionsOfUserBadRequest with default headers values
func NewGetAvailableEditQuestionsOfUserBadRequest() *GetAvailableEditQuestionsOfUserBadRequest {

	return &GetAvailableEditQuestionsOfUserBadRequest{}
}

// WithPayload adds the payload to the get available edit questions of user bad request response
func (o *GetAvailableEditQuestionsOfUserBadRequest) WithPayload(payload *models.Error) *GetAvailableEditQuestionsOfUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available edit questions of user bad request response
func (o *GetAvailableEditQuestionsOfUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableEditQuestionsOfUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableEditQuestionsOfUserForbiddenCode is the HTTP code returned for type GetAvailableEditQuestionsOfUserForbidden
const GetAvailableEditQuestionsOfUserForbiddenCode int = 403

/*GetAvailableEditQuestionsOfUserForbidden Not authorized to this content

swagger:response getAvailableEditQuestionsOfUserForbidden
*/
type GetAvailableEditQuestionsOfUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableEditQuestionsOfUserForbidden creates GetAvailableEditQuestionsOfUserForbidden with default headers values
func NewGetAvailableEditQuestionsOfUserForbidden() *GetAvailableEditQuestionsOfUserForbidden {

	return &GetAvailableEditQuestionsOfUserForbidden{}
}

// WithPayload adds the payload to the get available edit questions of user forbidden response
func (o *GetAvailableEditQuestionsOfUserForbidden) WithPayload(payload *models.Error) *GetAvailableEditQuestionsOfUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available edit questions of user forbidden response
func (o *GetAvailableEditQuestionsOfUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableEditQuestionsOfUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableEditQuestionsOfUserGoneCode is the HTTP code returned for type GetAvailableEditQuestionsOfUserGone
const GetAvailableEditQuestionsOfUserGoneCode int = 410

/*GetAvailableEditQuestionsOfUserGone That resource does not exist

swagger:response getAvailableEditQuestionsOfUserGone
*/
type GetAvailableEditQuestionsOfUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableEditQuestionsOfUserGone creates GetAvailableEditQuestionsOfUserGone with default headers values
func NewGetAvailableEditQuestionsOfUserGone() *GetAvailableEditQuestionsOfUserGone {

	return &GetAvailableEditQuestionsOfUserGone{}
}

// WithPayload adds the payload to the get available edit questions of user gone response
func (o *GetAvailableEditQuestionsOfUserGone) WithPayload(payload *models.Error) *GetAvailableEditQuestionsOfUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available edit questions of user gone response
func (o *GetAvailableEditQuestionsOfUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableEditQuestionsOfUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableEditQuestionsOfUserInternalServerErrorCode is the HTTP code returned for type GetAvailableEditQuestionsOfUserInternalServerError
const GetAvailableEditQuestionsOfUserInternalServerErrorCode int = 500

/*GetAvailableEditQuestionsOfUserInternalServerError Internal error

swagger:response getAvailableEditQuestionsOfUserInternalServerError
*/
type GetAvailableEditQuestionsOfUserInternalServerError struct {
}

// NewGetAvailableEditQuestionsOfUserInternalServerError creates GetAvailableEditQuestionsOfUserInternalServerError with default headers values
func NewGetAvailableEditQuestionsOfUserInternalServerError() *GetAvailableEditQuestionsOfUserInternalServerError {

	return &GetAvailableEditQuestionsOfUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetAvailableEditQuestionsOfUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
