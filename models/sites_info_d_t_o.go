// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SitesInfoDTO sites info d t o
// swagger:model SitesInfoDTO
type SitesInfoDTO struct {

	// site uuids
	SiteUuids []string `json:"siteUuids"`
}

// Validate validates this sites info d t o
func (m *SitesInfoDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SitesInfoDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SitesInfoDTO) UnmarshalBinary(b []byte) error {
	var res SitesInfoDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}