// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoLabelsOKCode is the HTTP code returned for type GetReposOwnerRepoLabelsOK
const GetReposOwnerRepoLabelsOKCode int = 200

/*GetReposOwnerRepoLabelsOK OK

swagger:response getReposOwnerRepoLabelsOK
*/
type GetReposOwnerRepoLabelsOK struct {

	/*
	  In: Body
	*/
	Payload models.Labels `json:"body,omitempty"`
}

// NewGetReposOwnerRepoLabelsOK creates GetReposOwnerRepoLabelsOK with default headers values
func NewGetReposOwnerRepoLabelsOK() *GetReposOwnerRepoLabelsOK {
	return &GetReposOwnerRepoLabelsOK{}
}

// WithPayload adds the payload to the get repos owner repo labels o k response
func (o *GetReposOwnerRepoLabelsOK) WithPayload(payload models.Labels) *GetReposOwnerRepoLabelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo labels o k response
func (o *GetReposOwnerRepoLabelsOK) SetPayload(payload models.Labels) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoLabelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Labels, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReposOwnerRepoLabelsForbiddenCode is the HTTP code returned for type GetReposOwnerRepoLabelsForbidden
const GetReposOwnerRepoLabelsForbiddenCode int = 403

/*GetReposOwnerRepoLabelsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoLabelsForbidden
*/
type GetReposOwnerRepoLabelsForbidden struct {
}

// NewGetReposOwnerRepoLabelsForbidden creates GetReposOwnerRepoLabelsForbidden with default headers values
func NewGetReposOwnerRepoLabelsForbidden() *GetReposOwnerRepoLabelsForbidden {
	return &GetReposOwnerRepoLabelsForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoLabelsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
