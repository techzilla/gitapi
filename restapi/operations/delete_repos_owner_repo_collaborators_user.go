// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReposOwnerRepoCollaboratorsUserHandlerFunc turns a function with the right signature into a delete repos owner repo collaborators user handler
type DeleteReposOwnerRepoCollaboratorsUserHandlerFunc func(DeleteReposOwnerRepoCollaboratorsUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReposOwnerRepoCollaboratorsUserHandlerFunc) Handle(params DeleteReposOwnerRepoCollaboratorsUserParams) middleware.Responder {
	return fn(params)
}

// DeleteReposOwnerRepoCollaboratorsUserHandler interface for that can handle valid delete repos owner repo collaborators user params
type DeleteReposOwnerRepoCollaboratorsUserHandler interface {
	Handle(DeleteReposOwnerRepoCollaboratorsUserParams) middleware.Responder
}

// NewDeleteReposOwnerRepoCollaboratorsUser creates a new http.Handler for the delete repos owner repo collaborators user operation
func NewDeleteReposOwnerRepoCollaboratorsUser(ctx *middleware.Context, handler DeleteReposOwnerRepoCollaboratorsUserHandler) *DeleteReposOwnerRepoCollaboratorsUser {
	return &DeleteReposOwnerRepoCollaboratorsUser{Context: ctx, Handler: handler}
}

/*DeleteReposOwnerRepoCollaboratorsUser swagger:route DELETE /repos/{owner}/{repo}/collaborators/{user} deleteReposOwnerRepoCollaboratorsUser

Remove collaborator.

*/
type DeleteReposOwnerRepoCollaboratorsUser struct {
	Context *middleware.Context
	Handler DeleteReposOwnerRepoCollaboratorsUserHandler
}

func (o *DeleteReposOwnerRepoCollaboratorsUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReposOwnerRepoCollaboratorsUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
