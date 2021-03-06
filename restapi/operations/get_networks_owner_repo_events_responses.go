// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetNetworksOwnerRepoEventsOKCode is the HTTP code returned for type GetNetworksOwnerRepoEventsOK
const GetNetworksOwnerRepoEventsOKCode int = 200

/*GetNetworksOwnerRepoEventsOK OK

swagger:response getNetworksOwnerRepoEventsOK
*/
type GetNetworksOwnerRepoEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Events `json:"body,omitempty"`
}

// NewGetNetworksOwnerRepoEventsOK creates GetNetworksOwnerRepoEventsOK with default headers values
func NewGetNetworksOwnerRepoEventsOK() *GetNetworksOwnerRepoEventsOK {
	return &GetNetworksOwnerRepoEventsOK{}
}

// WithPayload adds the payload to the get networks owner repo events o k response
func (o *GetNetworksOwnerRepoEventsOK) WithPayload(payload *models.Events) *GetNetworksOwnerRepoEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get networks owner repo events o k response
func (o *GetNetworksOwnerRepoEventsOK) SetPayload(payload *models.Events) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNetworksOwnerRepoEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNetworksOwnerRepoEventsForbiddenCode is the HTTP code returned for type GetNetworksOwnerRepoEventsForbidden
const GetNetworksOwnerRepoEventsForbiddenCode int = 403

/*GetNetworksOwnerRepoEventsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getNetworksOwnerRepoEventsForbidden
*/
type GetNetworksOwnerRepoEventsForbidden struct {
}

// NewGetNetworksOwnerRepoEventsForbidden creates GetNetworksOwnerRepoEventsForbidden with default headers values
func NewGetNetworksOwnerRepoEventsForbidden() *GetNetworksOwnerRepoEventsForbidden {
	return &GetNetworksOwnerRepoEventsForbidden{}
}

// WriteResponse to the client
func (o *GetNetworksOwnerRepoEventsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
