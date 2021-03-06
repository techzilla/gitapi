// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUsersUsernameReceivedEventsPublicForbiddenCode is the HTTP code returned for type GetUsersUsernameReceivedEventsPublicForbidden
const GetUsersUsernameReceivedEventsPublicForbiddenCode int = 403

/*GetUsersUsernameReceivedEventsPublicForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUsersUsernameReceivedEventsPublicForbidden
*/
type GetUsersUsernameReceivedEventsPublicForbidden struct {
}

// NewGetUsersUsernameReceivedEventsPublicForbidden creates GetUsersUsernameReceivedEventsPublicForbidden with default headers values
func NewGetUsersUsernameReceivedEventsPublicForbidden() *GetUsersUsernameReceivedEventsPublicForbidden {
	return &GetUsersUsernameReceivedEventsPublicForbidden{}
}

// WriteResponse to the client
func (o *GetUsersUsernameReceivedEventsPublicForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
