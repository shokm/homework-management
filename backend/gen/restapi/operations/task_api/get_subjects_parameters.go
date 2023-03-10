// Code generated by go-swagger; DO NOT EDIT.

package task_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSubjectsParams creates a new GetSubjectsParams object
//
// There are no default values defined in the spec.
func NewGetSubjectsParams() GetSubjectsParams {

	return GetSubjectsParams{}
}

// GetSubjectsParams contains all the bound params for the get subjects operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSubjects
type GetSubjectsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*アーカイブされた教科だけ取得するか（デフォルトでは含まれない）
	  In: query
	*/
	IsArchived *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSubjectsParams() beforehand.
func (o *GetSubjectsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qIsArchived, qhkIsArchived, _ := qs.GetOK("isArchived")
	if err := o.bindIsArchived(qIsArchived, qhkIsArchived, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIsArchived binds and validates parameter IsArchived from query.
func (o *GetSubjectsParams) bindIsArchived(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("isArchived", "query", "bool", raw)
	}
	o.IsArchived = &value

	return nil
}
