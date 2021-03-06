// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoReadmeHandlerFunc turns a function with the right signature into a get repos owner repo readme handler
type GetReposOwnerRepoReadmeHandlerFunc func(GetReposOwnerRepoReadmeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoReadmeHandlerFunc) Handle(params GetReposOwnerRepoReadmeParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoReadmeHandler interface for that can handle valid get repos owner repo readme params
type GetReposOwnerRepoReadmeHandler interface {
	Handle(GetReposOwnerRepoReadmeParams) middleware.Responder
}

// NewGetReposOwnerRepoReadme creates a new http.Handler for the get repos owner repo readme operation
func NewGetReposOwnerRepoReadme(ctx *middleware.Context, handler GetReposOwnerRepoReadmeHandler) *GetReposOwnerRepoReadme {
	return &GetReposOwnerRepoReadme{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoReadme swagger:route GET /repos/{owner}/{repo}/readme getReposOwnerRepoReadme

Get the README.
This method returns the preferred README for a repository.


*/
type GetReposOwnerRepoReadme struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoReadmeHandler
}

func (o *GetReposOwnerRepoReadme) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoReadmeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
