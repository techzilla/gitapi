// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetEventsOKCode is the HTTP code returned for type GetEventsOK
const GetEventsOKCode int = 200

/*GetEventsOK OK

swagger:response getEventsOK
*/
type GetEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Events `json:"body,omitempty"`
}

// NewGetEventsOK creates GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {
	return &GetEventsOK{}
}

// WithPayload adds the payload to the get events o k response
func (o *GetEventsOK) WithPayload(payload *models.Events) *GetEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events o k response
func (o *GetEventsOK) SetPayload(payload *models.Events) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsForbiddenCode is the HTTP code returned for type GetEventsForbidden
const GetEventsForbiddenCode int = 403

/*GetEventsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getEventsForbidden
*/
type GetEventsForbidden struct {
}

// NewGetEventsForbidden creates GetEventsForbidden with default headers values
func NewGetEventsForbidden() *GetEventsForbidden {
	return &GetEventsForbidden{}
}

// WriteResponse to the client
func (o *GetEventsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
