// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOrgsOrgMembersHandlerFunc turns a function with the right signature into a get orgs org members handler
type GetOrgsOrgMembersHandlerFunc func(GetOrgsOrgMembersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOrgsOrgMembersHandlerFunc) Handle(params GetOrgsOrgMembersParams) middleware.Responder {
	return fn(params)
}

// GetOrgsOrgMembersHandler interface for that can handle valid get orgs org members params
type GetOrgsOrgMembersHandler interface {
	Handle(GetOrgsOrgMembersParams) middleware.Responder
}

// NewGetOrgsOrgMembers creates a new http.Handler for the get orgs org members operation
func NewGetOrgsOrgMembers(ctx *middleware.Context, handler GetOrgsOrgMembersHandler) *GetOrgsOrgMembers {
	return &GetOrgsOrgMembers{Context: ctx, Handler: handler}
}

/*GetOrgsOrgMembers swagger:route GET /orgs/{org}/members getOrgsOrgMembers

Members list.
List all users who are members of an organization. A member is a user tha
belongs to at least 1 team in the organization. If the authenticated user
is also an owner of this organization then both concealed and public members
will be returned. If the requester is not an owner of the organization the
query will be redirected to the public members list.


*/
type GetOrgsOrgMembers struct {
	Context *middleware.Context
	Handler GetOrgsOrgMembersHandler
}

func (o *GetOrgsOrgMembers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOrgsOrgMembersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
