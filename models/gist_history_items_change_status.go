// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GistHistoryItemsChangeStatus gist history items change status
// swagger:model gistHistoryItemsChangeStatus
type GistHistoryItemsChangeStatus struct {

	// additions
	Additions int64 `json:"additions,omitempty"`

	// deletions
	Deletions int64 `json:"deletions,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this gist history items change status
func (m *GistHistoryItemsChangeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GistHistoryItemsChangeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GistHistoryItemsChangeStatus) UnmarshalBinary(b []byte) error {
	var res GistHistoryItemsChangeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}