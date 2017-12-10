// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoMilestonesNumberLabelsOKCode is the HTTP code returned for type GetReposOwnerRepoMilestonesNumberLabelsOK
const GetReposOwnerRepoMilestonesNumberLabelsOKCode int = 200

/*GetReposOwnerRepoMilestonesNumberLabelsOK OK

swagger:response getReposOwnerRepoMilestonesNumberLabelsOK
*/
type GetReposOwnerRepoMilestonesNumberLabelsOK struct {

	/*
	  In: Body
	*/
	Payload models.Labels `json:"body,omitempty"`
}

// NewGetReposOwnerRepoMilestonesNumberLabelsOK creates GetReposOwnerRepoMilestonesNumberLabelsOK with default headers values
func NewGetReposOwnerRepoMilestonesNumberLabelsOK() *GetReposOwnerRepoMilestonesNumberLabelsOK {
	return &GetReposOwnerRepoMilestonesNumberLabelsOK{}
}

// WithPayload adds the payload to the get repos owner repo milestones number labels o k response
func (o *GetReposOwnerRepoMilestonesNumberLabelsOK) WithPayload(payload models.Labels) *GetReposOwnerRepoMilestonesNumberLabelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo milestones number labels o k response
func (o *GetReposOwnerRepoMilestonesNumberLabelsOK) SetPayload(payload models.Labels) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoMilestonesNumberLabelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Labels, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReposOwnerRepoMilestonesNumberLabelsForbiddenCode is the HTTP code returned for type GetReposOwnerRepoMilestonesNumberLabelsForbidden
const GetReposOwnerRepoMilestonesNumberLabelsForbiddenCode int = 403

/*GetReposOwnerRepoMilestonesNumberLabelsForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoMilestonesNumberLabelsForbidden
*/
type GetReposOwnerRepoMilestonesNumberLabelsForbidden struct {
}

// NewGetReposOwnerRepoMilestonesNumberLabelsForbidden creates GetReposOwnerRepoMilestonesNumberLabelsForbidden with default headers values
func NewGetReposOwnerRepoMilestonesNumberLabelsForbidden() *GetReposOwnerRepoMilestonesNumberLabelsForbidden {
	return &GetReposOwnerRepoMilestonesNumberLabelsForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoMilestonesNumberLabelsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
