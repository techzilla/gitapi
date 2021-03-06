// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteNotificationsThreadsIDSubscriptionNoContentCode is the HTTP code returned for type DeleteNotificationsThreadsIDSubscriptionNoContent
const DeleteNotificationsThreadsIDSubscriptionNoContentCode int = 204

/*DeleteNotificationsThreadsIDSubscriptionNoContent No Content


swagger:response deleteNotificationsThreadsIdSubscriptionNoContent
*/
type DeleteNotificationsThreadsIDSubscriptionNoContent struct {
}

// NewDeleteNotificationsThreadsIDSubscriptionNoContent creates DeleteNotificationsThreadsIDSubscriptionNoContent with default headers values
func NewDeleteNotificationsThreadsIDSubscriptionNoContent() *DeleteNotificationsThreadsIDSubscriptionNoContent {
	return &DeleteNotificationsThreadsIDSubscriptionNoContent{}
}

// WriteResponse to the client
func (o *DeleteNotificationsThreadsIDSubscriptionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteNotificationsThreadsIDSubscriptionForbiddenCode is the HTTP code returned for type DeleteNotificationsThreadsIDSubscriptionForbidden
const DeleteNotificationsThreadsIDSubscriptionForbiddenCode int = 403

/*DeleteNotificationsThreadsIDSubscriptionForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response deleteNotificationsThreadsIdSubscriptionForbidden
*/
type DeleteNotificationsThreadsIDSubscriptionForbidden struct {
}

// NewDeleteNotificationsThreadsIDSubscriptionForbidden creates DeleteNotificationsThreadsIDSubscriptionForbidden with default headers values
func NewDeleteNotificationsThreadsIDSubscriptionForbidden() *DeleteNotificationsThreadsIDSubscriptionForbidden {
	return &DeleteNotificationsThreadsIDSubscriptionForbidden{}
}

// WriteResponse to the client
func (o *DeleteNotificationsThreadsIDSubscriptionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
