// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRepositoriesHandlerFunc turns a function with the right signature into a get repositories handler
type GetRepositoriesHandlerFunc func(GetRepositoriesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRepositoriesHandlerFunc) Handle(params GetRepositoriesParams) middleware.Responder {
	return fn(params)
}

// GetRepositoriesHandler interface for that can handle valid get repositories params
type GetRepositoriesHandler interface {
	Handle(GetRepositoriesParams) middleware.Responder
}

// NewGetRepositories creates a new http.Handler for the get repositories operation
func NewGetRepositories(ctx *middleware.Context, handler GetRepositoriesHandler) *GetRepositories {
	return &GetRepositories{Context: ctx, Handler: handler}
}

/*GetRepositories swagger:route GET /repositories getRepositories

List all public repositories.
This provides a dump of every public repository, in the order that they
were created.
Note: Pagination is powered exclusively by the since parameter. is the
Link header to get the URL for the next page of repositories.


*/
type GetRepositories struct {
	Context *middleware.Context
	Handler GetRepositoriesHandler
}

func (o *GetRepositories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRepositoriesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
