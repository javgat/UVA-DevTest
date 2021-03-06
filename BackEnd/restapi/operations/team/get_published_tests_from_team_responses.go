// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublishedTestsFromTeamOKCode is the HTTP code returned for type GetPublishedTestsFromTeamOK
const GetPublishedTestsFromTeamOKCode int = 200

/*GetPublishedTestsFromTeamOK publishedTests found

swagger:response getPublishedTestsFromTeamOK
*/
type GetPublishedTestsFromTeamOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTeamOK creates GetPublishedTestsFromTeamOK with default headers values
func NewGetPublishedTestsFromTeamOK() *GetPublishedTestsFromTeamOK {

	return &GetPublishedTestsFromTeamOK{}
}

// WithPayload adds the payload to the get published tests from team o k response
func (o *GetPublishedTestsFromTeamOK) WithPayload(payload []*models.Test) *GetPublishedTestsFromTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from team o k response
func (o *GetPublishedTestsFromTeamOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Test, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPublishedTestsFromTeamBadRequestCode is the HTTP code returned for type GetPublishedTestsFromTeamBadRequest
const GetPublishedTestsFromTeamBadRequestCode int = 400

/*GetPublishedTestsFromTeamBadRequest Incorrect Request, or invalida data

swagger:response getPublishedTestsFromTeamBadRequest
*/
type GetPublishedTestsFromTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTeamBadRequest creates GetPublishedTestsFromTeamBadRequest with default headers values
func NewGetPublishedTestsFromTeamBadRequest() *GetPublishedTestsFromTeamBadRequest {

	return &GetPublishedTestsFromTeamBadRequest{}
}

// WithPayload adds the payload to the get published tests from team bad request response
func (o *GetPublishedTestsFromTeamBadRequest) WithPayload(payload *models.Error) *GetPublishedTestsFromTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from team bad request response
func (o *GetPublishedTestsFromTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsFromTeamForbiddenCode is the HTTP code returned for type GetPublishedTestsFromTeamForbidden
const GetPublishedTestsFromTeamForbiddenCode int = 403

/*GetPublishedTestsFromTeamForbidden Not authorized to this content

swagger:response getPublishedTestsFromTeamForbidden
*/
type GetPublishedTestsFromTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTeamForbidden creates GetPublishedTestsFromTeamForbidden with default headers values
func NewGetPublishedTestsFromTeamForbidden() *GetPublishedTestsFromTeamForbidden {

	return &GetPublishedTestsFromTeamForbidden{}
}

// WithPayload adds the payload to the get published tests from team forbidden response
func (o *GetPublishedTestsFromTeamForbidden) WithPayload(payload *models.Error) *GetPublishedTestsFromTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from team forbidden response
func (o *GetPublishedTestsFromTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsFromTeamGoneCode is the HTTP code returned for type GetPublishedTestsFromTeamGone
const GetPublishedTestsFromTeamGoneCode int = 410

/*GetPublishedTestsFromTeamGone That user (password and name) does not exist

swagger:response getPublishedTestsFromTeamGone
*/
type GetPublishedTestsFromTeamGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublishedTestsFromTeamGone creates GetPublishedTestsFromTeamGone with default headers values
func NewGetPublishedTestsFromTeamGone() *GetPublishedTestsFromTeamGone {

	return &GetPublishedTestsFromTeamGone{}
}

// WithPayload adds the payload to the get published tests from team gone response
func (o *GetPublishedTestsFromTeamGone) WithPayload(payload *models.Error) *GetPublishedTestsFromTeamGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get published tests from team gone response
func (o *GetPublishedTestsFromTeamGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTeamGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublishedTestsFromTeamInternalServerErrorCode is the HTTP code returned for type GetPublishedTestsFromTeamInternalServerError
const GetPublishedTestsFromTeamInternalServerErrorCode int = 500

/*GetPublishedTestsFromTeamInternalServerError Internal error

swagger:response getPublishedTestsFromTeamInternalServerError
*/
type GetPublishedTestsFromTeamInternalServerError struct {
}

// NewGetPublishedTestsFromTeamInternalServerError creates GetPublishedTestsFromTeamInternalServerError with default headers values
func NewGetPublishedTestsFromTeamInternalServerError() *GetPublishedTestsFromTeamInternalServerError {

	return &GetPublishedTestsFromTeamInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublishedTestsFromTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
