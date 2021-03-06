// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetLegacyUserEmailEmailOKCode is the HTTP code returned for type GetLegacyUserEmailEmailOK
const GetLegacyUserEmailEmailOKCode int = 200

/*GetLegacyUserEmailEmailOK OK

swagger:response getLegacyUserEmailEmailOK
*/
type GetLegacyUserEmailEmailOK struct {

	/*
	  In: Body
	*/
	Payload *models.SearchUserByEmail `json:"body,omitempty"`
}

// NewGetLegacyUserEmailEmailOK creates GetLegacyUserEmailEmailOK with default headers values
func NewGetLegacyUserEmailEmailOK() *GetLegacyUserEmailEmailOK {
	return &GetLegacyUserEmailEmailOK{}
}

// WithPayload adds the payload to the get legacy user email email o k response
func (o *GetLegacyUserEmailEmailOK) WithPayload(payload *models.SearchUserByEmail) *GetLegacyUserEmailEmailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get legacy user email email o k response
func (o *GetLegacyUserEmailEmailOK) SetPayload(payload *models.SearchUserByEmail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLegacyUserEmailEmailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLegacyUserEmailEmailForbiddenCode is the HTTP code returned for type GetLegacyUserEmailEmailForbidden
const GetLegacyUserEmailEmailForbiddenCode int = 403

/*GetLegacyUserEmailEmailForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getLegacyUserEmailEmailForbidden
*/
type GetLegacyUserEmailEmailForbidden struct {
}

// NewGetLegacyUserEmailEmailForbidden creates GetLegacyUserEmailEmailForbidden with default headers values
func NewGetLegacyUserEmailEmailForbidden() *GetLegacyUserEmailEmailForbidden {
	return &GetLegacyUserEmailEmailForbidden{}
}

// WriteResponse to the client
func (o *GetLegacyUserEmailEmailForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
