// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "medicarePartD/gen/models"
)

// GetSingleRecordOKCode is the HTTP code returned for type GetSingleRecordOK
const GetSingleRecordOKCode int = 200

/*GetSingleRecordOK A JSON object containing a record

swagger:response getSingleRecordOK
*/
type GetSingleRecordOK struct {

	/*
	  In: Body
	*/
	Payload *models.Record `json:"body,omitempty"`
}

// NewGetSingleRecordOK creates GetSingleRecordOK with default headers values
func NewGetSingleRecordOK() *GetSingleRecordOK {

	return &GetSingleRecordOK{}
}

// WithPayload adds the payload to the get single record o k response
func (o *GetSingleRecordOK) WithPayload(payload *models.Record) *GetSingleRecordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get single record o k response
func (o *GetSingleRecordOK) SetPayload(payload *models.Record) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSingleRecordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSingleRecordForbiddenCode is the HTTP code returned for type GetSingleRecordForbidden
const GetSingleRecordForbiddenCode int = 403

/*GetSingleRecordForbidden Forbidden Error

swagger:response getSingleRecordForbidden
*/
type GetSingleRecordForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewGetSingleRecordForbidden creates GetSingleRecordForbidden with default headers values
func NewGetSingleRecordForbidden() *GetSingleRecordForbidden {

	return &GetSingleRecordForbidden{}
}

// WithPayload adds the payload to the get single record forbidden response
func (o *GetSingleRecordForbidden) WithPayload(payload *models.ResponseCode) *GetSingleRecordForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get single record forbidden response
func (o *GetSingleRecordForbidden) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSingleRecordForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSingleRecordNotFoundCode is the HTTP code returned for type GetSingleRecordNotFound
const GetSingleRecordNotFoundCode int = 404

/*GetSingleRecordNotFound Not Found

swagger:response getSingleRecordNotFound
*/
type GetSingleRecordNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewGetSingleRecordNotFound creates GetSingleRecordNotFound with default headers values
func NewGetSingleRecordNotFound() *GetSingleRecordNotFound {

	return &GetSingleRecordNotFound{}
}

// WithPayload adds the payload to the get single record not found response
func (o *GetSingleRecordNotFound) WithPayload(payload *models.ResponseCode) *GetSingleRecordNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get single record not found response
func (o *GetSingleRecordNotFound) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSingleRecordNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSingleRecordInternalServerErrorCode is the HTTP code returned for type GetSingleRecordInternalServerError
const GetSingleRecordInternalServerErrorCode int = 500

/*GetSingleRecordInternalServerError Internal Server Error

swagger:response getSingleRecordInternalServerError
*/
type GetSingleRecordInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewGetSingleRecordInternalServerError creates GetSingleRecordInternalServerError with default headers values
func NewGetSingleRecordInternalServerError() *GetSingleRecordInternalServerError {

	return &GetSingleRecordInternalServerError{}
}

// WithPayload adds the payload to the get single record internal server error response
func (o *GetSingleRecordInternalServerError) WithPayload(payload *models.ResponseCode) *GetSingleRecordInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get single record internal server error response
func (o *GetSingleRecordInternalServerError) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSingleRecordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
