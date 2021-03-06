// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IssuesItemsPullRequest issues items pull request
// swagger:model issuesItemsPullRequest
type IssuesItemsPullRequest struct {

	// diff url
	DiffURL string `json:"diff_url,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// patch url
	PatchURL string `json:"patch_url,omitempty"`
}

// Validate validates this issues items pull request
func (m *IssuesItemsPullRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IssuesItemsPullRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssuesItemsPullRequest) UnmarshalBinary(b []byte) error {
	var res IssuesItemsPullRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
