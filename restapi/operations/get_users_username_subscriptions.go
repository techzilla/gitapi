// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUsersUsernameSubscriptionsHandlerFunc turns a function with the right signature into a get users username subscriptions handler
type GetUsersUsernameSubscriptionsHandlerFunc func(GetUsersUsernameSubscriptionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersUsernameSubscriptionsHandlerFunc) Handle(params GetUsersUsernameSubscriptionsParams) middleware.Responder {
	return fn(params)
}

// GetUsersUsernameSubscriptionsHandler interface for that can handle valid get users username subscriptions params
type GetUsersUsernameSubscriptionsHandler interface {
	Handle(GetUsersUsernameSubscriptionsParams) middleware.Responder
}

// NewGetUsersUsernameSubscriptions creates a new http.Handler for the get users username subscriptions operation
func NewGetUsersUsernameSubscriptions(ctx *middleware.Context, handler GetUsersUsernameSubscriptionsHandler) *GetUsersUsernameSubscriptions {
	return &GetUsersUsernameSubscriptions{Context: ctx, Handler: handler}
}

/*GetUsersUsernameSubscriptions swagger:route GET /users/{username}/subscriptions getUsersUsernameSubscriptions

List repositories being watched by a user.

*/
type GetUsersUsernameSubscriptions struct {
	Context *middleware.Context
	Handler GetUsersUsernameSubscriptionsHandler
}

func (o *GetUsersUsernameSubscriptions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersUsernameSubscriptionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}