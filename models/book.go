package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Book book

swagger:model Book
*/
type Book struct {

	/* Book author.
	 */
	Author string `json:"author,omitempty"`

	/* Unique identifier for the book.
	 */
	ID int64 `json:"id,omitempty"`

	/* IDs for resources linked to this book.
	 */
	Resources []int64 `json:"resources,omitempty"`

	/* Book title.
	 */
	Title string `json:"title,omitempty"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	return nil
}
