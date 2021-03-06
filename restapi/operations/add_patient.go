// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddPatientHandlerFunc turns a function with the right signature into a add patient handler
type AddPatientHandlerFunc func(AddPatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPatientHandlerFunc) Handle(params AddPatientParams) middleware.Responder {
	return fn(params)
}

// AddPatientHandler interface for that can handle valid add patient params
type AddPatientHandler interface {
	Handle(AddPatientParams) middleware.Responder
}

// NewAddPatient creates a new http.Handler for the add patient operation
func NewAddPatient(ctx *middleware.Context, handler AddPatientHandler) *AddPatient {
	return &AddPatient{Context: ctx, Handler: handler}
}

/*AddPatient swagger:route POST /patient addPatient

adds a patient item

Adds a patient to the system

*/
type AddPatient struct {
	Context *middleware.Context
	Handler AddPatientHandler
}

func (o *AddPatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddPatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
