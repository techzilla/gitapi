// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SearchCodeItemsItemsRepository search code items items repository
// swagger:model searchCodeItemsItemsRepository
type SearchCodeItemsItemsRepository struct {

	// archive url
	ArchiveURL string `json:"archive_url,omitempty"`

	// assignees url
	AssigneesURL string `json:"assignees_url,omitempty"`

	// blobs url
	BlobsURL string `json:"blobs_url,omitempty"`

	// branches url
	BranchesURL string `json:"branches_url,omitempty"`

	// collaborators url
	CollaboratorsURL string `json:"collaborators_url,omitempty"`

	// comments url
	CommentsURL string `json:"comments_url,omitempty"`

	// commits url
	CommitsURL string `json:"commits_url,omitempty"`

	// compare url
	CompareURL string `json:"compare_url,omitempty"`

	// contents url
	ContentsURL string `json:"contents_url,omitempty"`

	// contributors url
	ContributorsURL string `json:"contributors_url,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// downloads url
	DownloadsURL string `json:"downloads_url,omitempty"`

	// events url
	EventsURL string `json:"events_url,omitempty"`

	// fork
	Fork bool `json:"fork,omitempty"`

	// forks url
	ForksURL string `json:"forks_url,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// git commits url
	GitCommitsURL string `json:"git_commits_url,omitempty"`

	// git refs url
	GitRefsURL string `json:"git_refs_url,omitempty"`

	// git tags url
	GitTagsURL string `json:"git_tags_url,omitempty"`

	// hooks url
	HooksURL string `json:"hooks_url,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// issue comment url
	IssueCommentURL string `json:"issue_comment_url,omitempty"`

	// issue events url
	IssueEventsURL string `json:"issue_events_url,omitempty"`

	// issues url
	IssuesURL string `json:"issues_url,omitempty"`

	// keys url
	KeysURL string `json:"keys_url,omitempty"`

	// labels url
	LabelsURL string `json:"labels_url,omitempty"`

	// languages url
	LanguagesURL string `json:"languages_url,omitempty"`

	// merges url
	MergesURL string `json:"merges_url,omitempty"`

	// milestones url
	MilestonesURL string `json:"milestones_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notifications url
	NotificationsURL string `json:"notifications_url,omitempty"`

	// owner
	Owner *SearchCodeItemsItemsRepositoryOwner `json:"owner,omitempty"`

	// private
	Private bool `json:"private,omitempty"`

	// pulls url
	PullsURL string `json:"pulls_url,omitempty"`

	// stargazers url
	StargazersURL string `json:"stargazers_url,omitempty"`

	// statuses url
	StatusesURL string `json:"statuses_url,omitempty"`

	// subscribers url
	SubscribersURL string `json:"subscribers_url,omitempty"`

	// subscription url
	SubscriptionURL string `json:"subscription_url,omitempty"`

	// tags url
	TagsURL string `json:"tags_url,omitempty"`

	// teams url
	TeamsURL string `json:"teams_url,omitempty"`

	// trees url
	TreesURL string `json:"trees_url,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this search code items items repository
func (m *SearchCodeItemsItemsRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchCodeItemsItemsRepository) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {

		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchCodeItemsItemsRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchCodeItemsItemsRepository) UnmarshalBinary(b []byte) error {
	var res SearchCodeItemsItemsRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
