// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HookBody hook body
// swagger:model hookBody
type HookBody struct {

	// active
	Active bool `json:"active,omitempty"`

	// add events
	AddEvents []string `json:"add_events"`
}

// Validate validates this hook body
func (m *HookBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HookBody) validateAddEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.AddEvents) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HookBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HookBody) UnmarshalBinary(b []byte) error {
	var res HookBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
