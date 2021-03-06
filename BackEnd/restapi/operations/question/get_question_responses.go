// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetQuestionOKCode is the HTTP code returned for type GetQuestionOK
const GetQuestionOKCode int = 200

/*GetQuestionOK Question found

swagger:response getQuestionOK
*/
type GetQuestionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Question `json:"body,omitempty"`
}

// NewGetQuestionOK creates GetQuestionOK with default headers values
func NewGetQuestionOK() *GetQuestionOK {

	return &GetQuestionOK{}
}

// WithPayload adds the payload to the get question o k response
func (o *GetQuestionOK) WithPayload(payload *models.Question) *GetQuestionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question o k response
func (o *GetQuestionOK) SetPayload(payload *models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionBadRequestCode is the HTTP code returned for type GetQuestionBadRequest
const GetQuestionBadRequestCode int = 400

/*GetQuestionBadRequest Incorrect Request, or invalida data

swagger:response getQuestionBadRequest
*/
type GetQuestionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionBadRequest creates GetQuestionBadRequest with default headers values
func NewGetQuestionBadRequest() *GetQuestionBadRequest {

	return &GetQuestionBadRequest{}
}

// WithPayload adds the payload to the get question bad request response
func (o *GetQuestionBadRequest) WithPayload(payload *models.Error) *GetQuestionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question bad request response
func (o *GetQuestionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionForbiddenCode is the HTTP code returned for type GetQuestionForbidden
const GetQuestionForbiddenCode int = 403

/*GetQuestionForbidden Not authorized to this content

swagger:response getQuestionForbidden
*/
type GetQuestionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionForbidden creates GetQuestionForbidden with default headers values
func NewGetQuestionForbidden() *GetQuestionForbidden {

	return &GetQuestionForbidden{}
}

// WithPayload adds the payload to the get question forbidden response
func (o *GetQuestionForbidden) WithPayload(payload *models.Error) *GetQuestionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question forbidden response
func (o *GetQuestionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionGoneCode is the HTTP code returned for type GetQuestionGone
const GetQuestionGoneCode int = 410

/*GetQuestionGone That resource does not exist

swagger:response getQuestionGone
*/
type GetQuestionGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionGone creates GetQuestionGone with default headers values
func NewGetQuestionGone() *GetQuestionGone {

	return &GetQuestionGone{}
}

// WithPayload adds the payload to the get question gone response
func (o *GetQuestionGone) WithPayload(payload *models.Error) *GetQuestionGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question gone response
func (o *GetQuestionGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionInternalServerErrorCode is the HTTP code returned for type GetQuestionInternalServerError
const GetQuestionInternalServerErrorCode int = 500

/*GetQuestionInternalServerError Internal error

swagger:response getQuestionInternalServerError
*/
type GetQuestionInternalServerError struct {
}

// NewGetQuestionInternalServerError creates GetQuestionInternalServerError with default headers values
func NewGetQuestionInternalServerError() *GetQuestionInternalServerError {

	return &GetQuestionInternalServerError{}
}

// WriteResponse to the client
func (o *GetQuestionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
