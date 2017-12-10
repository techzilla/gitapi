// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoStargazersHandlerFunc turns a function with the right signature into a get repos owner repo stargazers handler
type GetReposOwnerRepoStargazersHandlerFunc func(GetReposOwnerRepoStargazersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoStargazersHandlerFunc) Handle(params GetReposOwnerRepoStargazersParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoStargazersHandler interface for that can handle valid get repos owner repo stargazers params
type GetReposOwnerRepoStargazersHandler interface {
	Handle(GetReposOwnerRepoStargazersParams) middleware.Responder
}

// NewGetReposOwnerRepoStargazers creates a new http.Handler for the get repos owner repo stargazers operation
func NewGetReposOwnerRepoStargazers(ctx *middleware.Context, handler GetReposOwnerRepoStargazersHandler) *GetReposOwnerRepoStargazers {
	return &GetReposOwnerRepoStargazers{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoStargazers swagger:route GET /repos/{owner}/{repo}/stargazers getReposOwnerRepoStargazers

List Stargazers.

*/
type GetReposOwnerRepoStargazers struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoStargazersHandler
}

func (o *GetReposOwnerRepoStargazers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoStargazersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
