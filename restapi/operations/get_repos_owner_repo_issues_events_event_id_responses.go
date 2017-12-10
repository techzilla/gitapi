// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoIssuesEventsEventIDOKCode is the HTTP code returned for type GetReposOwnerRepoIssuesEventsEventIDOK
const GetReposOwnerRepoIssuesEventsEventIDOKCode int = 200

/*GetReposOwnerRepoIssuesEventsEventIDOK OK

swagger:response getReposOwnerRepoIssuesEventsEventIdOK
*/
type GetReposOwnerRepoIssuesEventsEventIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Event `json:"body,omitempty"`
}

// NewGetReposOwnerRepoIssuesEventsEventIDOK creates GetReposOwnerRepoIssuesEventsEventIDOK with default headers values
func NewGetReposOwnerRepoIssuesEventsEventIDOK() *GetReposOwnerRepoIssuesEventsEventIDOK {
	return &GetReposOwnerRepoIssuesEventsEventIDOK{}
}

// WithPayload adds the payload to the get repos owner repo issues events event Id o k response
func (o *GetReposOwnerRepoIssuesEventsEventIDOK) WithPayload(payload *models.Event) *GetReposOwnerRepoIssuesEventsEventIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo issues events event Id o k response
func (o *GetReposOwnerRepoIssuesEventsEventIDOK) SetPayload(payload *models.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesEventsEventIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoIssuesEventsEventIDForbiddenCode is the HTTP code returned for type GetReposOwnerRepoIssuesEventsEventIDForbidden
const GetReposOwnerRepoIssuesEventsEventIDForbiddenCode int = 403

/*GetReposOwnerRepoIssuesEventsEventIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoIssuesEventsEventIdForbidden
*/
type GetReposOwnerRepoIssuesEventsEventIDForbidden struct {
}

// NewGetReposOwnerRepoIssuesEventsEventIDForbidden creates GetReposOwnerRepoIssuesEventsEventIDForbidden with default headers values
func NewGetReposOwnerRepoIssuesEventsEventIDForbidden() *GetReposOwnerRepoIssuesEventsEventIDForbidden {
	return &GetReposOwnerRepoIssuesEventsEventIDForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesEventsEventIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}