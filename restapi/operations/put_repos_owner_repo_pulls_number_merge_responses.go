// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PutReposOwnerRepoPullsNumberMergeOKCode is the HTTP code returned for type PutReposOwnerRepoPullsNumberMergeOK
const PutReposOwnerRepoPullsNumberMergeOKCode int = 200

/*PutReposOwnerRepoPullsNumberMergeOK Response if merge was successful.

swagger:response putReposOwnerRepoPullsNumberMergeOK
*/
type PutReposOwnerRepoPullsNumberMergeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Merge `json:"body,omitempty"`
}

// NewPutReposOwnerRepoPullsNumberMergeOK creates PutReposOwnerRepoPullsNumberMergeOK with default headers values
func NewPutReposOwnerRepoPullsNumberMergeOK() *PutReposOwnerRepoPullsNumberMergeOK {
	return &PutReposOwnerRepoPullsNumberMergeOK{}
}

// WithPayload adds the payload to the put repos owner repo pulls number merge o k response
func (o *PutReposOwnerRepoPullsNumberMergeOK) WithPayload(payload *models.Merge) *PutReposOwnerRepoPullsNumberMergeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put repos owner repo pulls number merge o k response
func (o *PutReposOwnerRepoPullsNumberMergeOK) SetPayload(payload *models.Merge) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutReposOwnerRepoPullsNumberMergeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutReposOwnerRepoPullsNumberMergeForbiddenCode is the HTTP code returned for type PutReposOwnerRepoPullsNumberMergeForbidden
const PutReposOwnerRepoPullsNumberMergeForbiddenCode int = 403

/*PutReposOwnerRepoPullsNumberMergeForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response putReposOwnerRepoPullsNumberMergeForbidden
*/
type PutReposOwnerRepoPullsNumberMergeForbidden struct {
}

// NewPutReposOwnerRepoPullsNumberMergeForbidden creates PutReposOwnerRepoPullsNumberMergeForbidden with default headers values
func NewPutReposOwnerRepoPullsNumberMergeForbidden() *PutReposOwnerRepoPullsNumberMergeForbidden {
	return &PutReposOwnerRepoPullsNumberMergeForbidden{}
}

// WriteResponse to the client
func (o *PutReposOwnerRepoPullsNumberMergeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PutReposOwnerRepoPullsNumberMergeMethodNotAllowedCode is the HTTP code returned for type PutReposOwnerRepoPullsNumberMergeMethodNotAllowed
const PutReposOwnerRepoPullsNumberMergeMethodNotAllowedCode int = 405

/*PutReposOwnerRepoPullsNumberMergeMethodNotAllowed Response if merge cannot be performed.

swagger:response putReposOwnerRepoPullsNumberMergeMethodNotAllowed
*/
type PutReposOwnerRepoPullsNumberMergeMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Merge `json:"body,omitempty"`
}

// NewPutReposOwnerRepoPullsNumberMergeMethodNotAllowed creates PutReposOwnerRepoPullsNumberMergeMethodNotAllowed with default headers values
func NewPutReposOwnerRepoPullsNumberMergeMethodNotAllowed() *PutReposOwnerRepoPullsNumberMergeMethodNotAllowed {
	return &PutReposOwnerRepoPullsNumberMergeMethodNotAllowed{}
}

// WithPayload adds the payload to the put repos owner repo pulls number merge method not allowed response
func (o *PutReposOwnerRepoPullsNumberMergeMethodNotAllowed) WithPayload(payload *models.Merge) *PutReposOwnerRepoPullsNumberMergeMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put repos owner repo pulls number merge method not allowed response
func (o *PutReposOwnerRepoPullsNumberMergeMethodNotAllowed) SetPayload(payload *models.Merge) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutReposOwnerRepoPullsNumberMergeMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
