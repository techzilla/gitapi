// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetNotificationsThreadsIDSubscriptionOKCode is the HTTP code returned for type GetNotificationsThreadsIDSubscriptionOK
const GetNotificationsThreadsIDSubscriptionOKCode int = 200

/*GetNotificationsThreadsIDSubscriptionOK OK

swagger:response getNotificationsThreadsIdSubscriptionOK
*/
type GetNotificationsThreadsIDSubscriptionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Subscription `json:"body,omitempty"`
}

// NewGetNotificationsThreadsIDSubscriptionOK creates GetNotificationsThreadsIDSubscriptionOK with default headers values
func NewGetNotificationsThreadsIDSubscriptionOK() *GetNotificationsThreadsIDSubscriptionOK {
	return &GetNotificationsThreadsIDSubscriptionOK{}
}

// WithPayload adds the payload to the get notifications threads Id subscription o k response
func (o *GetNotificationsThreadsIDSubscriptionOK) WithPayload(payload *models.Subscription) *GetNotificationsThreadsIDSubscriptionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get notifications threads Id subscription o k response
func (o *GetNotificationsThreadsIDSubscriptionOK) SetPayload(payload *models.Subscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNotificationsThreadsIDSubscriptionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNotificationsThreadsIDSubscriptionForbiddenCode is the HTTP code returned for type GetNotificationsThreadsIDSubscriptionForbidden
const GetNotificationsThreadsIDSubscriptionForbiddenCode int = 403

/*GetNotificationsThreadsIDSubscriptionForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getNotificationsThreadsIdSubscriptionForbidden
*/
type GetNotificationsThreadsIDSubscriptionForbidden struct {
}

// NewGetNotificationsThreadsIDSubscriptionForbidden creates GetNotificationsThreadsIDSubscriptionForbidden with default headers values
func NewGetNotificationsThreadsIDSubscriptionForbidden() *GetNotificationsThreadsIDSubscriptionForbidden {
	return &GetNotificationsThreadsIDSubscriptionForbidden{}
}

// WriteResponse to the client
func (o *GetNotificationsThreadsIDSubscriptionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
