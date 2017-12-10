// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserKeysKeyIDHandlerFunc turns a function with the right signature into a delete user keys key ID handler
type DeleteUserKeysKeyIDHandlerFunc func(DeleteUserKeysKeyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserKeysKeyIDHandlerFunc) Handle(params DeleteUserKeysKeyIDParams) middleware.Responder {
	return fn(params)
}

// DeleteUserKeysKeyIDHandler interface for that can handle valid delete user keys key ID params
type DeleteUserKeysKeyIDHandler interface {
	Handle(DeleteUserKeysKeyIDParams) middleware.Responder
}

// NewDeleteUserKeysKeyID creates a new http.Handler for the delete user keys key ID operation
func NewDeleteUserKeysKeyID(ctx *middleware.Context, handler DeleteUserKeysKeyIDHandler) *DeleteUserKeysKeyID {
	return &DeleteUserKeysKeyID{Context: ctx, Handler: handler}
}

/*DeleteUserKeysKeyID swagger:route DELETE /user/keys/{keyId} deleteUserKeysKeyId

Delete a public key. Removes a public key. Requires that you are authenticated via Basic Auth or via OAuth with at least admin:public_key scope.

*/
type DeleteUserKeysKeyID struct {
	Context *middleware.Context
	Handler DeleteUserKeysKeyIDHandler
}

func (o *DeleteUserKeysKeyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserKeysKeyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}