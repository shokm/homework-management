// Code generated by go-swagger; DO NOT EDIT.

package task_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTasksBySubjectsHandlerFunc turns a function with the right signature into a get tasks by subjects handler
type GetTasksBySubjectsHandlerFunc func(GetTasksBySubjectsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTasksBySubjectsHandlerFunc) Handle(params GetTasksBySubjectsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTasksBySubjectsHandler interface for that can handle valid get tasks by subjects params
type GetTasksBySubjectsHandler interface {
	Handle(GetTasksBySubjectsParams, interface{}) middleware.Responder
}

// NewGetTasksBySubjects creates a new http.Handler for the get tasks by subjects operation
func NewGetTasksBySubjects(ctx *middleware.Context, handler GetTasksBySubjectsHandler) *GetTasksBySubjects {
	return &GetTasksBySubjects{Context: ctx, Handler: handler}
}

/* GetTasksBySubjects swagger:route GET /subjects/tasks TaskApi getTasksBySubjects

該当ユーザの全ての有効な課題を教科ごとに取得する

*/
type GetTasksBySubjects struct {
	Context *middleware.Context
	Handler GetTasksBySubjectsHandler
}

func (o *GetTasksBySubjects) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTasksBySubjectsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
