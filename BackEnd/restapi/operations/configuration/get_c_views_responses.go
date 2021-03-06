// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetCViewsOKCode is the HTTP code returned for type GetCViewsOK
const GetCViewsOKCode int = 200

/*GetCViewsOK CustomizedViews Found

swagger:response getCViewsOK
*/
type GetCViewsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CustomizedView `json:"body,omitempty"`
}

// NewGetCViewsOK creates GetCViewsOK with default headers values
func NewGetCViewsOK() *GetCViewsOK {

	return &GetCViewsOK{}
}

// WithPayload adds the payload to the get c views o k response
func (o *GetCViewsOK) WithPayload(payload []*models.CustomizedView) *GetCViewsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get c views o k response
func (o *GetCViewsOK) SetPayload(payload []*models.CustomizedView) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCViewsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CustomizedView, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCViewsForbiddenCode is the HTTP code returned for type GetCViewsForbidden
const GetCViewsForbiddenCode int = 403

/*GetCViewsForbidden Not authorized to this content

swagger:response getCViewsForbidden
*/
type GetCViewsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCViewsForbidden creates GetCViewsForbidden with default headers values
func NewGetCViewsForbidden() *GetCViewsForbidden {

	return &GetCViewsForbidden{}
}

// WithPayload adds the payload to the get c views forbidden response
func (o *GetCViewsForbidden) WithPayload(payload *models.Error) *GetCViewsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get c views forbidden response
func (o *GetCViewsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCViewsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCViewsInternalServerErrorCode is the HTTP code returned for type GetCViewsInternalServerError
const GetCViewsInternalServerErrorCode int = 500

/*GetCViewsInternalServerError Internal error

swagger:response getCViewsInternalServerError
*/
type GetCViewsInternalServerError struct {
}

// NewGetCViewsInternalServerError creates GetCViewsInternalServerError with default headers values
func NewGetCViewsInternalServerError() *GetCViewsInternalServerError {

	return &GetCViewsInternalServerError{}
}

// WriteResponse to the client
func (o *GetCViewsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
