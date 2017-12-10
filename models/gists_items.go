// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GistsItems gists items
// swagger:model gistsItems
type GistsItems struct {

	// comments
	Comments int64 `json:"comments,omitempty"`

	// comments url
	CommentsURL string `json:"comments_url,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// files
	Files *GistsItemsFiles `json:"files,omitempty"`

	// git pull url
	GitPullURL string `json:"git_pull_url,omitempty"`

	// git push url
	GitPushURL string `json:"git_push_url,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// public
	Public bool `json:"public,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *GistsItemsUser `json:"user,omitempty"`
}

// Validate validates this gists items
func (m *GistsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GistsItems) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	if m.Files != nil {

		if err := m.Files.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("files")
			}
			return err
		}
	}

	return nil
}

func (m *GistsItems) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GistsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GistsItems) UnmarshalBinary(b []byte) error {
	var res GistsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
