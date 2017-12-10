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

// Milestone milestone
// swagger:model milestone
type Milestone struct {

	// closed issues
	ClosedIssues int64 `json:"closed_issues,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt string `json:"created_at,omitempty"`

	// creator
	Creator *MilestoneCreator `json:"creator,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	DueOn string `json:"due_on,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues,omitempty"`

	// state
	State interface{} `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this milestone
func (m *Milestone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Milestone) validateCreator(formats strfmt.Registry) error {

	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {

		if err := m.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

var milestoneTypeStatePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["open","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		milestoneTypeStatePropEnum = append(milestoneTypeStatePropEnum, v)
	}
}

// prop value enum
func (m *Milestone) validateStateEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, milestoneTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Milestone) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Milestone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Milestone) UnmarshalBinary(b []byte) error {
	var res Milestone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}