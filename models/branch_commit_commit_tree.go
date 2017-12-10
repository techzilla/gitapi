// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BranchCommitCommitTree branch commit commit tree
// swagger:model branchCommitCommitTree
type BranchCommitCommitTree struct {

	// sha
	Sha string `json:"sha,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this branch commit commit tree
func (m *BranchCommitCommitTree) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BranchCommitCommitTree) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchCommitCommitTree) UnmarshalBinary(b []byte) error {
	var res BranchCommitCommitTree
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}