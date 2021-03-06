// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostUserKeysCreatedCode is the HTTP code returned for type PostUserKeysCreated
const PostUserKeysCreatedCode int = 201

/*PostUserKeysCreated Created

swagger:response postUserKeysCreated
*/
type PostUserKeysCreated struct {

	/*
	  In: Body
	*/
	Payload *models.UserKeysKeyID `json:"body,omitempty"`
}

// NewPostUserKeysCreated creates PostUserKeysCreated with default headers values
func NewPostUserKeysCreated() *PostUserKeysCreated {
	return &PostUserKeysCreated{}
}

// WithPayload adds the payload to the post user keys created response
func (o *PostUserKeysCreated) WithPayload(payload *models.UserKeysKeyID) *PostUserKeysCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user keys created response
func (o *PostUserKeysCreated) SetPayload(payload *models.UserKeysKeyID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserKeysCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserKeysForbiddenCode is the HTTP code returned for type PostUserKeysForbidden
const PostUserKeysForbiddenCode int = 403

/*PostUserKeysForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postUserKeysForbidden
*/
type PostUserKeysForbidden struct {
}

// NewPostUserKeysForbidden creates PostUserKeysForbidden with default headers values
func NewPostUserKeysForbidden() *PostUserKeysForbidden {
	return &PostUserKeysForbidden{}
}

// WriteResponse to the client
func (o *PostUserKeysForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
