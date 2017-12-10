// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetUserTeamsOKCode is the HTTP code returned for type GetUserTeamsOK
const GetUserTeamsOKCode int = 200

/*GetUserTeamsOK OK

swagger:response getUserTeamsOK
*/
type GetUserTeamsOK struct {

	/*
	  In: Body
	*/
	Payload models.TeamsList `json:"body,omitempty"`
}

// NewGetUserTeamsOK creates GetUserTeamsOK with default headers values
func NewGetUserTeamsOK() *GetUserTeamsOK {
	return &GetUserTeamsOK{}
}

// WithPayload adds the payload to the get user teams o k response
func (o *GetUserTeamsOK) WithPayload(payload models.TeamsList) *GetUserTeamsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user teams o k response
func (o *GetUserTeamsOK) SetPayload(payload models.TeamsList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserTeamsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.TeamsList, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUserTeamsForbiddenCode is the HTTP code returned for type GetUserTeamsForbidden
const GetUserTeamsForbiddenCode int = 403

/*GetUserTeamsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUserTeamsForbidden
*/
type GetUserTeamsForbidden struct {
}

// NewGetUserTeamsForbidden creates GetUserTeamsForbidden with default headers values
func NewGetUserTeamsForbidden() *GetUserTeamsForbidden {
	return &GetUserTeamsForbidden{}
}

// WriteResponse to the client
func (o *GetUserTeamsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
