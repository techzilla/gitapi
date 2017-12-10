// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutReposOwnerRepoPullsNumberMergeHandlerFunc turns a function with the right signature into a put repos owner repo pulls number merge handler
type PutReposOwnerRepoPullsNumberMergeHandlerFunc func(PutReposOwnerRepoPullsNumberMergeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutReposOwnerRepoPullsNumberMergeHandlerFunc) Handle(params PutReposOwnerRepoPullsNumberMergeParams) middleware.Responder {
	return fn(params)
}

// PutReposOwnerRepoPullsNumberMergeHandler interface for that can handle valid put repos owner repo pulls number merge params
type PutReposOwnerRepoPullsNumberMergeHandler interface {
	Handle(PutReposOwnerRepoPullsNumberMergeParams) middleware.Responder
}

// NewPutReposOwnerRepoPullsNumberMerge creates a new http.Handler for the put repos owner repo pulls number merge operation
func NewPutReposOwnerRepoPullsNumberMerge(ctx *middleware.Context, handler PutReposOwnerRepoPullsNumberMergeHandler) *PutReposOwnerRepoPullsNumberMerge {
	return &PutReposOwnerRepoPullsNumberMerge{Context: ctx, Handler: handler}
}

/*PutReposOwnerRepoPullsNumberMerge swagger:route PUT /repos/{owner}/{repo}/pulls/{number}/merge putReposOwnerRepoPullsNumberMerge

Merge a pull request (Merge Button's)

*/
type PutReposOwnerRepoPullsNumberMerge struct {
	Context *middleware.Context
	Handler PutReposOwnerRepoPullsNumberMergeHandler
}

func (o *PutReposOwnerRepoPullsNumberMerge) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutReposOwnerRepoPullsNumberMergeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}