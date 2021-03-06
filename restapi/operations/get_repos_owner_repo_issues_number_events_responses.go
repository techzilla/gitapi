// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoIssuesNumberEventsOKCode is the HTTP code returned for type GetReposOwnerRepoIssuesNumberEventsOK
const GetReposOwnerRepoIssuesNumberEventsOKCode int = 200

/*GetReposOwnerRepoIssuesNumberEventsOK OK

swagger:response getReposOwnerRepoIssuesNumberEventsOK
*/
type GetReposOwnerRepoIssuesNumberEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Events `json:"body,omitempty"`
}

// NewGetReposOwnerRepoIssuesNumberEventsOK creates GetReposOwnerRepoIssuesNumberEventsOK with default headers values
func NewGetReposOwnerRepoIssuesNumberEventsOK() *GetReposOwnerRepoIssuesNumberEventsOK {
	return &GetReposOwnerRepoIssuesNumberEventsOK{}
}

// WithPayload adds the payload to the get repos owner repo issues number events o k response
func (o *GetReposOwnerRepoIssuesNumberEventsOK) WithPayload(payload *models.Events) *GetReposOwnerRepoIssuesNumberEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo issues number events o k response
func (o *GetReposOwnerRepoIssuesNumberEventsOK) SetPayload(payload *models.Events) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesNumberEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoIssuesNumberEventsForbiddenCode is the HTTP code returned for type GetReposOwnerRepoIssuesNumberEventsForbidden
const GetReposOwnerRepoIssuesNumberEventsForbiddenCode int = 403

/*GetReposOwnerRepoIssuesNumberEventsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoIssuesNumberEventsForbidden
*/
type GetReposOwnerRepoIssuesNumberEventsForbidden struct {
}

// NewGetReposOwnerRepoIssuesNumberEventsForbidden creates GetReposOwnerRepoIssuesNumberEventsForbidden with default headers values
func NewGetReposOwnerRepoIssuesNumberEventsForbidden() *GetReposOwnerRepoIssuesNumberEventsForbidden {
	return &GetReposOwnerRepoIssuesNumberEventsForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoIssuesNumberEventsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
