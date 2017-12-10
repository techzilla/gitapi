// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoCommentsCommentIDOKCode is the HTTP code returned for type GetReposOwnerRepoCommentsCommentIDOK
const GetReposOwnerRepoCommentsCommentIDOKCode int = 200

/*GetReposOwnerRepoCommentsCommentIDOK OK

swagger:response getReposOwnerRepoCommentsCommentIdOK
*/
type GetReposOwnerRepoCommentsCommentIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CommitComments `json:"body,omitempty"`
}

// NewGetReposOwnerRepoCommentsCommentIDOK creates GetReposOwnerRepoCommentsCommentIDOK with default headers values
func NewGetReposOwnerRepoCommentsCommentIDOK() *GetReposOwnerRepoCommentsCommentIDOK {
	return &GetReposOwnerRepoCommentsCommentIDOK{}
}

// WithPayload adds the payload to the get repos owner repo comments comment Id o k response
func (o *GetReposOwnerRepoCommentsCommentIDOK) WithPayload(payload *models.CommitComments) *GetReposOwnerRepoCommentsCommentIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo comments comment Id o k response
func (o *GetReposOwnerRepoCommentsCommentIDOK) SetPayload(payload *models.CommitComments) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoCommentsCommentIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoCommentsCommentIDForbiddenCode is the HTTP code returned for type GetReposOwnerRepoCommentsCommentIDForbidden
const GetReposOwnerRepoCommentsCommentIDForbiddenCode int = 403

/*GetReposOwnerRepoCommentsCommentIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoCommentsCommentIdForbidden
*/
type GetReposOwnerRepoCommentsCommentIDForbidden struct {
}

// NewGetReposOwnerRepoCommentsCommentIDForbidden creates GetReposOwnerRepoCommentsCommentIDForbidden with default headers values
func NewGetReposOwnerRepoCommentsCommentIDForbidden() *GetReposOwnerRepoCommentsCommentIDForbidden {
	return &GetReposOwnerRepoCommentsCommentIDForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoCommentsCommentIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}