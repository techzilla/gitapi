// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommentsItems comments items
// swagger:model commentsItems
type CommentsItems struct {

	// body
	Body string `json:"body,omitempty"`

	// ISO 8601.
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User *CommentsItemsUser `json:"user,omitempty"`
}

// Validate validates this comments items
func (m *CommentsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentsItems) validateUser(formats strfmt.Registry) error {

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
func (m *CommentsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentsItems) UnmarshalBinary(b []byte) error {
	var res CommentsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
