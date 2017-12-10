// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateFileContent create file content
// swagger:model createFileContent
type CreateFileContent struct {

	// links
	Links *CreateFileContentLinks `json:"_links,omitempty"`

	// git url
	GitURL string `json:"git_url,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this create file content
func (m *CreateFileContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFileContent) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFileContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFileContent) UnmarshalBinary(b []byte) error {
	var res CreateFileContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
