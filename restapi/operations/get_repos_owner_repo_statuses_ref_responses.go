// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoStatusesRefOKCode is the HTTP code returned for type GetReposOwnerRepoStatusesRefOK
const GetReposOwnerRepoStatusesRefOKCode int = 200

/*GetReposOwnerRepoStatusesRefOK OK

swagger:response getReposOwnerRepoStatusesRefOK
*/
type GetReposOwnerRepoStatusesRefOK struct {

	/*
	  In: Body
	*/
	Payload models.Ref `json:"body,omitempty"`
}

// NewGetReposOwnerRepoStatusesRefOK creates GetReposOwnerRepoStatusesRefOK with default headers values
func NewGetReposOwnerRepoStatusesRefOK() *GetReposOwnerRepoStatusesRefOK {
	return &GetReposOwnerRepoStatusesRefOK{}
}

// WithPayload adds the payload to the get repos owner repo statuses ref o k response
func (o *GetReposOwnerRepoStatusesRefOK) WithPayload(payload models.Ref) *GetReposOwnerRepoStatusesRefOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo statuses ref o k response
func (o *GetReposOwnerRepoStatusesRefOK) SetPayload(payload models.Ref) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoStatusesRefOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Ref, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReposOwnerRepoStatusesRefForbiddenCode is the HTTP code returned for type GetReposOwnerRepoStatusesRefForbidden
const GetReposOwnerRepoStatusesRefForbiddenCode int = 403

/*GetReposOwnerRepoStatusesRefForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoStatusesRefForbidden
*/
type GetReposOwnerRepoStatusesRefForbidden struct {
}

// NewGetReposOwnerRepoStatusesRefForbidden creates GetReposOwnerRepoStatusesRefForbidden with default headers values
func NewGetReposOwnerRepoStatusesRefForbidden() *GetReposOwnerRepoStatusesRefForbidden {
	return &GetReposOwnerRepoStatusesRefForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoStatusesRefForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}