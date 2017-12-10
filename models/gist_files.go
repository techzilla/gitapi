// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GistFiles gist files
// swagger:model gistFiles
type GistFiles struct {

	// ring erl
	RingErl *GistFilesRingErl `json:"ring.erl,omitempty"`
}

// Validate validates this gist files
func (m *GistFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRingErl(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GistFiles) validateRingErl(formats strfmt.Registry) error {

	if swag.IsZero(m.RingErl) { // not required
		return nil
	}

	if m.RingErl != nil {

		if err := m.RingErl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ring.erl")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GistFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GistFiles) UnmarshalBinary(b []byte) error {
	var res GistFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
