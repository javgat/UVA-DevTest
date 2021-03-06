// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// PutEmailConfigurationOKCode is the HTTP code returned for type PutEmailConfigurationOK
const PutEmailConfigurationOKCode int = 200

/*PutEmailConfigurationOK Email configuration updated

swagger:response putEmailConfigurationOK
*/
type PutEmailConfigurationOK struct {
}

// NewPutEmailConfigurationOK creates PutEmailConfigurationOK with default headers values
func NewPutEmailConfigurationOK() *PutEmailConfigurationOK {

	return &PutEmailConfigurationOK{}
}

// WriteResponse to the client
func (o *PutEmailConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutEmailConfigurationForbiddenCode is the HTTP code returned for type PutEmailConfigurationForbidden
const PutEmailConfigurationForbiddenCode int = 403

/*PutEmailConfigurationForbidden Not authorized to this content

swagger:response putEmailConfigurationForbidden
*/
type PutEmailConfigurationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutEmailConfigurationForbidden creates PutEmailConfigurationForbidden with default headers values
func NewPutEmailConfigurationForbidden() *PutEmailConfigurationForbidden {

	return &PutEmailConfigurationForbidden{}
}

// WithPayload adds the payload to the put email configuration forbidden response
func (o *PutEmailConfigurationForbidden) WithPayload(payload *models.Error) *PutEmailConfigurationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put email configuration forbidden response
func (o *PutEmailConfigurationForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEmailConfigurationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutEmailConfigurationInternalServerErrorCode is the HTTP code returned for type PutEmailConfigurationInternalServerError
const PutEmailConfigurationInternalServerErrorCode int = 500

/*PutEmailConfigurationInternalServerError Internal error

swagger:response putEmailConfigurationInternalServerError
*/
type PutEmailConfigurationInternalServerError struct {
}

// NewPutEmailConfigurationInternalServerError creates PutEmailConfigurationInternalServerError with default headers values
func NewPutEmailConfigurationInternalServerError() *PutEmailConfigurationInternalServerError {

	return &PutEmailConfigurationInternalServerError{}
}

// WriteResponse to the client
func (o *PutEmailConfigurationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
