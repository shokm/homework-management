// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthUserReq auth user req
//
// swagger:model AuthUserReq
type AuthUserReq struct {

	// password
	Password string `json:"password,omitempty"`

	// screen name
	ScreenName string `json:"screen_name,omitempty"`
}

// Validate validates this auth user req
func (m *AuthUserReq) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth user req based on context it is used
func (m *AuthUserReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthUserReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthUserReq) UnmarshalBinary(b []byte) error {
	var res AuthUserReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
