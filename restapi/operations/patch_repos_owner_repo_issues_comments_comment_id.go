// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchReposOwnerRepoIssuesCommentsCommentIDHandlerFunc turns a function with the right signature into a patch repos owner repo issues comments comment ID handler
type PatchReposOwnerRepoIssuesCommentsCommentIDHandlerFunc func(PatchReposOwnerRepoIssuesCommentsCommentIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchReposOwnerRepoIssuesCommentsCommentIDHandlerFunc) Handle(params PatchReposOwnerRepoIssuesCommentsCommentIDParams) middleware.Responder {
	return fn(params)
}

// PatchReposOwnerRepoIssuesCommentsCommentIDHandler interface for that can handle valid patch repos owner repo issues comments comment ID params
type PatchReposOwnerRepoIssuesCommentsCommentIDHandler interface {
	Handle(PatchReposOwnerRepoIssuesCommentsCommentIDParams) middleware.Responder
}

// NewPatchReposOwnerRepoIssuesCommentsCommentID creates a new http.Handler for the patch repos owner repo issues comments comment ID operation
func NewPatchReposOwnerRepoIssuesCommentsCommentID(ctx *middleware.Context, handler PatchReposOwnerRepoIssuesCommentsCommentIDHandler) *PatchReposOwnerRepoIssuesCommentsCommentID {
	return &PatchReposOwnerRepoIssuesCommentsCommentID{Context: ctx, Handler: handler}
}

/*PatchReposOwnerRepoIssuesCommentsCommentID swagger:route PATCH /repos/{owner}/{repo}/issues/comments/{commentId} patchReposOwnerRepoIssuesCommentsCommentId

Edit a comment.

*/
type PatchReposOwnerRepoIssuesCommentsCommentID struct {
	Context *middleware.Context
	Handler PatchReposOwnerRepoIssuesCommentsCommentIDHandler
}

func (o *PatchReposOwnerRepoIssuesCommentsCommentID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchReposOwnerRepoIssuesCommentsCommentIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}