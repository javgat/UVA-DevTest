// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublicEditTestsFromUserOKCode is the HTTP code returned for type GetPublicEditTestsFromUserOK
const GetPublicEditTestsFromUserOKCode int = 200

/*GetPublicEditTestsFromUserOK tests found

swagger:response getPublicEditTestsFromUserOK
*/
type GetPublicEditTestsFromUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetPublicEditTestsFromUserOK creates GetPublicEditTestsFromUserOK with default headers values
func NewGetPublicEditTestsFromUserOK() *GetPublicEditTestsFromUserOK {

	return &GetPublicEditTestsFromUserOK{}
}

// WithPayload adds the payload to the get public edit tests from user o k response
func (o *GetPublicEditTestsFromUserOK) WithPayload(payload []*models.Test) *GetPublicEditTestsFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public edit tests from user o k response
func (o *GetPublicEditTestsFromUserOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicEditTestsFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPublicEditTestsFromUserBadRequestCode is the HTTP code returned for type GetPublicEditTestsFromUserBadRequest
const GetPublicEditTestsFromUserBadRequestCode int = 400

/*GetPublicEditTestsFromUserBadRequest Incorrect Request, or invalida data

swagger:response getPublicEditTestsFromUserBadRequest
*/
type GetPublicEditTestsFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicEditTestsFromUserBadRequest creates GetPublicEditTestsFromUserBadRequest with default headers values
func NewGetPublicEditTestsFromUserBadRequest() *GetPublicEditTestsFromUserBadRequest {

	return &GetPublicEditTestsFromUserBadRequest{}
}

// WithPayload adds the payload to the get public edit tests from user bad request response
func (o *GetPublicEditTestsFromUserBadRequest) WithPayload(payload *models.Error) *GetPublicEditTestsFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public edit tests from user bad request response
func (o *GetPublicEditTestsFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicEditTestsFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicEditTestsFromUserForbiddenCode is the HTTP code returned for type GetPublicEditTestsFromUserForbidden
const GetPublicEditTestsFromUserForbiddenCode int = 403

/*GetPublicEditTestsFromUserForbidden Not authorized to this content

swagger:response getPublicEditTestsFromUserForbidden
*/
type GetPublicEditTestsFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicEditTestsFromUserForbidden creates GetPublicEditTestsFromUserForbidden with default headers values
func NewGetPublicEditTestsFromUserForbidden() *GetPublicEditTestsFromUserForbidden {

	return &GetPublicEditTestsFromUserForbidden{}
}

// WithPayload adds the payload to the get public edit tests from user forbidden response
func (o *GetPublicEditTestsFromUserForbidden) WithPayload(payload *models.Error) *GetPublicEditTestsFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public edit tests from user forbidden response
func (o *GetPublicEditTestsFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicEditTestsFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicEditTestsFromUserGoneCode is the HTTP code returned for type GetPublicEditTestsFromUserGone
const GetPublicEditTestsFromUserGoneCode int = 410

/*GetPublicEditTestsFromUserGone That resource does not exist

swagger:response getPublicEditTestsFromUserGone
*/
type GetPublicEditTestsFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicEditTestsFromUserGone creates GetPublicEditTestsFromUserGone with default headers values
func NewGetPublicEditTestsFromUserGone() *GetPublicEditTestsFromUserGone {

	return &GetPublicEditTestsFromUserGone{}
}

// WithPayload adds the payload to the get public edit tests from user gone response
func (o *GetPublicEditTestsFromUserGone) WithPayload(payload *models.Error) *GetPublicEditTestsFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public edit tests from user gone response
func (o *GetPublicEditTestsFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicEditTestsFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicEditTestsFromUserInternalServerErrorCode is the HTTP code returned for type GetPublicEditTestsFromUserInternalServerError
const GetPublicEditTestsFromUserInternalServerErrorCode int = 500

/*GetPublicEditTestsFromUserInternalServerError Internal error

swagger:response getPublicEditTestsFromUserInternalServerError
*/
type GetPublicEditTestsFromUserInternalServerError struct {
}

// NewGetPublicEditTestsFromUserInternalServerError creates GetPublicEditTestsFromUserInternalServerError with default headers values
func NewGetPublicEditTestsFromUserInternalServerError() *GetPublicEditTestsFromUserInternalServerError {

	return &GetPublicEditTestsFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublicEditTestsFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}