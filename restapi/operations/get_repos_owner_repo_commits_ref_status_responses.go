// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoCommitsRefStatusOKCode is the HTTP code returned for type GetReposOwnerRepoCommitsRefStatusOK
const GetReposOwnerRepoCommitsRefStatusOKCode int = 200

/*GetReposOwnerRepoCommitsRefStatusOK OK

swagger:response getReposOwnerRepoCommitsRefStatusOK
*/
type GetReposOwnerRepoCommitsRefStatusOK struct {

	/*
	  In: Body
	*/
	Payload models.RefStatus `json:"body,omitempty"`
}

// NewGetReposOwnerRepoCommitsRefStatusOK creates GetReposOwnerRepoCommitsRefStatusOK with default headers values
func NewGetReposOwnerRepoCommitsRefStatusOK() *GetReposOwnerRepoCommitsRefStatusOK {
	return &GetReposOwnerRepoCommitsRefStatusOK{}
}

// WithPayload adds the payload to the get repos owner repo commits ref status o k response
func (o *GetReposOwnerRepoCommitsRefStatusOK) WithPayload(payload models.RefStatus) *GetReposOwnerRepoCommitsRefStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo commits ref status o k response
func (o *GetReposOwnerRepoCommitsRefStatusOK) SetPayload(payload models.RefStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoCommitsRefStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.RefStatus, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReposOwnerRepoCommitsRefStatusForbiddenCode is the HTTP code returned for type GetReposOwnerRepoCommitsRefStatusForbidden
const GetReposOwnerRepoCommitsRefStatusForbiddenCode int = 403

/*GetReposOwnerRepoCommitsRefStatusForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoCommitsRefStatusForbidden
*/
type GetReposOwnerRepoCommitsRefStatusForbidden struct {
}

// NewGetReposOwnerRepoCommitsRefStatusForbidden creates GetReposOwnerRepoCommitsRefStatusForbidden with default headers values
func NewGetReposOwnerRepoCommitsRefStatusForbidden() *GetReposOwnerRepoCommitsRefStatusForbidden {
	return &GetReposOwnerRepoCommitsRefStatusForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoCommitsRefStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
