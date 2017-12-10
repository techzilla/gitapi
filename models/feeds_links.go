// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FeedsLinks feeds links
// swagger:model feedsLinks
type FeedsLinks struct {

	// current user
	CurrentUser *FeedsLinksCurrentUser `json:"current_user,omitempty"`

	// current user actor
	CurrentUserActor *FeedsLinksCurrentUserActor `json:"current_user_actor,omitempty"`

	// current user organization
	CurrentUserOrganization *FeedsLinksCurrentUserOrganization `json:"current_user_organization,omitempty"`

	// current user public
	CurrentUserPublic *FeedsLinksCurrentUserPublic `json:"current_user_public,omitempty"`

	// timeline
	Timeline *FeedsLinksTimeline `json:"timeline,omitempty"`

	// user
	User *FeedsLinksUser `json:"user,omitempty"`
}

// Validate validates this feeds links
func (m *FeedsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentUserActor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentUserOrganization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentUserPublic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeline(formats); err != nil {
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

func (m *FeedsLinks) validateCurrentUser(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentUser) { // not required
		return nil
	}

	if m.CurrentUser != nil {

		if err := m.CurrentUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_user")
			}
			return err
		}
	}

	return nil
}

func (m *FeedsLinks) validateCurrentUserActor(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentUserActor) { // not required
		return nil
	}

	if m.CurrentUserActor != nil {

		if err := m.CurrentUserActor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_user_actor")
			}
			return err
		}
	}

	return nil
}

func (m *FeedsLinks) validateCurrentUserOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentUserOrganization) { // not required
		return nil
	}

	if m.CurrentUserOrganization != nil {

		if err := m.CurrentUserOrganization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_user_organization")
			}
			return err
		}
	}

	return nil
}

func (m *FeedsLinks) validateCurrentUserPublic(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentUserPublic) { // not required
		return nil
	}

	if m.CurrentUserPublic != nil {

		if err := m.CurrentUserPublic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_user_public")
			}
			return err
		}
	}

	return nil
}

func (m *FeedsLinks) validateTimeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeline) { // not required
		return nil
	}

	if m.Timeline != nil {

		if err := m.Timeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeline")
			}
			return err
		}
	}

	return nil
}

func (m *FeedsLinks) validateUser(formats strfmt.Registry) error {

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
func (m *FeedsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedsLinks) UnmarshalBinary(b []byte) error {
	var res FeedsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
