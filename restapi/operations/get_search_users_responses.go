// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetSearchUsersOKCode is the HTTP code returned for type GetSearchUsersOK
const GetSearchUsersOKCode int = 200

/*GetSearchUsersOK OK

swagger:response getSearchUsersOK
*/
type GetSearchUsersOK struct {

	/*
	  In: Body
	*/
	Payload *models.SearchUsers `json:"body,omitempty"`
}

// NewGetSearchUsersOK creates GetSearchUsersOK with default headers values
func NewGetSearchUsersOK() *GetSearchUsersOK {
	return &GetSearchUsersOK{}
}

// WithPayload adds the payload to the get search users o k response
func (o *GetSearchUsersOK) WithPayload(payload *models.SearchUsers) *GetSearchUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get search users o k response
func (o *GetSearchUsersOK) SetPayload(payload *models.SearchUsers) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSearchUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSearchUsersForbiddenCode is the HTTP code returned for type GetSearchUsersForbidden
const GetSearchUsersForbiddenCode int = 403

/*GetSearchUsersForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getSearchUsersForbidden
*/
type GetSearchUsersForbidden struct {
}

// NewGetSearchUsersForbidden creates GetSearchUsersForbidden with default headers values
func NewGetSearchUsersForbidden() *GetSearchUsersForbidden {
	return &GetSearchUsersForbidden{}
}

// WriteResponse to the client
func (o *GetSearchUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}