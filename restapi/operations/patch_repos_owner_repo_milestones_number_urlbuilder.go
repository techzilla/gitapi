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

// PatchReposOwnerRepoMilestonesNumberURL generates an URL for the patch repos owner repo milestones number operation
type PatchReposOwnerRepoMilestonesNumberURL struct {
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
func (o *PatchReposOwnerRepoMilestonesNumberURL) WithBasePath(bp string) *PatchReposOwnerRepoMilestonesNumberURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PatchReposOwnerRepoMilestonesNumberURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PatchReposOwnerRepoMilestonesNumberURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/repos/{owner}/{repo}/milestones/{number}"

	number := swag.FormatInt64(o.Number)
	if number != "" {
		_path = strings.Replace(_path, "{number}", number, -1)
	} else {
		return nil, errors.New("Number is required on PatchReposOwnerRepoMilestonesNumberURL")
	}
	owner := o.Owner
	if owner != "" {
		_path = strings.Replace(_path, "{owner}", owner, -1)
	} else {
		return nil, errors.New("Owner is required on PatchReposOwnerRepoMilestonesNumberURL")
	}
	repo := o.Repo
	if repo != "" {
		_path = strings.Replace(_path, "{repo}", repo, -1)
	} else {
		return nil, errors.New("Repo is required on PatchReposOwnerRepoMilestonesNumberURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PatchReposOwnerRepoMilestonesNumberURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PatchReposOwnerRepoMilestonesNumberURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PatchReposOwnerRepoMilestonesNumberURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PatchReposOwnerRepoMilestonesNumberURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PatchReposOwnerRepoMilestonesNumberURL")
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
func (o *PatchReposOwnerRepoMilestonesNumberURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
