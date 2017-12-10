// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Trees trees
// swagger:model trees
type Trees struct {

	// base tree
	BaseTree string `json:"base_tree,omitempty"`

	// SHA1 checksum ID of the object in the tree.
	Sha string `json:"sha,omitempty"`

	// tree
	Tree TreesTree `json:"tree"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this trees
func (m *Trees) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Trees) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Trees) UnmarshalBinary(b []byte) error {
	var res Trees
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
