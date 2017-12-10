// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostReposOwnerRepoDeploymentsIDStatusesHandlerFunc turns a function with the right signature into a post repos owner repo deployments ID statuses handler
type PostReposOwnerRepoDeploymentsIDStatusesHandlerFunc func(PostReposOwnerRepoDeploymentsIDStatusesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostReposOwnerRepoDeploymentsIDStatusesHandlerFunc) Handle(params PostReposOwnerRepoDeploymentsIDStatusesParams) middleware.Responder {
	return fn(params)
}

// PostReposOwnerRepoDeploymentsIDStatusesHandler interface for that can handle valid post repos owner repo deployments ID statuses params
type PostReposOwnerRepoDeploymentsIDStatusesHandler interface {
	Handle(PostReposOwnerRepoDeploymentsIDStatusesParams) middleware.Responder
}

// NewPostReposOwnerRepoDeploymentsIDStatuses creates a new http.Handler for the post repos owner repo deployments ID statuses operation
func NewPostReposOwnerRepoDeploymentsIDStatuses(ctx *middleware.Context, handler PostReposOwnerRepoDeploymentsIDStatusesHandler) *PostReposOwnerRepoDeploymentsIDStatuses {
	return &PostReposOwnerRepoDeploymentsIDStatuses{Context: ctx, Handler: handler}
}

/*PostReposOwnerRepoDeploymentsIDStatuses swagger:route POST /repos/{owner}/{repo}/deployments/{id}/statuses postReposOwnerRepoDeploymentsIdStatuses

Create a Deployment Status
Users with push access can create deployment statuses for a given deployment:


*/
type PostReposOwnerRepoDeploymentsIDStatuses struct {
	Context *middleware.Context
	Handler PostReposOwnerRepoDeploymentsIDStatusesHandler
}

func (o *PostReposOwnerRepoDeploymentsIDStatuses) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostReposOwnerRepoDeploymentsIDStatusesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
