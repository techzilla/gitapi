// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RepoCommitBodyAuthor repo commit body author
// swagger:model repoCommitBodyAuthor
type RepoCommitBodyAuthor struct {

	// date
	Date string `json:"date,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this repo commit body author
func (m *RepoCommitBodyAuthor) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RepoCommitBodyAuthor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepoCommitBodyAuthor) UnmarshalBinary(b []byte) error {
	var res RepoCommitBodyAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
