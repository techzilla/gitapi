// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostReposOwnerRepoMilestonesCreatedCode is the HTTP code returned for type PostReposOwnerRepoMilestonesCreated
const PostReposOwnerRepoMilestonesCreatedCode int = 201

/*PostReposOwnerRepoMilestonesCreated Created

swagger:response postReposOwnerRepoMilestonesCreated
*/
type PostReposOwnerRepoMilestonesCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Milestone `json:"body,omitempty"`
}

// NewPostReposOwnerRepoMilestonesCreated creates PostReposOwnerRepoMilestonesCreated with default headers values
func NewPostReposOwnerRepoMilestonesCreated() *PostReposOwnerRepoMilestonesCreated {
	return &PostReposOwnerRepoMilestonesCreated{}
}

// WithPayload adds the payload to the post repos owner repo milestones created response
func (o *PostReposOwnerRepoMilestonesCreated) WithPayload(payload *models.Milestone) *PostReposOwnerRepoMilestonesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post repos owner repo milestones created response
func (o *PostReposOwnerRepoMilestonesCreated) SetPayload(payload *models.Milestone) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReposOwnerRepoMilestonesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReposOwnerRepoMilestonesForbiddenCode is the HTTP code returned for type PostReposOwnerRepoMilestonesForbidden
const PostReposOwnerRepoMilestonesForbiddenCode int = 403

/*PostReposOwnerRepoMilestonesForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postReposOwnerRepoMilestonesForbidden
*/
type PostReposOwnerRepoMilestonesForbidden struct {
}

// NewPostReposOwnerRepoMilestonesForbidden creates PostReposOwnerRepoMilestonesForbidden with default headers values
func NewPostReposOwnerRepoMilestonesForbidden() *PostReposOwnerRepoMilestonesForbidden {
	return &PostReposOwnerRepoMilestonesForbidden{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoMilestonesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
