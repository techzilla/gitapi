// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteReposOwnerRepoIssuesNumberLabelsNameURL generates an URL for the delete repos owner repo issues number labels name operation
type DeleteReposOwnerRepoIssuesNumberLabelsNameURL struct {
	Name   string
	Number int64
	Owner  string
	Repo   string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) WithBasePath(bp string) *DeleteReposOwnerRepoIssuesNumberLabelsNameURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/repos/{owner}/{repo}/issues/{number}/labels/{name}"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("Name is required on DeleteReposOwnerRepoIssuesNumberLabelsNameURL")
	}
	number := swag.FormatInt64(o.Number)
	if number != "" {
		_path = strings.Replace(_path, "{number}", number, -1)
	} else {
		return nil, errors.New("Number is required on DeleteReposOwnerRepoIssuesNumberLabelsNameURL")
	}
	owner := o.Owner
	if owner != "" {
		_path = strings.Replace(_path, "{owner}", owner, -1)
	} else {
		return nil, errors.New("Owner is required on DeleteReposOwnerRepoIssuesNumberLabelsNameURL")
	}
	repo := o.Repo
	if repo != "" {
		_path = strings.Replace(_path, "{repo}", repo, -1)
	} else {
		return nil, errors.New("Repo is required on DeleteReposOwnerRepoIssuesNumberLabelsNameURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteReposOwnerRepoIssuesNumberLabelsNameURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteReposOwnerRepoIssuesNumberLabelsNameURL")
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
func (o *DeleteReposOwnerRepoIssuesNumberLabelsNameURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}