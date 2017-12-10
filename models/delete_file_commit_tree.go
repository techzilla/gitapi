// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteFileCommitTree delete file commit tree
// swagger:model deleteFileCommitTree
type DeleteFileCommitTree struct {

	// sha
	Sha string `json:"sha,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this delete file commit tree
func (m *DeleteFileCommitTree) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFileCommitTree) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFileCommitTree) UnmarshalBinary(b []byte) error {
	var res DeleteFileCommitTree
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
