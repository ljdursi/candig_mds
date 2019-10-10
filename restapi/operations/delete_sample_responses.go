// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteSampleOKCode is the HTTP code returned for type DeleteSampleOK
const DeleteSampleOKCode int = 200

/*DeleteSampleOK Deleted sample

swagger:response deleteSampleOK
*/
type DeleteSampleOK struct {
	Payload string `json:"body,omitempty"`
}

// NewDeleteSampleOK creates DeleteSampleOK with default headers values
func NewDeleteSampleOK() *DeleteSampleOK {

	return &DeleteSampleOK{}
}

// WithPayload adds the payload to the delete sample o k response
func (o *DeleteSampleOK) WithPayload(payload string) *DeleteSampleOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *DeleteSampleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
