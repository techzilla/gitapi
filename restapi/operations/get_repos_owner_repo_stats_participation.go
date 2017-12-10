// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoStatsParticipationHandlerFunc turns a function with the right signature into a get repos owner repo stats participation handler
type GetReposOwnerRepoStatsParticipationHandlerFunc func(GetReposOwnerRepoStatsParticipationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoStatsParticipationHandlerFunc) Handle(params GetReposOwnerRepoStatsParticipationParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoStatsParticipationHandler interface for that can handle valid get repos owner repo stats participation params
type GetReposOwnerRepoStatsParticipationHandler interface {
	Handle(GetReposOwnerRepoStatsParticipationParams) middleware.Responder
}

// NewGetReposOwnerRepoStatsParticipation creates a new http.Handler for the get repos owner repo stats participation operation
func NewGetReposOwnerRepoStatsParticipation(ctx *middleware.Context, handler GetReposOwnerRepoStatsParticipationHandler) *GetReposOwnerRepoStatsParticipation {
	return &GetReposOwnerRepoStatsParticipation{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoStatsParticipation swagger:route GET /repos/{owner}/{repo}/stats/participation getReposOwnerRepoStatsParticipation

Get the weekly commit count for the repo owner and everyone else.

*/
type GetReposOwnerRepoStatsParticipation struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoStatsParticipationHandler
}

func (o *GetReposOwnerRepoStatsParticipation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoStatsParticipationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
