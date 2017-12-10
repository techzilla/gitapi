// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGistsIDStarHandlerFunc turns a function with the right signature into a get gists ID star handler
type GetGistsIDStarHandlerFunc func(GetGistsIDStarParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGistsIDStarHandlerFunc) Handle(params GetGistsIDStarParams) middleware.Responder {
	return fn(params)
}

// GetGistsIDStarHandler interface for that can handle valid get gists ID star params
type GetGistsIDStarHandler interface {
	Handle(GetGistsIDStarParams) middleware.Responder
}

// NewGetGistsIDStar creates a new http.Handler for the get gists ID star operation
func NewGetGistsIDStar(ctx *middleware.Context, handler GetGistsIDStarHandler) *GetGistsIDStar {
	return &GetGistsIDStar{Context: ctx, Handler: handler}
}

/*GetGistsIDStar swagger:route GET /gists/{id}/star getGistsIdStar

Check if a gist is starred.

*/
type GetGistsIDStar struct {
	Context *middleware.Context
	Handler GetGistsIDStarHandler
}

func (o *GetGistsIDStar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGistsIDStarParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}