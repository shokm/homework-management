// Code generated by go-swagger; DO NOT EDIT.

package task_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"backend/gen/models"
)

// PostTaskByTaskIDOKCode is the HTTP code returned for type PostTaskByTaskIDOK
const PostTaskByTaskIDOKCode int = 200

/*PostTaskByTaskIDOK 成功

swagger:response postTaskByTaskIdOK
*/
type PostTaskByTaskIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.TaskSingle `json:"body,omitempty"`
}

// NewPostTaskByTaskIDOK creates PostTaskByTaskIDOK with default headers values
func NewPostTaskByTaskIDOK() *PostTaskByTaskIDOK {

	return &PostTaskByTaskIDOK{}
}

// WithPayload adds the payload to the post task by task Id o k response
func (o *PostTaskByTaskIDOK) WithPayload(payload *models.TaskSingle) *PostTaskByTaskIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post task by task Id o k response
func (o *PostTaskByTaskIDOK) SetPayload(payload *models.TaskSingle) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTaskByTaskIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTaskByTaskIDUnauthorizedCode is the HTTP code returned for type PostTaskByTaskIDUnauthorized
const PostTaskByTaskIDUnauthorizedCode int = 401

/*PostTaskByTaskIDUnauthorized Unauthorized

swagger:response postTaskByTaskIdUnauthorized
*/
type PostTaskByTaskIDUnauthorized struct {
}

// NewPostTaskByTaskIDUnauthorized creates PostTaskByTaskIDUnauthorized with default headers values
func NewPostTaskByTaskIDUnauthorized() *PostTaskByTaskIDUnauthorized {

	return &PostTaskByTaskIDUnauthorized{}
}

// WriteResponse to the client
func (o *PostTaskByTaskIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostTaskByTaskIDNotFoundCode is the HTTP code returned for type PostTaskByTaskIDNotFound
const PostTaskByTaskIDNotFoundCode int = 404

/*PostTaskByTaskIDNotFound Not Found

swagger:response postTaskByTaskIdNotFound
*/
type PostTaskByTaskIDNotFound struct {
}

// NewPostTaskByTaskIDNotFound creates PostTaskByTaskIDNotFound with default headers values
func NewPostTaskByTaskIDNotFound() *PostTaskByTaskIDNotFound {

	return &PostTaskByTaskIDNotFound{}
}

// WriteResponse to the client
func (o *PostTaskByTaskIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostTaskByTaskIDInternalServerErrorCode is the HTTP code returned for type PostTaskByTaskIDInternalServerError
const PostTaskByTaskIDInternalServerErrorCode int = 500

/*PostTaskByTaskIDInternalServerError Internal Server Error

swagger:response postTaskByTaskIdInternalServerError
*/
type PostTaskByTaskIDInternalServerError struct {
}

// NewPostTaskByTaskIDInternalServerError creates PostTaskByTaskIDInternalServerError with default headers values
func NewPostTaskByTaskIDInternalServerError() *PostTaskByTaskIDInternalServerError {

	return &PostTaskByTaskIDInternalServerError{}
}

// WriteResponse to the client
func (o *PostTaskByTaskIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// PostTaskByTaskIDServiceUnavailableCode is the HTTP code returned for type PostTaskByTaskIDServiceUnavailable
const PostTaskByTaskIDServiceUnavailableCode int = 503

/*PostTaskByTaskIDServiceUnavailable Service Unavailable

swagger:response postTaskByTaskIdServiceUnavailable
*/
type PostTaskByTaskIDServiceUnavailable struct {
}

// NewPostTaskByTaskIDServiceUnavailable creates PostTaskByTaskIDServiceUnavailable with default headers values
func NewPostTaskByTaskIDServiceUnavailable() *PostTaskByTaskIDServiceUnavailable {

	return &PostTaskByTaskIDServiceUnavailable{}
}

// WriteResponse to the client
func (o *PostTaskByTaskIDServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}