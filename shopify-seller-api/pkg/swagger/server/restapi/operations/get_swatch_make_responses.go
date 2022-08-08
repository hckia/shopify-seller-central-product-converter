// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

// GetSwatchMakeOKCode is the HTTP code returned for type GetSwatchMakeOK
const GetSwatchMakeOKCode int = 200

/*GetSwatchMakeOK A JSON object containing the data necessary to make all color swatches.

swagger:response getSwatchMakeOK
*/
type GetSwatchMakeOK struct {

	/*
	  In: Body
	*/
	Payload models.MakeSwatches `json:"body,omitempty"`
}

// NewGetSwatchMakeOK creates GetSwatchMakeOK with default headers values
func NewGetSwatchMakeOK() *GetSwatchMakeOK {

	return &GetSwatchMakeOK{}
}

// WithPayload adds the payload to the get swatch make o k response
func (o *GetSwatchMakeOK) WithPayload(payload models.MakeSwatches) *GetSwatchMakeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get swatch make o k response
func (o *GetSwatchMakeOK) SetPayload(payload models.MakeSwatches) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSwatchMakeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.MakeSwatches{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSwatchMakeBadRequestCode is the HTTP code returned for type GetSwatchMakeBadRequest
const GetSwatchMakeBadRequestCode int = 400

/*GetSwatchMakeBadRequest Invalid "make" provided or not available.

swagger:response getSwatchMakeBadRequest
*/
type GetSwatchMakeBadRequest struct {
}

// NewGetSwatchMakeBadRequest creates GetSwatchMakeBadRequest with default headers values
func NewGetSwatchMakeBadRequest() *GetSwatchMakeBadRequest {

	return &GetSwatchMakeBadRequest{}
}

// WriteResponse to the client
func (o *GetSwatchMakeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}