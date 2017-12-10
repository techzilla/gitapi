// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RefStatusItems ref status items
// swagger:model refStatusItems
type RefStatusItems struct {

	// commit url
	CommitURL string `json:"commit_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// repository url
	RepositoryURL string `json:"repository_url,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// statuses
	Statuses RefStatusItemsStatuses `json:"statuses"`
}

// Validate validates this ref status items
func (m *RefStatusItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RefStatusItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefStatusItems) UnmarshalBinary(b []byte) error {
	var res RefStatusItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
