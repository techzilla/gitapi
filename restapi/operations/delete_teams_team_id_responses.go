// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteTeamsTeamIDNoContentCode is the HTTP code returned for type DeleteTeamsTeamIDNoContent
const DeleteTeamsTeamIDNoContentCode int = 204

/*DeleteTeamsTeamIDNoContent No content.


swagger:response deleteTeamsTeamIdNoContent
*/
type DeleteTeamsTeamIDNoContent struct {
}

// NewDeleteTeamsTeamIDNoContent creates DeleteTeamsTeamIDNoContent with default headers values
func NewDeleteTeamsTeamIDNoContent() *DeleteTeamsTeamIDNoContent {
	return &DeleteTeamsTeamIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteTeamsTeamIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTeamsTeamIDForbiddenCode is the HTTP code returned for type DeleteTeamsTeamIDForbidden
const DeleteTeamsTeamIDForbiddenCode int = 403

/*DeleteTeamsTeamIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response deleteTeamsTeamIdForbidden
*/
type DeleteTeamsTeamIDForbidden struct {
}

// NewDeleteTeamsTeamIDForbidden creates DeleteTeamsTeamIDForbidden with default headers values
func NewDeleteTeamsTeamIDForbidden() *DeleteTeamsTeamIDForbidden {
	return &DeleteTeamsTeamIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteTeamsTeamIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}