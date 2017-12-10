// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutTeamsTeamIDMembershipsUsernameHandlerFunc turns a function with the right signature into a put teams team ID memberships username handler
type PutTeamsTeamIDMembershipsUsernameHandlerFunc func(PutTeamsTeamIDMembershipsUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutTeamsTeamIDMembershipsUsernameHandlerFunc) Handle(params PutTeamsTeamIDMembershipsUsernameParams) middleware.Responder {
	return fn(params)
}

// PutTeamsTeamIDMembershipsUsernameHandler interface for that can handle valid put teams team ID memberships username params
type PutTeamsTeamIDMembershipsUsernameHandler interface {
	Handle(PutTeamsTeamIDMembershipsUsernameParams) middleware.Responder
}

// NewPutTeamsTeamIDMembershipsUsername creates a new http.Handler for the put teams team ID memberships username operation
func NewPutTeamsTeamIDMembershipsUsername(ctx *middleware.Context, handler PutTeamsTeamIDMembershipsUsernameHandler) *PutTeamsTeamIDMembershipsUsername {
	return &PutTeamsTeamIDMembershipsUsername{Context: ctx, Handler: handler}
}

/*PutTeamsTeamIDMembershipsUsername swagger:route PUT /teams/{teamId}/memberships/{username} putTeamsTeamIdMembershipsUsername

Add team membership.
In order to add a membership between a user and a team, the authenticated user must have 'admin' permissions to the team or be an owner of the organization that the team is associated with.

If the user is already a part of the team's organization (meaning they're on at least one other team in the organization), this endpoint will add the user to the team.

If the user is completely unaffiliated with the team's organization (meaning they're on none of the organization's teams), this endpoint will send an invitation to the user via email. This newly-created membership will be in the 'pending' state until the user accepts the invitation, at which point the membership will transition to the 'active' state and the user will be added as a member of the team.


*/
type PutTeamsTeamIDMembershipsUsername struct {
	Context *middleware.Context
	Handler PutTeamsTeamIDMembershipsUsernameHandler
}

func (o *PutTeamsTeamIDMembershipsUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutTeamsTeamIDMembershipsUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}