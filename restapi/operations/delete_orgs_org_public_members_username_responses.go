// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteOrgsOrgPublicMembersUsernameNoContentCode is the HTTP code returned for type DeleteOrgsOrgPublicMembersUsernameNoContent
const DeleteOrgsOrgPublicMembersUsernameNoContentCode int = 204

/*DeleteOrgsOrgPublicMembersUsernameNoContent Concealed.

swagger:response deleteOrgsOrgPublicMembersUsernameNoContent
*/
type DeleteOrgsOrgPublicMembersUsernameNoContent struct {
}

// NewDeleteOrgsOrgPublicMembersUsernameNoContent creates DeleteOrgsOrgPublicMembersUsernameNoContent with default headers values
func NewDeleteOrgsOrgPublicMembersUsernameNoContent() *DeleteOrgsOrgPublicMembersUsernameNoContent {
	return &DeleteOrgsOrgPublicMembersUsernameNoContent{}
}

// WriteResponse to the client
func (o *DeleteOrgsOrgPublicMembersUsernameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteOrgsOrgPublicMembersUsernameForbiddenCode is the HTTP code returned for type DeleteOrgsOrgPublicMembersUsernameForbidden
const DeleteOrgsOrgPublicMembersUsernameForbiddenCode int = 403

/*DeleteOrgsOrgPublicMembersUsernameForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response deleteOrgsOrgPublicMembersUsernameForbidden
*/
type DeleteOrgsOrgPublicMembersUsernameForbidden struct {
}

// NewDeleteOrgsOrgPublicMembersUsernameForbidden creates DeleteOrgsOrgPublicMembersUsernameForbidden with default headers values
func NewDeleteOrgsOrgPublicMembersUsernameForbidden() *DeleteOrgsOrgPublicMembersUsernameForbidden {
	return &DeleteOrgsOrgPublicMembersUsernameForbidden{}
}

// WriteResponse to the client
func (o *DeleteOrgsOrgPublicMembersUsernameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}