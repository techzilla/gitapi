// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoIssuesNumberHandlerFunc turns a function with the right signature into a get repos owner repo issues number handler
type GetReposOwnerRepoIssuesNumberHandlerFunc func(GetReposOwnerRepoIssuesNumberParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoIssuesNumberHandlerFunc) Handle(params GetReposOwnerRepoIssuesNumberParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoIssuesNumberHandler interface for that can handle valid get repos owner repo issues number params
type GetReposOwnerRepoIssuesNumberHandler interface {
	Handle(GetReposOwnerRepoIssuesNumberParams) middleware.Responder
}

// NewGetReposOwnerRepoIssuesNumber creates a new http.Handler for the get repos owner repo issues number operation
func NewGetReposOwnerRepoIssuesNumber(ctx *middleware.Context, handler GetReposOwnerRepoIssuesNumberHandler) *GetReposOwnerRepoIssuesNumber {
	return &GetReposOwnerRepoIssuesNumber{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoIssuesNumber swagger:route GET /repos/{owner}/{repo}/issues/{number} getReposOwnerRepoIssuesNumber

Get a single issue

*/
type GetReposOwnerRepoIssuesNumber struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoIssuesNumberHandler
}

func (o *GetReposOwnerRepoIssuesNumber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoIssuesNumberParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
