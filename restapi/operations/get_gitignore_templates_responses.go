// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetGitignoreTemplatesOKCode is the HTTP code returned for type GetGitignoreTemplatesOK
const GetGitignoreTemplatesOKCode int = 200

/*GetGitignoreTemplatesOK OK

swagger:response getGitignoreTemplatesOK
*/
type GetGitignoreTemplatesOK struct {

	/*
	  In: Body
	*/
	Payload models.Gitignore `json:"body,omitempty"`
}

// NewGetGitignoreTemplatesOK creates GetGitignoreTemplatesOK with default headers values
func NewGetGitignoreTemplatesOK() *GetGitignoreTemplatesOK {
	return &GetGitignoreTemplatesOK{}
}

// WithPayload adds the payload to the get gitignore templates o k response
func (o *GetGitignoreTemplatesOK) WithPayload(payload models.Gitignore) *GetGitignoreTemplatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gitignore templates o k response
func (o *GetGitignoreTemplatesOK) SetPayload(payload models.Gitignore) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGitignoreTemplatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Gitignore, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetGitignoreTemplatesForbiddenCode is the HTTP code returned for type GetGitignoreTemplatesForbidden
const GetGitignoreTemplatesForbiddenCode int = 403

/*GetGitignoreTemplatesForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getGitignoreTemplatesForbidden
*/
type GetGitignoreTemplatesForbidden struct {
}

// NewGetGitignoreTemplatesForbidden creates GetGitignoreTemplatesForbidden with default headers values
func NewGetGitignoreTemplatesForbidden() *GetGitignoreTemplatesForbidden {
	return &GetGitignoreTemplatesForbidden{}
}

// WriteResponse to the client
func (o *GetGitignoreTemplatesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}