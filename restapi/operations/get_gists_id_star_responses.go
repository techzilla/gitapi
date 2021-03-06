// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetGistsIDStarNoContentCode is the HTTP code returned for type GetGistsIDStarNoContent
const GetGistsIDStarNoContentCode int = 204

/*GetGistsIDStarNoContent Exists.

swagger:response getGistsIdStarNoContent
*/
type GetGistsIDStarNoContent struct {
}

// NewGetGistsIDStarNoContent creates GetGistsIDStarNoContent with default headers values
func NewGetGistsIDStarNoContent() *GetGistsIDStarNoContent {
	return &GetGistsIDStarNoContent{}
}

// WriteResponse to the client
func (o *GetGistsIDStarNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetGistsIDStarForbiddenCode is the HTTP code returned for type GetGistsIDStarForbidden
const GetGistsIDStarForbiddenCode int = 403

/*GetGistsIDStarForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response getGistsIdStarForbidden
*/
type GetGistsIDStarForbidden struct {
}

// NewGetGistsIDStarForbidden creates GetGistsIDStarForbidden with default headers values
func NewGetGistsIDStarForbidden() *GetGistsIDStarForbidden {
	return &GetGistsIDStarForbidden{}
}

// WriteResponse to the client
func (o *GetGistsIDStarForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetGistsIDStarNotFoundCode is the HTTP code returned for type GetGistsIDStarNotFound
const GetGistsIDStarNotFoundCode int = 404

/*GetGistsIDStarNotFound Not exists.

swagger:response getGistsIdStarNotFound
*/
type GetGistsIDStarNotFound struct {
}

// NewGetGistsIDStarNotFound creates GetGistsIDStarNotFound with default headers values
func NewGetGistsIDStarNotFound() *GetGistsIDStarNotFound {
	return &GetGistsIDStarNotFound{}
}

// WriteResponse to the client
func (o *GetGistsIDStarNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
