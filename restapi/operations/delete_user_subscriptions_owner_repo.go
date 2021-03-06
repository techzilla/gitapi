// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserSubscriptionsOwnerRepoHandlerFunc turns a function with the right signature into a delete user subscriptions owner repo handler
type DeleteUserSubscriptionsOwnerRepoHandlerFunc func(DeleteUserSubscriptionsOwnerRepoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserSubscriptionsOwnerRepoHandlerFunc) Handle(params DeleteUserSubscriptionsOwnerRepoParams) middleware.Responder {
	return fn(params)
}

// DeleteUserSubscriptionsOwnerRepoHandler interface for that can handle valid delete user subscriptions owner repo params
type DeleteUserSubscriptionsOwnerRepoHandler interface {
	Handle(DeleteUserSubscriptionsOwnerRepoParams) middleware.Responder
}

// NewDeleteUserSubscriptionsOwnerRepo creates a new http.Handler for the delete user subscriptions owner repo operation
func NewDeleteUserSubscriptionsOwnerRepo(ctx *middleware.Context, handler DeleteUserSubscriptionsOwnerRepoHandler) *DeleteUserSubscriptionsOwnerRepo {
	return &DeleteUserSubscriptionsOwnerRepo{Context: ctx, Handler: handler}
}

/*DeleteUserSubscriptionsOwnerRepo swagger:route DELETE /user/subscriptions/{owner}/{repo} deleteUserSubscriptionsOwnerRepo

Stop watching a repository

*/
type DeleteUserSubscriptionsOwnerRepo struct {
	Context *middleware.Context
	Handler DeleteUserSubscriptionsOwnerRepoHandler
}

func (o *DeleteUserSubscriptionsOwnerRepo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserSubscriptionsOwnerRepoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
