// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "medicarePartD/gen/models"
)

// LoginOKCode is the HTTP code returned for type LoginOK
const LoginOKCode int = 200

/*LoginOK Succesful login

swagger:response loginOK
*/
type LoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.JwtCode `json:"body,omitempty"`
}

// NewLoginOK creates LoginOK with default headers values
func NewLoginOK() *LoginOK {

	return &LoginOK{}
}

// WithPayload adds the payload to the login o k response
func (o *LoginOK) WithPayload(payload *models.JwtCode) *LoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login o k response
func (o *LoginOK) SetPayload(payload *models.JwtCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUnauthorizedCode is the HTTP code returned for type LoginUnauthorized
const LoginUnauthorizedCode int = 401

/*LoginUnauthorized Unauthorized Error

swagger:response loginUnauthorized
*/
type LoginUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewLoginUnauthorized creates LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {

	return &LoginUnauthorized{}
}

// WithPayload adds the payload to the login unauthorized response
func (o *LoginUnauthorized) WithPayload(payload *models.ResponseCode) *LoginUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login unauthorized response
func (o *LoginUnauthorized) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginForbiddenCode is the HTTP code returned for type LoginForbidden
const LoginForbiddenCode int = 403

/*LoginForbidden Forbidden Error

swagger:response loginForbidden
*/
type LoginForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewLoginForbidden creates LoginForbidden with default headers values
func NewLoginForbidden() *LoginForbidden {

	return &LoginForbidden{}
}

// WithPayload adds the payload to the login forbidden response
func (o *LoginForbidden) WithPayload(payload *models.ResponseCode) *LoginForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login forbidden response
func (o *LoginForbidden) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginNotAcceptableCode is the HTTP code returned for type LoginNotAcceptable
const LoginNotAcceptableCode int = 406

/*LoginNotAcceptable Not Acceptable Error

swagger:response loginNotAcceptable
*/
type LoginNotAcceptable struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewLoginNotAcceptable creates LoginNotAcceptable with default headers values
func NewLoginNotAcceptable() *LoginNotAcceptable {

	return &LoginNotAcceptable{}
}

// WithPayload adds the payload to the login not acceptable response
func (o *LoginNotAcceptable) WithPayload(payload *models.ResponseCode) *LoginNotAcceptable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login not acceptable response
func (o *LoginNotAcceptable) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginNotAcceptable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(406)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginInternalServerErrorCode is the HTTP code returned for type LoginInternalServerError
const LoginInternalServerErrorCode int = 500

/*LoginInternalServerError Error when logging in

swagger:response loginInternalServerError
*/
type LoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseCode `json:"body,omitempty"`
}

// NewLoginInternalServerError creates LoginInternalServerError with default headers values
func NewLoginInternalServerError() *LoginInternalServerError {

	return &LoginInternalServerError{}
}

// WithPayload adds the payload to the login internal server error response
func (o *LoginInternalServerError) WithPayload(payload *models.ResponseCode) *LoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login internal server error response
func (o *LoginInternalServerError) SetPayload(payload *models.ResponseCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}