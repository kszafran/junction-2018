// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModuleListResult module list result
// swagger:model ModuleListResult
type ModuleListResult struct {

	// response
	Response []*ModuleListResultResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this module list result
func (m *ModuleListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleListResult) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	for i := 0; i < len(m.Response); i++ {
		if swag.IsZero(m.Response[i]) { // not required
			continue
		}

		if m.Response[i] != nil {
			if err := m.Response[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModuleListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleListResult) UnmarshalBinary(b []byte) error {
	var res ModuleListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModuleListResultResponseItems0 module list result response items0
// swagger:model ModuleListResultResponseItems0
type ModuleListResultResponseItems0 struct {

	// assembly number
	AssemblyNumber string `json:"assemblyNumber,omitempty"`

	// assembly revision
	AssemblyRevision string `json:"assemblyRevision,omitempty"`

	// attribute info
	AttributeInfo interface{} `json:"attributeInfo,omitempty"`

	// containment entity
	ContainmentEntity string `json:"containmentEntity,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// entity physical index
	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is field replaceable
	// Enum: [UNKNOWN TRUE FALSE NOT_APPLICABLE]
	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"`

	// is reporting alarms allowed
	// Enum: [UNKNOWN TRUE FALSE NOT_APPLICABLE]
	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"`

	// manufacturer
	Manufacturer string `json:"manufacturer,omitempty"`

	// module index
	ModuleIndex int64 `json:"moduleIndex,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// operational state code
	OperationalStateCode string `json:"operationalStateCode,omitempty"`

	// part number
	PartNumber string `json:"partNumber,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

	// vendor equipment type
	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"`
}

// Validate validates this module list result response items0
func (m *ModuleListResultResponseItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsFieldReplaceable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsReportingAlarmsAllowed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var moduleListResultResponseItems0TypeIsFieldReplaceablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","TRUE","FALSE","NOT_APPLICABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moduleListResultResponseItems0TypeIsFieldReplaceablePropEnum = append(moduleListResultResponseItems0TypeIsFieldReplaceablePropEnum, v)
	}
}

const (

	// ModuleListResultResponseItems0IsFieldReplaceableUNKNOWN captures enum value "UNKNOWN"
	ModuleListResultResponseItems0IsFieldReplaceableUNKNOWN string = "UNKNOWN"

	// ModuleListResultResponseItems0IsFieldReplaceableTRUE captures enum value "TRUE"
	ModuleListResultResponseItems0IsFieldReplaceableTRUE string = "TRUE"

	// ModuleListResultResponseItems0IsFieldReplaceableFALSE captures enum value "FALSE"
	ModuleListResultResponseItems0IsFieldReplaceableFALSE string = "FALSE"

	// ModuleListResultResponseItems0IsFieldReplaceableNOTAPPLICABLE captures enum value "NOT_APPLICABLE"
	ModuleListResultResponseItems0IsFieldReplaceableNOTAPPLICABLE string = "NOT_APPLICABLE"
)

// prop value enum
func (m *ModuleListResultResponseItems0) validateIsFieldReplaceableEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moduleListResultResponseItems0TypeIsFieldReplaceablePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ModuleListResultResponseItems0) validateIsFieldReplaceable(formats strfmt.Registry) error {

	if swag.IsZero(m.IsFieldReplaceable) { // not required
		return nil
	}

	// value enum
	if err := m.validateIsFieldReplaceableEnum("isFieldReplaceable", "body", m.IsFieldReplaceable); err != nil {
		return err
	}

	return nil
}

var moduleListResultResponseItems0TypeIsReportingAlarmsAllowedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","TRUE","FALSE","NOT_APPLICABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moduleListResultResponseItems0TypeIsReportingAlarmsAllowedPropEnum = append(moduleListResultResponseItems0TypeIsReportingAlarmsAllowedPropEnum, v)
	}
}

const (

	// ModuleListResultResponseItems0IsReportingAlarmsAllowedUNKNOWN captures enum value "UNKNOWN"
	ModuleListResultResponseItems0IsReportingAlarmsAllowedUNKNOWN string = "UNKNOWN"

	// ModuleListResultResponseItems0IsReportingAlarmsAllowedTRUE captures enum value "TRUE"
	ModuleListResultResponseItems0IsReportingAlarmsAllowedTRUE string = "TRUE"

	// ModuleListResultResponseItems0IsReportingAlarmsAllowedFALSE captures enum value "FALSE"
	ModuleListResultResponseItems0IsReportingAlarmsAllowedFALSE string = "FALSE"

	// ModuleListResultResponseItems0IsReportingAlarmsAllowedNOTAPPLICABLE captures enum value "NOT_APPLICABLE"
	ModuleListResultResponseItems0IsReportingAlarmsAllowedNOTAPPLICABLE string = "NOT_APPLICABLE"
)

// prop value enum
func (m *ModuleListResultResponseItems0) validateIsReportingAlarmsAllowedEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moduleListResultResponseItems0TypeIsReportingAlarmsAllowedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ModuleListResultResponseItems0) validateIsReportingAlarmsAllowed(formats strfmt.Registry) error {

	if swag.IsZero(m.IsReportingAlarmsAllowed) { // not required
		return nil
	}

	// value enum
	if err := m.validateIsReportingAlarmsAllowedEnum("isReportingAlarmsAllowed", "body", m.IsReportingAlarmsAllowed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModuleListResultResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleListResultResponseItems0) UnmarshalBinary(b []byte) error {
	var res ModuleListResultResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}