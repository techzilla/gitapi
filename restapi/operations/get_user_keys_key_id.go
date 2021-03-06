// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserKeysKeyIDHandlerFunc turns a function with the right signature into a get user keys key ID handler
type GetUserKeysKeyIDHandlerFunc func(GetUserKeysKeyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserKeysKeyIDHandlerFunc) Handle(params GetUserKeysKeyIDParams) middleware.Responder {
	return fn(params)
}

// GetUserKeysKeyIDHandler interface for that can handle valid get user keys key ID params
type GetUserKeysKeyIDHandler interface {
	Handle(GetUserKeysKeyIDParams) middleware.Responder
}

// NewGetUserKeysKeyID creates a new http.Handler for the get user keys key ID operation
func NewGetUserKeysKeyID(ctx *middleware.Context, handler GetUserKeysKeyIDHandler) *GetUserKeysKeyID {
	return &GetUserKeysKeyID{Context: ctx, Handler: handler}
}

/*GetUserKeysKeyID swagger:route GET /user/keys/{keyId} getUserKeysKeyId

Get a single public key.

*/
type GetUserKeysKeyID struct {
	Context *middleware.Context
	Handler GetUserKeysKeyIDHandler
}

func (o *GetUserKeysKeyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserKeysKeyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
