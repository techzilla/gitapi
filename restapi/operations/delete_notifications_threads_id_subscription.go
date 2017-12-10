// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteNotificationsThreadsIDSubscriptionHandlerFunc turns a function with the right signature into a delete notifications threads ID subscription handler
type DeleteNotificationsThreadsIDSubscriptionHandlerFunc func(DeleteNotificationsThreadsIDSubscriptionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNotificationsThreadsIDSubscriptionHandlerFunc) Handle(params DeleteNotificationsThreadsIDSubscriptionParams) middleware.Responder {
	return fn(params)
}

// DeleteNotificationsThreadsIDSubscriptionHandler interface for that can handle valid delete notifications threads ID subscription params
type DeleteNotificationsThreadsIDSubscriptionHandler interface {
	Handle(DeleteNotificationsThreadsIDSubscriptionParams) middleware.Responder
}

// NewDeleteNotificationsThreadsIDSubscription creates a new http.Handler for the delete notifications threads ID subscription operation
func NewDeleteNotificationsThreadsIDSubscription(ctx *middleware.Context, handler DeleteNotificationsThreadsIDSubscriptionHandler) *DeleteNotificationsThreadsIDSubscription {
	return &DeleteNotificationsThreadsIDSubscription{Context: ctx, Handler: handler}
}

/*DeleteNotificationsThreadsIDSubscription swagger:route DELETE /notifications/threads/{id}/subscription deleteNotificationsThreadsIdSubscription

Delete a Thread Subscription.

*/
type DeleteNotificationsThreadsIDSubscription struct {
	Context *middleware.Context
	Handler DeleteNotificationsThreadsIDSubscriptionHandler
}

func (o *DeleteNotificationsThreadsIDSubscription) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNotificationsThreadsIDSubscriptionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
