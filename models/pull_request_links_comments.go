// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PullRequestLinksComments pull request links comments
// swagger:model pullRequestLinksComments
type PullRequestLinksComments struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this pull request links comments
func (m *PullRequestLinksComments) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PullRequestLinksComments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PullRequestLinksComments) UnmarshalBinary(b []byte) error {
	var res PullRequestLinksComments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}