// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PutTeamsTeamIDMembersUsernameNoContentCode is the HTTP code returned for type PutTeamsTeamIDMembersUsernameNoContent
const PutTeamsTeamIDMembersUsernameNoContentCode int = 204

/*PutTeamsTeamIDMembersUsernameNoContent Team member added.

swagger:response putTeamsTeamIdMembersUsernameNoContent
*/
type PutTeamsTeamIDMembersUsernameNoContent struct {
}

// NewPutTeamsTeamIDMembersUsernameNoContent creates PutTeamsTeamIDMembersUsernameNoContent with default headers values
func NewPutTeamsTeamIDMembersUsernameNoContent() *PutTeamsTeamIDMembersUsernameNoContent {
	return &PutTeamsTeamIDMembersUsernameNoContent{}
}

// WriteResponse to the client
func (o *PutTeamsTeamIDMembersUsernameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PutTeamsTeamIDMembersUsernameForbiddenCode is the HTTP code returned for type PutTeamsTeamIDMembersUsernameForbidden
const PutTeamsTeamIDMembersUsernameForbiddenCode int = 403

/*PutTeamsTeamIDMembersUsernameForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response putTeamsTeamIdMembersUsernameForbidden
*/
type PutTeamsTeamIDMembersUsernameForbidden struct {
}

// NewPutTeamsTeamIDMembersUsernameForbidden creates PutTeamsTeamIDMembersUsernameForbidden with default headers values
func NewPutTeamsTeamIDMembersUsernameForbidden() *PutTeamsTeamIDMembersUsernameForbidden {
	return &PutTeamsTeamIDMembersUsernameForbidden{}
}

// WriteResponse to the client
func (o *PutTeamsTeamIDMembersUsernameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PutTeamsTeamIDMembersUsernameUnprocessableEntityCode is the HTTP code returned for type PutTeamsTeamIDMembersUsernameUnprocessableEntity
const PutTeamsTeamIDMembersUsernameUnprocessableEntityCode int = 422

/*PutTeamsTeamIDMembersUsernameUnprocessableEntity If you attempt to add an organization to a team, you will get this.

swagger:response putTeamsTeamIdMembersUsernameUnprocessableEntity
*/
type PutTeamsTeamIDMembersUsernameUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.OrganizationAsTeamMember `json:"body,omitempty"`
}

// NewPutTeamsTeamIDMembersUsernameUnprocessableEntity creates PutTeamsTeamIDMembersUsernameUnprocessableEntity with default headers values
func NewPutTeamsTeamIDMembersUsernameUnprocessableEntity() *PutTeamsTeamIDMembersUsernameUnprocessableEntity {
	return &PutTeamsTeamIDMembersUsernameUnprocessableEntity{}
}

// WithPayload adds the payload to the put teams team Id members username unprocessable entity response
func (o *PutTeamsTeamIDMembersUsernameUnprocessableEntity) WithPayload(payload *models.OrganizationAsTeamMember) *PutTeamsTeamIDMembersUsernameUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put teams team Id members username unprocessable entity response
func (o *PutTeamsTeamIDMembersUsernameUnprocessableEntity) SetPayload(payload *models.OrganizationAsTeamMember) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTeamsTeamIDMembersUsernameUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
