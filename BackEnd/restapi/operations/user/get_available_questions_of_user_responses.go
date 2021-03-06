// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetAvailableQuestionsOfUserOKCode is the HTTP code returned for type GetAvailableQuestionsOfUserOK
const GetAvailableQuestionsOfUserOKCode int = 200

/*GetAvailableQuestionsOfUserOK questions found

swagger:response getAvailableQuestionsOfUserOK
*/
type GetAvailableQuestionsOfUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Question `json:"body,omitempty"`
}

// NewGetAvailableQuestionsOfUserOK creates GetAvailableQuestionsOfUserOK with default headers values
func NewGetAvailableQuestionsOfUserOK() *GetAvailableQuestionsOfUserOK {

	return &GetAvailableQuestionsOfUserOK{}
}

// WithPayload adds the payload to the get available questions of user o k response
func (o *GetAvailableQuestionsOfUserOK) WithPayload(payload []*models.Question) *GetAvailableQuestionsOfUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available questions of user o k response
func (o *GetAvailableQuestionsOfUserOK) SetPayload(payload []*models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableQuestionsOfUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetAvailableQuestionsOfUserBadRequestCode is the HTTP code returned for type GetAvailableQuestionsOfUserBadRequest
const GetAvailableQuestionsOfUserBadRequestCode int = 400

/*GetAvailableQuestionsOfUserBadRequest Incorrect Request, or invalida data

swagger:response getAvailableQuestionsOfUserBadRequest
*/
type GetAvailableQuestionsOfUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableQuestionsOfUserBadRequest creates GetAvailableQuestionsOfUserBadRequest with default headers values
func NewGetAvailableQuestionsOfUserBadRequest() *GetAvailableQuestionsOfUserBadRequest {

	return &GetAvailableQuestionsOfUserBadRequest{}
}

// WithPayload adds the payload to the get available questions of user bad request response
func (o *GetAvailableQuestionsOfUserBadRequest) WithPayload(payload *models.Error) *GetAvailableQuestionsOfUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available questions of user bad request response
func (o *GetAvailableQuestionsOfUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableQuestionsOfUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableQuestionsOfUserForbiddenCode is the HTTP code returned for type GetAvailableQuestionsOfUserForbidden
const GetAvailableQuestionsOfUserForbiddenCode int = 403

/*GetAvailableQuestionsOfUserForbidden Not authorized to this content

swagger:response getAvailableQuestionsOfUserForbidden
*/
type GetAvailableQuestionsOfUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableQuestionsOfUserForbidden creates GetAvailableQuestionsOfUserForbidden with default headers values
func NewGetAvailableQuestionsOfUserForbidden() *GetAvailableQuestionsOfUserForbidden {

	return &GetAvailableQuestionsOfUserForbidden{}
}

// WithPayload adds the payload to the get available questions of user forbidden response
func (o *GetAvailableQuestionsOfUserForbidden) WithPayload(payload *models.Error) *GetAvailableQuestionsOfUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available questions of user forbidden response
func (o *GetAvailableQuestionsOfUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableQuestionsOfUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableQuestionsOfUserGoneCode is the HTTP code returned for type GetAvailableQuestionsOfUserGone
const GetAvailableQuestionsOfUserGoneCode int = 410

/*GetAvailableQuestionsOfUserGone That resource does not exist

swagger:response getAvailableQuestionsOfUserGone
*/
type GetAvailableQuestionsOfUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableQuestionsOfUserGone creates GetAvailableQuestionsOfUserGone with default headers values
func NewGetAvailableQuestionsOfUserGone() *GetAvailableQuestionsOfUserGone {

	return &GetAvailableQuestionsOfUserGone{}
}

// WithPayload adds the payload to the get available questions of user gone response
func (o *GetAvailableQuestionsOfUserGone) WithPayload(payload *models.Error) *GetAvailableQuestionsOfUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available questions of user gone response
func (o *GetAvailableQuestionsOfUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableQuestionsOfUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableQuestionsOfUserInternalServerErrorCode is the HTTP code returned for type GetAvailableQuestionsOfUserInternalServerError
const GetAvailableQuestionsOfUserInternalServerErrorCode int = 500

/*GetAvailableQuestionsOfUserInternalServerError Internal error

swagger:response getAvailableQuestionsOfUserInternalServerError
*/
type GetAvailableQuestionsOfUserInternalServerError struct {
}

// NewGetAvailableQuestionsOfUserInternalServerError creates GetAvailableQuestionsOfUserInternalServerError with default headers values
func NewGetAvailableQuestionsOfUserInternalServerError() *GetAvailableQuestionsOfUserInternalServerError {

	return &GetAvailableQuestionsOfUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetAvailableQuestionsOfUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
