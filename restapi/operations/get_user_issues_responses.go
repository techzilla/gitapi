// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetUserIssuesOKCode is the HTTP code returned for type GetUserIssuesOK
const GetUserIssuesOKCode int = 200

/*GetUserIssuesOK OK

swagger:response getUserIssuesOK
*/
type GetUserIssuesOK struct {

	/*
	  In: Body
	*/
	Payload models.Issues `json:"body,omitempty"`
}

// NewGetUserIssuesOK creates GetUserIssuesOK with default headers values
func NewGetUserIssuesOK() *GetUserIssuesOK {
	return &GetUserIssuesOK{}
}

// WithPayload adds the payload to the get user issues o k response
func (o *GetUserIssuesOK) WithPayload(payload models.Issues) *GetUserIssuesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user issues o k response
func (o *GetUserIssuesOK) SetPayload(payload models.Issues) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIssuesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Issues, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUserIssuesForbiddenCode is the HTTP code returned for type GetUserIssuesForbidden
const GetUserIssuesForbiddenCode int = 403

/*GetUserIssuesForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUserIssuesForbidden
*/
type GetUserIssuesForbidden struct {
}

// NewGetUserIssuesForbidden creates GetUserIssuesForbidden with default headers values
func NewGetUserIssuesForbidden() *GetUserIssuesForbidden {
	return &GetUserIssuesForbidden{}
}

// WriteResponse to the client
func (o *GetUserIssuesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
