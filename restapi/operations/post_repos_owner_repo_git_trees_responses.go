// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostReposOwnerRepoGitTreesCreatedCode is the HTTP code returned for type PostReposOwnerRepoGitTreesCreated
const PostReposOwnerRepoGitTreesCreatedCode int = 201

/*PostReposOwnerRepoGitTreesCreated Created

swagger:response postReposOwnerRepoGitTreesCreated
*/
type PostReposOwnerRepoGitTreesCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Trees `json:"body,omitempty"`
}

// NewPostReposOwnerRepoGitTreesCreated creates PostReposOwnerRepoGitTreesCreated with default headers values
func NewPostReposOwnerRepoGitTreesCreated() *PostReposOwnerRepoGitTreesCreated {
	return &PostReposOwnerRepoGitTreesCreated{}
}

// WithPayload adds the payload to the post repos owner repo git trees created response
func (o *PostReposOwnerRepoGitTreesCreated) WithPayload(payload *models.Trees) *PostReposOwnerRepoGitTreesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post repos owner repo git trees created response
func (o *PostReposOwnerRepoGitTreesCreated) SetPayload(payload *models.Trees) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReposOwnerRepoGitTreesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReposOwnerRepoGitTreesForbiddenCode is the HTTP code returned for type PostReposOwnerRepoGitTreesForbidden
const PostReposOwnerRepoGitTreesForbiddenCode int = 403

/*PostReposOwnerRepoGitTreesForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postReposOwnerRepoGitTreesForbidden
*/
type PostReposOwnerRepoGitTreesForbidden struct {
}

// NewPostReposOwnerRepoGitTreesForbidden creates PostReposOwnerRepoGitTreesForbidden with default headers values
func NewPostReposOwnerRepoGitTreesForbidden() *PostReposOwnerRepoGitTreesForbidden {
	return &PostReposOwnerRepoGitTreesForbidden{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoGitTreesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}