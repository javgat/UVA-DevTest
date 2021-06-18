// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublicOwnedPTestsFromUserOKCode is the HTTP code returned for type GetPublicOwnedPTestsFromUserOK
const GetPublicOwnedPTestsFromUserOKCode int = 200

/*GetPublicOwnedPTestsFromUserOK publishedTests found

swagger:response getPublicOwnedPTestsFromUserOK
*/
type GetPublicOwnedPTestsFromUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetPublicOwnedPTestsFromUserOK creates GetPublicOwnedPTestsFromUserOK with default headers values
func NewGetPublicOwnedPTestsFromUserOK() *GetPublicOwnedPTestsFromUserOK {

	return &GetPublicOwnedPTestsFromUserOK{}
}

// WithPayload adds the payload to the get public owned p tests from user o k response
func (o *GetPublicOwnedPTestsFromUserOK) WithPayload(payload []*models.Test) *GetPublicOwnedPTestsFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public owned p tests from user o k response
func (o *GetPublicOwnedPTestsFromUserOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicOwnedPTestsFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPublicOwnedPTestsFromUserBadRequestCode is the HTTP code returned for type GetPublicOwnedPTestsFromUserBadRequest
const GetPublicOwnedPTestsFromUserBadRequestCode int = 400

/*GetPublicOwnedPTestsFromUserBadRequest Incorrect Request, or invalida data

swagger:response getPublicOwnedPTestsFromUserBadRequest
*/
type GetPublicOwnedPTestsFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicOwnedPTestsFromUserBadRequest creates GetPublicOwnedPTestsFromUserBadRequest with default headers values
func NewGetPublicOwnedPTestsFromUserBadRequest() *GetPublicOwnedPTestsFromUserBadRequest {

	return &GetPublicOwnedPTestsFromUserBadRequest{}
}

// WithPayload adds the payload to the get public owned p tests from user bad request response
func (o *GetPublicOwnedPTestsFromUserBadRequest) WithPayload(payload *models.Error) *GetPublicOwnedPTestsFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public owned p tests from user bad request response
func (o *GetPublicOwnedPTestsFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicOwnedPTestsFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicOwnedPTestsFromUserForbiddenCode is the HTTP code returned for type GetPublicOwnedPTestsFromUserForbidden
const GetPublicOwnedPTestsFromUserForbiddenCode int = 403

/*GetPublicOwnedPTestsFromUserForbidden Not authorized to this content

swagger:response getPublicOwnedPTestsFromUserForbidden
*/
type GetPublicOwnedPTestsFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicOwnedPTestsFromUserForbidden creates GetPublicOwnedPTestsFromUserForbidden with default headers values
func NewGetPublicOwnedPTestsFromUserForbidden() *GetPublicOwnedPTestsFromUserForbidden {

	return &GetPublicOwnedPTestsFromUserForbidden{}
}

// WithPayload adds the payload to the get public owned p tests from user forbidden response
func (o *GetPublicOwnedPTestsFromUserForbidden) WithPayload(payload *models.Error) *GetPublicOwnedPTestsFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public owned p tests from user forbidden response
func (o *GetPublicOwnedPTestsFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicOwnedPTestsFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicOwnedPTestsFromUserGoneCode is the HTTP code returned for type GetPublicOwnedPTestsFromUserGone
const GetPublicOwnedPTestsFromUserGoneCode int = 410

/*GetPublicOwnedPTestsFromUserGone That resource does not exist

swagger:response getPublicOwnedPTestsFromUserGone
*/
type GetPublicOwnedPTestsFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicOwnedPTestsFromUserGone creates GetPublicOwnedPTestsFromUserGone with default headers values
func NewGetPublicOwnedPTestsFromUserGone() *GetPublicOwnedPTestsFromUserGone {

	return &GetPublicOwnedPTestsFromUserGone{}
}

// WithPayload adds the payload to the get public owned p tests from user gone response
func (o *GetPublicOwnedPTestsFromUserGone) WithPayload(payload *models.Error) *GetPublicOwnedPTestsFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public owned p tests from user gone response
func (o *GetPublicOwnedPTestsFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicOwnedPTestsFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicOwnedPTestsFromUserInternalServerErrorCode is the HTTP code returned for type GetPublicOwnedPTestsFromUserInternalServerError
const GetPublicOwnedPTestsFromUserInternalServerErrorCode int = 500

/*GetPublicOwnedPTestsFromUserInternalServerError Internal error

swagger:response getPublicOwnedPTestsFromUserInternalServerError
*/
type GetPublicOwnedPTestsFromUserInternalServerError struct {
}

// NewGetPublicOwnedPTestsFromUserInternalServerError creates GetPublicOwnedPTestsFromUserInternalServerError with default headers values
func NewGetPublicOwnedPTestsFromUserInternalServerError() *GetPublicOwnedPTestsFromUserInternalServerError {

	return &GetPublicOwnedPTestsFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublicOwnedPTestsFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
