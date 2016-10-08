package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetResourcesIDParams creates a new GetResourcesIDParams object
// with the default values initialized.
func NewGetResourcesIDParams() GetResourcesIDParams {
	var ()
	return GetResourcesIDParams{}
}

// GetResourcesIDParams contains all the bound params for the get resources ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetResourcesID
type GetResourcesIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The ID of the resource to get
	  Required: true
	  Minimum: 1
	  In: path
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetResourcesIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourcesIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetResourcesIDParams) validateID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("id", "path", int64(o.ID), 1, false); err != nil {
		return err
	}

	return nil
}