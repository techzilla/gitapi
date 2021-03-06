// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostReposOwnerRepoIssuesNumberCommentsCreatedCode is the HTTP code returned for type PostReposOwnerRepoIssuesNumberCommentsCreated
const PostReposOwnerRepoIssuesNumberCommentsCreatedCode int = 201

/*PostReposOwnerRepoIssuesNumberCommentsCreated Created

swagger:response postReposOwnerRepoIssuesNumberCommentsCreated
*/
type PostReposOwnerRepoIssuesNumberCommentsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IssuesComment `json:"body,omitempty"`
}

// NewPostReposOwnerRepoIssuesNumberCommentsCreated creates PostReposOwnerRepoIssuesNumberCommentsCreated with default headers values
func NewPostReposOwnerRepoIssuesNumberCommentsCreated() *PostReposOwnerRepoIssuesNumberCommentsCreated {
	return &PostReposOwnerRepoIssuesNumberCommentsCreated{}
}

// WithPayload adds the payload to the post repos owner repo issues number comments created response
func (o *PostReposOwnerRepoIssuesNumberCommentsCreated) WithPayload(payload *models.IssuesComment) *PostReposOwnerRepoIssuesNumberCommentsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post repos owner repo issues number comments created response
func (o *PostReposOwnerRepoIssuesNumberCommentsCreated) SetPayload(payload *models.IssuesComment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReposOwnerRepoIssuesNumberCommentsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReposOwnerRepoIssuesNumberCommentsForbiddenCode is the HTTP code returned for type PostReposOwnerRepoIssuesNumberCommentsForbidden
const PostReposOwnerRepoIssuesNumberCommentsForbiddenCode int = 403

/*PostReposOwnerRepoIssuesNumberCommentsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postReposOwnerRepoIssuesNumberCommentsForbidden
*/
type PostReposOwnerRepoIssuesNumberCommentsForbidden struct {
}

// NewPostReposOwnerRepoIssuesNumberCommentsForbidden creates PostReposOwnerRepoIssuesNumberCommentsForbidden with default headers values
func NewPostReposOwnerRepoIssuesNumberCommentsForbidden() *PostReposOwnerRepoIssuesNumberCommentsForbidden {
	return &PostReposOwnerRepoIssuesNumberCommentsForbidden{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoIssuesNumberCommentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
