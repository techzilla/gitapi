// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateDownload create download
// swagger:model createDownload
type CreateDownload struct {

	// accesskeyid
	Accesskeyid string `json:"accesskeyid,omitempty"`

	// acl
	ACL string `json:"acl,omitempty"`

	// bucket
	Bucket string `json:"bucket,omitempty"`

	// content type
	ContentType string `json:"content_type,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// download count
	DownloadCount int64 `json:"download_count,omitempty"`

	// ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
	Expirationdate string `json:"expirationdate,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// mime type
	MimeType string `json:"mime_type,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// policy
	Policy string `json:"policy,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`

	// redirect
	Redirect bool `json:"redirect,omitempty"`

	// s3 url
	S3URL string `json:"s3_url,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this create download
func (m *CreateDownload) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateDownload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDownload) UnmarshalBinary(b []byte) error {
	var res CreateDownload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
