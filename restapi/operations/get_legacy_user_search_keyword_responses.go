// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetLegacyUserSearchKeywordOKCode is the HTTP code returned for type GetLegacyUserSearchKeywordOK
const GetLegacyUserSearchKeywordOKCode int = 200

/*GetLegacyUserSearchKeywordOK OK

swagger:response getLegacyUserSearchKeywordOK
*/
type GetLegacyUserSearchKeywordOK struct {

	/*
	  In: Body
	*/
	Payload *models.SearchUsersByKeyword `json:"body,omitempty"`
}

// NewGetLegacyUserSearchKeywordOK creates GetLegacyUserSearchKeywordOK with default headers values
func NewGetLegacyUserSearchKeywordOK() *GetLegacyUserSearchKeywordOK {
	return &GetLegacyUserSearchKeywordOK{}
}

// WithPayload adds the payload to the get legacy user search keyword o k response
func (o *GetLegacyUserSearchKeywordOK) WithPayload(payload *models.SearchUsersByKeyword) *GetLegacyUserSearchKeywordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get legacy user search keyword o k response
func (o *GetLegacyUserSearchKeywordOK) SetPayload(payload *models.SearchUsersByKeyword) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLegacyUserSearchKeywordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLegacyUserSearchKeywordForbiddenCode is the HTTP code returned for type GetLegacyUserSearchKeywordForbidden
const GetLegacyUserSearchKeywordForbiddenCode int = 403

/*GetLegacyUserSearchKeywordForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getLegacyUserSearchKeywordForbidden
*/
type GetLegacyUserSearchKeywordForbidden struct {
}

// NewGetLegacyUserSearchKeywordForbidden creates GetLegacyUserSearchKeywordForbidden with default headers values
func NewGetLegacyUserSearchKeywordForbidden() *GetLegacyUserSearchKeywordForbidden {
	return &GetLegacyUserSearchKeywordForbidden{}
}

// WriteResponse to the client
func (o *GetLegacyUserSearchKeywordForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
