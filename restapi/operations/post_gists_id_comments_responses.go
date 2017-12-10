// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// PostGistsIDCommentsCreatedCode is the HTTP code returned for type PostGistsIDCommentsCreated
const PostGistsIDCommentsCreatedCode int = 201

/*PostGistsIDCommentsCreated Created

swagger:response postGistsIdCommentsCreated
*/
type PostGistsIDCommentsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Comment `json:"body,omitempty"`
}

// NewPostGistsIDCommentsCreated creates PostGistsIDCommentsCreated with default headers values
func NewPostGistsIDCommentsCreated() *PostGistsIDCommentsCreated {
	return &PostGistsIDCommentsCreated{}
}

// WithPayload adds the payload to the post gists Id comments created response
func (o *PostGistsIDCommentsCreated) WithPayload(payload *models.Comment) *PostGistsIDCommentsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post gists Id comments created response
func (o *PostGistsIDCommentsCreated) SetPayload(payload *models.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGistsIDCommentsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostGistsIDCommentsForbiddenCode is the HTTP code returned for type PostGistsIDCommentsForbidden
const PostGistsIDCommentsForbiddenCode int = 403

/*PostGistsIDCommentsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response postGistsIdCommentsForbidden
*/
type PostGistsIDCommentsForbidden struct {
}

// NewPostGistsIDCommentsForbidden creates PostGistsIDCommentsForbidden with default headers values
func NewPostGistsIDCommentsForbidden() *PostGistsIDCommentsForbidden {
	return &PostGistsIDCommentsForbidden{}
}

// WriteResponse to the client
func (o *PostGistsIDCommentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}