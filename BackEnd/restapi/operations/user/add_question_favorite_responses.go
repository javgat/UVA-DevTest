// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// AddQuestionFavoriteOKCode is the HTTP code returned for type AddQuestionFavoriteOK
const AddQuestionFavoriteOKCode int = 200

/*AddQuestionFavoriteOK question marked

swagger:response addQuestionFavoriteOK
*/
type AddQuestionFavoriteOK struct {
}

// NewAddQuestionFavoriteOK creates AddQuestionFavoriteOK with default headers values
func NewAddQuestionFavoriteOK() *AddQuestionFavoriteOK {

	return &AddQuestionFavoriteOK{}
}

// WriteResponse to the client
func (o *AddQuestionFavoriteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AddQuestionFavoriteBadRequestCode is the HTTP code returned for type AddQuestionFavoriteBadRequest
const AddQuestionFavoriteBadRequestCode int = 400

/*AddQuestionFavoriteBadRequest Incorrect Request, or invalida data

swagger:response addQuestionFavoriteBadRequest
*/
type AddQuestionFavoriteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddQuestionFavoriteBadRequest creates AddQuestionFavoriteBadRequest with default headers values
func NewAddQuestionFavoriteBadRequest() *AddQuestionFavoriteBadRequest {

	return &AddQuestionFavoriteBadRequest{}
}

// WithPayload adds the payload to the add question favorite bad request response
func (o *AddQuestionFavoriteBadRequest) WithPayload(payload *models.Error) *AddQuestionFavoriteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add question favorite bad request response
func (o *AddQuestionFavoriteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddQuestionFavoriteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddQuestionFavoriteForbiddenCode is the HTTP code returned for type AddQuestionFavoriteForbidden
const AddQuestionFavoriteForbiddenCode int = 403

/*AddQuestionFavoriteForbidden Not authorized to this content

swagger:response addQuestionFavoriteForbidden
*/
type AddQuestionFavoriteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddQuestionFavoriteForbidden creates AddQuestionFavoriteForbidden with default headers values
func NewAddQuestionFavoriteForbidden() *AddQuestionFavoriteForbidden {

	return &AddQuestionFavoriteForbidden{}
}

// WithPayload adds the payload to the add question favorite forbidden response
func (o *AddQuestionFavoriteForbidden) WithPayload(payload *models.Error) *AddQuestionFavoriteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add question favorite forbidden response
func (o *AddQuestionFavoriteForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddQuestionFavoriteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddQuestionFavoriteGoneCode is the HTTP code returned for type AddQuestionFavoriteGone
const AddQuestionFavoriteGoneCode int = 410

/*AddQuestionFavoriteGone That resource does not exist

swagger:response addQuestionFavoriteGone
*/
type AddQuestionFavoriteGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddQuestionFavoriteGone creates AddQuestionFavoriteGone with default headers values
func NewAddQuestionFavoriteGone() *AddQuestionFavoriteGone {

	return &AddQuestionFavoriteGone{}
}

// WithPayload adds the payload to the add question favorite gone response
func (o *AddQuestionFavoriteGone) WithPayload(payload *models.Error) *AddQuestionFavoriteGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add question favorite gone response
func (o *AddQuestionFavoriteGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddQuestionFavoriteGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddQuestionFavoriteInternalServerErrorCode is the HTTP code returned for type AddQuestionFavoriteInternalServerError
const AddQuestionFavoriteInternalServerErrorCode int = 500

/*AddQuestionFavoriteInternalServerError Internal error

swagger:response addQuestionFavoriteInternalServerError
*/
type AddQuestionFavoriteInternalServerError struct {
}

// NewAddQuestionFavoriteInternalServerError creates AddQuestionFavoriteInternalServerError with default headers values
func NewAddQuestionFavoriteInternalServerError() *AddQuestionFavoriteInternalServerError {

	return &AddQuestionFavoriteInternalServerError{}
}

// WriteResponse to the client
func (o *AddQuestionFavoriteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
