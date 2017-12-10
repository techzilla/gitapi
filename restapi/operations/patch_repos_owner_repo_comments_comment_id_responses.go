// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PatchReposOwnerRepoCommentsCommentIDOKCode is the HTTP code returned for type PatchReposOwnerRepoCommentsCommentIDOK
const PatchReposOwnerRepoCommentsCommentIDOKCode int = 200

/*PatchReposOwnerRepoCommentsCommentIDOK OK

swagger:response patchReposOwnerRepoCommentsCommentIdOK
*/
type PatchReposOwnerRepoCommentsCommentIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CommitComments `json:"body,omitempty"`
}

// NewPatchReposOwnerRepoCommentsCommentIDOK creates PatchReposOwnerRepoCommentsCommentIDOK with default headers values
func NewPatchReposOwnerRepoCommentsCommentIDOK() *PatchReposOwnerRepoCommentsCommentIDOK {
	return &PatchReposOwnerRepoCommentsCommentIDOK{}
}

// WithPayload adds the payload to the patch repos owner repo comments comment Id o k response
func (o *PatchReposOwnerRepoCommentsCommentIDOK) WithPayload(payload *models.CommitComments) *PatchReposOwnerRepoCommentsCommentIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch repos owner repo comments comment Id o k response
func (o *PatchReposOwnerRepoCommentsCommentIDOK) SetPayload(payload *models.CommitComments) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoCommentsCommentIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchReposOwnerRepoCommentsCommentIDForbiddenCode is the HTTP code returned for type PatchReposOwnerRepoCommentsCommentIDForbidden
const PatchReposOwnerRepoCommentsCommentIDForbiddenCode int = 403

/*PatchReposOwnerRepoCommentsCommentIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response patchReposOwnerRepoCommentsCommentIdForbidden
*/
type PatchReposOwnerRepoCommentsCommentIDForbidden struct {
}

// NewPatchReposOwnerRepoCommentsCommentIDForbidden creates PatchReposOwnerRepoCommentsCommentIDForbidden with default headers values
func NewPatchReposOwnerRepoCommentsCommentIDForbidden() *PatchReposOwnerRepoCommentsCommentIDForbidden {
	return &PatchReposOwnerRepoCommentsCommentIDForbidden{}
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoCommentsCommentIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}