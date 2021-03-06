// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// RemoveTestFavoriteOKCode is the HTTP code returned for type RemoveTestFavoriteOK
const RemoveTestFavoriteOKCode int = 200

/*RemoveTestFavoriteOK test unmarked

swagger:response removeTestFavoriteOK
*/
type RemoveTestFavoriteOK struct {
}

// NewRemoveTestFavoriteOK creates RemoveTestFavoriteOK with default headers values
func NewRemoveTestFavoriteOK() *RemoveTestFavoriteOK {

	return &RemoveTestFavoriteOK{}
}

// WriteResponse to the client
func (o *RemoveTestFavoriteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RemoveTestFavoriteBadRequestCode is the HTTP code returned for type RemoveTestFavoriteBadRequest
const RemoveTestFavoriteBadRequestCode int = 400

/*RemoveTestFavoriteBadRequest Incorrect Request, or invalida data

swagger:response removeTestFavoriteBadRequest
*/
type RemoveTestFavoriteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTestFavoriteBadRequest creates RemoveTestFavoriteBadRequest with default headers values
func NewRemoveTestFavoriteBadRequest() *RemoveTestFavoriteBadRequest {

	return &RemoveTestFavoriteBadRequest{}
}

// WithPayload adds the payload to the remove test favorite bad request response
func (o *RemoveTestFavoriteBadRequest) WithPayload(payload *models.Error) *RemoveTestFavoriteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove test favorite bad request response
func (o *RemoveTestFavoriteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTestFavoriteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTestFavoriteForbiddenCode is the HTTP code returned for type RemoveTestFavoriteForbidden
const RemoveTestFavoriteForbiddenCode int = 403

/*RemoveTestFavoriteForbidden Not authorized to this content

swagger:response removeTestFavoriteForbidden
*/
type RemoveTestFavoriteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTestFavoriteForbidden creates RemoveTestFavoriteForbidden with default headers values
func NewRemoveTestFavoriteForbidden() *RemoveTestFavoriteForbidden {

	return &RemoveTestFavoriteForbidden{}
}

// WithPayload adds the payload to the remove test favorite forbidden response
func (o *RemoveTestFavoriteForbidden) WithPayload(payload *models.Error) *RemoveTestFavoriteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove test favorite forbidden response
func (o *RemoveTestFavoriteForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTestFavoriteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTestFavoriteGoneCode is the HTTP code returned for type RemoveTestFavoriteGone
const RemoveTestFavoriteGoneCode int = 410

/*RemoveTestFavoriteGone That resource does not exist

swagger:response removeTestFavoriteGone
*/
type RemoveTestFavoriteGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTestFavoriteGone creates RemoveTestFavoriteGone with default headers values
func NewRemoveTestFavoriteGone() *RemoveTestFavoriteGone {

	return &RemoveTestFavoriteGone{}
}

// WithPayload adds the payload to the remove test favorite gone response
func (o *RemoveTestFavoriteGone) WithPayload(payload *models.Error) *RemoveTestFavoriteGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove test favorite gone response
func (o *RemoveTestFavoriteGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTestFavoriteGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTestFavoriteInternalServerErrorCode is the HTTP code returned for type RemoveTestFavoriteInternalServerError
const RemoveTestFavoriteInternalServerErrorCode int = 500

/*RemoveTestFavoriteInternalServerError Internal error

swagger:response removeTestFavoriteInternalServerError
*/
type RemoveTestFavoriteInternalServerError struct {
}

// NewRemoveTestFavoriteInternalServerError creates RemoveTestFavoriteInternalServerError with default headers values
func NewRemoveTestFavoriteInternalServerError() *RemoveTestFavoriteInternalServerError {

	return &RemoveTestFavoriteInternalServerError{}
}

// WriteResponse to the client
func (o *RemoveTestFavoriteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
