// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoPullsNumberFilesHandlerFunc turns a function with the right signature into a get repos owner repo pulls number files handler
type GetReposOwnerRepoPullsNumberFilesHandlerFunc func(GetReposOwnerRepoPullsNumberFilesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoPullsNumberFilesHandlerFunc) Handle(params GetReposOwnerRepoPullsNumberFilesParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoPullsNumberFilesHandler interface for that can handle valid get repos owner repo pulls number files params
type GetReposOwnerRepoPullsNumberFilesHandler interface {
	Handle(GetReposOwnerRepoPullsNumberFilesParams) middleware.Responder
}

// NewGetReposOwnerRepoPullsNumberFiles creates a new http.Handler for the get repos owner repo pulls number files operation
func NewGetReposOwnerRepoPullsNumberFiles(ctx *middleware.Context, handler GetReposOwnerRepoPullsNumberFilesHandler) *GetReposOwnerRepoPullsNumberFiles {
	return &GetReposOwnerRepoPullsNumberFiles{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoPullsNumberFiles swagger:route GET /repos/{owner}/{repo}/pulls/{number}/files getReposOwnerRepoPullsNumberFiles

List pull requests files.

*/
type GetReposOwnerRepoPullsNumberFiles struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoPullsNumberFilesHandler
}

func (o *GetReposOwnerRepoPullsNumberFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoPullsNumberFilesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
