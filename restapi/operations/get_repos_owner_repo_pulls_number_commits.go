// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoPullsNumberCommitsHandlerFunc turns a function with the right signature into a get repos owner repo pulls number commits handler
type GetReposOwnerRepoPullsNumberCommitsHandlerFunc func(GetReposOwnerRepoPullsNumberCommitsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoPullsNumberCommitsHandlerFunc) Handle(params GetReposOwnerRepoPullsNumberCommitsParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoPullsNumberCommitsHandler interface for that can handle valid get repos owner repo pulls number commits params
type GetReposOwnerRepoPullsNumberCommitsHandler interface {
	Handle(GetReposOwnerRepoPullsNumberCommitsParams) middleware.Responder
}

// NewGetReposOwnerRepoPullsNumberCommits creates a new http.Handler for the get repos owner repo pulls number commits operation
func NewGetReposOwnerRepoPullsNumberCommits(ctx *middleware.Context, handler GetReposOwnerRepoPullsNumberCommitsHandler) *GetReposOwnerRepoPullsNumberCommits {
	return &GetReposOwnerRepoPullsNumberCommits{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoPullsNumberCommits swagger:route GET /repos/{owner}/{repo}/pulls/{number}/commits getReposOwnerRepoPullsNumberCommits

List commits on a pull request.

*/
type GetReposOwnerRepoPullsNumberCommits struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoPullsNumberCommitsHandler
}

func (o *GetReposOwnerRepoPullsNumberCommits) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoPullsNumberCommitsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
