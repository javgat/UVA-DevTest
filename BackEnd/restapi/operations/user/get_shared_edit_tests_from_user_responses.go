// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetSharedEditTestsFromUserOKCode is the HTTP code returned for type GetSharedEditTestsFromUserOK
const GetSharedEditTestsFromUserOKCode int = 200

/*GetSharedEditTestsFromUserOK tests found

swagger:response getSharedEditTestsFromUserOK
*/
type GetSharedEditTestsFromUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetSharedEditTestsFromUserOK creates GetSharedEditTestsFromUserOK with default headers values
func NewGetSharedEditTestsFromUserOK() *GetSharedEditTestsFromUserOK {

	return &GetSharedEditTestsFromUserOK{}
}

// WithPayload adds the payload to the get shared edit tests from user o k response
func (o *GetSharedEditTestsFromUserOK) WithPayload(payload []*models.Test) *GetSharedEditTestsFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared edit tests from user o k response
func (o *GetSharedEditTestsFromUserOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedEditTestsFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetSharedEditTestsFromUserBadRequestCode is the HTTP code returned for type GetSharedEditTestsFromUserBadRequest
const GetSharedEditTestsFromUserBadRequestCode int = 400

/*GetSharedEditTestsFromUserBadRequest Incorrect Request, or invalida data

swagger:response getSharedEditTestsFromUserBadRequest
*/
type GetSharedEditTestsFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSharedEditTestsFromUserBadRequest creates GetSharedEditTestsFromUserBadRequest with default headers values
func NewGetSharedEditTestsFromUserBadRequest() *GetSharedEditTestsFromUserBadRequest {

	return &GetSharedEditTestsFromUserBadRequest{}
}

// WithPayload adds the payload to the get shared edit tests from user bad request response
func (o *GetSharedEditTestsFromUserBadRequest) WithPayload(payload *models.Error) *GetSharedEditTestsFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared edit tests from user bad request response
func (o *GetSharedEditTestsFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedEditTestsFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedEditTestsFromUserForbiddenCode is the HTTP code returned for type GetSharedEditTestsFromUserForbidden
const GetSharedEditTestsFromUserForbiddenCode int = 403

/*GetSharedEditTestsFromUserForbidden Not authorized to this content

swagger:response getSharedEditTestsFromUserForbidden
*/
type GetSharedEditTestsFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSharedEditTestsFromUserForbidden creates GetSharedEditTestsFromUserForbidden with default headers values
func NewGetSharedEditTestsFromUserForbidden() *GetSharedEditTestsFromUserForbidden {

	return &GetSharedEditTestsFromUserForbidden{}
}

// WithPayload adds the payload to the get shared edit tests from user forbidden response
func (o *GetSharedEditTestsFromUserForbidden) WithPayload(payload *models.Error) *GetSharedEditTestsFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared edit tests from user forbidden response
func (o *GetSharedEditTestsFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedEditTestsFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedEditTestsFromUserGoneCode is the HTTP code returned for type GetSharedEditTestsFromUserGone
const GetSharedEditTestsFromUserGoneCode int = 410

/*GetSharedEditTestsFromUserGone That resource does not exist

swagger:response getSharedEditTestsFromUserGone
*/
type GetSharedEditTestsFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSharedEditTestsFromUserGone creates GetSharedEditTestsFromUserGone with default headers values
func NewGetSharedEditTestsFromUserGone() *GetSharedEditTestsFromUserGone {

	return &GetSharedEditTestsFromUserGone{}
}

// WithPayload adds the payload to the get shared edit tests from user gone response
func (o *GetSharedEditTestsFromUserGone) WithPayload(payload *models.Error) *GetSharedEditTestsFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get shared edit tests from user gone response
func (o *GetSharedEditTestsFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSharedEditTestsFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSharedEditTestsFromUserInternalServerErrorCode is the HTTP code returned for type GetSharedEditTestsFromUserInternalServerError
const GetSharedEditTestsFromUserInternalServerErrorCode int = 500

/*GetSharedEditTestsFromUserInternalServerError Internal error

swagger:response getSharedEditTestsFromUserInternalServerError
*/
type GetSharedEditTestsFromUserInternalServerError struct {
}

// NewGetSharedEditTestsFromUserInternalServerError creates GetSharedEditTestsFromUserInternalServerError with default headers values
func NewGetSharedEditTestsFromUserInternalServerError() *GetSharedEditTestsFromUserInternalServerError {

	return &GetSharedEditTestsFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetSharedEditTestsFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
