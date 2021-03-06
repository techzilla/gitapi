// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFileBody delete file body
// swagger:model deleteFileBody
type DeleteFileBody struct {

	// committer
	Committer *DeleteFileBodyCommitter `json:"committer,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this delete file body
func (m *DeleteFileBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommitter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteFileBody) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {

		if err := m.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFileBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFileBody) UnmarshalBinary(b []byte) error {
	var res DeleteFileBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
