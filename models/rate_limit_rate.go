// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RateLimitRate rate limit rate
// swagger:model rateLimitRate
type RateLimitRate struct {

	// limit
	Limit int64 `json:"limit,omitempty"`

	// remaining
	Remaining int64 `json:"remaining,omitempty"`

	// reset
	Reset int64 `json:"reset,omitempty"`
}

// Validate validates this rate limit rate
func (m *RateLimitRate) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RateLimitRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateLimitRate) UnmarshalBinary(b []byte) error {
	var res RateLimitRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}