// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateSiteRequest create site request
// swagger:model CreateSiteRequest
type CreateSiteRequest struct {

	// site
	Site *CreateSiteRequestSite `json:"site,omitempty"`

	// type
	// Enum: [area building floor]
	Type string `json:"type,omitempty"`
}

// Validate validates this create site request
func (m *CreateSiteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSiteRequest) validateSite(formats strfmt.Registry) error {

	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

var createSiteRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["area","building","floor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createSiteRequestTypeTypePropEnum = append(createSiteRequestTypeTypePropEnum, v)
	}
}

const (

	// CreateSiteRequestTypeArea captures enum value "area"
	CreateSiteRequestTypeArea string = "area"

	// CreateSiteRequestTypeBuilding captures enum value "building"
	CreateSiteRequestTypeBuilding string = "building"

	// CreateSiteRequestTypeFloor captures enum value "floor"
	CreateSiteRequestTypeFloor string = "floor"
)

// prop value enum
func (m *CreateSiteRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createSiteRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateSiteRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSiteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSiteRequest) UnmarshalBinary(b []byte) error {
	var res CreateSiteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateSiteRequestSite create site request site
// swagger:model CreateSiteRequestSite
type CreateSiteRequestSite struct {

	// area
	Area *CreateSiteRequestSiteArea `json:"area,omitempty"`

	// building
	Building *CreateSiteRequestSiteBuilding `json:"building,omitempty"`

	// floor
	Floor *CreateSiteRequestSiteFloor `json:"floor,omitempty"`
}

// Validate validates this create site request site
func (m *CreateSiteRequestSite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuilding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSiteRequestSite) validateArea(formats strfmt.Registry) error {

	if swag.IsZero(m.Area) { // not required
		return nil
	}

	if m.Area != nil {
		if err := m.Area.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site" + "." + "area")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSiteRequestSite) validateBuilding(formats strfmt.Registry) error {

	if swag.IsZero(m.Building) { // not required
		return nil
	}

	if m.Building != nil {
		if err := m.Building.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site" + "." + "building")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSiteRequestSite) validateFloor(formats strfmt.Registry) error {

	if swag.IsZero(m.Floor) { // not required
		return nil
	}

	if m.Floor != nil {
		if err := m.Floor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site" + "." + "floor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSiteRequestSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSiteRequestSite) UnmarshalBinary(b []byte) error {
	var res CreateSiteRequestSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateSiteRequestSiteArea create site request site area
// swagger:model CreateSiteRequestSiteArea
type CreateSiteRequestSiteArea struct {

	// name
	Name string `json:"name,omitempty"`

	// parent name
	ParentName string `json:"parentName,omitempty"`
}

// Validate validates this create site request site area
func (m *CreateSiteRequestSiteArea) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSiteRequestSiteArea) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSiteRequestSiteArea) UnmarshalBinary(b []byte) error {
	var res CreateSiteRequestSiteArea
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateSiteRequestSiteBuilding create site request site building
// swagger:model CreateSiteRequestSiteBuilding
type CreateSiteRequestSiteBuilding struct {

	// address
	Address string `json:"address,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create site request site building
func (m *CreateSiteRequestSiteBuilding) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSiteRequestSiteBuilding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSiteRequestSiteBuilding) UnmarshalBinary(b []byte) error {
	var res CreateSiteRequestSiteBuilding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateSiteRequestSiteFloor create site request site floor
// swagger:model CreateSiteRequestSiteFloor
type CreateSiteRequestSiteFloor struct {

	// name
	Name string `json:"name,omitempty"`

	// parent name
	ParentName string `json:"parentName,omitempty"`
}

// Validate validates this create site request site floor
func (m *CreateSiteRequestSiteFloor) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSiteRequestSiteFloor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSiteRequestSiteFloor) UnmarshalBinary(b []byte) error {
	var res CreateSiteRequestSiteFloor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
