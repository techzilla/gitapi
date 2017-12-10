// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostReposOwnerRepoDeploymentsCreatedCode is the HTTP code returned for type PostReposOwnerRepoDeploymentsCreated
const PostReposOwnerRepoDeploymentsCreatedCode int = 201

/*PostReposOwnerRepoDeploymentsCreated Created

swagger:response postReposOwnerRepoDeploymentsCreated
*/
type PostReposOwnerRepoDeploymentsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DeploymentResp `json:"body,omitempty"`
}

// NewPostReposOwnerRepoDeploymentsCreated creates PostReposOwnerRepoDeploymentsCreated with default headers values
func NewPostReposOwnerRepoDeploymentsCreated() *PostReposOwnerRepoDeploymentsCreated {
	return &PostReposOwnerRepoDeploymentsCreated{}
}

// WithPayload adds the payload to the post repos owner repo deployments created response
func (o *PostReposOwnerRepoDeploymentsCreated) WithPayload(payload *models.DeploymentResp) *PostReposOwnerRepoDeploymentsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post repos owner repo deployments created response
func (o *PostReposOwnerRepoDeploymentsCreated) SetPayload(payload *models.DeploymentResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostReposOwnerRepoDeploymentsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostReposOwnerRepoDeploymentsForbiddenCode is the HTTP code returned for type PostReposOwnerRepoDeploymentsForbidden
const PostReposOwnerRepoDeploymentsForbiddenCode int = 403

/*PostReposOwnerRepoDeploymentsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postReposOwnerRepoDeploymentsForbidden
*/
type PostReposOwnerRepoDeploymentsForbidden struct {
}

// NewPostReposOwnerRepoDeploymentsForbidden creates PostReposOwnerRepoDeploymentsForbidden with default headers values
func NewPostReposOwnerRepoDeploymentsForbidden() *PostReposOwnerRepoDeploymentsForbidden {
	return &PostReposOwnerRepoDeploymentsForbidden{}
}

// WriteResponse to the client
func (o *PostReposOwnerRepoDeploymentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
