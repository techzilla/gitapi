// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoGitBlobsShaCodeOKCode is the HTTP code returned for type GetReposOwnerRepoGitBlobsShaCodeOK
const GetReposOwnerRepoGitBlobsShaCodeOKCode int = 200

/*GetReposOwnerRepoGitBlobsShaCodeOK OK

swagger:response getReposOwnerRepoGitBlobsShaCodeOK
*/
type GetReposOwnerRepoGitBlobsShaCodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Blob `json:"body,omitempty"`
}

// NewGetReposOwnerRepoGitBlobsShaCodeOK creates GetReposOwnerRepoGitBlobsShaCodeOK with default headers values
func NewGetReposOwnerRepoGitBlobsShaCodeOK() *GetReposOwnerRepoGitBlobsShaCodeOK {
	return &GetReposOwnerRepoGitBlobsShaCodeOK{}
}

// WithPayload adds the payload to the get repos owner repo git blobs sha code o k response
func (o *GetReposOwnerRepoGitBlobsShaCodeOK) WithPayload(payload *models.Blob) *GetReposOwnerRepoGitBlobsShaCodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo git blobs sha code o k response
func (o *GetReposOwnerRepoGitBlobsShaCodeOK) SetPayload(payload *models.Blob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoGitBlobsShaCodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoGitBlobsShaCodeForbiddenCode is the HTTP code returned for type GetReposOwnerRepoGitBlobsShaCodeForbidden
const GetReposOwnerRepoGitBlobsShaCodeForbiddenCode int = 403

/*GetReposOwnerRepoGitBlobsShaCodeForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoGitBlobsShaCodeForbidden
*/
type GetReposOwnerRepoGitBlobsShaCodeForbidden struct {
}

// NewGetReposOwnerRepoGitBlobsShaCodeForbidden creates GetReposOwnerRepoGitBlobsShaCodeForbidden with default headers values
func NewGetReposOwnerRepoGitBlobsShaCodeForbidden() *GetReposOwnerRepoGitBlobsShaCodeForbidden {
	return &GetReposOwnerRepoGitBlobsShaCodeForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoGitBlobsShaCodeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}