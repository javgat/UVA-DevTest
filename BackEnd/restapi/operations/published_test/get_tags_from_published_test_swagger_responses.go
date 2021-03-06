// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTagsFromPublishedTestOKCode is the HTTP code returned for type GetTagsFromPublishedTestOK
const GetTagsFromPublishedTestOKCode int = 200

/*GetTagsFromPublishedTestOK tags found

swagger:response getTagsFromPublishedTestOK
*/
type GetTagsFromPublishedTestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Tag `json:"body,omitempty"`
}

// NewGetTagsFromPublishedTestOK creates GetTagsFromPublishedTestOK with default headers values
func NewGetTagsFromPublishedTestOK() *GetTagsFromPublishedTestOK {

	return &GetTagsFromPublishedTestOK{}
}

// WithPayload adds the payload to the get tags from published test o k response
func (o *GetTagsFromPublishedTestOK) WithPayload(payload []*models.Tag) *GetTagsFromPublishedTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from published test o k response
func (o *GetTagsFromPublishedTestOK) SetPayload(payload []*models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromPublishedTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTagsFromPublishedTestBadRequestCode is the HTTP code returned for type GetTagsFromPublishedTestBadRequest
const GetTagsFromPublishedTestBadRequestCode int = 400

/*GetTagsFromPublishedTestBadRequest Incorrect Request, or invalida data

swagger:response getTagsFromPublishedTestBadRequest
*/
type GetTagsFromPublishedTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromPublishedTestBadRequest creates GetTagsFromPublishedTestBadRequest with default headers values
func NewGetTagsFromPublishedTestBadRequest() *GetTagsFromPublishedTestBadRequest {

	return &GetTagsFromPublishedTestBadRequest{}
}

// WithPayload adds the payload to the get tags from published test bad request response
func (o *GetTagsFromPublishedTestBadRequest) WithPayload(payload *models.Error) *GetTagsFromPublishedTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from published test bad request response
func (o *GetTagsFromPublishedTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromPublishedTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromPublishedTestForbiddenCode is the HTTP code returned for type GetTagsFromPublishedTestForbidden
const GetTagsFromPublishedTestForbiddenCode int = 403

/*GetTagsFromPublishedTestForbidden Not authorized to this content

swagger:response getTagsFromPublishedTestForbidden
*/
type GetTagsFromPublishedTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromPublishedTestForbidden creates GetTagsFromPublishedTestForbidden with default headers values
func NewGetTagsFromPublishedTestForbidden() *GetTagsFromPublishedTestForbidden {

	return &GetTagsFromPublishedTestForbidden{}
}

// WithPayload adds the payload to the get tags from published test forbidden response
func (o *GetTagsFromPublishedTestForbidden) WithPayload(payload *models.Error) *GetTagsFromPublishedTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from published test forbidden response
func (o *GetTagsFromPublishedTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromPublishedTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromPublishedTestGoneCode is the HTTP code returned for type GetTagsFromPublishedTestGone
const GetTagsFromPublishedTestGoneCode int = 410

/*GetTagsFromPublishedTestGone That resource does not exist

swagger:response getTagsFromPublishedTestGone
*/
type GetTagsFromPublishedTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromPublishedTestGone creates GetTagsFromPublishedTestGone with default headers values
func NewGetTagsFromPublishedTestGone() *GetTagsFromPublishedTestGone {

	return &GetTagsFromPublishedTestGone{}
}

// WithPayload adds the payload to the get tags from published test gone response
func (o *GetTagsFromPublishedTestGone) WithPayload(payload *models.Error) *GetTagsFromPublishedTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from published test gone response
func (o *GetTagsFromPublishedTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromPublishedTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromPublishedTestInternalServerErrorCode is the HTTP code returned for type GetTagsFromPublishedTestInternalServerError
const GetTagsFromPublishedTestInternalServerErrorCode int = 500

/*GetTagsFromPublishedTestInternalServerError Internal error

swagger:response getTagsFromPublishedTestInternalServerError
*/
type GetTagsFromPublishedTestInternalServerError struct {
}

// NewGetTagsFromPublishedTestInternalServerError creates GetTagsFromPublishedTestInternalServerError with default headers values
func NewGetTagsFromPublishedTestInternalServerError() *GetTagsFromPublishedTestInternalServerError {

	return &GetTagsFromPublishedTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetTagsFromPublishedTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
