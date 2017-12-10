// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Meta meta
// swagger:model meta
type Meta struct {

	// git
	Git []string `json:"git"`

	// hooks
	Hooks []string `json:"hooks"`
}

// Validate validates this meta
func (m *Meta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHooks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Meta) validateGit(formats strfmt.Registry) error {

	if swag.IsZero(m.Git) { // not required
		return nil
	}

	return nil
}

func (m *Meta) validateHooks(formats strfmt.Registry) error {

	if swag.IsZero(m.Hooks) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Meta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Meta) UnmarshalBinary(b []byte) error {
	var res Meta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
