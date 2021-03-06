// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Asset asset
// swagger:model asset
type Asset struct {

	// content type
	ContentType string `json:"content_type,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// download count
	DownloadCount float64 `json:"download_count,omitempty"`

	// id
	ID float64 `json:"id,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size float64 `json:"size,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// uploader
	Uploader *AssetUploader `json:"uploader,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this asset
func (m *Asset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUploader(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Asset) validateUploader(formats strfmt.Registry) error {

	if swag.IsZero(m.Uploader) { // not required
		return nil
	}

	if m.Uploader != nil {

		if err := m.Uploader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploader")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Asset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Asset) UnmarshalBinary(b []byte) error {
	var res Asset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
