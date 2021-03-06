// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFile delete file
// swagger:model deleteFile
type DeleteFile struct {

	// commit
	Commit *DeleteFileCommit `json:"commit,omitempty"`

	// content
	Content string `json:"content,omitempty"`
}

// Validate validates this delete file
func (m *DeleteFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteFile) validateCommit(formats strfmt.Registry) error {

	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {

		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFile) UnmarshalBinary(b []byte) error {
	var res DeleteFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
