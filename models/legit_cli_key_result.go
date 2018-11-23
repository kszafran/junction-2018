// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LegitCliKeyResult legit cli key result
// swagger:model LegitCliKeyResult
type LegitCliKeyResult struct {

	// response
	Response []string `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this legit cli key result
func (m *LegitCliKeyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LegitCliKeyResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegitCliKeyResult) UnmarshalBinary(b []byte) error {
	var res LegitCliKeyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
