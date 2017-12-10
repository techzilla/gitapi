// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PullsCommentLinksSelf pulls comment links self
// swagger:model pullsCommentLinksSelf
type PullsCommentLinksSelf struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pulls comment links self
func (m *PullsCommentLinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PullsCommentLinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PullsCommentLinksSelf) UnmarshalBinary(b []byte) error {
	var res PullsCommentLinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}