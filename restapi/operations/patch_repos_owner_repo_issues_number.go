// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchReposOwnerRepoIssuesNumberHandlerFunc turns a function with the right signature into a patch repos owner repo issues number handler
type PatchReposOwnerRepoIssuesNumberHandlerFunc func(PatchReposOwnerRepoIssuesNumberParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchReposOwnerRepoIssuesNumberHandlerFunc) Handle(params PatchReposOwnerRepoIssuesNumberParams) middleware.Responder {
	return fn(params)
}

// PatchReposOwnerRepoIssuesNumberHandler interface for that can handle valid patch repos owner repo issues number params
type PatchReposOwnerRepoIssuesNumberHandler interface {
	Handle(PatchReposOwnerRepoIssuesNumberParams) middleware.Responder
}

// NewPatchReposOwnerRepoIssuesNumber creates a new http.Handler for the patch repos owner repo issues number operation
func NewPatchReposOwnerRepoIssuesNumber(ctx *middleware.Context, handler PatchReposOwnerRepoIssuesNumberHandler) *PatchReposOwnerRepoIssuesNumber {
	return &PatchReposOwnerRepoIssuesNumber{Context: ctx, Handler: handler}
}

/*PatchReposOwnerRepoIssuesNumber swagger:route PATCH /repos/{owner}/{repo}/issues/{number} patchReposOwnerRepoIssuesNumber

Edit an issue.
Issue owners and users with push access can edit an issue.


*/
type PatchReposOwnerRepoIssuesNumber struct {
	Context *middleware.Context
	Handler PatchReposOwnerRepoIssuesNumberHandler
}

func (o *PatchReposOwnerRepoIssuesNumber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchReposOwnerRepoIssuesNumberParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
