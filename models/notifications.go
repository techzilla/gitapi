// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Notifications notifications
// swagger:model notifications
type Notifications struct {

	// id
	ID int64 `json:"id,omitempty"`

	// last read at
	LastReadAt string `json:"last_read_at,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// repository
	Repository *NotificationsRepository `json:"repository,omitempty"`

	// subject
	Subject *NotificationsSubject `json:"subject,omitempty"`

	// unread
	Unread bool `json:"unread,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this notifications
func (m *Notifications) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notifications) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {

		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *Notifications) validateSubject(formats strfmt.Registry) error {

	if swag.IsZero(m.Subject) { // not required
		return nil
	}

	if m.Subject != nil {

		if err := m.Subject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notifications) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notifications) UnmarshalBinary(b []byte) error {
	var res Notifications
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
