// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoPullsNumberCommitsOKCode is the HTTP code returned for type GetReposOwnerRepoPullsNumberCommitsOK
const GetReposOwnerRepoPullsNumberCommitsOKCode int = 200

/*GetReposOwnerRepoPullsNumberCommitsOK OK

swagger:response getReposOwnerRepoPullsNumberCommitsOK
*/
type GetReposOwnerRepoPullsNumberCommitsOK struct {

	/*
	  In: Body
	*/
	Payload models.Commits `json:"body,omitempty"`
}

// NewGetReposOwnerRepoPullsNumberCommitsOK creates GetReposOwnerRepoPullsNumberCommitsOK with default headers values
func NewGetReposOwnerRepoPullsNumberCommitsOK() *GetReposOwnerRepoPullsNumberCommitsOK {
	return &GetReposOwnerRepoPullsNumberCommitsOK{}
}

// WithPayload adds the payload to the get repos owner repo pulls number commits o k response
func (o *GetReposOwnerRepoPullsNumberCommitsOK) WithPayload(payload models.Commits) *GetReposOwnerRepoPullsNumberCommitsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo pulls number commits o k response
func (o *GetReposOwnerRepoPullsNumberCommitsOK) SetPayload(payload models.Commits) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoPullsNumberCommitsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Commits, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReposOwnerRepoPullsNumberCommitsForbiddenCode is the HTTP code returned for type GetReposOwnerRepoPullsNumberCommitsForbidden
const GetReposOwnerRepoPullsNumberCommitsForbiddenCode int = 403

/*GetReposOwnerRepoPullsNumberCommitsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoPullsNumberCommitsForbidden
*/
type GetReposOwnerRepoPullsNumberCommitsForbidden struct {
}

// NewGetReposOwnerRepoPullsNumberCommitsForbidden creates GetReposOwnerRepoPullsNumberCommitsForbidden with default headers values
func NewGetReposOwnerRepoPullsNumberCommitsForbidden() *GetReposOwnerRepoPullsNumberCommitsForbidden {
	return &GetReposOwnerRepoPullsNumberCommitsForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoPullsNumberCommitsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
