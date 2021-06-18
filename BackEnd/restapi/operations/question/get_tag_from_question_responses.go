// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTagFromQuestionOKCode is the HTTP code returned for type GetTagFromQuestionOK
const GetTagFromQuestionOKCode int = 200

/*GetTagFromQuestionOK tag found

swagger:response getTagFromQuestionOK
*/
type GetTagFromQuestionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Tag `json:"body,omitempty"`
}

// NewGetTagFromQuestionOK creates GetTagFromQuestionOK with default headers values
func NewGetTagFromQuestionOK() *GetTagFromQuestionOK {

	return &GetTagFromQuestionOK{}
}

// WithPayload adds the payload to the get tag from question o k response
func (o *GetTagFromQuestionOK) WithPayload(payload *models.Tag) *GetTagFromQuestionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag from question o k response
func (o *GetTagFromQuestionOK) SetPayload(payload *models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagFromQuestionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagFromQuestionBadRequestCode is the HTTP code returned for type GetTagFromQuestionBadRequest
const GetTagFromQuestionBadRequestCode int = 400

/*GetTagFromQuestionBadRequest Incorrect Request, or invalida data

swagger:response getTagFromQuestionBadRequest
*/
type GetTagFromQuestionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagFromQuestionBadRequest creates GetTagFromQuestionBadRequest with default headers values
func NewGetTagFromQuestionBadRequest() *GetTagFromQuestionBadRequest {

	return &GetTagFromQuestionBadRequest{}
}

// WithPayload adds the payload to the get tag from question bad request response
func (o *GetTagFromQuestionBadRequest) WithPayload(payload *models.Error) *GetTagFromQuestionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag from question bad request response
func (o *GetTagFromQuestionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagFromQuestionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagFromQuestionForbiddenCode is the HTTP code returned for type GetTagFromQuestionForbidden
const GetTagFromQuestionForbiddenCode int = 403

/*GetTagFromQuestionForbidden Not authorized to this content

swagger:response getTagFromQuestionForbidden
*/
type GetTagFromQuestionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagFromQuestionForbidden creates GetTagFromQuestionForbidden with default headers values
func NewGetTagFromQuestionForbidden() *GetTagFromQuestionForbidden {

	return &GetTagFromQuestionForbidden{}
}

// WithPayload adds the payload to the get tag from question forbidden response
func (o *GetTagFromQuestionForbidden) WithPayload(payload *models.Error) *GetTagFromQuestionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag from question forbidden response
func (o *GetTagFromQuestionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagFromQuestionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagFromQuestionGoneCode is the HTTP code returned for type GetTagFromQuestionGone
const GetTagFromQuestionGoneCode int = 410

/*GetTagFromQuestionGone That resource does not exist

swagger:response getTagFromQuestionGone
*/
type GetTagFromQuestionGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagFromQuestionGone creates GetTagFromQuestionGone with default headers values
func NewGetTagFromQuestionGone() *GetTagFromQuestionGone {

	return &GetTagFromQuestionGone{}
}

// WithPayload adds the payload to the get tag from question gone response
func (o *GetTagFromQuestionGone) WithPayload(payload *models.Error) *GetTagFromQuestionGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag from question gone response
func (o *GetTagFromQuestionGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagFromQuestionGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTagFromQuestionInternalServerErrorCode is the HTTP code returned for type GetTagFromQuestionInternalServerError
const GetTagFromQuestionInternalServerErrorCode int = 500

/*GetTagFromQuestionInternalServerError Internal error

swagger:response getTagFromQuestionInternalServerError
*/
type GetTagFromQuestionInternalServerError struct {
}

// NewGetTagFromQuestionInternalServerError creates GetTagFromQuestionInternalServerError with default headers values
func NewGetTagFromQuestionInternalServerError() *GetTagFromQuestionInternalServerError {

	return &GetTagFromQuestionInternalServerError{}
}

// WriteResponse to the client
func (o *GetTagFromQuestionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
