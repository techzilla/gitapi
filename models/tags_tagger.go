// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TagsTagger tags tagger
// swagger:model tagsTagger
type TagsTagger struct {

	// Timestamp of when this object was tagged.
	Date string `json:"date,omitempty"`

	// String of the email of the author of the tag.
	Email string `json:"email,omitempty"`

	// String of the name of the author of the tag.
	Name string `json:"name,omitempty"`
}

// Validate validates this tags tagger
func (m *TagsTagger) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TagsTagger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagsTagger) UnmarshalBinary(b []byte) error {
	var res TagsTagger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
