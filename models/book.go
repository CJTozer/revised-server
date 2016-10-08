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
	BookID float64 `json:"book_id,omitempty"`
	// @@@ Need this to be an int

	/* resources
	 */
	Resources []*Resource `json:"resources,omitempty"`

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

	for i := 0; i < len(m.Resources); i++ {

		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {

			if err := m.Resources[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
