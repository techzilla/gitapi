// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PatchReposOwnerRepoPullsCommentsCommentIDOKCode is the HTTP code returned for type PatchReposOwnerRepoPullsCommentsCommentIDOK
const PatchReposOwnerRepoPullsCommentsCommentIDOKCode int = 200

/*PatchReposOwnerRepoPullsCommentsCommentIDOK OK

swagger:response patchReposOwnerRepoPullsCommentsCommentIdOK
*/
type PatchReposOwnerRepoPullsCommentsCommentIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.PullsComment `json:"body,omitempty"`
}

// NewPatchReposOwnerRepoPullsCommentsCommentIDOK creates PatchReposOwnerRepoPullsCommentsCommentIDOK with default headers values
func NewPatchReposOwnerRepoPullsCommentsCommentIDOK() *PatchReposOwnerRepoPullsCommentsCommentIDOK {
	return &PatchReposOwnerRepoPullsCommentsCommentIDOK{}
}

// WithPayload adds the payload to the patch repos owner repo pulls comments comment Id o k response
func (o *PatchReposOwnerRepoPullsCommentsCommentIDOK) WithPayload(payload *models.PullsComment) *PatchReposOwnerRepoPullsCommentsCommentIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch repos owner repo pulls comments comment Id o k response
func (o *PatchReposOwnerRepoPullsCommentsCommentIDOK) SetPayload(payload *models.PullsComment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoPullsCommentsCommentIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchReposOwnerRepoPullsCommentsCommentIDForbiddenCode is the HTTP code returned for type PatchReposOwnerRepoPullsCommentsCommentIDForbidden
const PatchReposOwnerRepoPullsCommentsCommentIDForbiddenCode int = 403

/*PatchReposOwnerRepoPullsCommentsCommentIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response patchReposOwnerRepoPullsCommentsCommentIdForbidden
*/
type PatchReposOwnerRepoPullsCommentsCommentIDForbidden struct {
}

// NewPatchReposOwnerRepoPullsCommentsCommentIDForbidden creates PatchReposOwnerRepoPullsCommentsCommentIDForbidden with default headers values
func NewPatchReposOwnerRepoPullsCommentsCommentIDForbidden() *PatchReposOwnerRepoPullsCommentsCommentIDForbidden {
	return &PatchReposOwnerRepoPullsCommentsCommentIDForbidden{}
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoPullsCommentsCommentIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}