// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUneditableOKCode is the HTTP code returned for type GetUneditableOK
const GetUneditableOKCode int = 200

/*GetUneditableOK columns

swagger:response getUneditableOK
*/
type GetUneditableOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetUneditableOK creates GetUneditableOK with default headers values
func NewGetUneditableOK() *GetUneditableOK {
	return &GetUneditableOK{}
}

// WithPayload adds the payload to the get uneditable o k response
func (o *GetUneditableOK) WithPayload(payload []string) *GetUneditableOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get uneditable o k response
func (o *GetUneditableOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUneditableOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUneditableBadRequestCode is the HTTP code returned for type GetUneditableBadRequest
const GetUneditableBadRequestCode int = 400

/*GetUneditableBadRequest bad input parameter

swagger:response getUneditableBadRequest
*/
type GetUneditableBadRequest struct {
}

// NewGetUneditableBadRequest creates GetUneditableBadRequest with default headers values
func NewGetUneditableBadRequest() *GetUneditableBadRequest {
	return &GetUneditableBadRequest{}
}

// WriteResponse to the client
func (o *GetUneditableBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}