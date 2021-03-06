// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateFileCommit create file commit
// swagger:model createFileCommit
type CreateFileCommit struct {

	// author
	Author *CreateFileCommitAuthor `json:"author,omitempty"`

	// committer
	Committer *CreateFileCommitCommitter `json:"committer,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// parents
	Parents CreateFileCommitParents `json:"parents"`

	// sha
	Sha string `json:"sha,omitempty"`

	// tree
	Tree *CreateFileCommitTree `json:"tree,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this create file commit
func (m *CreateFileCommit) Validate(formats strfmt.Registry) error {
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

func (m *CreateFileCommit) validateAuthor(formats strfmt.Registry) error {

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

func (m *CreateFileCommit) validateCommitter(formats strfmt.Registry) error {

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

func (m *CreateFileCommit) validateTree(formats strfmt.Registry) error {

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
func (m *CreateFileCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFileCommit) UnmarshalBinary(b []byte) error {
	var res CreateFileCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
