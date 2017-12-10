// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoIssuesCommentsHandlerFunc turns a function with the right signature into a get repos owner repo issues comments handler
type GetReposOwnerRepoIssuesCommentsHandlerFunc func(GetReposOwnerRepoIssuesCommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoIssuesCommentsHandlerFunc) Handle(params GetReposOwnerRepoIssuesCommentsParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoIssuesCommentsHandler interface for that can handle valid get repos owner repo issues comments params
type GetReposOwnerRepoIssuesCommentsHandler interface {
	Handle(GetReposOwnerRepoIssuesCommentsParams) middleware.Responder
}

// NewGetReposOwnerRepoIssuesComments creates a new http.Handler for the get repos owner repo issues comments operation
func NewGetReposOwnerRepoIssuesComments(ctx *middleware.Context, handler GetReposOwnerRepoIssuesCommentsHandler) *GetReposOwnerRepoIssuesComments {
	return &GetReposOwnerRepoIssuesComments{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoIssuesComments swagger:route GET /repos/{owner}/{repo}/issues/comments getReposOwnerRepoIssuesComments

List comments in a repository.

*/
type GetReposOwnerRepoIssuesComments struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoIssuesCommentsHandler
}

func (o *GetReposOwnerRepoIssuesComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoIssuesCommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}