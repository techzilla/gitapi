// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutUserFollowingUsernameHandlerFunc turns a function with the right signature into a put user following username handler
type PutUserFollowingUsernameHandlerFunc func(PutUserFollowingUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUserFollowingUsernameHandlerFunc) Handle(params PutUserFollowingUsernameParams) middleware.Responder {
	return fn(params)
}

// PutUserFollowingUsernameHandler interface for that can handle valid put user following username params
type PutUserFollowingUsernameHandler interface {
	Handle(PutUserFollowingUsernameParams) middleware.Responder
}

// NewPutUserFollowingUsername creates a new http.Handler for the put user following username operation
func NewPutUserFollowingUsername(ctx *middleware.Context, handler PutUserFollowingUsernameHandler) *PutUserFollowingUsername {
	return &PutUserFollowingUsername{Context: ctx, Handler: handler}
}

/*PutUserFollowingUsername swagger:route PUT /user/following/{username} putUserFollowingUsername

Follow a user.
Following a user requires the user to be logged in and authenticated with
basic auth or OAuth with the user:follow scope.


*/
type PutUserFollowingUsername struct {
	Context *middleware.Context
	Handler PutUserFollowingUsernameHandler
}

func (o *PutUserFollowingUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutUserFollowingUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
