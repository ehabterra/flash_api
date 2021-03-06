// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ehabterra/flash_api/api/models"
)

// UploadOKCode is the HTTP code returned for type UploadOK
const UploadOKCode int = 200

/*UploadOK Success

swagger:response uploadOK
*/
type UploadOK struct {
}

// NewUploadOK creates UploadOK with default headers values
func NewUploadOK() *UploadOK {

	return &UploadOK{}
}

// WriteResponse to the client
func (o *UploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*UploadDefault generic error response

swagger:response uploadDefault
*/
type UploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadDefault creates UploadDefault with default headers values
func NewUploadDefault(code int) *UploadDefault {
	if code <= 0 {
		code = 500
	}

	return &UploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the upload default response
func (o *UploadDefault) WithStatusCode(code int) *UploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the upload default response
func (o *UploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the upload default response
func (o *UploadDefault) WithPayload(payload *models.Error) *UploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload default response
func (o *UploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
