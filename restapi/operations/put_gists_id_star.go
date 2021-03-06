// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutGistsIDStarHandlerFunc turns a function with the right signature into a put gists ID star handler
type PutGistsIDStarHandlerFunc func(PutGistsIDStarParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutGistsIDStarHandlerFunc) Handle(params PutGistsIDStarParams) middleware.Responder {
	return fn(params)
}

// PutGistsIDStarHandler interface for that can handle valid put gists ID star params
type PutGistsIDStarHandler interface {
	Handle(PutGistsIDStarParams) middleware.Responder
}

// NewPutGistsIDStar creates a new http.Handler for the put gists ID star operation
func NewPutGistsIDStar(ctx *middleware.Context, handler PutGistsIDStarHandler) *PutGistsIDStar {
	return &PutGistsIDStar{Context: ctx, Handler: handler}
}

/*PutGistsIDStar swagger:route PUT /gists/{id}/star putGistsIdStar

Star a gist.

*/
type PutGistsIDStar struct {
	Context *middleware.Context
	Handler PutGistsIDStarHandler
}

func (o *PutGistsIDStar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutGistsIDStarParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
