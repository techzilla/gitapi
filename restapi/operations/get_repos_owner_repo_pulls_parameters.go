// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReposOwnerRepoPullsParams creates a new GetReposOwnerRepoPullsParams object
// with the default values initialized.
func NewGetReposOwnerRepoPullsParams() GetReposOwnerRepoPullsParams {
	var (
		stateDefault = string("open")
	)
	return GetReposOwnerRepoPullsParams{
		State: &stateDefault,
	}
}

// GetReposOwnerRepoPullsParams contains all the bound params for the get repos owner repo pulls operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetReposOwnerRepoPulls
type GetReposOwnerRepoPullsParams struct {

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
	/*Filter pulls by base branch name. Example - gh-pages.
	  In: query
	*/
	Base *string
	/*Filter pulls by head user and branch name in the format of 'user:ref-name'.
	Example: github:new-script-format.

	  In: query
	*/
	Head *string
	/*Name of repository owner.
	  Required: true
	  In: path
	*/
	Owner string
	/*Name of repository.
	  Required: true
	  In: path
	*/
	Repo string
	/*String to filter by state.
	  In: query
	  Default: "open"
	*/
	State *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetReposOwnerRepoPullsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

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

	qBase, qhkBase, _ := qs.GetOK("base")
	if err := o.bindBase(qBase, qhkBase, route.Formats); err != nil {
		res = append(res, err)
	}

	qHead, qhkHead, _ := qs.GetOK("head")
	if err := o.bindHead(qHead, qhkHead, route.Formats); err != nil {
		res = append(res, err)
	}

	rOwner, rhkOwner, _ := route.Params.GetOK("owner")
	if err := o.bindOwner(rOwner, rhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}

	rRepo, rhkRepo, _ := route.Params.GetOK("repo")
	if err := o.bindRepo(rRepo, rhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}

	qState, qhkState, _ := qs.GetOK("state")
	if err := o.bindState(qState, qhkState, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReposOwnerRepoPullsParams) bindAccept(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoPullsParams) bindXGitHubMediaType(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoPullsParams) bindXGitHubRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoPullsParams) bindXRateLimitLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoPullsParams) bindXRateLimitRemaining(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoPullsParams) bindXRateLimitReset(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoPullsParams) bindBase(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Base = &raw

	return nil
}

func (o *GetReposOwnerRepoPullsParams) bindHead(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Head = &raw

	return nil
}

func (o *GetReposOwnerRepoPullsParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Owner = raw

	return nil
}

func (o *GetReposOwnerRepoPullsParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Repo = raw

	return nil
}

func (o *GetReposOwnerRepoPullsParams) bindState(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var stateDefault string = string("open")
		o.State = &stateDefault
		return nil
	}

	o.State = &raw

	if err := o.validateState(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetReposOwnerRepoPullsParams) validateState(formats strfmt.Registry) error {

	if err := validate.Enum("state", "query", *o.State, []interface{}{"open", "closed"}); err != nil {
		return err
	}

	return nil
}