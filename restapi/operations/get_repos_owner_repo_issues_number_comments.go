// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoIssuesNumberCommentsHandlerFunc turns a function with the right signature into a get repos owner repo issues number comments handler
type GetReposOwnerRepoIssuesNumberCommentsHandlerFunc func(GetReposOwnerRepoIssuesNumberCommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoIssuesNumberCommentsHandlerFunc) Handle(params GetReposOwnerRepoIssuesNumberCommentsParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoIssuesNumberCommentsHandler interface for that can handle valid get repos owner repo issues number comments params
type GetReposOwnerRepoIssuesNumberCommentsHandler interface {
	Handle(GetReposOwnerRepoIssuesNumberCommentsParams) middleware.Responder
}

// NewGetReposOwnerRepoIssuesNumberComments creates a new http.Handler for the get repos owner repo issues number comments operation
func NewGetReposOwnerRepoIssuesNumberComments(ctx *middleware.Context, handler GetReposOwnerRepoIssuesNumberCommentsHandler) *GetReposOwnerRepoIssuesNumberComments {
	return &GetReposOwnerRepoIssuesNumberComments{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoIssuesNumberComments swagger:route GET /repos/{owner}/{repo}/issues/{number}/comments getReposOwnerRepoIssuesNumberComments

List comments on an issue.

*/
type GetReposOwnerRepoIssuesNumberComments struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoIssuesNumberCommentsHandler
}

func (o *GetReposOwnerRepoIssuesNumberComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoIssuesNumberCommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}