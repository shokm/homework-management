// Code generated by go-swagger; DO NOT EDIT.

package task_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"backend/gen/models"
)

// GetTasksBySubjectsOKCode is the HTTP code returned for type GetTasksBySubjectsOK
const GetTasksBySubjectsOKCode int = 200

/*GetTasksBySubjectsOK 成功

swagger:response getTasksBySubjectsOK
*/
type GetTasksBySubjectsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnTasksBySubjects `json:"body,omitempty"`
}

// NewGetTasksBySubjectsOK creates GetTasksBySubjectsOK with default headers values
func NewGetTasksBySubjectsOK() *GetTasksBySubjectsOK {

	return &GetTasksBySubjectsOK{}
}

// WithPayload adds the payload to the get tasks by subjects o k response
func (o *GetTasksBySubjectsOK) WithPayload(payload *models.ReturnTasksBySubjects) *GetTasksBySubjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tasks by subjects o k response
func (o *GetTasksBySubjectsOK) SetPayload(payload *models.ReturnTasksBySubjects) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTasksBySubjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTasksBySubjectsUnauthorizedCode is the HTTP code returned for type GetTasksBySubjectsUnauthorized
const GetTasksBySubjectsUnauthorizedCode int = 401

/*GetTasksBySubjectsUnauthorized Unauthorized

swagger:response getTasksBySubjectsUnauthorized
*/
type GetTasksBySubjectsUnauthorized struct {
}

// NewGetTasksBySubjectsUnauthorized creates GetTasksBySubjectsUnauthorized with default headers values
func NewGetTasksBySubjectsUnauthorized() *GetTasksBySubjectsUnauthorized {

	return &GetTasksBySubjectsUnauthorized{}
}

// WriteResponse to the client
func (o *GetTasksBySubjectsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetTasksBySubjectsNotFoundCode is the HTTP code returned for type GetTasksBySubjectsNotFound
const GetTasksBySubjectsNotFoundCode int = 404

/*GetTasksBySubjectsNotFound Not Found

swagger:response getTasksBySubjectsNotFound
*/
type GetTasksBySubjectsNotFound struct {
}

// NewGetTasksBySubjectsNotFound creates GetTasksBySubjectsNotFound with default headers values
func NewGetTasksBySubjectsNotFound() *GetTasksBySubjectsNotFound {

	return &GetTasksBySubjectsNotFound{}
}

// WriteResponse to the client
func (o *GetTasksBySubjectsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTasksBySubjectsInternalServerErrorCode is the HTTP code returned for type GetTasksBySubjectsInternalServerError
const GetTasksBySubjectsInternalServerErrorCode int = 500

/*GetTasksBySubjectsInternalServerError Internal Server Error

swagger:response getTasksBySubjectsInternalServerError
*/
type GetTasksBySubjectsInternalServerError struct {
}

// NewGetTasksBySubjectsInternalServerError creates GetTasksBySubjectsInternalServerError with default headers values
func NewGetTasksBySubjectsInternalServerError() *GetTasksBySubjectsInternalServerError {

	return &GetTasksBySubjectsInternalServerError{}
}

// WriteResponse to the client
func (o *GetTasksBySubjectsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// GetTasksBySubjectsServiceUnavailableCode is the HTTP code returned for type GetTasksBySubjectsServiceUnavailable
const GetTasksBySubjectsServiceUnavailableCode int = 503

/*GetTasksBySubjectsServiceUnavailable Service Unavailable

swagger:response getTasksBySubjectsServiceUnavailable
*/
type GetTasksBySubjectsServiceUnavailable struct {
}

// NewGetTasksBySubjectsServiceUnavailable creates GetTasksBySubjectsServiceUnavailable with default headers values
func NewGetTasksBySubjectsServiceUnavailable() *GetTasksBySubjectsServiceUnavailable {

	return &GetTasksBySubjectsServiceUnavailable{}
}

// WriteResponse to the client
func (o *GetTasksBySubjectsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}
