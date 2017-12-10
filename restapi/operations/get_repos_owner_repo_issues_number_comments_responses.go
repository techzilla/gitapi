// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoIssuesNumberCommentsOKCode is the HTTP code returned for type GetReposOwnerRepoIssuesNumberCommentsOK
const GetReposOwnerRepoIssuesNumberCommentsOKCode int = 200

/*GetReposOwnerRepoIssuesNumberCommentsOK OK

swagger:response getReposOwnerRepoIssuesNumberCommentsOK
*/
type GetReposOwnerRepoIssuesNumberCommentsOK struct {

	/*
	  In: Body
	*/
	Payload models.IssuesComments `json:"body,omitempty"`
}

// NewGetReposOwnerRepoIssuesNumberCommentsOK creates GetReposOwnerRepoIssuesNumberCommentsOK with default headers values
func NewGetReposOwnerRepoIssuesNumberCommentsOK() *GetReposOwnerRepoIssuesNumberCommentsOK {
	return &GetReposOwnerRepoIssuesNumberCommentsOK{}
}

// WithPayload adds the payload to the get repos owner repo issues number comments o k response
func (o *GetReposOwnerRepoIssuesNumberCommentsOK) WithPayload(payload models.IssuesComments) *GetReposOwnerRepoIssuesNumberCommentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo issues number comments o k response
func (o *GetReposOwnerRepoIssuesNumberCommentsOK) SetPayload(payload models.IssuesComments) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesNumberCommentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.IssuesComments, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReposOwnerRepoIssuesNumberCommentsForbiddenCode is the HTTP code returned for type GetReposOwnerRepoIssuesNumberCommentsForbidden
const GetReposOwnerRepoIssuesNumberCommentsForbiddenCode int = 403

/*GetReposOwnerRepoIssuesNumberCommentsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoIssuesNumberCommentsForbidden
*/
type GetReposOwnerRepoIssuesNumberCommentsForbidden struct {
}

// NewGetReposOwnerRepoIssuesNumberCommentsForbidden creates GetReposOwnerRepoIssuesNumberCommentsForbidden with default headers values
func NewGetReposOwnerRepoIssuesNumberCommentsForbidden() *GetReposOwnerRepoIssuesNumberCommentsForbidden {
	return &GetReposOwnerRepoIssuesNumberCommentsForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesNumberCommentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}