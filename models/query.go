package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Query query
// swagger:model Query
type Query struct {

	// selected condition
	SelectedCondition [][]string `json:"selected_conditions"`

	// selected fields
	SelectedFields []string `json:"selected_fields"`

	// selected tables
	SelectedTables []string `json:"selected_tables"`
}

// Validate validates this query
func (m *Query) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelectedCondition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelectedFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelectedTables(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Query) validateSelectedCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedCondition) { // not required
		return nil
	}

	return nil
}

func (m *Query) validateSelectedFields(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedFields) { // not required
		return nil
	}

	return nil
}

func (m *Query) validateSelectedTables(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedTables) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Query) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Query) UnmarshalBinary(b []byte) error {
	var res Query
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
