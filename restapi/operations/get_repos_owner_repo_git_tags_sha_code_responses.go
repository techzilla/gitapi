// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoGitTagsShaCodeOKCode is the HTTP code returned for type GetReposOwnerRepoGitTagsShaCodeOK
const GetReposOwnerRepoGitTagsShaCodeOKCode int = 200

/*GetReposOwnerRepoGitTagsShaCodeOK OK

swagger:response getReposOwnerRepoGitTagsShaCodeOK
*/
type GetReposOwnerRepoGitTagsShaCodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Tag `json:"body,omitempty"`
}

// NewGetReposOwnerRepoGitTagsShaCodeOK creates GetReposOwnerRepoGitTagsShaCodeOK with default headers values
func NewGetReposOwnerRepoGitTagsShaCodeOK() *GetReposOwnerRepoGitTagsShaCodeOK {
	return &GetReposOwnerRepoGitTagsShaCodeOK{}
}

// WithPayload adds the payload to the get repos owner repo git tags sha code o k response
func (o *GetReposOwnerRepoGitTagsShaCodeOK) WithPayload(payload *models.Tag) *GetReposOwnerRepoGitTagsShaCodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo git tags sha code o k response
func (o *GetReposOwnerRepoGitTagsShaCodeOK) SetPayload(payload *models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoGitTagsShaCodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoGitTagsShaCodeForbiddenCode is the HTTP code returned for type GetReposOwnerRepoGitTagsShaCodeForbidden
const GetReposOwnerRepoGitTagsShaCodeForbiddenCode int = 403

/*GetReposOwnerRepoGitTagsShaCodeForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoGitTagsShaCodeForbidden
*/
type GetReposOwnerRepoGitTagsShaCodeForbidden struct {
}

// NewGetReposOwnerRepoGitTagsShaCodeForbidden creates GetReposOwnerRepoGitTagsShaCodeForbidden with default headers values
func NewGetReposOwnerRepoGitTagsShaCodeForbidden() *GetReposOwnerRepoGitTagsShaCodeForbidden {
	return &GetReposOwnerRepoGitTagsShaCodeForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoGitTagsShaCodeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}