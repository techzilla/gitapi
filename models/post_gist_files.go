// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostGistFiles post gist files
// swagger:model postGistFiles
type PostGistFiles struct {

	// file1 txt
	File1Txt *PostGistFilesFile1Txt `json:"file1.txt,omitempty"`
}

// Validate validates this post gist files
func (m *PostGistFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile1Txt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostGistFiles) validateFile1Txt(formats strfmt.Registry) error {

	if swag.IsZero(m.File1Txt) { // not required
		return nil
	}

	if m.File1Txt != nil {

		if err := m.File1Txt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file1.txt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostGistFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostGistFiles) UnmarshalBinary(b []byte) error {
	var res PostGistFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
