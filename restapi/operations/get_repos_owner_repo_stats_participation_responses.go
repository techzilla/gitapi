// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/techzilla/gitapi/models"
)

// GetReposOwnerRepoStatsParticipationOKCode is the HTTP code returned for type GetReposOwnerRepoStatsParticipationOK
const GetReposOwnerRepoStatsParticipationOKCode int = 200

/*GetReposOwnerRepoStatsParticipationOK OK

swagger:response getReposOwnerRepoStatsParticipationOK
*/
type GetReposOwnerRepoStatsParticipationOK struct {

	/*
	  In: Body
	*/
	Payload *models.ParticipationStats `json:"body,omitempty"`
}

// NewGetReposOwnerRepoStatsParticipationOK creates GetReposOwnerRepoStatsParticipationOK with default headers values
func NewGetReposOwnerRepoStatsParticipationOK() *GetReposOwnerRepoStatsParticipationOK {
	return &GetReposOwnerRepoStatsParticipationOK{}
}

// WithPayload adds the payload to the get repos owner repo stats participation o k response
func (o *GetReposOwnerRepoStatsParticipationOK) WithPayload(payload *models.ParticipationStats) *GetReposOwnerRepoStatsParticipationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos owner repo stats participation o k response
func (o *GetReposOwnerRepoStatsParticipationOK) SetPayload(payload *models.ParticipationStats) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOwnerRepoStatsParticipationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposOwnerRepoStatsParticipationForbiddenCode is the HTTP code returned for type GetReposOwnerRepoStatsParticipationForbidden
const GetReposOwnerRepoStatsParticipationForbiddenCode int = 403

/*GetReposOwnerRepoStatsParticipationForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getReposOwnerRepoStatsParticipationForbidden
*/
type GetReposOwnerRepoStatsParticipationForbidden struct {
}

// NewGetReposOwnerRepoStatsParticipationForbidden creates GetReposOwnerRepoStatsParticipationForbidden with default headers values
func NewGetReposOwnerRepoStatsParticipationForbidden() *GetReposOwnerRepoStatsParticipationForbidden {
	return &GetReposOwnerRepoStatsParticipationForbidden{}
}

// WriteResponse to the client
func (o *GetReposOwnerRepoStatsParticipationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}