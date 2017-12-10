// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetUserReposOKCode is the HTTP code returned for type GetUserReposOK
const GetUserReposOKCode int = 200

/*GetUserReposOK OK

swagger:response getUserReposOK
*/
type GetUserReposOK struct {

	/*
	  In: Body
	*/
	Payload models.Repos `json:"body,omitempty"`
}

// NewGetUserReposOK creates GetUserReposOK with default headers values
func NewGetUserReposOK() *GetUserReposOK {
	return &GetUserReposOK{}
}

// WithPayload adds the payload to the get user repos o k response
func (o *GetUserReposOK) WithPayload(payload models.Repos) *GetUserReposOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user repos o k response
func (o *GetUserReposOK) SetPayload(payload models.Repos) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserReposOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Repos, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUserReposForbiddenCode is the HTTP code returned for type GetUserReposForbidden
const GetUserReposForbiddenCode int = 403

/*GetUserReposForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUserReposForbidden
*/
type GetUserReposForbidden struct {
}

// NewGetUserReposForbidden creates GetUserReposForbidden with default headers values
func NewGetUserReposForbidden() *GetUserReposForbidden {
	return &GetUserReposForbidden{}
}

// WriteResponse to the client
func (o *GetUserReposForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}