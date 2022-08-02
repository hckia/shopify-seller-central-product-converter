// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetProductMakeParams creates a new GetProductMakeParams object
//
// There are no default values defined in the spec.
func NewGetProductMakeParams() GetProductMakeParams {

	return GetProductMakeParams{}
}

// GetProductMakeParams contains all the bound params for the get product make operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetProductMake
type GetProductMakeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*name of the make we want on all product variants.
	  Required: true
	  In: path
	*/
	Make string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetProductMakeParams() beforehand.
func (o *GetProductMakeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMake, rhkMake, _ := route.Params.GetOK("make")
	if err := o.bindMake(rMake, rhkMake, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMake binds and validates parameter Make from path.
func (o *GetProductMakeParams) bindMake(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Make = raw

	return nil
}
