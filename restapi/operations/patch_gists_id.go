// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchGistsIDHandlerFunc turns a function with the right signature into a patch gists ID handler
type PatchGistsIDHandlerFunc func(PatchGistsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchGistsIDHandlerFunc) Handle(params PatchGistsIDParams) middleware.Responder {
	return fn(params)
}

// PatchGistsIDHandler interface for that can handle valid patch gists ID params
type PatchGistsIDHandler interface {
	Handle(PatchGistsIDParams) middleware.Responder
}

// NewPatchGistsID creates a new http.Handler for the patch gists ID operation
func NewPatchGistsID(ctx *middleware.Context, handler PatchGistsIDHandler) *PatchGistsID {
	return &PatchGistsID{Context: ctx, Handler: handler}
}

/*PatchGistsID swagger:route PATCH /gists/{id} patchGistsId

Edit a gist.

*/
type PatchGistsID struct {
	Context *middleware.Context
	Handler PatchGistsIDHandler
}

func (o *PatchGistsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchGistsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}