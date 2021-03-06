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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReposOwnerRepoNotificationsParams creates a new GetReposOwnerRepoNotificationsParams object
// with the default values initialized.
func NewGetReposOwnerRepoNotificationsParams() GetReposOwnerRepoNotificationsParams {
	var ()
	return GetReposOwnerRepoNotificationsParams{}
}

// GetReposOwnerRepoNotificationsParams contains all the bound params for the get repos owner repo notifications operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetReposOwnerRepoNotifications
type GetReposOwnerRepoNotificationsParams struct {

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
	/*True to show notifications marked as read.
	  In: query
	*/
	All *bool
	/*Name of repository owner.
	  Required: true
	  In: path
	*/
	Owner string
	/*True to show only notifications in which the user is directly participating
	or mentioned.

	  In: query
	*/
	Participating *bool
	/*Name of repository.
	  Required: true
	  In: path
	*/
	Repo string
	/*The time should be passed in as UTC in the ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ.
	Example: "2012-10-09T23:39:01Z".

	  In: query
	*/
	Since *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetReposOwnerRepoNotificationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	qAll, qhkAll, _ := qs.GetOK("all")
	if err := o.bindAll(qAll, qhkAll, route.Formats); err != nil {
		res = append(res, err)
	}

	rOwner, rhkOwner, _ := route.Params.GetOK("owner")
	if err := o.bindOwner(rOwner, rhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}

	qParticipating, qhkParticipating, _ := qs.GetOK("participating")
	if err := o.bindParticipating(qParticipating, qhkParticipating, route.Formats); err != nil {
		res = append(res, err)
	}

	rRepo, rhkRepo, _ := route.Params.GetOK("repo")
	if err := o.bindRepo(rRepo, rhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReposOwnerRepoNotificationsParams) bindAccept(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoNotificationsParams) bindXGitHubMediaType(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoNotificationsParams) bindXGitHubRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoNotificationsParams) bindXRateLimitLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoNotificationsParams) bindXRateLimitRemaining(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoNotificationsParams) bindXRateLimitReset(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetReposOwnerRepoNotificationsParams) bindAll(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("all", "query", "bool", raw)
	}
	o.All = &value

	return nil
}

func (o *GetReposOwnerRepoNotificationsParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Owner = raw

	return nil
}

func (o *GetReposOwnerRepoNotificationsParams) bindParticipating(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("participating", "query", "bool", raw)
	}
	o.Participating = &value

	return nil
}

func (o *GetReposOwnerRepoNotificationsParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Repo = raw

	return nil
}

func (o *GetReposOwnerRepoNotificationsParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Since = &raw

	return nil
}
