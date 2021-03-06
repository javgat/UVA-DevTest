// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetFullTestingOKCode is the HTTP code returned for type GetFullTestingOK
const GetFullTestingOKCode int = 200

/*GetFullTestingOK Testing found

swagger:response getFullTestingOK
*/
type GetFullTestingOK struct {

	/*
	  In: Body
	*/
	Payload *models.Testing `json:"body,omitempty"`
}

// NewGetFullTestingOK creates GetFullTestingOK with default headers values
func NewGetFullTestingOK() *GetFullTestingOK {

	return &GetFullTestingOK{}
}

// WithPayload adds the payload to the get full testing o k response
func (o *GetFullTestingOK) WithPayload(payload *models.Testing) *GetFullTestingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get full testing o k response
func (o *GetFullTestingOK) SetPayload(payload *models.Testing) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFullTestingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFullTestingBadRequestCode is the HTTP code returned for type GetFullTestingBadRequest
const GetFullTestingBadRequestCode int = 400

/*GetFullTestingBadRequest Incorrect Request, or invalida data

swagger:response getFullTestingBadRequest
*/
type GetFullTestingBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFullTestingBadRequest creates GetFullTestingBadRequest with default headers values
func NewGetFullTestingBadRequest() *GetFullTestingBadRequest {

	return &GetFullTestingBadRequest{}
}

// WithPayload adds the payload to the get full testing bad request response
func (o *GetFullTestingBadRequest) WithPayload(payload *models.Error) *GetFullTestingBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get full testing bad request response
func (o *GetFullTestingBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFullTestingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFullTestingForbiddenCode is the HTTP code returned for type GetFullTestingForbidden
const GetFullTestingForbiddenCode int = 403

/*GetFullTestingForbidden Not authorized to this content

swagger:response getFullTestingForbidden
*/
type GetFullTestingForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFullTestingForbidden creates GetFullTestingForbidden with default headers values
func NewGetFullTestingForbidden() *GetFullTestingForbidden {

	return &GetFullTestingForbidden{}
}

// WithPayload adds the payload to the get full testing forbidden response
func (o *GetFullTestingForbidden) WithPayload(payload *models.Error) *GetFullTestingForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get full testing forbidden response
func (o *GetFullTestingForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFullTestingForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFullTestingGoneCode is the HTTP code returned for type GetFullTestingGone
const GetFullTestingGoneCode int = 410

/*GetFullTestingGone That resource does not exist

swagger:response getFullTestingGone
*/
type GetFullTestingGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFullTestingGone creates GetFullTestingGone with default headers values
func NewGetFullTestingGone() *GetFullTestingGone {

	return &GetFullTestingGone{}
}

// WithPayload adds the payload to the get full testing gone response
func (o *GetFullTestingGone) WithPayload(payload *models.Error) *GetFullTestingGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get full testing gone response
func (o *GetFullTestingGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFullTestingGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFullTestingInternalServerErrorCode is the HTTP code returned for type GetFullTestingInternalServerError
const GetFullTestingInternalServerErrorCode int = 500

/*GetFullTestingInternalServerError Internal error

swagger:response getFullTestingInternalServerError
*/
type GetFullTestingInternalServerError struct {
}

// NewGetFullTestingInternalServerError creates GetFullTestingInternalServerError with default headers values
func NewGetFullTestingInternalServerError() *GetFullTestingInternalServerError {

	return &GetFullTestingInternalServerError{}
}

// WriteResponse to the client
func (o *GetFullTestingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
