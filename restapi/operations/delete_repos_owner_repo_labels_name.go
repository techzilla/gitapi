// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReposOwnerRepoLabelsNameHandlerFunc turns a function with the right signature into a delete repos owner repo labels name handler
type DeleteReposOwnerRepoLabelsNameHandlerFunc func(DeleteReposOwnerRepoLabelsNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReposOwnerRepoLabelsNameHandlerFunc) Handle(params DeleteReposOwnerRepoLabelsNameParams) middleware.Responder {
	return fn(params)
}

// DeleteReposOwnerRepoLabelsNameHandler interface for that can handle valid delete repos owner repo labels name params
type DeleteReposOwnerRepoLabelsNameHandler interface {
	Handle(DeleteReposOwnerRepoLabelsNameParams) middleware.Responder
}

// NewDeleteReposOwnerRepoLabelsName creates a new http.Handler for the delete repos owner repo labels name operation
func NewDeleteReposOwnerRepoLabelsName(ctx *middleware.Context, handler DeleteReposOwnerRepoLabelsNameHandler) *DeleteReposOwnerRepoLabelsName {
	return &DeleteReposOwnerRepoLabelsName{Context: ctx, Handler: handler}
}

/*DeleteReposOwnerRepoLabelsName swagger:route DELETE /repos/{owner}/{repo}/labels/{name} deleteReposOwnerRepoLabelsName

Delete a label.

*/
type DeleteReposOwnerRepoLabelsName struct {
	Context *middleware.Context
	Handler DeleteReposOwnerRepoLabelsNameHandler
}

func (o *DeleteReposOwnerRepoLabelsName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReposOwnerRepoLabelsNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}