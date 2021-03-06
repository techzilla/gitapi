// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTeamsTeamIDMembersUsernameHandlerFunc turns a function with the right signature into a get teams team ID members username handler
type GetTeamsTeamIDMembersUsernameHandlerFunc func(GetTeamsTeamIDMembersUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeamsTeamIDMembersUsernameHandlerFunc) Handle(params GetTeamsTeamIDMembersUsernameParams) middleware.Responder {
	return fn(params)
}

// GetTeamsTeamIDMembersUsernameHandler interface for that can handle valid get teams team ID members username params
type GetTeamsTeamIDMembersUsernameHandler interface {
	Handle(GetTeamsTeamIDMembersUsernameParams) middleware.Responder
}

// NewGetTeamsTeamIDMembersUsername creates a new http.Handler for the get teams team ID members username operation
func NewGetTeamsTeamIDMembersUsername(ctx *middleware.Context, handler GetTeamsTeamIDMembersUsernameHandler) *GetTeamsTeamIDMembersUsername {
	return &GetTeamsTeamIDMembersUsername{Context: ctx, Handler: handler}
}

/*GetTeamsTeamIDMembersUsername swagger:route GET /teams/{teamId}/members/{username} getTeamsTeamIdMembersUsername

The "Get team member" API is deprecated and is scheduled for removal in the next major version of the API. We recommend using the Get team membership API instead. It allows you to get both active and pending memberships.

Get team member.
In order to get if a user is a member of a team, the authenticated user mus
be a member of the team.


*/
type GetTeamsTeamIDMembersUsername struct {
	Context *middleware.Context
	Handler GetTeamsTeamIDMembersUsernameHandler
}

func (o *GetTeamsTeamIDMembersUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTeamsTeamIDMembersUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
