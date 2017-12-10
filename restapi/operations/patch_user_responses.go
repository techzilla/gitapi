// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PatchUserOKCode is the HTTP code returned for type PatchUserOK
const PatchUserOKCode int = 200

/*PatchUserOK OK

swagger:response patchUserOK
*/
type PatchUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPatchUserOK creates PatchUserOK with default headers values
func NewPatchUserOK() *PatchUserOK {
	return &PatchUserOK{}
}

// WithPayload adds the payload to the patch user o k response
func (o *PatchUserOK) WithPayload(payload *models.User) *PatchUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user o k response
func (o *PatchUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUserForbiddenCode is the HTTP code returned for type PatchUserForbidden
const PatchUserForbiddenCode int = 403

/*PatchUserForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response patchUserForbidden
*/
type PatchUserForbidden struct {
}

// NewPatchUserForbidden creates PatchUserForbidden with default headers values
func NewPatchUserForbidden() *PatchUserForbidden {
	return &PatchUserForbidden{}
}

// WriteResponse to the client
func (o *PatchUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}