// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetGistsIDOKCode is the HTTP code returned for type GetGistsIDOK
const GetGistsIDOKCode int = 200

/*GetGistsIDOK OK

swagger:response getGistsIdOK
*/
type GetGistsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Gist `json:"body,omitempty"`
}

// NewGetGistsIDOK creates GetGistsIDOK with default headers values
func NewGetGistsIDOK() *GetGistsIDOK {
	return &GetGistsIDOK{}
}

// WithPayload adds the payload to the get gists Id o k response
func (o *GetGistsIDOK) WithPayload(payload *models.Gist) *GetGistsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gists Id o k response
func (o *GetGistsIDOK) SetPayload(payload *models.Gist) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGistsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGistsIDForbiddenCode is the HTTP code returned for type GetGistsIDForbidden
const GetGistsIDForbiddenCode int = 403

/*GetGistsIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getGistsIdForbidden
*/
type GetGistsIDForbidden struct {
}

// NewGetGistsIDForbidden creates GetGistsIDForbidden with default headers values
func NewGetGistsIDForbidden() *GetGistsIDForbidden {
	return &GetGistsIDForbidden{}
}

// WriteResponse to the client
func (o *GetGistsIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
