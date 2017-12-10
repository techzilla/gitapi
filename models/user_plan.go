// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserPlan user plan
// swagger:model userPlan
type UserPlan struct {

	// collaborators
	Collaborators int64 `json:"collaborators,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private repos
	PrivateRepos int64 `json:"private_repos,omitempty"`

	// space
	Space int64 `json:"space,omitempty"`
}

// Validate validates this user plan
func (m *UserPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPlan) UnmarshalBinary(b []byte) error {
	var res UserPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
