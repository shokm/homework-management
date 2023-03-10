// Code generated by go-swagger; DO NOT EDIT.

package auth_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"backend/gen/models"
)

// PostAuthLoginOKCode is the HTTP code returned for type PostAuthLoginOK
const PostAuthLoginOKCode int = 200

/*PostAuthLoginOK ログイン成功

swagger:response postAuthLoginOK
*/
type PostAuthLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthReturnJWT `json:"body,omitempty"`
}

// NewPostAuthLoginOK creates PostAuthLoginOK with default headers values
func NewPostAuthLoginOK() *PostAuthLoginOK {

	return &PostAuthLoginOK{}
}

// WithPayload adds the payload to the post auth login o k response
func (o *PostAuthLoginOK) WithPayload(payload *models.AuthReturnJWT) *PostAuthLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth login o k response
func (o *PostAuthLoginOK) SetPayload(payload *models.AuthReturnJWT) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthLoginUnauthorizedCode is the HTTP code returned for type PostAuthLoginUnauthorized
const PostAuthLoginUnauthorizedCode int = 401

/*PostAuthLoginUnauthorized Unauthorized

swagger:response postAuthLoginUnauthorized
*/
type PostAuthLoginUnauthorized struct {
}

// NewPostAuthLoginUnauthorized creates PostAuthLoginUnauthorized with default headers values
func NewPostAuthLoginUnauthorized() *PostAuthLoginUnauthorized {

	return &PostAuthLoginUnauthorized{}
}

// WriteResponse to the client
func (o *PostAuthLoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAuthLoginInternalServerErrorCode is the HTTP code returned for type PostAuthLoginInternalServerError
const PostAuthLoginInternalServerErrorCode int = 500

/*PostAuthLoginInternalServerError Internal Server Error

swagger:response postAuthLoginInternalServerError
*/
type PostAuthLoginInternalServerError struct {
}

// NewPostAuthLoginInternalServerError creates PostAuthLoginInternalServerError with default headers values
func NewPostAuthLoginInternalServerError() *PostAuthLoginInternalServerError {

	return &PostAuthLoginInternalServerError{}
}

// WriteResponse to the client
func (o *PostAuthLoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// PostAuthLoginServiceUnavailableCode is the HTTP code returned for type PostAuthLoginServiceUnavailable
const PostAuthLoginServiceUnavailableCode int = 503

/*PostAuthLoginServiceUnavailable Service Unavailable

swagger:response postAuthLoginServiceUnavailable
*/
type PostAuthLoginServiceUnavailable struct {
}

// NewPostAuthLoginServiceUnavailable creates PostAuthLoginServiceUnavailable with default headers values
func NewPostAuthLoginServiceUnavailable() *PostAuthLoginServiceUnavailable {

	return &PostAuthLoginServiceUnavailable{}
}

// WriteResponse to the client
func (o *PostAuthLoginServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}
