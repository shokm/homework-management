// Code generated by go-swagger; DO NOT EDIT.

package auth_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"backend/gen/models"
)

// GetAuthUserOKCode is the HTTP code returned for type GetAuthUserOK
const GetAuthUserOKCode int = 200

/*GetAuthUserOK ユーザー登録成功

swagger:response getAuthUserOK
*/
type GetAuthUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthReturnUser `json:"body,omitempty"`
}

// NewGetAuthUserOK creates GetAuthUserOK with default headers values
func NewGetAuthUserOK() *GetAuthUserOK {

	return &GetAuthUserOK{}
}

// WithPayload adds the payload to the get auth user o k response
func (o *GetAuthUserOK) WithPayload(payload *models.AuthReturnUser) *GetAuthUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth user o k response
func (o *GetAuthUserOK) SetPayload(payload *models.AuthReturnUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAuthUserUnauthorizedCode is the HTTP code returned for type GetAuthUserUnauthorized
const GetAuthUserUnauthorizedCode int = 401

/*GetAuthUserUnauthorized Unauthorized

swagger:response getAuthUserUnauthorized
*/
type GetAuthUserUnauthorized struct {
}

// NewGetAuthUserUnauthorized creates GetAuthUserUnauthorized with default headers values
func NewGetAuthUserUnauthorized() *GetAuthUserUnauthorized {

	return &GetAuthUserUnauthorized{}
}

// WriteResponse to the client
func (o *GetAuthUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetAuthUserInternalServerErrorCode is the HTTP code returned for type GetAuthUserInternalServerError
const GetAuthUserInternalServerErrorCode int = 500

/*GetAuthUserInternalServerError Internal Server Error

swagger:response getAuthUserInternalServerError
*/
type GetAuthUserInternalServerError struct {
}

// NewGetAuthUserInternalServerError creates GetAuthUserInternalServerError with default headers values
func NewGetAuthUserInternalServerError() *GetAuthUserInternalServerError {

	return &GetAuthUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetAuthUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// GetAuthUserServiceUnavailableCode is the HTTP code returned for type GetAuthUserServiceUnavailable
const GetAuthUserServiceUnavailableCode int = 503

/*GetAuthUserServiceUnavailable Service Unavailable

swagger:response getAuthUserServiceUnavailable
*/
type GetAuthUserServiceUnavailable struct {
}

// NewGetAuthUserServiceUnavailable creates GetAuthUserServiceUnavailable with default headers values
func NewGetAuthUserServiceUnavailable() *GetAuthUserServiceUnavailable {

	return &GetAuthUserServiceUnavailable{}
}

// WriteResponse to the client
func (o *GetAuthUserServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}
