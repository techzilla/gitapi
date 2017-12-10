// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetMetaOKCode is the HTTP code returned for type GetMetaOK
const GetMetaOKCode int = 200

/*GetMetaOK OK

swagger:response getMetaOK
*/
type GetMetaOK struct {

	/*
	  In: Body
	*/
	Payload *models.Meta `json:"body,omitempty"`
}

// NewGetMetaOK creates GetMetaOK with default headers values
func NewGetMetaOK() *GetMetaOK {
	return &GetMetaOK{}
}

// WithPayload adds the payload to the get meta o k response
func (o *GetMetaOK) WithPayload(payload *models.Meta) *GetMetaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get meta o k response
func (o *GetMetaOK) SetPayload(payload *models.Meta) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMetaForbiddenCode is the HTTP code returned for type GetMetaForbidden
const GetMetaForbiddenCode int = 403

/*GetMetaForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getMetaForbidden
*/
type GetMetaForbidden struct {
}

// NewGetMetaForbidden creates GetMetaForbidden with default headers values
func NewGetMetaForbidden() *GetMetaForbidden {
	return &GetMetaForbidden{}
}

// WriteResponse to the client
func (o *GetMetaForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
