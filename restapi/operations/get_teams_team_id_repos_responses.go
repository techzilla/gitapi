// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetTeamsTeamIDReposOKCode is the HTTP code returned for type GetTeamsTeamIDReposOK
const GetTeamsTeamIDReposOKCode int = 200

/*GetTeamsTeamIDReposOK OK

swagger:response getTeamsTeamIdReposOK
*/
type GetTeamsTeamIDReposOK struct {

	/*
	  In: Body
	*/
	Payload models.TeamRepos `json:"body,omitempty"`
}

// NewGetTeamsTeamIDReposOK creates GetTeamsTeamIDReposOK with default headers values
func NewGetTeamsTeamIDReposOK() *GetTeamsTeamIDReposOK {
	return &GetTeamsTeamIDReposOK{}
}

// WithPayload adds the payload to the get teams team Id repos o k response
func (o *GetTeamsTeamIDReposOK) WithPayload(payload models.TeamRepos) *GetTeamsTeamIDReposOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams team Id repos o k response
func (o *GetTeamsTeamIDReposOK) SetPayload(payload models.TeamRepos) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsTeamIDReposOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.TeamRepos, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTeamsTeamIDReposForbiddenCode is the HTTP code returned for type GetTeamsTeamIDReposForbidden
const GetTeamsTeamIDReposForbiddenCode int = 403

/*GetTeamsTeamIDReposForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getTeamsTeamIdReposForbidden
*/
type GetTeamsTeamIDReposForbidden struct {
}

// NewGetTeamsTeamIDReposForbidden creates GetTeamsTeamIDReposForbidden with default headers values
func NewGetTeamsTeamIDReposForbidden() *GetTeamsTeamIDReposForbidden {
	return &GetTeamsTeamIDReposForbidden{}
}

// WriteResponse to the client
func (o *GetTeamsTeamIDReposForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
