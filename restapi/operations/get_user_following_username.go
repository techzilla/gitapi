// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserFollowingUsernameHandlerFunc turns a function with the right signature into a get user following username handler
type GetUserFollowingUsernameHandlerFunc func(GetUserFollowingUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserFollowingUsernameHandlerFunc) Handle(params GetUserFollowingUsernameParams) middleware.Responder {
	return fn(params)
}

// GetUserFollowingUsernameHandler interface for that can handle valid get user following username params
type GetUserFollowingUsernameHandler interface {
	Handle(GetUserFollowingUsernameParams) middleware.Responder
}

// NewGetUserFollowingUsername creates a new http.Handler for the get user following username operation
func NewGetUserFollowingUsername(ctx *middleware.Context, handler GetUserFollowingUsernameHandler) *GetUserFollowingUsername {
	return &GetUserFollowingUsername{Context: ctx, Handler: handler}
}

/*GetUserFollowingUsername swagger:route GET /user/following/{username} getUserFollowingUsername

Check if you are following a user.

*/
type GetUserFollowingUsername struct {
	Context *middleware.Context
	Handler GetUserFollowingUsernameHandler
}

func (o *GetUserFollowingUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserFollowingUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}