// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGitignoreTemplatesLanguageHandlerFunc turns a function with the right signature into a get gitignore templates language handler
type GetGitignoreTemplatesLanguageHandlerFunc func(GetGitignoreTemplatesLanguageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGitignoreTemplatesLanguageHandlerFunc) Handle(params GetGitignoreTemplatesLanguageParams) middleware.Responder {
	return fn(params)
}

// GetGitignoreTemplatesLanguageHandler interface for that can handle valid get gitignore templates language params
type GetGitignoreTemplatesLanguageHandler interface {
	Handle(GetGitignoreTemplatesLanguageParams) middleware.Responder
}

// NewGetGitignoreTemplatesLanguage creates a new http.Handler for the get gitignore templates language operation
func NewGetGitignoreTemplatesLanguage(ctx *middleware.Context, handler GetGitignoreTemplatesLanguageHandler) *GetGitignoreTemplatesLanguage {
	return &GetGitignoreTemplatesLanguage{Context: ctx, Handler: handler}
}

/*GetGitignoreTemplatesLanguage swagger:route GET /gitignore/templates/{language} getGitignoreTemplatesLanguage

Get a single template.

*/
type GetGitignoreTemplatesLanguage struct {
	Context *middleware.Context
	Handler GetGitignoreTemplatesLanguageHandler
}

func (o *GetGitignoreTemplatesLanguage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGitignoreTemplatesLanguageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
