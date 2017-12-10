// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetLegacyReposSearchKeywordURL generates an URL for the get legacy repos search keyword operation
type GetLegacyReposSearchKeywordURL struct {
	Keyword string

	Language  *string
	Order     *string
	Sort      *string
	StartPage *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetLegacyReposSearchKeywordURL) WithBasePath(bp string) *GetLegacyReposSearchKeywordURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetLegacyReposSearchKeywordURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetLegacyReposSearchKeywordURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/legacy/repos/search/{keyword}"

	keyword := o.Keyword
	if keyword != "" {
		_path = strings.Replace(_path, "{keyword}", keyword, -1)
	} else {
		return nil, errors.New("Keyword is required on GetLegacyReposSearchKeywordURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var language string
	if o.Language != nil {
		language = *o.Language
	}
	if language != "" {
		qs.Set("language", language)
	}

	var order string
	if o.Order != nil {
		order = *o.Order
	}
	if order != "" {
		qs.Set("order", order)
	}

	var sort string
	if o.Sort != nil {
		sort = *o.Sort
	}
	if sort != "" {
		qs.Set("sort", sort)
	}

	var startPage string
	if o.StartPage != nil {
		startPage = *o.StartPage
	}
	if startPage != "" {
		qs.Set("start_page", startPage)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetLegacyReposSearchKeywordURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetLegacyReposSearchKeywordURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetLegacyReposSearchKeywordURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetLegacyReposSearchKeywordURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetLegacyReposSearchKeywordURL")
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
func (o *GetLegacyReposSearchKeywordURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
