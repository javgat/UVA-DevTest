// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTagsFromQuestionOKCode is the HTTP code returned for type GetTagsFromQuestionOK
const GetTagsFromQuestionOKCode int = 200

/*GetTagsFromQuestionOK tags found

swagger:response getTagsFromQuestionOK
*/
type GetTagsFromQuestionOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Tag `json:"body,omitempty"`
}

// NewGetTagsFromQuestionOK creates GetTagsFromQuestionOK with default headers values
func NewGetTagsFromQuestionOK() *GetTagsFromQuestionOK {

	return &GetTagsFromQuestionOK{}
}

// WithPayload adds the payload to the get tags from question o k response
func (o *GetTagsFromQuestionOK) WithPayload(payload []*models.Tag) *GetTagsFromQuestionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from question o k response
func (o *GetTagsFromQuestionOK) SetPayload(payload []*models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromQuestionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetTagsFromQuestionBadRequestCode is the HTTP code returned for type GetTagsFromQuestionBadRequest
const GetTagsFromQuestionBadRequestCode int = 400

/*GetTagsFromQuestionBadRequest Incorrect Request, or invalida data

swagger:response getTagsFromQuestionBadRequest
*/
type GetTagsFromQuestionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromQuestionBadRequest creates GetTagsFromQuestionBadRequest with default headers values
func NewGetTagsFromQuestionBadRequest() *GetTagsFromQuestionBadRequest {

	return &GetTagsFromQuestionBadRequest{}
}

// WithPayload adds the payload to the get tags from question bad request response
func (o *GetTagsFromQuestionBadRequest) WithPayload(payload *models.Error) *GetTagsFromQuestionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from question bad request response
func (o *GetTagsFromQuestionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromQuestionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromQuestionForbiddenCode is the HTTP code returned for type GetTagsFromQuestionForbidden
const GetTagsFromQuestionForbiddenCode int = 403

/*GetTagsFromQuestionForbidden Not authorized to this content

swagger:response getTagsFromQuestionForbidden
*/
type GetTagsFromQuestionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromQuestionForbidden creates GetTagsFromQuestionForbidden with default headers values
func NewGetTagsFromQuestionForbidden() *GetTagsFromQuestionForbidden {

	return &GetTagsFromQuestionForbidden{}
}

// WithPayload adds the payload to the get tags from question forbidden response
func (o *GetTagsFromQuestionForbidden) WithPayload(payload *models.Error) *GetTagsFromQuestionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from question forbidden response
func (o *GetTagsFromQuestionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromQuestionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromQuestionGoneCode is the HTTP code returned for type GetTagsFromQuestionGone
const GetTagsFromQuestionGoneCode int = 410

/*GetTagsFromQuestionGone That user (password and name) does not exist

swagger:response getTagsFromQuestionGone
*/
type GetTagsFromQuestionGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagsFromQuestionGone creates GetTagsFromQuestionGone with default headers values
func NewGetTagsFromQuestionGone() *GetTagsFromQuestionGone {

	return &GetTagsFromQuestionGone{}
}

// WithPayload adds the payload to the get tags from question gone response
func (o *GetTagsFromQuestionGone) WithPayload(payload *models.Error) *GetTagsFromQuestionGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tags from question gone response
func (o *GetTagsFromQuestionGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagsFromQuestionGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagsFromQuestionInternalServerErrorCode is the HTTP code returned for type GetTagsFromQuestionInternalServerError
const GetTagsFromQuestionInternalServerErrorCode int = 500

/*GetTagsFromQuestionInternalServerError Internal error

swagger:response getTagsFromQuestionInternalServerError
*/
type GetTagsFromQuestionInternalServerError struct {
}

// NewGetTagsFromQuestionInternalServerError creates GetTagsFromQuestionInternalServerError with default headers values
func NewGetTagsFromQuestionInternalServerError() *GetTagsFromQuestionInternalServerError {

	return &GetTagsFromQuestionInternalServerError{}
}

// WriteResponse to the client
func (o *GetTagsFromQuestionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}