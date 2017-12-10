// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoGitRefsHandlerFunc turns a function with the right signature into a get repos owner repo git refs handler
type GetReposOwnerRepoGitRefsHandlerFunc func(GetReposOwnerRepoGitRefsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoGitRefsHandlerFunc) Handle(params GetReposOwnerRepoGitRefsParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoGitRefsHandler interface for that can handle valid get repos owner repo git refs params
type GetReposOwnerRepoGitRefsHandler interface {
	Handle(GetReposOwnerRepoGitRefsParams) middleware.Responder
}

// NewGetReposOwnerRepoGitRefs creates a new http.Handler for the get repos owner repo git refs operation
func NewGetReposOwnerRepoGitRefs(ctx *middleware.Context, handler GetReposOwnerRepoGitRefsHandler) *GetReposOwnerRepoGitRefs {
	return &GetReposOwnerRepoGitRefs{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoGitRefs swagger:route GET /repos/{owner}/{repo}/git/refs getReposOwnerRepoGitRefs

Get all References

*/
type GetReposOwnerRepoGitRefs struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoGitRefsHandler
}

func (o *GetReposOwnerRepoGitRefs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoGitRefsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}