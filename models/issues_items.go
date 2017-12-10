// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IssuesItems issues items
// swagger:model issuesItems
type IssuesItems struct {

	// assignee
	Assignee *IssuesItemsAssignee `json:"assignee,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	ClosedAt string `json:"closed_at,omitempty"`

	// comments
	Comments int64 `json:"comments,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt string `json:"created_at,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// labels
	Labels IssuesItemsLabels `json:"labels"`

	// milestone
	Milestone *IssuesItemsMilestone `json:"milestone,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`

	// pull request
	PullRequest *IssuesItemsPullRequest `json:"pull_request,omitempty"`

	// state
	State interface{} `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *IssuesItemsUser `json:"user,omitempty"`
}

// Validate validates this issues items
func (m *IssuesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignee(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMilestone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePullRequest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *IssuesItems) validateAssignee(formats strfmt.Registry) error {

	if swag.IsZero(m.Assignee) { // not required
		return nil
	}

	if m.Assignee != nil {

		if err := m.Assignee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

func (m *IssuesItems) validateMilestone(formats strfmt.Registry) error {

	if swag.IsZero(m.Milestone) { // not required
		return nil
	}

	if m.Milestone != nil {

		if err := m.Milestone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("milestone")
			}
			return err
		}
	}

	return nil
}

func (m *IssuesItems) validatePullRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.PullRequest) { // not required
		return nil
	}

	if m.PullRequest != nil {

		if err := m.PullRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request")
			}
			return err
		}
	}

	return nil
}

var issuesItemsTypeStatePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["open","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		issuesItemsTypeStatePropEnum = append(issuesItemsTypeStatePropEnum, v)
	}
}

// prop value enum
func (m *IssuesItems) validateStateEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, issuesItemsTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IssuesItems) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	return nil
}

func (m *IssuesItems) validateUser(formats strfmt.Registry) error {

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
func (m *IssuesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssuesItems) UnmarshalBinary(b []byte) error {
	var res IssuesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
