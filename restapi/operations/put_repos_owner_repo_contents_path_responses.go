// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PutReposOwnerRepoContentsPathOKCode is the HTTP code returned for type PutReposOwnerRepoContentsPathOK
const PutReposOwnerRepoContentsPathOKCode int = 200

/*PutReposOwnerRepoContentsPathOK OK

swagger:response putReposOwnerRepoContentsPathOK
*/
type PutReposOwnerRepoContentsPathOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateFile `json:"body,omitempty"`
}

// NewPutReposOwnerRepoContentsPathOK creates PutReposOwnerRepoContentsPathOK with default headers values
func NewPutReposOwnerRepoContentsPathOK() *PutReposOwnerRepoContentsPathOK {
	return &PutReposOwnerRepoContentsPathOK{}
}

// WithPayload adds the payload to the put repos owner repo contents path o k response
func (o *PutReposOwnerRepoContentsPathOK) WithPayload(payload *models.CreateFile) *PutReposOwnerRepoContentsPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put repos owner repo contents path o k response
func (o *PutReposOwnerRepoContentsPathOK) SetPayload(payload *models.CreateFile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutReposOwnerRepoContentsPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutReposOwnerRepoContentsPathForbiddenCode is the HTTP code returned for type PutReposOwnerRepoContentsPathForbidden
const PutReposOwnerRepoContentsPathForbiddenCode int = 403

/*PutReposOwnerRepoContentsPathForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response putReposOwnerRepoContentsPathForbidden
*/
type PutReposOwnerRepoContentsPathForbidden struct {
}

// NewPutReposOwnerRepoContentsPathForbidden creates PutReposOwnerRepoContentsPathForbidden with default headers values
func NewPutReposOwnerRepoContentsPathForbidden() *PutReposOwnerRepoContentsPathForbidden {
	return &PutReposOwnerRepoContentsPathForbidden{}
}

// WriteResponse to the client
func (o *PutReposOwnerRepoContentsPathForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}