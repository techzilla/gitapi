// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrgsOrgPublicMembersUsernameParams creates a new GetOrgsOrgPublicMembersUsernameParams object
// with the default values initialized.
func NewGetOrgsOrgPublicMembersUsernameParams() GetOrgsOrgPublicMembersUsernameParams {
	var ()
	return GetOrgsOrgPublicMembersUsernameParams{}
}

// GetOrgsOrgPublicMembersUsernameParams contains all the bound params for the get orgs org public members username operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetOrgsOrgPublicMembersUsername
type GetOrgsOrgPublicMembersUsernameParams struct {

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
	/*Name of organisation.
	  Required: true
	  In: path
	*/
	Org string
	/*Name of the user.
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetOrgsOrgPublicMembersUsernameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	rOrg, rhkOrg, _ := route.Params.GetOK("org")
	if err := o.bindOrg(rOrg, rhkOrg, route.Formats); err != nil {
		res = append(res, err)
	}

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgsOrgPublicMembersUsernameParams) bindAccept(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOrgsOrgPublicMembersUsernameParams) bindXGitHubMediaType(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOrgsOrgPublicMembersUsernameParams) bindXGitHubRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOrgsOrgPublicMembersUsernameParams) bindXRateLimitLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOrgsOrgPublicMembersUsernameParams) bindXRateLimitRemaining(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOrgsOrgPublicMembersUsernameParams) bindXRateLimitReset(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOrgsOrgPublicMembersUsernameParams) bindOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Org = raw

	return nil
}

func (o *GetOrgsOrgPublicMembersUsernameParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Username = raw

	return nil
}
