// Code generated by go-swagger; DO NOT EDIT.

package rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetRatesParams creates a new GetRatesParams object
//
// There are no default values defined in the spec.
func NewGetRatesParams() GetRatesParams {

	return GetRatesParams{}
}

// GetRatesParams contains all the bound params for the get rates operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetRates
type GetRatesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Max Length: 5
	  Min Length: 2
	  In: path
	*/
	Base string
	/*
	  Required: true
	  Max Length: 5
	  Min Length: 2
	  In: path
	*/
	Target string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRatesParams() beforehand.
func (o *GetRatesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBase, rhkBase, _ := route.Params.GetOK("base")
	if err := o.bindBase(rBase, rhkBase, route.Formats); err != nil {
		res = append(res, err)
	}

	rTarget, rhkTarget, _ := route.Params.GetOK("target")
	if err := o.bindTarget(rTarget, rhkTarget, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBase binds and validates parameter Base from path.
func (o *GetRatesParams) bindBase(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Base = raw

	if err := o.validateBase(formats); err != nil {
		return err
	}

	return nil
}

// validateBase carries on validations for parameter Base
func (o *GetRatesParams) validateBase(formats strfmt.Registry) error {

	if err := validate.MinLength("base", "path", o.Base, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("base", "path", o.Base, 5); err != nil {
		return err
	}

	return nil
}

// bindTarget binds and validates parameter Target from path.
func (o *GetRatesParams) bindTarget(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Target = raw

	if err := o.validateTarget(formats); err != nil {
		return err
	}

	return nil
}

// validateTarget carries on validations for parameter Target
func (o *GetRatesParams) validateTarget(formats strfmt.Registry) error {

	if err := validate.MinLength("target", "path", o.Target, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("target", "path", o.Target, 5); err != nil {
		return err
	}

	return nil
}
