// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoIssuesEventsHandlerFunc turns a function with the right signature into a get repos owner repo issues events handler
type GetReposOwnerRepoIssuesEventsHandlerFunc func(GetReposOwnerRepoIssuesEventsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoIssuesEventsHandlerFunc) Handle(params GetReposOwnerRepoIssuesEventsParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoIssuesEventsHandler interface for that can handle valid get repos owner repo issues events params
type GetReposOwnerRepoIssuesEventsHandler interface {
	Handle(GetReposOwnerRepoIssuesEventsParams) middleware.Responder
}

// NewGetReposOwnerRepoIssuesEvents creates a new http.Handler for the get repos owner repo issues events operation
func NewGetReposOwnerRepoIssuesEvents(ctx *middleware.Context, handler GetReposOwnerRepoIssuesEventsHandler) *GetReposOwnerRepoIssuesEvents {
	return &GetReposOwnerRepoIssuesEvents{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoIssuesEvents swagger:route GET /repos/{owner}/{repo}/issues/events getReposOwnerRepoIssuesEvents

List issue events for a repository.

*/
type GetReposOwnerRepoIssuesEvents struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoIssuesEventsHandler
}

func (o *GetReposOwnerRepoIssuesEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoIssuesEventsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
