// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetUsersFromPublishedTestOKCode is the HTTP code returned for type GetUsersFromPublishedTestOK
const GetUsersFromPublishedTestOKCode int = 200

/*GetUsersFromPublishedTestOK users found

swagger:response getUsersFromPublishedTestOK
*/
type GetUsersFromPublishedTestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewGetUsersFromPublishedTestOK creates GetUsersFromPublishedTestOK with default headers values
func NewGetUsersFromPublishedTestOK() *GetUsersFromPublishedTestOK {

	return &GetUsersFromPublishedTestOK{}
}

// WithPayload adds the payload to the get users from published test o k response
func (o *GetUsersFromPublishedTestOK) WithPayload(payload []*models.User) *GetUsersFromPublishedTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users from published test o k response
func (o *GetUsersFromPublishedTestOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersFromPublishedTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetUsersFromPublishedTestBadRequestCode is the HTTP code returned for type GetUsersFromPublishedTestBadRequest
const GetUsersFromPublishedTestBadRequestCode int = 400

/*GetUsersFromPublishedTestBadRequest Incorrect Request, or invalida data

swagger:response getUsersFromPublishedTestBadRequest
*/
type GetUsersFromPublishedTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersFromPublishedTestBadRequest creates GetUsersFromPublishedTestBadRequest with default headers values
func NewGetUsersFromPublishedTestBadRequest() *GetUsersFromPublishedTestBadRequest {

	return &GetUsersFromPublishedTestBadRequest{}
}

// WithPayload adds the payload to the get users from published test bad request response
func (o *GetUsersFromPublishedTestBadRequest) WithPayload(payload *models.Error) *GetUsersFromPublishedTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users from published test bad request response
func (o *GetUsersFromPublishedTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersFromPublishedTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersFromPublishedTestForbiddenCode is the HTTP code returned for type GetUsersFromPublishedTestForbidden
const GetUsersFromPublishedTestForbiddenCode int = 403

/*GetUsersFromPublishedTestForbidden Not authorized to this content

swagger:response getUsersFromPublishedTestForbidden
*/
type GetUsersFromPublishedTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersFromPublishedTestForbidden creates GetUsersFromPublishedTestForbidden with default headers values
func NewGetUsersFromPublishedTestForbidden() *GetUsersFromPublishedTestForbidden {

	return &GetUsersFromPublishedTestForbidden{}
}

// WithPayload adds the payload to the get users from published test forbidden response
func (o *GetUsersFromPublishedTestForbidden) WithPayload(payload *models.Error) *GetUsersFromPublishedTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users from published test forbidden response
func (o *GetUsersFromPublishedTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersFromPublishedTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersFromPublishedTestGoneCode is the HTTP code returned for type GetUsersFromPublishedTestGone
const GetUsersFromPublishedTestGoneCode int = 410

/*GetUsersFromPublishedTestGone That user (password and name) does not exist

swagger:response getUsersFromPublishedTestGone
*/
type GetUsersFromPublishedTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUsersFromPublishedTestGone creates GetUsersFromPublishedTestGone with default headers values
func NewGetUsersFromPublishedTestGone() *GetUsersFromPublishedTestGone {

	return &GetUsersFromPublishedTestGone{}
}

// WithPayload adds the payload to the get users from published test gone response
func (o *GetUsersFromPublishedTestGone) WithPayload(payload *models.Error) *GetUsersFromPublishedTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users from published test gone response
func (o *GetUsersFromPublishedTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersFromPublishedTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersFromPublishedTestInternalServerErrorCode is the HTTP code returned for type GetUsersFromPublishedTestInternalServerError
const GetUsersFromPublishedTestInternalServerErrorCode int = 500

/*GetUsersFromPublishedTestInternalServerError Internal error

swagger:response getUsersFromPublishedTestInternalServerError
*/
type GetUsersFromPublishedTestInternalServerError struct {
}

// NewGetUsersFromPublishedTestInternalServerError creates GetUsersFromPublishedTestInternalServerError with default headers values
func NewGetUsersFromPublishedTestInternalServerError() *GetUsersFromPublishedTestInternalServerError {

	return &GetUsersFromPublishedTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetUsersFromPublishedTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
