// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTagsFromTestOKCode is the HTTP code returned for type GetTagsFromTestOK
const GetTagsFromTestOKCode int = 200

/*GetTagsFromTestOK tags found

swagger:response getTagsFromTestOK
*/
type GetTagsFromTestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Tag `json:"body,omitempty"`
}

// NewGetTagsFromTestOK creates GetTagsFromTestOK with default headers values
func NewGetTagsFromTestOK() *GetTagsFromTestOK {

	return &GetTagsFromTestOK{}
}

// WithPayload adds the payload to the get tags from test o k response
func (o *GetTagsFromTestOK) WithPayload(payload []*models.Tag) *GetTagsFromTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from test o k response
func (o *GetTagsFromTestOK) SetPayload(payload []*models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Tag, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTagsFromTestBadRequestCode is the HTTP code returned for type GetTagsFromTestBadRequest
const GetTagsFromTestBadRequestCode int = 400

/*GetTagsFromTestBadRequest Incorrect Request, or invalida data

swagger:response getTagsFromTestBadRequest
*/
type GetTagsFromTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromTestBadRequest creates GetTagsFromTestBadRequest with default headers values
func NewGetTagsFromTestBadRequest() *GetTagsFromTestBadRequest {

	return &GetTagsFromTestBadRequest{}
}

// WithPayload adds the payload to the get tags from test bad request response
func (o *GetTagsFromTestBadRequest) WithPayload(payload *models.Error) *GetTagsFromTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from test bad request response
func (o *GetTagsFromTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromTestForbiddenCode is the HTTP code returned for type GetTagsFromTestForbidden
const GetTagsFromTestForbiddenCode int = 403

/*GetTagsFromTestForbidden Not authorized to this content

swagger:response getTagsFromTestForbidden
*/
type GetTagsFromTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromTestForbidden creates GetTagsFromTestForbidden with default headers values
func NewGetTagsFromTestForbidden() *GetTagsFromTestForbidden {

	return &GetTagsFromTestForbidden{}
}

// WithPayload adds the payload to the get tags from test forbidden response
func (o *GetTagsFromTestForbidden) WithPayload(payload *models.Error) *GetTagsFromTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from test forbidden response
func (o *GetTagsFromTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromTestGoneCode is the HTTP code returned for type GetTagsFromTestGone
const GetTagsFromTestGoneCode int = 410

/*GetTagsFromTestGone That user (password and name) does not exist

swagger:response getTagsFromTestGone
*/
type GetTagsFromTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromTestGone creates GetTagsFromTestGone with default headers values
func NewGetTagsFromTestGone() *GetTagsFromTestGone {

	return &GetTagsFromTestGone{}
}

// WithPayload adds the payload to the get tags from test gone response
func (o *GetTagsFromTestGone) WithPayload(payload *models.Error) *GetTagsFromTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from test gone response
func (o *GetTagsFromTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromTestInternalServerErrorCode is the HTTP code returned for type GetTagsFromTestInternalServerError
const GetTagsFromTestInternalServerErrorCode int = 500

/*GetTagsFromTestInternalServerError Internal error

swagger:response getTagsFromTestInternalServerError
*/
type GetTagsFromTestInternalServerError struct {
}

// NewGetTagsFromTestInternalServerError creates GetTagsFromTestInternalServerError with default headers values
func NewGetTagsFromTestInternalServerError() *GetTagsFromTestInternalServerError {

	return &GetTagsFromTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetTagsFromTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
