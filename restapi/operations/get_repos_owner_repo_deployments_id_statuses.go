// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoDeploymentsIDStatusesHandlerFunc turns a function with the right signature into a get repos owner repo deployments ID statuses handler
type GetReposOwnerRepoDeploymentsIDStatusesHandlerFunc func(GetReposOwnerRepoDeploymentsIDStatusesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoDeploymentsIDStatusesHandlerFunc) Handle(params GetReposOwnerRepoDeploymentsIDStatusesParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoDeploymentsIDStatusesHandler interface for that can handle valid get repos owner repo deployments ID statuses params
type GetReposOwnerRepoDeploymentsIDStatusesHandler interface {
	Handle(GetReposOwnerRepoDeploymentsIDStatusesParams) middleware.Responder
}

// NewGetReposOwnerRepoDeploymentsIDStatuses creates a new http.Handler for the get repos owner repo deployments ID statuses operation
func NewGetReposOwnerRepoDeploymentsIDStatuses(ctx *middleware.Context, handler GetReposOwnerRepoDeploymentsIDStatusesHandler) *GetReposOwnerRepoDeploymentsIDStatuses {
	return &GetReposOwnerRepoDeploymentsIDStatuses{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoDeploymentsIDStatuses swagger:route GET /repos/{owner}/{repo}/deployments/{id}/statuses getReposOwnerRepoDeploymentsIdStatuses

Users with pull access can view deployment statuses for a deployment

*/
type GetReposOwnerRepoDeploymentsIDStatuses struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoDeploymentsIDStatusesHandler
}

func (o *GetReposOwnerRepoDeploymentsIDStatuses) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoDeploymentsIDStatusesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}