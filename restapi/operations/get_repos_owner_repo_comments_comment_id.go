// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoCommentsCommentIDHandlerFunc turns a function with the right signature into a get repos owner repo comments comment ID handler
type GetReposOwnerRepoCommentsCommentIDHandlerFunc func(GetReposOwnerRepoCommentsCommentIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoCommentsCommentIDHandlerFunc) Handle(params GetReposOwnerRepoCommentsCommentIDParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoCommentsCommentIDHandler interface for that can handle valid get repos owner repo comments comment ID params
type GetReposOwnerRepoCommentsCommentIDHandler interface {
	Handle(GetReposOwnerRepoCommentsCommentIDParams) middleware.Responder
}

// NewGetReposOwnerRepoCommentsCommentID creates a new http.Handler for the get repos owner repo comments comment ID operation
func NewGetReposOwnerRepoCommentsCommentID(ctx *middleware.Context, handler GetReposOwnerRepoCommentsCommentIDHandler) *GetReposOwnerRepoCommentsCommentID {
	return &GetReposOwnerRepoCommentsCommentID{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoCommentsCommentID swagger:route GET /repos/{owner}/{repo}/comments/{commentId} getReposOwnerRepoCommentsCommentId

Get a single commit comment.

*/
type GetReposOwnerRepoCommentsCommentID struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoCommentsCommentIDHandler
}

func (o *GetReposOwnerRepoCommentsCommentID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoCommentsCommentIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
