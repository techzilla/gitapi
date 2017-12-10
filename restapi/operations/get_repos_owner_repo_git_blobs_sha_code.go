// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposOwnerRepoGitBlobsShaCodeHandlerFunc turns a function with the right signature into a get repos owner repo git blobs sha code handler
type GetReposOwnerRepoGitBlobsShaCodeHandlerFunc func(GetReposOwnerRepoGitBlobsShaCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposOwnerRepoGitBlobsShaCodeHandlerFunc) Handle(params GetReposOwnerRepoGitBlobsShaCodeParams) middleware.Responder {
	return fn(params)
}

// GetReposOwnerRepoGitBlobsShaCodeHandler interface for that can handle valid get repos owner repo git blobs sha code params
type GetReposOwnerRepoGitBlobsShaCodeHandler interface {
	Handle(GetReposOwnerRepoGitBlobsShaCodeParams) middleware.Responder
}

// NewGetReposOwnerRepoGitBlobsShaCode creates a new http.Handler for the get repos owner repo git blobs sha code operation
func NewGetReposOwnerRepoGitBlobsShaCode(ctx *middleware.Context, handler GetReposOwnerRepoGitBlobsShaCodeHandler) *GetReposOwnerRepoGitBlobsShaCode {
	return &GetReposOwnerRepoGitBlobsShaCode{Context: ctx, Handler: handler}
}

/*GetReposOwnerRepoGitBlobsShaCode swagger:route GET /repos/{owner}/{repo}/git/blobs/{shaCode} getReposOwnerRepoGitBlobsShaCode

Get a Blob.
Since blobs can be any arbitrary binary data, the input and responses for
the blob API takes an encoding parameter that can be either utf-8 or
base64. If your data cannot be losslessly sent as a UTF-8 string, you can
base64 encode it.


*/
type GetReposOwnerRepoGitBlobsShaCode struct {
	Context *middleware.Context
	Handler GetReposOwnerRepoGitBlobsShaCodeHandler
}

func (o *GetReposOwnerRepoGitBlobsShaCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposOwnerRepoGitBlobsShaCodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}