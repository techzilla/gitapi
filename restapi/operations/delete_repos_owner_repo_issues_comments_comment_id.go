// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReposOwnerRepoIssuesCommentsCommentIDHandlerFunc turns a function with the right signature into a delete repos owner repo issues comments comment ID handler
type DeleteReposOwnerRepoIssuesCommentsCommentIDHandlerFunc func(DeleteReposOwnerRepoIssuesCommentsCommentIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReposOwnerRepoIssuesCommentsCommentIDHandlerFunc) Handle(params DeleteReposOwnerRepoIssuesCommentsCommentIDParams) middleware.Responder {
	return fn(params)
}

// DeleteReposOwnerRepoIssuesCommentsCommentIDHandler interface for that can handle valid delete repos owner repo issues comments comment ID params
type DeleteReposOwnerRepoIssuesCommentsCommentIDHandler interface {
	Handle(DeleteReposOwnerRepoIssuesCommentsCommentIDParams) middleware.Responder
}

// NewDeleteReposOwnerRepoIssuesCommentsCommentID creates a new http.Handler for the delete repos owner repo issues comments comment ID operation
func NewDeleteReposOwnerRepoIssuesCommentsCommentID(ctx *middleware.Context, handler DeleteReposOwnerRepoIssuesCommentsCommentIDHandler) *DeleteReposOwnerRepoIssuesCommentsCommentID {
	return &DeleteReposOwnerRepoIssuesCommentsCommentID{Context: ctx, Handler: handler}
}

/*DeleteReposOwnerRepoIssuesCommentsCommentID swagger:route DELETE /repos/{owner}/{repo}/issues/comments/{commentId} deleteReposOwnerRepoIssuesCommentsCommentId

Delete a comment.

*/
type DeleteReposOwnerRepoIssuesCommentsCommentID struct {
	Context *middleware.Context
	Handler DeleteReposOwnerRepoIssuesCommentsCommentIDHandler
}

func (o *DeleteReposOwnerRepoIssuesCommentsCommentID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReposOwnerRepoIssuesCommentsCommentIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
