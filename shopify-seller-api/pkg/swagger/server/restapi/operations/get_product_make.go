// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProductMakeHandlerFunc turns a function with the right signature into a get product make handler
type GetProductMakeHandlerFunc func(GetProductMakeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProductMakeHandlerFunc) Handle(params GetProductMakeParams) middleware.Responder {
	return fn(params)
}

// GetProductMakeHandler interface for that can handle valid get product make params
type GetProductMakeHandler interface {
	Handle(GetProductMakeParams) middleware.Responder
}

// NewGetProductMake creates a new http.Handler for the get product make operation
func NewGetProductMake(ctx *middleware.Context, handler GetProductMakeHandler) *GetProductMake {
	return &GetProductMake{Context: ctx, Handler: handler}
}

/* GetProductMake swagger:route GET /product/{make} getProductMake

Return json file for a specific make that will help produces products for seller central.

*/
type GetProductMake struct {
	Context *middleware.Context
	Handler GetProductMakeHandler
}

func (o *GetProductMake) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProductMakeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
