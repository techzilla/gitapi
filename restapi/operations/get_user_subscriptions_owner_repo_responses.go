// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUserSubscriptionsOwnerRepoNoContentCode is the HTTP code returned for type GetUserSubscriptionsOwnerRepoNoContent
const GetUserSubscriptionsOwnerRepoNoContentCode int = 204

/*GetUserSubscriptionsOwnerRepoNoContent Repository is watched by you.

swagger:response getUserSubscriptionsOwnerRepoNoContent
*/
type GetUserSubscriptionsOwnerRepoNoContent struct {
}

// NewGetUserSubscriptionsOwnerRepoNoContent creates GetUserSubscriptionsOwnerRepoNoContent with default headers values
func NewGetUserSubscriptionsOwnerRepoNoContent() *GetUserSubscriptionsOwnerRepoNoContent {
	return &GetUserSubscriptionsOwnerRepoNoContent{}
}

// WriteResponse to the client
func (o *GetUserSubscriptionsOwnerRepoNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetUserSubscriptionsOwnerRepoForbiddenCode is the HTTP code returned for type GetUserSubscriptionsOwnerRepoForbidden
const GetUserSubscriptionsOwnerRepoForbiddenCode int = 403

/*GetUserSubscriptionsOwnerRepoForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUserSubscriptionsOwnerRepoForbidden
*/
type GetUserSubscriptionsOwnerRepoForbidden struct {
}

// NewGetUserSubscriptionsOwnerRepoForbidden creates GetUserSubscriptionsOwnerRepoForbidden with default headers values
func NewGetUserSubscriptionsOwnerRepoForbidden() *GetUserSubscriptionsOwnerRepoForbidden {
	return &GetUserSubscriptionsOwnerRepoForbidden{}
}

// WriteResponse to the client
func (o *GetUserSubscriptionsOwnerRepoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetUserSubscriptionsOwnerRepoNotFoundCode is the HTTP code returned for type GetUserSubscriptionsOwnerRepoNotFound
const GetUserSubscriptionsOwnerRepoNotFoundCode int = 404

/*GetUserSubscriptionsOwnerRepoNotFound Repository is not watched by you.

swagger:response getUserSubscriptionsOwnerRepoNotFound
*/
type GetUserSubscriptionsOwnerRepoNotFound struct {
}

// NewGetUserSubscriptionsOwnerRepoNotFound creates GetUserSubscriptionsOwnerRepoNotFound with default headers values
func NewGetUserSubscriptionsOwnerRepoNotFound() *GetUserSubscriptionsOwnerRepoNotFound {
	return &GetUserSubscriptionsOwnerRepoNotFound{}
}

// WriteResponse to the client
func (o *GetUserSubscriptionsOwnerRepoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}