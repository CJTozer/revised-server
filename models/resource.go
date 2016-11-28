package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Resource resource

swagger:model Resource
*/
type Resource struct {

	/* The resource's ID
	 */
	ID int64 `json:"id,omitempty"`

	/* The type of resource
	 */
	ResourceType string `json:"resource_type,omitempty"`

	/* The text describing this resource
	 */
	Text string `json:"text,omitempty"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceTypeResourceTypePropEnum []interface{}

// prop value enum
func (m *Resource) validateResourceTypeEnum(path, location string, value string) error {
	if resourceTypeResourceTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Text","Link","Image"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			resourceTypeResourceTypePropEnum = append(resourceTypeResourceTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, resourceTypeResourceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateResourceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceTypeEnum("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}
