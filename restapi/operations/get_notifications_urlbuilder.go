// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetNotificationsURL generates an URL for the get notifications operation
type GetNotificationsURL struct {
	All           *bool
	Participating *bool
	Since         *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNotificationsURL) WithBasePath(bp string) *GetNotificationsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNotificationsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetNotificationsURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/notifications"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var all string
	if o.All != nil {
		all = swag.FormatBool(*o.All)
	}
	if all != "" {
		qs.Set("all", all)
	}

	var participating string
	if o.Participating != nil {
		participating = swag.FormatBool(*o.Participating)
	}
	if participating != "" {
		qs.Set("participating", participating)
	}

	var since string
	if o.Since != nil {
		since = *o.Since
	}
	if since != "" {
		qs.Set("since", since)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetNotificationsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetNotificationsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetNotificationsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetNotificationsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetNotificationsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetNotificationsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
