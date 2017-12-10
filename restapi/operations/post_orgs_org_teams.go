// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOrgsOrgTeamsHandlerFunc turns a function with the right signature into a post orgs org teams handler
type PostOrgsOrgTeamsHandlerFunc func(PostOrgsOrgTeamsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOrgsOrgTeamsHandlerFunc) Handle(params PostOrgsOrgTeamsParams) middleware.Responder {
	return fn(params)
}

// PostOrgsOrgTeamsHandler interface for that can handle valid post orgs org teams params
type PostOrgsOrgTeamsHandler interface {
	Handle(PostOrgsOrgTeamsParams) middleware.Responder
}

// NewPostOrgsOrgTeams creates a new http.Handler for the post orgs org teams operation
func NewPostOrgsOrgTeams(ctx *middleware.Context, handler PostOrgsOrgTeamsHandler) *PostOrgsOrgTeams {
	return &PostOrgsOrgTeams{Context: ctx, Handler: handler}
}

/*PostOrgsOrgTeams swagger:route POST /orgs/{org}/teams postOrgsOrgTeams

Create team.
In order to create a team, the authenticated user must be an owner of organization.


*/
type PostOrgsOrgTeams struct {
	Context *middleware.Context
	Handler PostOrgsOrgTeamsHandler
}

func (o *PostOrgsOrgTeams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOrgsOrgTeamsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}