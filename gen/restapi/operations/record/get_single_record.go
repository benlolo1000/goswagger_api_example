// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSingleRecordHandlerFunc turns a function with the right signature into a get single record handler
type GetSingleRecordHandlerFunc func(GetSingleRecordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSingleRecordHandlerFunc) Handle(params GetSingleRecordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSingleRecordHandler interface for that can handle valid get single record params
type GetSingleRecordHandler interface {
	Handle(GetSingleRecordParams, interface{}) middleware.Responder
}

// NewGetSingleRecord creates a new http.Handler for the get single record operation
func NewGetSingleRecord(ctx *middleware.Context, handler GetSingleRecordHandler) *GetSingleRecord {
	return &GetSingleRecord{Context: ctx, Handler: handler}
}

/*GetSingleRecord swagger:route GET /record record getSingleRecord

returns a record specified by ID.

*/
type GetSingleRecord struct {
	Context *middleware.Context
	Handler GetSingleRecordHandler
}

func (o *GetSingleRecord) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSingleRecordParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
