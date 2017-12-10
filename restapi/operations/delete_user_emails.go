// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserEmailsHandlerFunc turns a function with the right signature into a delete user emails handler
type DeleteUserEmailsHandlerFunc func(DeleteUserEmailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserEmailsHandlerFunc) Handle(params DeleteUserEmailsParams) middleware.Responder {
	return fn(params)
}

// DeleteUserEmailsHandler interface for that can handle valid delete user emails params
type DeleteUserEmailsHandler interface {
	Handle(DeleteUserEmailsParams) middleware.Responder
}

// NewDeleteUserEmails creates a new http.Handler for the delete user emails operation
func NewDeleteUserEmails(ctx *middleware.Context, handler DeleteUserEmailsHandler) *DeleteUserEmails {
	return &DeleteUserEmails{Context: ctx, Handler: handler}
}

/*DeleteUserEmails swagger:route DELETE /user/emails deleteUserEmails

Delete email address(es).
You can include a single email address or an array of addresses.


*/
type DeleteUserEmails struct {
	Context *middleware.Context
	Handler DeleteUserEmailsHandler
}

func (o *DeleteUserEmails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserEmailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
