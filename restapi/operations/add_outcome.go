// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddOutcomeHandlerFunc turns a function with the right signature into a add outcome handler
type AddOutcomeHandlerFunc func(AddOutcomeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOutcomeHandlerFunc) Handle(params AddOutcomeParams) middleware.Responder {
	return fn(params)
}

// AddOutcomeHandler interface for that can handle valid add outcome params
type AddOutcomeHandler interface {
	Handle(AddOutcomeParams) middleware.Responder
}

// NewAddOutcome creates a new http.Handler for the add outcome operation
func NewAddOutcome(ctx *middleware.Context, handler AddOutcomeHandler) *AddOutcome {
	return &AddOutcome{Context: ctx, Handler: handler}
}

/*AddOutcome swagger:route POST /outcome addOutcome

adds an outcome

Adds an outcome to the system

*/
type AddOutcome struct {
	Context *middleware.Context
	Handler AddOutcomeHandler
}

func (o *AddOutcome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddOutcomeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
