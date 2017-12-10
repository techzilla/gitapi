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

// PutTeamsTeamIDReposOrgRepoURL generates an URL for the put teams team ID repos org repo operation
type PutTeamsTeamIDReposOrgRepoURL struct {
	Org    string
	Repo   string
	TeamID int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutTeamsTeamIDReposOrgRepoURL) WithBasePath(bp string) *PutTeamsTeamIDReposOrgRepoURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutTeamsTeamIDReposOrgRepoURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PutTeamsTeamIDReposOrgRepoURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/teams/{teamId}/repos/{org}/{repo}"

	org := o.Org
	if org != "" {
		_path = strings.Replace(_path, "{org}", org, -1)
	} else {
		return nil, errors.New("Org is required on PutTeamsTeamIDReposOrgRepoURL")
	}
	repo := o.Repo
	if repo != "" {
		_path = strings.Replace(_path, "{repo}", repo, -1)
	} else {
		return nil, errors.New("Repo is required on PutTeamsTeamIDReposOrgRepoURL")
	}
	teamID := swag.FormatInt64(o.TeamID)
	if teamID != "" {
		_path = strings.Replace(_path, "{teamId}", teamID, -1)
	} else {
		return nil, errors.New("TeamID is required on PutTeamsTeamIDReposOrgRepoURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PutTeamsTeamIDReposOrgRepoURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PutTeamsTeamIDReposOrgRepoURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PutTeamsTeamIDReposOrgRepoURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PutTeamsTeamIDReposOrgRepoURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PutTeamsTeamIDReposOrgRepoURL")
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
func (o *PutTeamsTeamIDReposOrgRepoURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}