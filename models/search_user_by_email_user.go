// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SearchUserByEmailUser search user by email user
// swagger:model searchUserByEmailUser
type SearchUserByEmailUser struct {

	// blog
	Blog string `json:"blog,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// followers count
	FollowersCount int64 `json:"followers_count,omitempty"`

	// following count
	FollowingCount int64 `json:"following_count,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// public gist count
	PublicGistCount int64 `json:"public_gist_count,omitempty"`

	// public repo count
	PublicRepoCount int64 `json:"public_repo_count,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this search user by email user
func (m *SearchUserByEmailUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SearchUserByEmailUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchUserByEmailUser) UnmarshalBinary(b []byte) error {
	var res SearchUserByEmailUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
