// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetUnfinishedAnswersFromPublishedTestsOKCode is the HTTP code returned for type GetUnfinishedAnswersFromPublishedTestsOK
const GetUnfinishedAnswersFromPublishedTestsOKCode int = 200

/*GetUnfinishedAnswersFromPublishedTestsOK Answers found

swagger:response getUnfinishedAnswersFromPublishedTestsOK
*/
type GetUnfinishedAnswersFromPublishedTestsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Answer `json:"body,omitempty"`
}

// NewGetUnfinishedAnswersFromPublishedTestsOK creates GetUnfinishedAnswersFromPublishedTestsOK with default headers values
func NewGetUnfinishedAnswersFromPublishedTestsOK() *GetUnfinishedAnswersFromPublishedTestsOK {

	return &GetUnfinishedAnswersFromPublishedTestsOK{}
}

// WithPayload adds the payload to the get unfinished answers from published tests o k response
func (o *GetUnfinishedAnswersFromPublishedTestsOK) WithPayload(payload []*models.Answer) *GetUnfinishedAnswersFromPublishedTestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unfinished answers from published tests o k response
func (o *GetUnfinishedAnswersFromPublishedTestsOK) SetPayload(payload []*models.Answer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnfinishedAnswersFromPublishedTestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUnfinishedAnswersFromPublishedTestsBadRequestCode is the HTTP code returned for type GetUnfinishedAnswersFromPublishedTestsBadRequest
const GetUnfinishedAnswersFromPublishedTestsBadRequestCode int = 400

/*GetUnfinishedAnswersFromPublishedTestsBadRequest Incorrect Request, or invalida data

swagger:response getUnfinishedAnswersFromPublishedTestsBadRequest
*/
type GetUnfinishedAnswersFromPublishedTestsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUnfinishedAnswersFromPublishedTestsBadRequest creates GetUnfinishedAnswersFromPublishedTestsBadRequest with default headers values
func NewGetUnfinishedAnswersFromPublishedTestsBadRequest() *GetUnfinishedAnswersFromPublishedTestsBadRequest {

	return &GetUnfinishedAnswersFromPublishedTestsBadRequest{}
}

// WithPayload adds the payload to the get unfinished answers from published tests bad request response
func (o *GetUnfinishedAnswersFromPublishedTestsBadRequest) WithPayload(payload *models.Error) *GetUnfinishedAnswersFromPublishedTestsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unfinished answers from published tests bad request response
func (o *GetUnfinishedAnswersFromPublishedTestsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnfinishedAnswersFromPublishedTestsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUnfinishedAnswersFromPublishedTestsForbiddenCode is the HTTP code returned for type GetUnfinishedAnswersFromPublishedTestsForbidden
const GetUnfinishedAnswersFromPublishedTestsForbiddenCode int = 403

/*GetUnfinishedAnswersFromPublishedTestsForbidden Not authorized to this content

swagger:response getUnfinishedAnswersFromPublishedTestsForbidden
*/
type GetUnfinishedAnswersFromPublishedTestsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUnfinishedAnswersFromPublishedTestsForbidden creates GetUnfinishedAnswersFromPublishedTestsForbidden with default headers values
func NewGetUnfinishedAnswersFromPublishedTestsForbidden() *GetUnfinishedAnswersFromPublishedTestsForbidden {

	return &GetUnfinishedAnswersFromPublishedTestsForbidden{}
}

// WithPayload adds the payload to the get unfinished answers from published tests forbidden response
func (o *GetUnfinishedAnswersFromPublishedTestsForbidden) WithPayload(payload *models.Error) *GetUnfinishedAnswersFromPublishedTestsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unfinished answers from published tests forbidden response
func (o *GetUnfinishedAnswersFromPublishedTestsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnfinishedAnswersFromPublishedTestsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUnfinishedAnswersFromPublishedTestsGoneCode is the HTTP code returned for type GetUnfinishedAnswersFromPublishedTestsGone
const GetUnfinishedAnswersFromPublishedTestsGoneCode int = 410

/*GetUnfinishedAnswersFromPublishedTestsGone That resource does not exist

swagger:response getUnfinishedAnswersFromPublishedTestsGone
*/
type GetUnfinishedAnswersFromPublishedTestsGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUnfinishedAnswersFromPublishedTestsGone creates GetUnfinishedAnswersFromPublishedTestsGone with default headers values
func NewGetUnfinishedAnswersFromPublishedTestsGone() *GetUnfinishedAnswersFromPublishedTestsGone {

	return &GetUnfinishedAnswersFromPublishedTestsGone{}
}

// WithPayload adds the payload to the get unfinished answers from published tests gone response
func (o *GetUnfinishedAnswersFromPublishedTestsGone) WithPayload(payload *models.Error) *GetUnfinishedAnswersFromPublishedTestsGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unfinished answers from published tests gone response
func (o *GetUnfinishedAnswersFromPublishedTestsGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnfinishedAnswersFromPublishedTestsGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUnfinishedAnswersFromPublishedTestsInternalServerErrorCode is the HTTP code returned for type GetUnfinishedAnswersFromPublishedTestsInternalServerError
const GetUnfinishedAnswersFromPublishedTestsInternalServerErrorCode int = 500

/*GetUnfinishedAnswersFromPublishedTestsInternalServerError Internal error

swagger:response getUnfinishedAnswersFromPublishedTestsInternalServerError
*/
type GetUnfinishedAnswersFromPublishedTestsInternalServerError struct {
}

// NewGetUnfinishedAnswersFromPublishedTestsInternalServerError creates GetUnfinishedAnswersFromPublishedTestsInternalServerError with default headers values
func NewGetUnfinishedAnswersFromPublishedTestsInternalServerError() *GetUnfinishedAnswersFromPublishedTestsInternalServerError {

	return &GetUnfinishedAnswersFromPublishedTestsInternalServerError{}
}

// WriteResponse to the client
func (o *GetUnfinishedAnswersFromPublishedTestsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
