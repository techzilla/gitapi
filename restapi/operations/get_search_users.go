// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSearchUsersHandlerFunc turns a function with the right signature into a get search users handler
type GetSearchUsersHandlerFunc func(GetSearchUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSearchUsersHandlerFunc) Handle(params GetSearchUsersParams) middleware.Responder {
	return fn(params)
}

// GetSearchUsersHandler interface for that can handle valid get search users params
type GetSearchUsersHandler interface {
	Handle(GetSearchUsersParams) middleware.Responder
}

// NewGetSearchUsers creates a new http.Handler for the get search users operation
func NewGetSearchUsers(ctx *middleware.Context, handler GetSearchUsersHandler) *GetSearchUsers {
	return &GetSearchUsers{Context: ctx, Handler: handler}
}

/*GetSearchUsers swagger:route GET /search/users getSearchUsers

Search users.

*/
type GetSearchUsers struct {
	Context *middleware.Context
	Handler GetSearchUsersHandler
}

func (o *GetSearchUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSearchUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
