// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetFavoriteTestsOKCode is the HTTP code returned for type GetFavoriteTestsOK
const GetFavoriteTestsOKCode int = 200

/*GetFavoriteTestsOK tests found

swagger:response getFavoriteTestsOK
*/
type GetFavoriteTestsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetFavoriteTestsOK creates GetFavoriteTestsOK with default headers values
func NewGetFavoriteTestsOK() *GetFavoriteTestsOK {

	return &GetFavoriteTestsOK{}
}

// WithPayload adds the payload to the get favorite tests o k response
func (o *GetFavoriteTestsOK) WithPayload(payload []*models.Test) *GetFavoriteTestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite tests o k response
func (o *GetFavoriteTestsOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteTestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetFavoriteTestsBadRequestCode is the HTTP code returned for type GetFavoriteTestsBadRequest
const GetFavoriteTestsBadRequestCode int = 400

/*GetFavoriteTestsBadRequest Incorrect Request, or invalida data

swagger:response getFavoriteTestsBadRequest
*/
type GetFavoriteTestsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFavoriteTestsBadRequest creates GetFavoriteTestsBadRequest with default headers values
func NewGetFavoriteTestsBadRequest() *GetFavoriteTestsBadRequest {

	return &GetFavoriteTestsBadRequest{}
}

// WithPayload adds the payload to the get favorite tests bad request response
func (o *GetFavoriteTestsBadRequest) WithPayload(payload *models.Error) *GetFavoriteTestsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite tests bad request response
func (o *GetFavoriteTestsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteTestsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteTestsForbiddenCode is the HTTP code returned for type GetFavoriteTestsForbidden
const GetFavoriteTestsForbiddenCode int = 403

/*GetFavoriteTestsForbidden Not authorized to this content

swagger:response getFavoriteTestsForbidden
*/
type GetFavoriteTestsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFavoriteTestsForbidden creates GetFavoriteTestsForbidden with default headers values
func NewGetFavoriteTestsForbidden() *GetFavoriteTestsForbidden {

	return &GetFavoriteTestsForbidden{}
}

// WithPayload adds the payload to the get favorite tests forbidden response
func (o *GetFavoriteTestsForbidden) WithPayload(payload *models.Error) *GetFavoriteTestsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite tests forbidden response
func (o *GetFavoriteTestsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteTestsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteTestsGoneCode is the HTTP code returned for type GetFavoriteTestsGone
const GetFavoriteTestsGoneCode int = 410

/*GetFavoriteTestsGone That resource does not exist

swagger:response getFavoriteTestsGone
*/
type GetFavoriteTestsGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFavoriteTestsGone creates GetFavoriteTestsGone with default headers values
func NewGetFavoriteTestsGone() *GetFavoriteTestsGone {

	return &GetFavoriteTestsGone{}
}

// WithPayload adds the payload to the get favorite tests gone response
func (o *GetFavoriteTestsGone) WithPayload(payload *models.Error) *GetFavoriteTestsGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite tests gone response
func (o *GetFavoriteTestsGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteTestsGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteTestsInternalServerErrorCode is the HTTP code returned for type GetFavoriteTestsInternalServerError
const GetFavoriteTestsInternalServerErrorCode int = 500

/*GetFavoriteTestsInternalServerError Internal error

swagger:response getFavoriteTestsInternalServerError
*/
type GetFavoriteTestsInternalServerError struct {
}

// NewGetFavoriteTestsInternalServerError creates GetFavoriteTestsInternalServerError with default headers values
func NewGetFavoriteTestsInternalServerError() *GetFavoriteTestsInternalServerError {

	return &GetFavoriteTestsInternalServerError{}
}

// WriteResponse to the client
func (o *GetFavoriteTestsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
