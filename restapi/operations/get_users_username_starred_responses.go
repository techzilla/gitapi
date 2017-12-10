// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUsersUsernameStarredForbiddenCode is the HTTP code returned for type GetUsersUsernameStarredForbidden
const GetUsersUsernameStarredForbiddenCode int = 403

/*GetUsersUsernameStarredForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUsersUsernameStarredForbidden
*/
type GetUsersUsernameStarredForbidden struct {
}

// NewGetUsersUsernameStarredForbidden creates GetUsersUsernameStarredForbidden with default headers values
func NewGetUsersUsernameStarredForbidden() *GetUsersUsernameStarredForbidden {
	return &GetUsersUsernameStarredForbidden{}
}

// WriteResponse to the client
func (o *GetUsersUsernameStarredForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}