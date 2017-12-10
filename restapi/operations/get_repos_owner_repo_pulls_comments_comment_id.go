// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoPullsCommentsCommentIDHandlerFunc turns a function with the right signature into a get repos owner repo pulls comments comment ID handler
type GetReposOwnerRepoPullsCommentsCommentIDHandlerFunc func(GetReposOwnerRepoPullsCommentsCommentIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoPullsCommentsCommentIDHandlerFunc) Handle(params GetReposOwnerRepoPullsCommentsCommentIDParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoPullsCommentsCommentIDHandler interface for that can handle valid get repos owner repo pulls comments comment ID params
type GetReposOwnerRepoPullsCommentsCommentIDHandler interface {
	Handle(GetReposOwnerRepoPullsCommentsCommentIDParams) middleware.Responder
}

// NewGetReposOwnerRepoPullsCommentsCommentID creates a new http.Handler for the get repos owner repo pulls comments comment ID operation
func NewGetReposOwnerRepoPullsCommentsCommentID(ctx *middleware.Context, handler GetReposOwnerRepoPullsCommentsCommentIDHandler) *GetReposOwnerRepoPullsCommentsCommentID {
	return &GetReposOwnerRepoPullsCommentsCommentID{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoPullsCommentsCommentID swagger:route GET /repos/{owner}/{repo}/pulls/comments/{commentId} getReposOwnerRepoPullsCommentsCommentId

Get a single comment.

*/
type GetReposOwnerRepoPullsCommentsCommentID struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoPullsCommentsCommentIDHandler
}

func (o *GetReposOwnerRepoPullsCommentsCommentID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoPullsCommentsCommentIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}