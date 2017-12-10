// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostReposOwnerRepoDeploymentsHandlerFunc turns a function with the right signature into a post repos owner repo deployments handler
type PostReposOwnerRepoDeploymentsHandlerFunc func(PostReposOwnerRepoDeploymentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostReposOwnerRepoDeploymentsHandlerFunc) Handle(params PostReposOwnerRepoDeploymentsParams) middleware.Responder {
	return fn(params)
}

// PostReposOwnerRepoDeploymentsHandler interface for that can handle valid post repos owner repo deployments params
type PostReposOwnerRepoDeploymentsHandler interface {
	Handle(PostReposOwnerRepoDeploymentsParams) middleware.Responder
}

// NewPostReposOwnerRepoDeployments creates a new http.Handler for the post repos owner repo deployments operation
func NewPostReposOwnerRepoDeployments(ctx *middleware.Context, handler PostReposOwnerRepoDeploymentsHandler) *PostReposOwnerRepoDeployments {
	return &PostReposOwnerRepoDeployments{Context: ctx, Handler: handler}
}

/*PostReposOwnerRepoDeployments swagger:route POST /repos/{owner}/{repo}/deployments postReposOwnerRepoDeployments

Users with push access can create a deployment for a given ref

*/
type PostReposOwnerRepoDeployments struct {
	Context *middleware.Context
	Handler PostReposOwnerRepoDeploymentsHandler
}

func (o *PostReposOwnerRepoDeployments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostReposOwnerRepoDeploymentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
