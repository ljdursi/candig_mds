// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ljdursi/candig_mds/models"
)

// NewAddBiosampleParams creates a new AddBiosampleParams object
// with the default values initialized.
func NewAddBiosampleParams() AddBiosampleParams {
	var ()
	return AddBiosampleParams{}
}

// AddBiosampleParams contains all the bound params for the add biosample operation
// typically these are obtained from a http.Request
//
// swagger:parameters addBiosample
type AddBiosampleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Biosample
	  In: body
	*/
	Biosample *models.Biosample
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddBiosampleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Biosample
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("biosample", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Biosample = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
