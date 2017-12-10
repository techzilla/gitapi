// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUsersUsernameFollowingTargetUserHandlerFunc turns a function with the right signature into a get users username following target user handler
type GetUsersUsernameFollowingTargetUserHandlerFunc func(GetUsersUsernameFollowingTargetUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersUsernameFollowingTargetUserHandlerFunc) Handle(params GetUsersUsernameFollowingTargetUserParams) middleware.Responder {
	return fn(params)
}

// GetUsersUsernameFollowingTargetUserHandler interface for that can handle valid get users username following target user params
type GetUsersUsernameFollowingTargetUserHandler interface {
	Handle(GetUsersUsernameFollowingTargetUserParams) middleware.Responder
}

// NewGetUsersUsernameFollowingTargetUser creates a new http.Handler for the get users username following target user operation
func NewGetUsersUsernameFollowingTargetUser(ctx *middleware.Context, handler GetUsersUsernameFollowingTargetUserHandler) *GetUsersUsernameFollowingTargetUser {
	return &GetUsersUsernameFollowingTargetUser{Context: ctx, Handler: handler}
}

/*GetUsersUsernameFollowingTargetUser swagger:route GET /users/{username}/following/{targetUser} getUsersUsernameFollowingTargetUser

Check if one user follows another.

*/
type GetUsersUsernameFollowingTargetUser struct {
	Context *middleware.Context
	Handler GetUsersUsernameFollowingTargetUserHandler
}

func (o *GetUsersUsernameFollowingTargetUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersUsernameFollowingTargetUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
