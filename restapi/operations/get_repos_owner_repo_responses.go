// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoOKCode is the HTTP code returned for type GetReposOwnerRepoOK
const GetReposOwnerRepoOKCode int = 200

/*GetReposOwnerRepoOK OK

swagger:response getReposOwnerRepoOK
*/
type GetReposOwnerRepoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Repo `json:"body,omitempty"`
}

// NewGetReposOwnerRepoOK creates GetReposOwnerRepoOK with default headers values
func NewGetReposOwnerRepoOK() *GetReposOwnerRepoOK {
	return &GetReposOwnerRepoOK{}
}

// WithPayload adds the payload to the get repos owner repo o k response
func (o *GetReposOwnerRepoOK) WithPayload(payload *models.Repo) *GetReposOwnerRepoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo o k response
func (o *GetReposOwnerRepoOK) SetPayload(payload *models.Repo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoForbiddenCode is the HTTP code returned for type GetReposOwnerRepoForbidden
const GetReposOwnerRepoForbiddenCode int = 403

/*GetReposOwnerRepoForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoForbidden
*/
type GetReposOwnerRepoForbidden struct {
}

// NewGetReposOwnerRepoForbidden creates GetReposOwnerRepoForbidden with default headers values
func NewGetReposOwnerRepoForbidden() *GetReposOwnerRepoForbidden {
	return &GetReposOwnerRepoForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
