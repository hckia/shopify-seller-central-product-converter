// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSwatchMakeHandlerFunc turns a function with the right signature into a get swatch make handler
type GetSwatchMakeHandlerFunc func(GetSwatchMakeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSwatchMakeHandlerFunc) Handle(params GetSwatchMakeParams) middleware.Responder {
	return fn(params)
}

// GetSwatchMakeHandler interface for that can handle valid get swatch make params
type GetSwatchMakeHandler interface {
	Handle(GetSwatchMakeParams) middleware.Responder
}

// NewGetSwatchMake creates a new http.Handler for the get swatch make operation
func NewGetSwatchMake(ctx *middleware.Context, handler GetSwatchMakeHandler) *GetSwatchMake {
	return &GetSwatchMake{Context: ctx, Handler: handler}
}

/* GetSwatchMake swagger:route GET /swatch/{make} getSwatchMake

Return an array of objects for a specific make that will help produce color swatches for Amazon seller central.

*/
type GetSwatchMake struct {
	Context *middleware.Context
	Handler GetSwatchMakeHandler
}

func (o *GetSwatchMake) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSwatchMakeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
