// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Magicking/dao_tracabilite/models"
)

// QrcodeOKCode is the HTTP code returned for type QrcodeOK
const QrcodeOKCode int = 200

/*QrcodeOK QRcode identifiant interne

swagger:response qrcodeOK
*/
type QrcodeOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewQrcodeOK creates QrcodeOK with default headers values
func NewQrcodeOK() *QrcodeOK {
	return &QrcodeOK{}
}

// WithPayload adds the payload to the qrcode o k response
func (o *QrcodeOK) WithPayload(payload string) *QrcodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the qrcode o k response
func (o *QrcodeOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QrcodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*QrcodeDefault Unexpected error

swagger:response qrcodeDefault
*/
type QrcodeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewQrcodeDefault creates QrcodeDefault with default headers values
func NewQrcodeDefault(code int) *QrcodeDefault {
	if code <= 0 {
		code = 500
	}

	return &QrcodeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the qrcode default response
func (o *QrcodeDefault) WithStatusCode(code int) *QrcodeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the qrcode default response
func (o *QrcodeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the qrcode default response
func (o *QrcodeDefault) WithPayload(payload *models.Error) *QrcodeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the qrcode default response
func (o *QrcodeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QrcodeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}