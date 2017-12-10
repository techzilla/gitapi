// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostMarkdownHandlerFunc turns a function with the right signature into a post markdown handler
type PostMarkdownHandlerFunc func(PostMarkdownParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMarkdownHandlerFunc) Handle(params PostMarkdownParams) middleware.Responder {
	return fn(params)
}

// PostMarkdownHandler interface for that can handle valid post markdown params
type PostMarkdownHandler interface {
	Handle(PostMarkdownParams) middleware.Responder
}

// NewPostMarkdown creates a new http.Handler for the post markdown operation
func NewPostMarkdown(ctx *middleware.Context, handler PostMarkdownHandler) *PostMarkdown {
	return &PostMarkdown{Context: ctx, Handler: handler}
}

/*PostMarkdown swagger:route POST /markdown postMarkdown

Render an arbitrary Markdown document

*/
type PostMarkdown struct {
	Context *middleware.Context
	Handler PostMarkdownHandler
}

func (o *PostMarkdown) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMarkdownParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
