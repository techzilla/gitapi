// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostReposOwnerRepoDeploymentsIDStatusesCreatedCode is the HTTP code returned for type PostReposOwnerRepoDeploymentsIDStatusesCreated
const PostReposOwnerRepoDeploymentsIDStatusesCreatedCode int = 201

/*PostReposOwnerRepoDeploymentsIDStatusesCreated ok

swagger:response postReposOwnerRepoDeploymentsIdStatusesCreated
*/
type PostReposOwnerRepoDeploymentsIDStatusesCreated struct {
}

// NewPostReposOwnerRepoDeploymentsIDStatusesCreated creates PostReposOwnerRepoDeploymentsIDStatusesCreated with default headers values
func NewPostReposOwnerRepoDeploymentsIDStatusesCreated() *PostReposOwnerRepoDeploymentsIDStatusesCreated {
	return &PostReposOwnerRepoDeploymentsIDStatusesCreated{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoDeploymentsIDStatusesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostReposOwnerRepoDeploymentsIDStatusesForbiddenCode is the HTTP code returned for type PostReposOwnerRepoDeploymentsIDStatusesForbidden
const PostReposOwnerRepoDeploymentsIDStatusesForbiddenCode int = 403

/*PostReposOwnerRepoDeploymentsIDStatusesForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postReposOwnerRepoDeploymentsIdStatusesForbidden
*/
type PostReposOwnerRepoDeploymentsIDStatusesForbidden struct {
}

// NewPostReposOwnerRepoDeploymentsIDStatusesForbidden creates PostReposOwnerRepoDeploymentsIDStatusesForbidden with default headers values
func NewPostReposOwnerRepoDeploymentsIDStatusesForbidden() *PostReposOwnerRepoDeploymentsIDStatusesForbidden {
	return &PostReposOwnerRepoDeploymentsIDStatusesForbidden{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoDeploymentsIDStatusesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}