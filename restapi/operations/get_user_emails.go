// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserEmailsHandlerFunc turns a function with the right signature into a get user emails handler
type GetUserEmailsHandlerFunc func(GetUserEmailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserEmailsHandlerFunc) Handle(params GetUserEmailsParams) middleware.Responder {
	return fn(params)
}

// GetUserEmailsHandler interface for that can handle valid get user emails params
type GetUserEmailsHandler interface {
	Handle(GetUserEmailsParams) middleware.Responder
}

// NewGetUserEmails creates a new http.Handler for the get user emails operation
func NewGetUserEmails(ctx *middleware.Context, handler GetUserEmailsHandler) *GetUserEmails {
	return &GetUserEmails{Context: ctx, Handler: handler}
}

/*GetUserEmails swagger:route GET /user/emails getUserEmails

List email addresses for a user.
In the final version of the API, this method will return an array of hashes
with extended information for each email address indicating if the address
has been verified and if it's primary email address for GitHub.
Until API v3 is finalized, use the application/vnd.github.v3 media type to
get other response format.


*/
type GetUserEmails struct {
	Context *middleware.Context
	Handler GetUserEmailsHandler
}

func (o *GetUserEmails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserEmailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
