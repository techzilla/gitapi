// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOrgsOrgReposHandlerFunc turns a function with the right signature into a get orgs org repos handler
type GetOrgsOrgReposHandlerFunc func(GetOrgsOrgReposParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOrgsOrgReposHandlerFunc) Handle(params GetOrgsOrgReposParams) middleware.Responder {
	return fn(params)
}

// GetOrgsOrgReposHandler interface for that can handle valid get orgs org repos params
type GetOrgsOrgReposHandler interface {
	Handle(GetOrgsOrgReposParams) middleware.Responder
}

// NewGetOrgsOrgRepos creates a new http.Handler for the get orgs org repos operation
func NewGetOrgsOrgRepos(ctx *middleware.Context, handler GetOrgsOrgReposHandler) *GetOrgsOrgRepos {
	return &GetOrgsOrgRepos{Context: ctx, Handler: handler}
}

/*GetOrgsOrgRepos swagger:route GET /orgs/{org}/repos getOrgsOrgRepos

List repositories for the specified org.

*/
type GetOrgsOrgRepos struct {
	Context *middleware.Context
	Handler GetOrgsOrgReposHandler
}

func (o *GetOrgsOrgRepos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOrgsOrgReposParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}