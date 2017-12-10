// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CompareCommitsCommitsItemsCommitter compare commits commits items committer
// swagger:model compareCommitsCommitsItemsCommitter
type CompareCommitsCommitsItemsCommitter struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// events url
	EventsURL string `json:"events_url,omitempty"`

	// followers url
	FollowersURL string `json:"followers_url,omitempty"`

	// following url
	FollowingURL string `json:"following_url,omitempty"`

	// gists url
	GistsURL string `json:"gists_url,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// organizations url
	OrganizationsURL string `json:"organizations_url,omitempty"`

	// received events url
	ReceivedEventsURL string `json:"received_events_url,omitempty"`

	// repos url
	ReposURL string `json:"repos_url,omitempty"`

	// site admin
	SiteAdmin bool `json:"site_admin,omitempty"`

	// starred url
	StarredURL string `json:"starred_url,omitempty"`

	// subscriptions url
	SubscriptionsURL string `json:"subscriptions_url,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this compare commits commits items committer
func (m *CompareCommitsCommitsItemsCommitter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CompareCommitsCommitsItemsCommitter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompareCommitsCommitsItemsCommitter) UnmarshalBinary(b []byte) error {
	var res CompareCommitsCommitsItemsCommitter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
