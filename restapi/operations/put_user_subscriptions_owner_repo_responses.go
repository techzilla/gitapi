// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutUserSubscriptionsOwnerRepoNoContentCode is the HTTP code returned for type PutUserSubscriptionsOwnerRepoNoContent
const PutUserSubscriptionsOwnerRepoNoContentCode int = 204

/*PutUserSubscriptionsOwnerRepoNoContent Repository is watched.

swagger:response putUserSubscriptionsOwnerRepoNoContent
*/
type PutUserSubscriptionsOwnerRepoNoContent struct {
}

// NewPutUserSubscriptionsOwnerRepoNoContent creates PutUserSubscriptionsOwnerRepoNoContent with default headers values
func NewPutUserSubscriptionsOwnerRepoNoContent() *PutUserSubscriptionsOwnerRepoNoContent {
	return &PutUserSubscriptionsOwnerRepoNoContent{}
}

// WriteResponse to the client
func (o *PutUserSubscriptionsOwnerRepoNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PutUserSubscriptionsOwnerRepoForbiddenCode is the HTTP code returned for type PutUserSubscriptionsOwnerRepoForbidden
const PutUserSubscriptionsOwnerRepoForbiddenCode int = 403

/*PutUserSubscriptionsOwnerRepoForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response putUserSubscriptionsOwnerRepoForbidden
*/
type PutUserSubscriptionsOwnerRepoForbidden struct {
}

// NewPutUserSubscriptionsOwnerRepoForbidden creates PutUserSubscriptionsOwnerRepoForbidden with default headers values
func NewPutUserSubscriptionsOwnerRepoForbidden() *PutUserSubscriptionsOwnerRepoForbidden {
	return &PutUserSubscriptionsOwnerRepoForbidden{}
}

// WriteResponse to the client
func (o *PutUserSubscriptionsOwnerRepoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}