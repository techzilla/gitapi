// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Gitignore gitignore
// swagger:model gitignore
type Gitignore []interface{}

// Validate validates this gitignore
func (m Gitignore) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
