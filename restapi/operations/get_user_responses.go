// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetUserOKCode is the HTTP code returned for type GetUserOK
const GetUserOKCode int = 200

/*GetUserOK OK

swagger:response getUserOK
*/
type GetUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserOK creates GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

// WithPayload adds the payload to the get user o k response
func (o *GetUserOK) WithPayload(payload *models.User) *GetUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user o k response
func (o *GetUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserForbiddenCode is the HTTP code returned for type GetUserForbidden
const GetUserForbiddenCode int = 403

/*GetUserForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUserForbidden
*/
type GetUserForbidden struct {
}

// NewGetUserForbidden creates GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

// WriteResponse to the client
func (o *GetUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
