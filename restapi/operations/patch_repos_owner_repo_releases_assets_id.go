// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchReposOwnerRepoReleasesAssetsIDHandlerFunc turns a function with the right signature into a patch repos owner repo releases assets ID handler
type PatchReposOwnerRepoReleasesAssetsIDHandlerFunc func(PatchReposOwnerRepoReleasesAssetsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchReposOwnerRepoReleasesAssetsIDHandlerFunc) Handle(params PatchReposOwnerRepoReleasesAssetsIDParams) middleware.Responder {
	return fn(params)
}

// PatchReposOwnerRepoReleasesAssetsIDHandler interface for that can handle valid patch repos owner repo releases assets ID params
type PatchReposOwnerRepoReleasesAssetsIDHandler interface {
	Handle(PatchReposOwnerRepoReleasesAssetsIDParams) middleware.Responder
}

// NewPatchReposOwnerRepoReleasesAssetsID creates a new http.Handler for the patch repos owner repo releases assets ID operation
func NewPatchReposOwnerRepoReleasesAssetsID(ctx *middleware.Context, handler PatchReposOwnerRepoReleasesAssetsIDHandler) *PatchReposOwnerRepoReleasesAssetsID {
	return &PatchReposOwnerRepoReleasesAssetsID{Context: ctx, Handler: handler}
}

/*PatchReposOwnerRepoReleasesAssetsID swagger:route PATCH /repos/{owner}/{repo}/releases/assets/{id} patchReposOwnerRepoReleasesAssetsId

Edit a release asset
Users with push access to the repository can edit a release asset.


*/
type PatchReposOwnerRepoReleasesAssetsID struct {
	Context *middleware.Context
	Handler PatchReposOwnerRepoReleasesAssetsIDHandler
}

func (o *PatchReposOwnerRepoReleasesAssetsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchReposOwnerRepoReleasesAssetsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
