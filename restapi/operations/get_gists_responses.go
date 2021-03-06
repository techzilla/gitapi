// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetGistsOKCode is the HTTP code returned for type GetGistsOK
const GetGistsOKCode int = 200

/*GetGistsOK OK

swagger:response getGistsOK
*/
type GetGistsOK struct {

	/*
	  In: Body
	*/
	Payload models.Gists `json:"body,omitempty"`
}

// NewGetGistsOK creates GetGistsOK with default headers values
func NewGetGistsOK() *GetGistsOK {
	return &GetGistsOK{}
}

// WithPayload adds the payload to the get gists o k response
func (o *GetGistsOK) WithPayload(payload models.Gists) *GetGistsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gists o k response
func (o *GetGistsOK) SetPayload(payload models.Gists) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGistsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Gists, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetGistsForbiddenCode is the HTTP code returned for type GetGistsForbidden
const GetGistsForbiddenCode int = 403

/*GetGistsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getGistsForbidden
*/
type GetGistsForbidden struct {
}

// NewGetGistsForbidden creates GetGistsForbidden with default headers values
func NewGetGistsForbidden() *GetGistsForbidden {
	return &GetGistsForbidden{}
}

// WriteResponse to the client
func (o *GetGistsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
