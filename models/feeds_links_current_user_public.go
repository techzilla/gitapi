// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FeedsLinksCurrentUserPublic feeds links current user public
// swagger:model feedsLinksCurrentUserPublic
type FeedsLinksCurrentUserPublic struct {

	// href
	Href string `json:"href,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this feeds links current user public
func (m *FeedsLinksCurrentUserPublic) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FeedsLinksCurrentUserPublic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedsLinksCurrentUserPublic) UnmarshalBinary(b []byte) error {
	var res FeedsLinksCurrentUserPublic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
