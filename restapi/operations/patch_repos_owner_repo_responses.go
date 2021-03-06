// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PatchReposOwnerRepoOKCode is the HTTP code returned for type PatchReposOwnerRepoOK
const PatchReposOwnerRepoOKCode int = 200

/*PatchReposOwnerRepoOK OK

swagger:response patchReposOwnerRepoOK
*/
type PatchReposOwnerRepoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Repo `json:"body,omitempty"`
}

// NewPatchReposOwnerRepoOK creates PatchReposOwnerRepoOK with default headers values
func NewPatchReposOwnerRepoOK() *PatchReposOwnerRepoOK {
	return &PatchReposOwnerRepoOK{}
}

// WithPayload adds the payload to the patch repos owner repo o k response
func (o *PatchReposOwnerRepoOK) WithPayload(payload *models.Repo) *PatchReposOwnerRepoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch repos owner repo o k response
func (o *PatchReposOwnerRepoOK) SetPayload(payload *models.Repo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchReposOwnerRepoForbiddenCode is the HTTP code returned for type PatchReposOwnerRepoForbidden
const PatchReposOwnerRepoForbiddenCode int = 403

/*PatchReposOwnerRepoForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response patchReposOwnerRepoForbidden
*/
type PatchReposOwnerRepoForbidden struct {
}

// NewPatchReposOwnerRepoForbidden creates PatchReposOwnerRepoForbidden with default headers values
func NewPatchReposOwnerRepoForbidden() *PatchReposOwnerRepoForbidden {
	return &PatchReposOwnerRepoForbidden{}
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
