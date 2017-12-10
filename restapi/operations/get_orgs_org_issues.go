// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOrgsOrgIssuesHandlerFunc turns a function with the right signature into a get orgs org issues handler
type GetOrgsOrgIssuesHandlerFunc func(GetOrgsOrgIssuesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOrgsOrgIssuesHandlerFunc) Handle(params GetOrgsOrgIssuesParams) middleware.Responder {
	return fn(params)
}

// GetOrgsOrgIssuesHandler interface for that can handle valid get orgs org issues params
type GetOrgsOrgIssuesHandler interface {
	Handle(GetOrgsOrgIssuesParams) middleware.Responder
}

// NewGetOrgsOrgIssues creates a new http.Handler for the get orgs org issues operation
func NewGetOrgsOrgIssues(ctx *middleware.Context, handler GetOrgsOrgIssuesHandler) *GetOrgsOrgIssues {
	return &GetOrgsOrgIssues{Context: ctx, Handler: handler}
}

/*GetOrgsOrgIssues swagger:route GET /orgs/{org}/issues getOrgsOrgIssues

List issues.
List all issues for a given organization for the authenticated user.


*/
type GetOrgsOrgIssues struct {
	Context *middleware.Context
	Handler GetOrgsOrgIssuesHandler
}

func (o *GetOrgsOrgIssues) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOrgsOrgIssuesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}