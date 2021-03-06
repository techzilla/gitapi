// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoKeysKeyIDHandlerFunc turns a function with the right signature into a get repos owner repo keys key ID handler
type GetReposOwnerRepoKeysKeyIDHandlerFunc func(GetReposOwnerRepoKeysKeyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoKeysKeyIDHandlerFunc) Handle(params GetReposOwnerRepoKeysKeyIDParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoKeysKeyIDHandler interface for that can handle valid get repos owner repo keys key ID params
type GetReposOwnerRepoKeysKeyIDHandler interface {
	Handle(GetReposOwnerRepoKeysKeyIDParams) middleware.Responder
}

// NewGetReposOwnerRepoKeysKeyID creates a new http.Handler for the get repos owner repo keys key ID operation
func NewGetReposOwnerRepoKeysKeyID(ctx *middleware.Context, handler GetReposOwnerRepoKeysKeyIDHandler) *GetReposOwnerRepoKeysKeyID {
	return &GetReposOwnerRepoKeysKeyID{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoKeysKeyID swagger:route GET /repos/{owner}/{repo}/keys/{keyId} getReposOwnerRepoKeysKeyId

Get a key

*/
type GetReposOwnerRepoKeysKeyID struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoKeysKeyIDHandler
}

func (o *GetReposOwnerRepoKeysKeyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoKeysKeyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
