// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserSubscriptionsOwnerRepoHandlerFunc turns a function with the right signature into a get user subscriptions owner repo handler
type GetUserSubscriptionsOwnerRepoHandlerFunc func(GetUserSubscriptionsOwnerRepoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserSubscriptionsOwnerRepoHandlerFunc) Handle(params GetUserSubscriptionsOwnerRepoParams) middleware.Responder {
	return fn(params)
}

// GetUserSubscriptionsOwnerRepoHandler interface for that can handle valid get user subscriptions owner repo params
type GetUserSubscriptionsOwnerRepoHandler interface {
	Handle(GetUserSubscriptionsOwnerRepoParams) middleware.Responder
}

// NewGetUserSubscriptionsOwnerRepo creates a new http.Handler for the get user subscriptions owner repo operation
func NewGetUserSubscriptionsOwnerRepo(ctx *middleware.Context, handler GetUserSubscriptionsOwnerRepoHandler) *GetUserSubscriptionsOwnerRepo {
	return &GetUserSubscriptionsOwnerRepo{Context: ctx, Handler: handler}
}

/*GetUserSubscriptionsOwnerRepo swagger:route GET /user/subscriptions/{owner}/{repo} getUserSubscriptionsOwnerRepo

Check if you are watching a repository.

*/
type GetUserSubscriptionsOwnerRepo struct {
	Context *middleware.Context
	Handler GetUserSubscriptionsOwnerRepoHandler
}

func (o *GetUserSubscriptionsOwnerRepo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserSubscriptionsOwnerRepoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}