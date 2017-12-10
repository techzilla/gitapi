// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGistsIDCommentsHandlerFunc turns a function with the right signature into a get gists ID comments handler
type GetGistsIDCommentsHandlerFunc func(GetGistsIDCommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGistsIDCommentsHandlerFunc) Handle(params GetGistsIDCommentsParams) middleware.Responder {
	return fn(params)
}

// GetGistsIDCommentsHandler interface for that can handle valid get gists ID comments params
type GetGistsIDCommentsHandler interface {
	Handle(GetGistsIDCommentsParams) middleware.Responder
}

// NewGetGistsIDComments creates a new http.Handler for the get gists ID comments operation
func NewGetGistsIDComments(ctx *middleware.Context, handler GetGistsIDCommentsHandler) *GetGistsIDComments {
	return &GetGistsIDComments{Context: ctx, Handler: handler}
}

/*GetGistsIDComments swagger:route GET /gists/{id}/comments getGistsIdComments

List comments on a gist.

*/
type GetGistsIDComments struct {
	Context *middleware.Context
	Handler GetGistsIDCommentsHandler
}

func (o *GetGistsIDComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGistsIDCommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
