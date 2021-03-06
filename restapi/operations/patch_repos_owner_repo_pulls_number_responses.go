// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PatchReposOwnerRepoPullsNumberOKCode is the HTTP code returned for type PatchReposOwnerRepoPullsNumberOK
const PatchReposOwnerRepoPullsNumberOKCode int = 200

/*PatchReposOwnerRepoPullsNumberOK OK

swagger:response patchReposOwnerRepoPullsNumberOK
*/
type PatchReposOwnerRepoPullsNumberOK struct {

	/*
	  In: Body
	*/
	Payload *models.Repo `json:"body,omitempty"`
}

// NewPatchReposOwnerRepoPullsNumberOK creates PatchReposOwnerRepoPullsNumberOK with default headers values
func NewPatchReposOwnerRepoPullsNumberOK() *PatchReposOwnerRepoPullsNumberOK {
	return &PatchReposOwnerRepoPullsNumberOK{}
}

// WithPayload adds the payload to the patch repos owner repo pulls number o k response
func (o *PatchReposOwnerRepoPullsNumberOK) WithPayload(payload *models.Repo) *PatchReposOwnerRepoPullsNumberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch repos owner repo pulls number o k response
func (o *PatchReposOwnerRepoPullsNumberOK) SetPayload(payload *models.Repo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoPullsNumberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchReposOwnerRepoPullsNumberForbiddenCode is the HTTP code returned for type PatchReposOwnerRepoPullsNumberForbidden
const PatchReposOwnerRepoPullsNumberForbiddenCode int = 403

/*PatchReposOwnerRepoPullsNumberForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response patchReposOwnerRepoPullsNumberForbidden
*/
type PatchReposOwnerRepoPullsNumberForbidden struct {
}

// NewPatchReposOwnerRepoPullsNumberForbidden creates PatchReposOwnerRepoPullsNumberForbidden with default headers values
func NewPatchReposOwnerRepoPullsNumberForbidden() *PatchReposOwnerRepoPullsNumberForbidden {
	return &PatchReposOwnerRepoPullsNumberForbidden{}
}

// WriteResponse to the client
func (o *PatchReposOwnerRepoPullsNumberForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
