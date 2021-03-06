// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLegacyUserSearchKeywordHandlerFunc turns a function with the right signature into a get legacy user search keyword handler
type GetLegacyUserSearchKeywordHandlerFunc func(GetLegacyUserSearchKeywordParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLegacyUserSearchKeywordHandlerFunc) Handle(params GetLegacyUserSearchKeywordParams) middleware.Responder {
	return fn(params)
}

// GetLegacyUserSearchKeywordHandler interface for that can handle valid get legacy user search keyword params
type GetLegacyUserSearchKeywordHandler interface {
	Handle(GetLegacyUserSearchKeywordParams) middleware.Responder
}

// NewGetLegacyUserSearchKeyword creates a new http.Handler for the get legacy user search keyword operation
func NewGetLegacyUserSearchKeyword(ctx *middleware.Context, handler GetLegacyUserSearchKeywordHandler) *GetLegacyUserSearchKeyword {
	return &GetLegacyUserSearchKeyword{Context: ctx, Handler: handler}
}

/*GetLegacyUserSearchKeyword swagger:route GET /legacy/user/search/{keyword} getLegacyUserSearchKeyword

Find users by keyword.

*/
type GetLegacyUserSearchKeyword struct {
	Context *middleware.Context
	Handler GetLegacyUserSearchKeywordHandler
}

func (o *GetLegacyUserSearchKeyword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLegacyUserSearchKeywordParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
