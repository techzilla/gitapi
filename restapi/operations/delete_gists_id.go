// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteGistsIDHandlerFunc turns a function with the right signature into a delete gists ID handler
type DeleteGistsIDHandlerFunc func(DeleteGistsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteGistsIDHandlerFunc) Handle(params DeleteGistsIDParams) middleware.Responder {
	return fn(params)
}

// DeleteGistsIDHandler interface for that can handle valid delete gists ID params
type DeleteGistsIDHandler interface {
	Handle(DeleteGistsIDParams) middleware.Responder
}

// NewDeleteGistsID creates a new http.Handler for the delete gists ID operation
func NewDeleteGistsID(ctx *middleware.Context, handler DeleteGistsIDHandler) *DeleteGistsID {
	return &DeleteGistsID{Context: ctx, Handler: handler}
}

/*DeleteGistsID swagger:route DELETE /gists/{id} deleteGistsId

Delete a gist.

*/
type DeleteGistsID struct {
	Context *middleware.Context
	Handler DeleteGistsIDHandler
}

func (o *DeleteGistsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteGistsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
