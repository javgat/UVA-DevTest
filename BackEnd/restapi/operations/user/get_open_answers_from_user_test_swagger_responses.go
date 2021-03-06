// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetOpenAnswersFromUserTestOKCode is the HTTP code returned for type GetOpenAnswersFromUserTestOK
const GetOpenAnswersFromUserTestOKCode int = 200

/*GetOpenAnswersFromUserTestOK answers found

swagger:response getOpenAnswersFromUserTestOK
*/
type GetOpenAnswersFromUserTestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Answer `json:"body,omitempty"`
}

// NewGetOpenAnswersFromUserTestOK creates GetOpenAnswersFromUserTestOK with default headers values
func NewGetOpenAnswersFromUserTestOK() *GetOpenAnswersFromUserTestOK {

	return &GetOpenAnswersFromUserTestOK{}
}

// WithPayload adds the payload to the get open answers from user test o k response
func (o *GetOpenAnswersFromUserTestOK) WithPayload(payload []*models.Answer) *GetOpenAnswersFromUserTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get open answers from user test o k response
func (o *GetOpenAnswersFromUserTestOK) SetPayload(payload []*models.Answer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOpenAnswersFromUserTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Answer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetOpenAnswersFromUserTestBadRequestCode is the HTTP code returned for type GetOpenAnswersFromUserTestBadRequest
const GetOpenAnswersFromUserTestBadRequestCode int = 400

/*GetOpenAnswersFromUserTestBadRequest Incorrect Request, or invalida data

swagger:response getOpenAnswersFromUserTestBadRequest
*/
type GetOpenAnswersFromUserTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOpenAnswersFromUserTestBadRequest creates GetOpenAnswersFromUserTestBadRequest with default headers values
func NewGetOpenAnswersFromUserTestBadRequest() *GetOpenAnswersFromUserTestBadRequest {

	return &GetOpenAnswersFromUserTestBadRequest{}
}

// WithPayload adds the payload to the get open answers from user test bad request response
func (o *GetOpenAnswersFromUserTestBadRequest) WithPayload(payload *models.Error) *GetOpenAnswersFromUserTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get open answers from user test bad request response
func (o *GetOpenAnswersFromUserTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOpenAnswersFromUserTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOpenAnswersFromUserTestForbiddenCode is the HTTP code returned for type GetOpenAnswersFromUserTestForbidden
const GetOpenAnswersFromUserTestForbiddenCode int = 403

/*GetOpenAnswersFromUserTestForbidden Not authorized to this content

swagger:response getOpenAnswersFromUserTestForbidden
*/
type GetOpenAnswersFromUserTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOpenAnswersFromUserTestForbidden creates GetOpenAnswersFromUserTestForbidden with default headers values
func NewGetOpenAnswersFromUserTestForbidden() *GetOpenAnswersFromUserTestForbidden {

	return &GetOpenAnswersFromUserTestForbidden{}
}

// WithPayload adds the payload to the get open answers from user test forbidden response
func (o *GetOpenAnswersFromUserTestForbidden) WithPayload(payload *models.Error) *GetOpenAnswersFromUserTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get open answers from user test forbidden response
func (o *GetOpenAnswersFromUserTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOpenAnswersFromUserTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOpenAnswersFromUserTestGoneCode is the HTTP code returned for type GetOpenAnswersFromUserTestGone
const GetOpenAnswersFromUserTestGoneCode int = 410

/*GetOpenAnswersFromUserTestGone That resource does not exist

swagger:response getOpenAnswersFromUserTestGone
*/
type GetOpenAnswersFromUserTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOpenAnswersFromUserTestGone creates GetOpenAnswersFromUserTestGone with default headers values
func NewGetOpenAnswersFromUserTestGone() *GetOpenAnswersFromUserTestGone {

	return &GetOpenAnswersFromUserTestGone{}
}

// WithPayload adds the payload to the get open answers from user test gone response
func (o *GetOpenAnswersFromUserTestGone) WithPayload(payload *models.Error) *GetOpenAnswersFromUserTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get open answers from user test gone response
func (o *GetOpenAnswersFromUserTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOpenAnswersFromUserTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOpenAnswersFromUserTestInternalServerErrorCode is the HTTP code returned for type GetOpenAnswersFromUserTestInternalServerError
const GetOpenAnswersFromUserTestInternalServerErrorCode int = 500

/*GetOpenAnswersFromUserTestInternalServerError Internal error

swagger:response getOpenAnswersFromUserTestInternalServerError
*/
type GetOpenAnswersFromUserTestInternalServerError struct {
}

// NewGetOpenAnswersFromUserTestInternalServerError creates GetOpenAnswersFromUserTestInternalServerError with default headers values
func NewGetOpenAnswersFromUserTestInternalServerError() *GetOpenAnswersFromUserTestInternalServerError {

	return &GetOpenAnswersFromUserTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetOpenAnswersFromUserTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
