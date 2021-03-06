// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoBranchesBranchOKCode is the HTTP code returned for type GetReposOwnerRepoBranchesBranchOK
const GetReposOwnerRepoBranchesBranchOKCode int = 200

/*GetReposOwnerRepoBranchesBranchOK OK

swagger:response getReposOwnerRepoBranchesBranchOK
*/
type GetReposOwnerRepoBranchesBranchOK struct {

	/*
	  In: Body
	*/
	Payload *models.Branch `json:"body,omitempty"`
}

// NewGetReposOwnerRepoBranchesBranchOK creates GetReposOwnerRepoBranchesBranchOK with default headers values
func NewGetReposOwnerRepoBranchesBranchOK() *GetReposOwnerRepoBranchesBranchOK {
	return &GetReposOwnerRepoBranchesBranchOK{}
}

// WithPayload adds the payload to the get repos owner repo branches branch o k response
func (o *GetReposOwnerRepoBranchesBranchOK) WithPayload(payload *models.Branch) *GetReposOwnerRepoBranchesBranchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo branches branch o k response
func (o *GetReposOwnerRepoBranchesBranchOK) SetPayload(payload *models.Branch) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoBranchesBranchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoBranchesBranchForbiddenCode is the HTTP code returned for type GetReposOwnerRepoBranchesBranchForbidden
const GetReposOwnerRepoBranchesBranchForbiddenCode int = 403

/*GetReposOwnerRepoBranchesBranchForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoBranchesBranchForbidden
*/
type GetReposOwnerRepoBranchesBranchForbidden struct {
}

// NewGetReposOwnerRepoBranchesBranchForbidden creates GetReposOwnerRepoBranchesBranchForbidden with default headers values
func NewGetReposOwnerRepoBranchesBranchForbidden() *GetReposOwnerRepoBranchesBranchForbidden {
	return &GetReposOwnerRepoBranchesBranchForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoBranchesBranchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
