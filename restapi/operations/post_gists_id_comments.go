// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostGistsIDCommentsHandlerFunc turns a function with the right signature into a post gists ID comments handler
type PostGistsIDCommentsHandlerFunc func(PostGistsIDCommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGistsIDCommentsHandlerFunc) Handle(params PostGistsIDCommentsParams) middleware.Responder {
	return fn(params)
}

// PostGistsIDCommentsHandler interface for that can handle valid post gists ID comments params
type PostGistsIDCommentsHandler interface {
	Handle(PostGistsIDCommentsParams) middleware.Responder
}

// NewPostGistsIDComments creates a new http.Handler for the post gists ID comments operation
func NewPostGistsIDComments(ctx *middleware.Context, handler PostGistsIDCommentsHandler) *PostGistsIDComments {
	return &PostGistsIDComments{Context: ctx, Handler: handler}
}

/*PostGistsIDComments swagger:route POST /gists/{id}/comments postGistsIdComments

Create a commen

*/
type PostGistsIDComments struct {
	Context *middleware.Context
	Handler PostGistsIDCommentsHandler
}

func (o *PostGistsIDComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostGistsIDCommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
