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

// RepoCommitBody repo commit body
// swagger:model repoCommitBody
type RepoCommitBody struct {

	// author
	Author *RepoCommitBodyAuthor `json:"author,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`

	// parents
	// Required: true
	Parents []string `json:"parents"`

	// tree
	// Required: true
	Tree *string `json:"tree"`
}

// Validate validates this repo commit body
func (m *RepoCommitBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTree(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoCommitBody) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {

		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

func (m *RepoCommitBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *RepoCommitBody) validateParents(formats strfmt.Registry) error {

	if err := validate.Required("parents", "body", m.Parents); err != nil {
		return err
	}

	return nil
}

func (m *RepoCommitBody) validateTree(formats strfmt.Registry) error {

	if err := validate.Required("tree", "body", m.Tree); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepoCommitBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepoCommitBody) UnmarshalBinary(b []byte) error {
	var res RepoCommitBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
