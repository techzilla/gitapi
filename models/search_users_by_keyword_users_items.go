// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SearchUsersByKeywordUsersItems search users by keyword users items
// swagger:model searchUsersByKeywordUsersItems
type SearchUsersByKeywordUsersItems struct {

	// created
	Created string `json:"created,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// followers
	Followers int64 `json:"followers,omitempty"`

	// followers count
	FollowersCount int64 `json:"followers_count,omitempty"`

	// fullname
	Fullname string `json:"fullname,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// public repo count
	PublicRepoCount int64 `json:"public_repo_count,omitempty"`

	// repos
	Repos int64 `json:"repos,omitempty"`

	// score
	Score float64 `json:"score,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this search users by keyword users items
func (m *SearchUsersByKeywordUsersItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SearchUsersByKeywordUsersItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchUsersByKeywordUsersItems) UnmarshalBinary(b []byte) error {
	var res SearchUsersByKeywordUsersItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
