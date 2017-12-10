// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostReposOwnerRepoIssuesNumberCommentsHandlerFunc turns a function with the right signature into a post repos owner repo issues number comments handler
type PostReposOwnerRepoIssuesNumberCommentsHandlerFunc func(PostReposOwnerRepoIssuesNumberCommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostReposOwnerRepoIssuesNumberCommentsHandlerFunc) Handle(params PostReposOwnerRepoIssuesNumberCommentsParams) middleware.Responder {
	return fn(params)
}

// PostReposOwnerRepoIssuesNumberCommentsHandler interface for that can handle valid post repos owner repo issues number comments params
type PostReposOwnerRepoIssuesNumberCommentsHandler interface {
	Handle(PostReposOwnerRepoIssuesNumberCommentsParams) middleware.Responder
}

// NewPostReposOwnerRepoIssuesNumberComments creates a new http.Handler for the post repos owner repo issues number comments operation
func NewPostReposOwnerRepoIssuesNumberComments(ctx *middleware.Context, handler PostReposOwnerRepoIssuesNumberCommentsHandler) *PostReposOwnerRepoIssuesNumberComments {
	return &PostReposOwnerRepoIssuesNumberComments{Context: ctx, Handler: handler}
}

/*PostReposOwnerRepoIssuesNumberComments swagger:route POST /repos/{owner}/{repo}/issues/{number}/comments postReposOwnerRepoIssuesNumberComments

Create a comment.

*/
type PostReposOwnerRepoIssuesNumberComments struct {
	Context *middleware.Context
	Handler PostReposOwnerRepoIssuesNumberCommentsHandler
}

func (o *PostReposOwnerRepoIssuesNumberComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostReposOwnerRepoIssuesNumberCommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}