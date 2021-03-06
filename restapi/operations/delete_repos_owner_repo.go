// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReposOwnerRepoHandlerFunc turns a function with the right signature into a delete repos owner repo handler
type DeleteReposOwnerRepoHandlerFunc func(DeleteReposOwnerRepoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReposOwnerRepoHandlerFunc) Handle(params DeleteReposOwnerRepoParams) middleware.Responder {
	return fn(params)
}

// DeleteReposOwnerRepoHandler interface for that can handle valid delete repos owner repo params
type DeleteReposOwnerRepoHandler interface {
	Handle(DeleteReposOwnerRepoParams) middleware.Responder
}

// NewDeleteReposOwnerRepo creates a new http.Handler for the delete repos owner repo operation
func NewDeleteReposOwnerRepo(ctx *middleware.Context, handler DeleteReposOwnerRepoHandler) *DeleteReposOwnerRepo {
	return &DeleteReposOwnerRepo{Context: ctx, Handler: handler}
}

/*DeleteReposOwnerRepo swagger:route DELETE /repos/{owner}/{repo} deleteReposOwnerRepo

Delete a Repository.
Deleting a repository requires admin access. If OAuth is used, the delete_repo
scope is required.


*/
type DeleteReposOwnerRepo struct {
	Context *middleware.Context
	Handler DeleteReposOwnerRepoHandler
}

func (o *DeleteReposOwnerRepo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReposOwnerRepoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
