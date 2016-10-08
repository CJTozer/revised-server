package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Resource resource

swagger:model Resource
*/
type Resource struct {

	/* The resource's ID
	 */
	ResourceID float64 `json:"resource_id,omitempty"`
	// @@@ Need this to be an int

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
