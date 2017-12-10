// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeploymentPayload deployment payload
// swagger:model deploymentPayload
type DeploymentPayload struct {

	// deploy user
	DeployUser string `json:"deploy_user,omitempty"`

	// environment
	Environment string `json:"environment,omitempty"`

	// room id
	RoomID float64 `json:"room_id,omitempty"`
}

// Validate validates this deployment payload
func (m *DeploymentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentPayload) UnmarshalBinary(b []byte) error {
	var res DeploymentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
