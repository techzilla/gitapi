// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Events events
// swagger:model events
type Events struct {

	// actor
	Actor *EventsActor `json:"actor,omitempty"`

	// created at
	CreatedAt interface{} `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// org
	Org *EventsOrg `json:"org,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// public
	Public bool `json:"public,omitempty"`

	// repo
	Repo *EventsRepo `json:"repo,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this events
func (m *Events) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrg(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Events) validateActor(formats strfmt.Registry) error {

	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {

		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *Events) validateOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.Org) { // not required
		return nil
	}

	if m.Org != nil {

		if err := m.Org.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

func (m *Events) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Events) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Events) UnmarshalBinary(b []byte) error {
	var res Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
