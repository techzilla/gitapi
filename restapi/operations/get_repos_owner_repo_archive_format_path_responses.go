// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetReposOwnerRepoArchiveFormatPathFoundCode is the HTTP code returned for type GetReposOwnerRepoArchiveFormatPathFound
const GetReposOwnerRepoArchiveFormatPathFoundCode int = 302

/*GetReposOwnerRepoArchiveFormatPathFound Found.

swagger:response getReposOwnerRepoArchiveFormatPathFound
*/
type GetReposOwnerRepoArchiveFormatPathFound struct {
}

// NewGetReposOwnerRepoArchiveFormatPathFound creates GetReposOwnerRepoArchiveFormatPathFound with default headers values
func NewGetReposOwnerRepoArchiveFormatPathFound() *GetReposOwnerRepoArchiveFormatPathFound {
	return &GetReposOwnerRepoArchiveFormatPathFound{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoArchiveFormatPathFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(302)
}

// GetReposOwnerRepoArchiveFormatPathForbiddenCode is the HTTP code returned for type GetReposOwnerRepoArchiveFormatPathForbidden
const GetReposOwnerRepoArchiveFormatPathForbiddenCode int = 403

/*GetReposOwnerRepoArchiveFormatPathForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoArchiveFormatPathForbidden
*/
type GetReposOwnerRepoArchiveFormatPathForbidden struct {
}

// NewGetReposOwnerRepoArchiveFormatPathForbidden creates GetReposOwnerRepoArchiveFormatPathForbidden with default headers values
func NewGetReposOwnerRepoArchiveFormatPathForbidden() *GetReposOwnerRepoArchiveFormatPathForbidden {
	return &GetReposOwnerRepoArchiveFormatPathForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoArchiveFormatPathForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
