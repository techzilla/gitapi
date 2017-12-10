// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostMarkdownRawHandlerFunc turns a function with the right signature into a post markdown raw handler
type PostMarkdownRawHandlerFunc func(PostMarkdownRawParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMarkdownRawHandlerFunc) Handle(params PostMarkdownRawParams) middleware.Responder {
	return fn(params)
}

// PostMarkdownRawHandler interface for that can handle valid post markdown raw params
type PostMarkdownRawHandler interface {
	Handle(PostMarkdownRawParams) middleware.Responder
}

// NewPostMarkdownRaw creates a new http.Handler for the post markdown raw operation
func NewPostMarkdownRaw(ctx *middleware.Context, handler PostMarkdownRawHandler) *PostMarkdownRaw {
	return &PostMarkdownRaw{Context: ctx, Handler: handler}
}

/*PostMarkdownRaw swagger:route POST /markdown/raw postMarkdownRaw

Render a Markdown document in raw mode

*/
type PostMarkdownRaw struct {
	Context *middleware.Context
	Handler PostMarkdownRawHandler
}

func (o *PostMarkdownRaw) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMarkdownRawParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
