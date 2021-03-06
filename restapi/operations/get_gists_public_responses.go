// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetGistsPublicOKCode is the HTTP code returned for type GetGistsPublicOK
const GetGistsPublicOKCode int = 200

/*GetGistsPublicOK OK

swagger:response getGistsPublicOK
*/
type GetGistsPublicOK struct {

	/*
	  In: Body
	*/
	Payload models.Gists `json:"body,omitempty"`
}

// NewGetGistsPublicOK creates GetGistsPublicOK with default headers values
func NewGetGistsPublicOK() *GetGistsPublicOK {
	return &GetGistsPublicOK{}
}

// WithPayload adds the payload to the get gists public o k response
func (o *GetGistsPublicOK) WithPayload(payload models.Gists) *GetGistsPublicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gists public o k response
func (o *GetGistsPublicOK) SetPayload(payload models.Gists) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGistsPublicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Gists, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetGistsPublicForbiddenCode is the HTTP code returned for type GetGistsPublicForbidden
const GetGistsPublicForbiddenCode int = 403

/*GetGistsPublicForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getGistsPublicForbidden
*/
type GetGistsPublicForbidden struct {
}

// NewGetGistsPublicForbidden creates GetGistsPublicForbidden with default headers values
func NewGetGistsPublicForbidden() *GetGistsPublicForbidden {
	return &GetGistsPublicForbidden{}
}

// WriteResponse to the client
func (o *GetGistsPublicForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
