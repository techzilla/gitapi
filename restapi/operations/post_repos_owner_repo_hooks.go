// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostReposOwnerRepoHooksHandlerFunc turns a function with the right signature into a post repos owner repo hooks handler
type PostReposOwnerRepoHooksHandlerFunc func(PostReposOwnerRepoHooksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostReposOwnerRepoHooksHandlerFunc) Handle(params PostReposOwnerRepoHooksParams) middleware.Responder {
	return fn(params)
}

// PostReposOwnerRepoHooksHandler interface for that can handle valid post repos owner repo hooks params
type PostReposOwnerRepoHooksHandler interface {
	Handle(PostReposOwnerRepoHooksParams) middleware.Responder
}

// NewPostReposOwnerRepoHooks creates a new http.Handler for the post repos owner repo hooks operation
func NewPostReposOwnerRepoHooks(ctx *middleware.Context, handler PostReposOwnerRepoHooksHandler) *PostReposOwnerRepoHooks {
	return &PostReposOwnerRepoHooks{Context: ctx, Handler: handler}
}

/*PostReposOwnerRepoHooks swagger:route POST /repos/{owner}/{repo}/hooks postReposOwnerRepoHooks

Create a hook.

*/
type PostReposOwnerRepoHooks struct {
	Context *middleware.Context
	Handler PostReposOwnerRepoHooksHandler
}

func (o *PostReposOwnerRepoHooks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostReposOwnerRepoHooksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}