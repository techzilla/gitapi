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

// GetTeamsTeamIDReposOwnerRepoURL generates an URL for the get teams team ID repos owner repo operation
type GetTeamsTeamIDReposOwnerRepoURL struct {
	Owner  string
	Repo   string
	TeamID int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetTeamsTeamIDReposOwnerRepoURL) WithBasePath(bp string) *GetTeamsTeamIDReposOwnerRepoURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetTeamsTeamIDReposOwnerRepoURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetTeamsTeamIDReposOwnerRepoURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/teams/{teamId}/repos/{owner}/{repo}"

	owner := o.Owner
	if owner != "" {
		_path = strings.Replace(_path, "{owner}", owner, -1)
	} else {
		return nil, errors.New("Owner is required on GetTeamsTeamIDReposOwnerRepoURL")
	}
	repo := o.Repo
	if repo != "" {
		_path = strings.Replace(_path, "{repo}", repo, -1)
	} else {
		return nil, errors.New("Repo is required on GetTeamsTeamIDReposOwnerRepoURL")
	}
	teamID := swag.FormatInt64(o.TeamID)
	if teamID != "" {
		_path = strings.Replace(_path, "{teamId}", teamID, -1)
	} else {
		return nil, errors.New("TeamID is required on GetTeamsTeamIDReposOwnerRepoURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetTeamsTeamIDReposOwnerRepoURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetTeamsTeamIDReposOwnerRepoURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetTeamsTeamIDReposOwnerRepoURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetTeamsTeamIDReposOwnerRepoURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetTeamsTeamIDReposOwnerRepoURL")
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
func (o *GetTeamsTeamIDReposOwnerRepoURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}