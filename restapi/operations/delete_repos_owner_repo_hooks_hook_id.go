// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReposOwnerRepoHooksHookIDHandlerFunc turns a function with the right signature into a delete repos owner repo hooks hook ID handler
type DeleteReposOwnerRepoHooksHookIDHandlerFunc func(DeleteReposOwnerRepoHooksHookIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReposOwnerRepoHooksHookIDHandlerFunc) Handle(params DeleteReposOwnerRepoHooksHookIDParams) middleware.Responder {
	return fn(params)
}

// DeleteReposOwnerRepoHooksHookIDHandler interface for that can handle valid delete repos owner repo hooks hook ID params
type DeleteReposOwnerRepoHooksHookIDHandler interface {
	Handle(DeleteReposOwnerRepoHooksHookIDParams) middleware.Responder
}

// NewDeleteReposOwnerRepoHooksHookID creates a new http.Handler for the delete repos owner repo hooks hook ID operation
func NewDeleteReposOwnerRepoHooksHookID(ctx *middleware.Context, handler DeleteReposOwnerRepoHooksHookIDHandler) *DeleteReposOwnerRepoHooksHookID {
	return &DeleteReposOwnerRepoHooksHookID{Context: ctx, Handler: handler}
}

/*DeleteReposOwnerRepoHooksHookID swagger:route DELETE /repos/{owner}/{repo}/hooks/{hookId} deleteReposOwnerRepoHooksHookId

Delete a hook.

*/
type DeleteReposOwnerRepoHooksHookID struct {
	Context *middleware.Context
	Handler DeleteReposOwnerRepoHooksHookIDHandler
}

func (o *DeleteReposOwnerRepoHooksHookID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReposOwnerRepoHooksHookIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}