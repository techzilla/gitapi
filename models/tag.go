// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Tag tag
// swagger:model tag
type Tag struct {

	// message
	Message string `json:"message,omitempty"`

	// object
	Object *TagObject `json:"object,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// tagger
	Tagger *TagTagger `json:"tagger,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTagger(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validateObject(formats strfmt.Registry) error {

	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {

		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) validateTagger(formats strfmt.Registry) error {

	if swag.IsZero(m.Tagger) { // not required
		return nil
	}

	if m.Tagger != nil {

		if err := m.Tagger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tagger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
