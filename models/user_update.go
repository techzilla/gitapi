// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserUpdate user update
// swagger:model user-update
type UserUpdate struct {

	// bio
	Bio string `json:"bio,omitempty"`

	// blog
	Blog string `json:"blog,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// hireable
	Hireable bool `json:"hireable,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user update
func (m *UserUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdate) UnmarshalBinary(b []byte) error {
	var res UserUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
