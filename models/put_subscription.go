// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PutSubscription put subscription
// swagger:model putSubscription
type PutSubscription struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// ignored
	Ignored bool `json:"ignored,omitempty"`

	// reason
	Reason interface{} `json:"reason,omitempty"`

	// subscribed
	Subscribed bool `json:"subscribed,omitempty"`

	// thread url
	ThreadURL string `json:"thread_url,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this put subscription
func (m *PutSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PutSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutSubscription) UnmarshalBinary(b []byte) error {
	var res PutSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}