// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteGistsIDStarNoContentCode is the HTTP code returned for type DeleteGistsIDStarNoContent
const DeleteGistsIDStarNoContentCode int = 204

/*DeleteGistsIDStarNoContent Item removed.

swagger:response deleteGistsIdStarNoContent
*/
type DeleteGistsIDStarNoContent struct {
}

// NewDeleteGistsIDStarNoContent creates DeleteGistsIDStarNoContent with default headers values
func NewDeleteGistsIDStarNoContent() *DeleteGistsIDStarNoContent {
	return &DeleteGistsIDStarNoContent{}
}

// WriteResponse to the client
func (o *DeleteGistsIDStarNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteGistsIDStarForbiddenCode is the HTTP code returned for type DeleteGistsIDStarForbidden
const DeleteGistsIDStarForbiddenCode int = 403

/*DeleteGistsIDStarForbidden API rate limit exceeded. See http://developer.github.com/v3/#rate-limiting
for details.


swagger:response deleteGistsIdStarForbidden
*/
type DeleteGistsIDStarForbidden struct {
}

// NewDeleteGistsIDStarForbidden creates DeleteGistsIDStarForbidden with default headers values
func NewDeleteGistsIDStarForbidden() *DeleteGistsIDStarForbidden {
	return &DeleteGistsIDStarForbidden{}
}

// WriteResponse to the client
func (o *DeleteGistsIDStarForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
