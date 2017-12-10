// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoBranchesBranchHandlerFunc turns a function with the right signature into a get repos owner repo branches branch handler
type GetReposOwnerRepoBranchesBranchHandlerFunc func(GetReposOwnerRepoBranchesBranchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoBranchesBranchHandlerFunc) Handle(params GetReposOwnerRepoBranchesBranchParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoBranchesBranchHandler interface for that can handle valid get repos owner repo branches branch params
type GetReposOwnerRepoBranchesBranchHandler interface {
	Handle(GetReposOwnerRepoBranchesBranchParams) middleware.Responder
}

// NewGetReposOwnerRepoBranchesBranch creates a new http.Handler for the get repos owner repo branches branch operation
func NewGetReposOwnerRepoBranchesBranch(ctx *middleware.Context, handler GetReposOwnerRepoBranchesBranchHandler) *GetReposOwnerRepoBranchesBranch {
	return &GetReposOwnerRepoBranchesBranch{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoBranchesBranch swagger:route GET /repos/{owner}/{repo}/branches/{branch} getReposOwnerRepoBranchesBranch

Get Branch

*/
type GetReposOwnerRepoBranchesBranch struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoBranchesBranchHandler
}

func (o *GetReposOwnerRepoBranchesBranch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoBranchesBranchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}