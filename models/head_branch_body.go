// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HeadBranchBody head branch body
// swagger:model headBranchBody
type HeadBranchBody struct {

	// Boolean indicating whether to force the update or to make sure the update is a fast-forward update. The default is false, so leaving this out or setting it to false will make sure you’re not overwriting work.
	// Required: true
	Force *bool `json:"force"`

	// String of the SHA1 value to set this reference to.
	// Required: true
	Sha *string `json:"sha"`
}

// Validate validates this head branch body
func (m *HeadBranchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForce(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSha(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeadBranchBody) validateForce(formats strfmt.Registry) error {

	if err := validate.Required("force", "body", m.Force); err != nil {
		return err
	}

	return nil
}

func (m *HeadBranchBody) validateSha(formats strfmt.Registry) error {

	if err := validate.Required("sha", "body", m.Sha); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeadBranchBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeadBranchBody) UnmarshalBinary(b []byte) error {
	var res HeadBranchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
