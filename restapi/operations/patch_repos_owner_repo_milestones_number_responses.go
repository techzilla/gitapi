// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PatchReposOwnerRepoMilestonesNumberOKCode is the HTTP code returned for type PatchReposOwnerRepoMilestonesNumberOK
const PatchReposOwnerRepoMilestonesNumberOKCode int = 200

/*PatchReposOwnerRepoMilestonesNumberOK OK

swagger:response patchReposOwnerRepoMilestonesNumberOK
*/
type PatchReposOwnerRepoMilestonesNumberOK struct {

	/*
	  In: Body
	*/
	Payload *models.Milestone `json:"body,omitempty"`
}

// NewPatchReposOwnerRepoMilestonesNumberOK creates PatchReposOwnerRepoMilestonesNumberOK with default headers values
func NewPatchReposOwnerRepoMilestonesNumberOK() *PatchReposOwnerRepoMilestonesNumberOK {
	return &PatchReposOwnerRepoMilestonesNumberOK{}
}

// WithPayload adds the payload to the patch repos owner repo milestones number o k response
func (o *PatchReposOwnerRepoMilestonesNumberOK) WithPayload(payload *models.Milestone) *PatchReposOwnerRepoMilestonesNumberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch repos owner repo milestones number o k response
func (o *PatchReposOwnerRepoMilestonesNumberOK) SetPayload(payload *models.Milestone) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoMilestonesNumberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchReposOwnerRepoMilestonesNumberForbiddenCode is the HTTP code returned for type PatchReposOwnerRepoMilestonesNumberForbidden
const PatchReposOwnerRepoMilestonesNumberForbiddenCode int = 403

/*PatchReposOwnerRepoMilestonesNumberForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response patchReposOwnerRepoMilestonesNumberForbidden
*/
type PatchReposOwnerRepoMilestonesNumberForbidden struct {
}

// NewPatchReposOwnerRepoMilestonesNumberForbidden creates PatchReposOwnerRepoMilestonesNumberForbidden with default headers values
func NewPatchReposOwnerRepoMilestonesNumberForbidden() *PatchReposOwnerRepoMilestonesNumberForbidden {
	return &PatchReposOwnerRepoMilestonesNumberForbidden{}
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoMilestonesNumberForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}