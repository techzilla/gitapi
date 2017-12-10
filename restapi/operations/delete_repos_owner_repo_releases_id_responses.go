// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteReposOwnerRepoReleasesIDNoContentCode is the HTTP code returned for type DeleteReposOwnerRepoReleasesIDNoContent
const DeleteReposOwnerRepoReleasesIDNoContentCode int = 204

/*DeleteReposOwnerRepoReleasesIDNoContent No Content

swagger:response deleteReposOwnerRepoReleasesIdNoContent
*/
type DeleteReposOwnerRepoReleasesIDNoContent struct {
}

// NewDeleteReposOwnerRepoReleasesIDNoContent creates DeleteReposOwnerRepoReleasesIDNoContent with default headers values
func NewDeleteReposOwnerRepoReleasesIDNoContent() *DeleteReposOwnerRepoReleasesIDNoContent {
	return &DeleteReposOwnerRepoReleasesIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteReposOwnerRepoReleasesIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteReposOwnerRepoReleasesIDForbiddenCode is the HTTP code returned for type DeleteReposOwnerRepoReleasesIDForbidden
const DeleteReposOwnerRepoReleasesIDForbiddenCode int = 403

/*DeleteReposOwnerRepoReleasesIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response deleteReposOwnerRepoReleasesIdForbidden
*/
type DeleteReposOwnerRepoReleasesIDForbidden struct {
}

// NewDeleteReposOwnerRepoReleasesIDForbidden creates DeleteReposOwnerRepoReleasesIDForbidden with default headers values
func NewDeleteReposOwnerRepoReleasesIDForbidden() *DeleteReposOwnerRepoReleasesIDForbidden {
	return &DeleteReposOwnerRepoReleasesIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteReposOwnerRepoReleasesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
