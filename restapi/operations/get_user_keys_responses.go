// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetUserKeysOKCode is the HTTP code returned for type GetUserKeysOK
const GetUserKeysOKCode int = 200

/*GetUserKeysOK OK

swagger:response getUserKeysOK
*/
type GetUserKeysOK struct {

	/*
	  In: Body
	*/
	Payload models.Gitignore `json:"body,omitempty"`
}

// NewGetUserKeysOK creates GetUserKeysOK with default headers values
func NewGetUserKeysOK() *GetUserKeysOK {
	return &GetUserKeysOK{}
}

// WithPayload adds the payload to the get user keys o k response
func (o *GetUserKeysOK) WithPayload(payload models.Gitignore) *GetUserKeysOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user keys o k response
func (o *GetUserKeysOK) SetPayload(payload models.Gitignore) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserKeysOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Gitignore, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUserKeysForbiddenCode is the HTTP code returned for type GetUserKeysForbidden
const GetUserKeysForbiddenCode int = 403

/*GetUserKeysForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getUserKeysForbidden
*/
type GetUserKeysForbidden struct {
}

// NewGetUserKeysForbidden creates GetUserKeysForbidden with default headers values
func NewGetUserKeysForbidden() *GetUserKeysForbidden {
	return &GetUserKeysForbidden{}
}

// WriteResponse to the client
func (o *GetUserKeysForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}