// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGitignoreTemplatesHandlerFunc turns a function with the right signature into a get gitignore templates handler
type GetGitignoreTemplatesHandlerFunc func(GetGitignoreTemplatesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGitignoreTemplatesHandlerFunc) Handle(params GetGitignoreTemplatesParams) middleware.Responder {
	return fn(params)
}

// GetGitignoreTemplatesHandler interface for that can handle valid get gitignore templates params
type GetGitignoreTemplatesHandler interface {
	Handle(GetGitignoreTemplatesParams) middleware.Responder
}

// NewGetGitignoreTemplates creates a new http.Handler for the get gitignore templates operation
func NewGetGitignoreTemplates(ctx *middleware.Context, handler GetGitignoreTemplatesHandler) *GetGitignoreTemplates {
	return &GetGitignoreTemplates{Context: ctx, Handler: handler}
}

/*GetGitignoreTemplates swagger:route GET /gitignore/templates getGitignoreTemplates

Listing available templates.
List all templates available to pass as an option when creating a repository.


*/
type GetGitignoreTemplates struct {
	Context *middleware.Context
	Handler GetGitignoreTemplatesHandler
}

func (o *GetGitignoreTemplates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGitignoreTemplatesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}