// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetIssuesHandlerFunc turns a function with the right signature into a get issues handler
type GetIssuesHandlerFunc func(GetIssuesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIssuesHandlerFunc) Handle(params GetIssuesParams) middleware.Responder {
	return fn(params)
}

// GetIssuesHandler interface for that can handle valid get issues params
type GetIssuesHandler interface {
	Handle(GetIssuesParams) middleware.Responder
}

// NewGetIssues creates a new http.Handler for the get issues operation
func NewGetIssues(ctx *middleware.Context, handler GetIssuesHandler) *GetIssues {
	return &GetIssues{Context: ctx, Handler: handler}
}

/*GetIssues swagger:route GET /issues getIssues

List issues.
List all issues across all the authenticated user's visible repositories.


*/
type GetIssues struct {
	Context *middleware.Context
	Handler GetIssuesHandler
}

func (o *GetIssues) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetIssuesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}