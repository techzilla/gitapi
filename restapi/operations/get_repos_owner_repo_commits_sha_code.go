// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoCommitsShaCodeHandlerFunc turns a function with the right signature into a get repos owner repo commits sha code handler
type GetReposOwnerRepoCommitsShaCodeHandlerFunc func(GetReposOwnerRepoCommitsShaCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoCommitsShaCodeHandlerFunc) Handle(params GetReposOwnerRepoCommitsShaCodeParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoCommitsShaCodeHandler interface for that can handle valid get repos owner repo commits sha code params
type GetReposOwnerRepoCommitsShaCodeHandler interface {
	Handle(GetReposOwnerRepoCommitsShaCodeParams) middleware.Responder
}

// NewGetReposOwnerRepoCommitsShaCode creates a new http.Handler for the get repos owner repo commits sha code operation
func NewGetReposOwnerRepoCommitsShaCode(ctx *middleware.Context, handler GetReposOwnerRepoCommitsShaCodeHandler) *GetReposOwnerRepoCommitsShaCode {
	return &GetReposOwnerRepoCommitsShaCode{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoCommitsShaCode swagger:route GET /repos/{owner}/{repo}/commits/{shaCode} getReposOwnerRepoCommitsShaCode

Get a single commit.

*/
type GetReposOwnerRepoCommitsShaCode struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoCommitsShaCodeHandler
}

func (o *GetReposOwnerRepoCommitsShaCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoCommitsShaCodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}