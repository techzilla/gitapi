// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoIssuesCommentsCommentIDOKCode is the HTTP code returned for type GetReposOwnerRepoIssuesCommentsCommentIDOK
const GetReposOwnerRepoIssuesCommentsCommentIDOKCode int = 200

/*GetReposOwnerRepoIssuesCommentsCommentIDOK OK

swagger:response getReposOwnerRepoIssuesCommentsCommentIdOK
*/
type GetReposOwnerRepoIssuesCommentsCommentIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.IssuesComment `json:"body,omitempty"`
}

// NewGetReposOwnerRepoIssuesCommentsCommentIDOK creates GetReposOwnerRepoIssuesCommentsCommentIDOK with default headers values
func NewGetReposOwnerRepoIssuesCommentsCommentIDOK() *GetReposOwnerRepoIssuesCommentsCommentIDOK {
	return &GetReposOwnerRepoIssuesCommentsCommentIDOK{}
}

// WithPayload adds the payload to the get repos owner repo issues comments comment Id o k response
func (o *GetReposOwnerRepoIssuesCommentsCommentIDOK) WithPayload(payload *models.IssuesComment) *GetReposOwnerRepoIssuesCommentsCommentIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo issues comments comment Id o k response
func (o *GetReposOwnerRepoIssuesCommentsCommentIDOK) SetPayload(payload *models.IssuesComment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesCommentsCommentIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoIssuesCommentsCommentIDForbiddenCode is the HTTP code returned for type GetReposOwnerRepoIssuesCommentsCommentIDForbidden
const GetReposOwnerRepoIssuesCommentsCommentIDForbiddenCode int = 403

/*GetReposOwnerRepoIssuesCommentsCommentIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoIssuesCommentsCommentIdForbidden
*/
type GetReposOwnerRepoIssuesCommentsCommentIDForbidden struct {
}

// NewGetReposOwnerRepoIssuesCommentsCommentIDForbidden creates GetReposOwnerRepoIssuesCommentsCommentIDForbidden with default headers values
func NewGetReposOwnerRepoIssuesCommentsCommentIDForbidden() *GetReposOwnerRepoIssuesCommentsCommentIDForbidden {
	return &GetReposOwnerRepoIssuesCommentsCommentIDForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesCommentsCommentIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
