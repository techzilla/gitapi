// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostReposOwnerRepoIssuesCreatedCode is the HTTP code returned for type PostReposOwnerRepoIssuesCreated
const PostReposOwnerRepoIssuesCreatedCode int = 201

/*PostReposOwnerRepoIssuesCreated Created

swagger:response postReposOwnerRepoIssuesCreated
*/
type PostReposOwnerRepoIssuesCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Issue `json:"body,omitempty"`
}

// NewPostReposOwnerRepoIssuesCreated creates PostReposOwnerRepoIssuesCreated with default headers values
func NewPostReposOwnerRepoIssuesCreated() *PostReposOwnerRepoIssuesCreated {
	return &PostReposOwnerRepoIssuesCreated{}
}

// WithPayload adds the payload to the post repos owner repo issues created response
func (o *PostReposOwnerRepoIssuesCreated) WithPayload(payload *models.Issue) *PostReposOwnerRepoIssuesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post repos owner repo issues created response
func (o *PostReposOwnerRepoIssuesCreated) SetPayload(payload *models.Issue) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReposOwnerRepoIssuesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReposOwnerRepoIssuesForbiddenCode is the HTTP code returned for type PostReposOwnerRepoIssuesForbidden
const PostReposOwnerRepoIssuesForbiddenCode int = 403

/*PostReposOwnerRepoIssuesForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postReposOwnerRepoIssuesForbidden
*/
type PostReposOwnerRepoIssuesForbidden struct {
}

// NewPostReposOwnerRepoIssuesForbidden creates PostReposOwnerRepoIssuesForbidden with default headers values
func NewPostReposOwnerRepoIssuesForbidden() *PostReposOwnerRepoIssuesForbidden {
	return &PostReposOwnerRepoIssuesForbidden{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoIssuesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}