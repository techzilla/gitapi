// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReposOwnerRepoArchiveFormatPathParams creates a new GetReposOwnerRepoArchiveFormatPathParams object
// with the default values initialized.
func NewGetReposOwnerRepoArchiveFormatPathParams() GetReposOwnerRepoArchiveFormatPathParams {
	var ()
	return GetReposOwnerRepoArchiveFormatPathParams{}
}

// GetReposOwnerRepoArchiveFormatPathParams contains all the bound params for the get repos owner repo archive format path operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetReposOwnerRepoArchiveFormatPath
type GetReposOwnerRepoArchiveFormatPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Is used to set specified media type.
	  In: header
	*/
	Accept *string
	/*You can check the current version of media type in responses.

	  In: header
	*/
	XGitHubMediaType *string
	/*
	  In: header
	*/
	XGitHubRequestID *int64
	/*
	  In: header
	*/
	XRateLimitLimit *int64
	/*
	  In: header
	*/
	XRateLimitRemaining *int64
	/*
	  In: header
	*/
	XRateLimitReset *int64
	/*
	  Required: true
	  In: path
	*/
	ArchiveFormat string
	/*Name of repository owner.
	  Required: true
	  In: path
	*/
	Owner string
	/*Valid Git reference, defaults to 'master'.
	  Required: true
	  In: path
	*/
	Path string
	/*Name of repository.
	  Required: true
	  In: path
	*/
	Repo string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetReposOwnerRepoArchiveFormatPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := o.bindAccept(r.Header[http.CanonicalHeaderKey("Accept")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXGitHubMediaType(r.Header[http.CanonicalHeaderKey("X-GitHub-Media-Type")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXGitHubRequestID(r.Header[http.CanonicalHeaderKey("X-GitHub-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRateLimitLimit(r.Header[http.CanonicalHeaderKey("X-RateLimit-Limit")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRateLimitRemaining(r.Header[http.CanonicalHeaderKey("X-RateLimit-Remaining")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRateLimitReset(r.Header[http.CanonicalHeaderKey("X-RateLimit-Reset")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rArchiveFormat, rhkArchiveFormat, _ := route.Params.GetOK("archive_format")
	if err := o.bindArchiveFormat(rArchiveFormat, rhkArchiveFormat, route.Formats); err != nil {
		res = append(res, err)
	}

	rOwner, rhkOwner, _ := route.Params.GetOK("owner")
	if err := o.bindOwner(rOwner, rhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	rRepo, rhkRepo, _ := route.Params.GetOK("repo")
	if err := o.bindRepo(rRepo, rhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindAccept(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Accept = &raw

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindXGitHubMediaType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XGitHubMediaType = &raw

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindXGitHubRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("X-GitHub-Request-Id", "header", "int64", raw)
	}
	o.XGitHubRequestID = &value

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindXRateLimitLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", raw)
	}
	o.XRateLimitLimit = &value

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindXRateLimitRemaining(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", raw)
	}
	o.XRateLimitRemaining = &value

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindXRateLimitReset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("X-RateLimit-Reset", "header", "int64", raw)
	}
	o.XRateLimitReset = &value

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindArchiveFormat(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ArchiveFormat = raw

	if err := o.validateArchiveFormat(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) validateArchiveFormat(formats strfmt.Registry) error {

	if err := validate.Enum("archive_format", "path", o.ArchiveFormat, []interface{}{"tarball", "zipball"}); err != nil {
		return err
	}

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Owner = raw

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Path = raw

	return nil
}

func (o *GetReposOwnerRepoArchiveFormatPathParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Repo = raw

	return nil
}