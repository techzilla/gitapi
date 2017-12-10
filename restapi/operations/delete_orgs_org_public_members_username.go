// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteOrgsOrgPublicMembersUsernameHandlerFunc turns a function with the right signature into a delete orgs org public members username handler
type DeleteOrgsOrgPublicMembersUsernameHandlerFunc func(DeleteOrgsOrgPublicMembersUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteOrgsOrgPublicMembersUsernameHandlerFunc) Handle(params DeleteOrgsOrgPublicMembersUsernameParams) middleware.Responder {
	return fn(params)
}

// DeleteOrgsOrgPublicMembersUsernameHandler interface for that can handle valid delete orgs org public members username params
type DeleteOrgsOrgPublicMembersUsernameHandler interface {
	Handle(DeleteOrgsOrgPublicMembersUsernameParams) middleware.Responder
}

// NewDeleteOrgsOrgPublicMembersUsername creates a new http.Handler for the delete orgs org public members username operation
func NewDeleteOrgsOrgPublicMembersUsername(ctx *middleware.Context, handler DeleteOrgsOrgPublicMembersUsernameHandler) *DeleteOrgsOrgPublicMembersUsername {
	return &DeleteOrgsOrgPublicMembersUsername{Context: ctx, Handler: handler}
}

/*DeleteOrgsOrgPublicMembersUsername swagger:route DELETE /orgs/{org}/public_members/{username} deleteOrgsOrgPublicMembersUsername

Conceal a user's membership.

*/
type DeleteOrgsOrgPublicMembersUsername struct {
	Context *middleware.Context
	Handler DeleteOrgsOrgPublicMembersUsernameHandler
}

func (o *DeleteOrgsOrgPublicMembersUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteOrgsOrgPublicMembersUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
