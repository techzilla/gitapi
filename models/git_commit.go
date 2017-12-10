// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GitCommit git commit
// swagger:model gitCommit
type GitCommit struct {

	// author
	Author *GitCommitAuthor `json:"author,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// parents
	Parents string `json:"parents,omitempty"`

	// tree
	Tree string `json:"tree,omitempty"`
}

// Validate validates this git commit
func (m *GitCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitCommit) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {

		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitCommit) UnmarshalBinary(b []byte) error {
	var res GitCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
