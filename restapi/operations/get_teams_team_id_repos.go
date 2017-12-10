// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTeamsTeamIDReposHandlerFunc turns a function with the right signature into a get teams team ID repos handler
type GetTeamsTeamIDReposHandlerFunc func(GetTeamsTeamIDReposParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeamsTeamIDReposHandlerFunc) Handle(params GetTeamsTeamIDReposParams) middleware.Responder {
	return fn(params)
}

// GetTeamsTeamIDReposHandler interface for that can handle valid get teams team ID repos params
type GetTeamsTeamIDReposHandler interface {
	Handle(GetTeamsTeamIDReposParams) middleware.Responder
}

// NewGetTeamsTeamIDRepos creates a new http.Handler for the get teams team ID repos operation
func NewGetTeamsTeamIDRepos(ctx *middleware.Context, handler GetTeamsTeamIDReposHandler) *GetTeamsTeamIDRepos {
	return &GetTeamsTeamIDRepos{Context: ctx, Handler: handler}
}

/*GetTeamsTeamIDRepos swagger:route GET /teams/{teamId}/repos getTeamsTeamIdRepos

List team repos

*/
type GetTeamsTeamIDRepos struct {
	Context *middleware.Context
	Handler GetTeamsTeamIDReposHandler
}

func (o *GetTeamsTeamIDRepos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTeamsTeamIDReposParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
