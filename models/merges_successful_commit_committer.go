// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MergesSuccessfulCommitCommitter merges successful commit committer
// swagger:model mergesSuccessfulCommitCommitter
type MergesSuccessfulCommitCommitter struct {

	// date
	Date string `json:"date,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this merges successful commit committer
func (m *MergesSuccessfulCommitCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MergesSuccessfulCommitCommitter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MergesSuccessfulCommitCommitter) UnmarshalBinary(b []byte) error {
	var res MergesSuccessfulCommitCommitter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
