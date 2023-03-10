// Code generated by go-swagger; DO NOT EDIT.

package task_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSubjectBySubjectIDHandlerFunc turns a function with the right signature into a get subject by subject Id handler
type GetSubjectBySubjectIDHandlerFunc func(GetSubjectBySubjectIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSubjectBySubjectIDHandlerFunc) Handle(params GetSubjectBySubjectIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSubjectBySubjectIDHandler interface for that can handle valid get subject by subject Id params
type GetSubjectBySubjectIDHandler interface {
	Handle(GetSubjectBySubjectIDParams, interface{}) middleware.Responder
}

// NewGetSubjectBySubjectID creates a new http.Handler for the get subject by subject Id operation
func NewGetSubjectBySubjectID(ctx *middleware.Context, handler GetSubjectBySubjectIDHandler) *GetSubjectBySubjectID {
	return &GetSubjectBySubjectID{Context: ctx, Handler: handler}
}

/* GetSubjectBySubjectID swagger:route GET /subject/{subject_id} TaskApi getSubjectBySubjectId

{subject_id}の教科を取得する

*/
type GetSubjectBySubjectID struct {
	Context *middleware.Context
	Handler GetSubjectBySubjectIDHandler
}

func (o *GetSubjectBySubjectID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSubjectBySubjectIDParams()
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
