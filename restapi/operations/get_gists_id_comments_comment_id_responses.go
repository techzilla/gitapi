// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetGistsIDCommentsCommentIDOKCode is the HTTP code returned for type GetGistsIDCommentsCommentIDOK
const GetGistsIDCommentsCommentIDOKCode int = 200

/*GetGistsIDCommentsCommentIDOK OK

swagger:response getGistsIdCommentsCommentIdOK
*/
type GetGistsIDCommentsCommentIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Comment `json:"body,omitempty"`
}

// NewGetGistsIDCommentsCommentIDOK creates GetGistsIDCommentsCommentIDOK with default headers values
func NewGetGistsIDCommentsCommentIDOK() *GetGistsIDCommentsCommentIDOK {
	return &GetGistsIDCommentsCommentIDOK{}
}

// WithPayload adds the payload to the get gists Id comments comment Id o k response
func (o *GetGistsIDCommentsCommentIDOK) WithPayload(payload *models.Comment) *GetGistsIDCommentsCommentIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gists Id comments comment Id o k response
func (o *GetGistsIDCommentsCommentIDOK) SetPayload(payload *models.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGistsIDCommentsCommentIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGistsIDCommentsCommentIDForbiddenCode is the HTTP code returned for type GetGistsIDCommentsCommentIDForbidden
const GetGistsIDCommentsCommentIDForbiddenCode int = 403

/*GetGistsIDCommentsCommentIDForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getGistsIdCommentsCommentIdForbidden
*/
type GetGistsIDCommentsCommentIDForbidden struct {
}

// NewGetGistsIDCommentsCommentIDForbidden creates GetGistsIDCommentsCommentIDForbidden with default headers values
func NewGetGistsIDCommentsCommentIDForbidden() *GetGistsIDCommentsCommentIDForbidden {
	return &GetGistsIDCommentsCommentIDForbidden{}
}

// WriteResponse to the client
func (o *GetGistsIDCommentsCommentIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
