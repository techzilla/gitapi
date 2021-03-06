// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommitsItemsCommit commits items commit
// swagger:model commitsItemsCommit
type CommitsItemsCommit struct {

	// author
	Author *CommitsItemsCommitAuthor `json:"author,omitempty"`

	// committer
	Committer *CommitsItemsCommitCommitter `json:"committer,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// tree
	Tree *CommitsItemsCommitTree `json:"tree,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this commits items commit
func (m *CommitsItemsCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
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

func (m *CommitsItemsCommit) validateAuthor(formats strfmt.Registry) error {

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

func (m *CommitsItemsCommit) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {

		if err := m.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

func (m *CommitsItemsCommit) validateTree(formats strfmt.Registry) error {

	if swag.IsZero(m.Tree) { // not required
		return nil
	}

	if m.Tree != nil {

		if err := m.Tree.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tree")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommitsItemsCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommitsItemsCommit) UnmarshalBinary(b []byte) error {
	var res CommitsItemsCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
