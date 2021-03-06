// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProjectDTO project d t o
// swagger:model ProjectDTO
type ProjectDTO struct {

	// create time
	CreateTime int64 `json:"createTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last update time
	LastUpdateTime int64 `json:"lastUpdateTime,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// templates
	Templates interface{} `json:"templates,omitempty"`
}

// Validate validates this project d t o
func (m *ProjectDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectDTO) UnmarshalBinary(b []byte) error {
	var res ProjectDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
