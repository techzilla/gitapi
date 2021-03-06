// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteTeamsTeamIDMembershipsUsernameNoContentCode is the HTTP code returned for type DeleteTeamsTeamIDMembershipsUsernameNoContent
const DeleteTeamsTeamIDMembershipsUsernameNoContentCode int = 204

/*DeleteTeamsTeamIDMembershipsUsernameNoContent Team member removed.

swagger:response deleteTeamsTeamIdMembershipsUsernameNoContent
*/
type DeleteTeamsTeamIDMembershipsUsernameNoContent struct {
}

// NewDeleteTeamsTeamIDMembershipsUsernameNoContent creates DeleteTeamsTeamIDMembershipsUsernameNoContent with default headers values
func NewDeleteTeamsTeamIDMembershipsUsernameNoContent() *DeleteTeamsTeamIDMembershipsUsernameNoContent {
	return &DeleteTeamsTeamIDMembershipsUsernameNoContent{}
}

// WriteResponse to the client
func (o *DeleteTeamsTeamIDMembershipsUsernameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTeamsTeamIDMembershipsUsernameForbiddenCode is the HTTP code returned for type DeleteTeamsTeamIDMembershipsUsernameForbidden
const DeleteTeamsTeamIDMembershipsUsernameForbiddenCode int = 403

/*DeleteTeamsTeamIDMembershipsUsernameForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response deleteTeamsTeamIdMembershipsUsernameForbidden
*/
type DeleteTeamsTeamIDMembershipsUsernameForbidden struct {
}

// NewDeleteTeamsTeamIDMembershipsUsernameForbidden creates DeleteTeamsTeamIDMembershipsUsernameForbidden with default headers values
func NewDeleteTeamsTeamIDMembershipsUsernameForbidden() *DeleteTeamsTeamIDMembershipsUsernameForbidden {
	return &DeleteTeamsTeamIDMembershipsUsernameForbidden{}
}

// WriteResponse to the client
func (o *DeleteTeamsTeamIDMembershipsUsernameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
