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

// TemplateDTO template d t o
// swagger:model TemplateDTO
type TemplateDTO struct {

	// author
	Author string `json:"author,omitempty"`

	// composite
	Composite bool `json:"composite,omitempty"`

	// containing templates
	ContainingTemplates []*TemplateDTOContainingTemplatesItems0 `json:"containingTemplates"`

	// create time
	CreateTime int64 `json:"createTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// device types
	DeviceTypes []*TemplateDTODeviceTypesItems0 `json:"deviceTypes"`

	// failure policy
	// Enum: [ABORT_ON_ERROR CONTINUE_ON_ERROR ROLLBACK_ON_ERROR ROLLBACK_TARGET_ON_ERROR ABORT_TARGET_ON_ERROR]
	FailurePolicy string `json:"failurePolicy,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last update time
	LastUpdateTime int64 `json:"lastUpdateTime,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parent template Id
	ParentTemplateID string `json:"parentTemplateId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// rollback template content
	RollbackTemplateContent string `json:"rollbackTemplateContent,omitempty"`

	// rollback template params
	RollbackTemplateParams []*TemplateDTORollbackTemplateParamsItems0 `json:"rollbackTemplateParams"`

	// software type
	SoftwareType string `json:"softwareType,omitempty"`

	// software variant
	SoftwareVariant string `json:"softwareVariant,omitempty"`

	// software version
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// template content
	TemplateContent string `json:"templateContent,omitempty"`

	// template params
	TemplateParams []*TemplateDTOTemplateParamsItems0 `json:"templateParams"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this template d t o
func (m *TemplateDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainingTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailurePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRollbackTemplateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateDTO) validateContainingTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainingTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.ContainingTemplates); i++ {
		if swag.IsZero(m.ContainingTemplates[i]) { // not required
			continue
		}

		if m.ContainingTemplates[i] != nil {
			if err := m.ContainingTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containingTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TemplateDTO) validateDeviceTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceTypes); i++ {
		if swag.IsZero(m.DeviceTypes[i]) { // not required
			continue
		}

		if m.DeviceTypes[i] != nil {
			if err := m.DeviceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deviceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var templateDTOTypeFailurePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ABORT_ON_ERROR","CONTINUE_ON_ERROR","ROLLBACK_ON_ERROR","ROLLBACK_TARGET_ON_ERROR","ABORT_TARGET_ON_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateDTOTypeFailurePolicyPropEnum = append(templateDTOTypeFailurePolicyPropEnum, v)
	}
}

const (

	// TemplateDTOFailurePolicyABORTONERROR captures enum value "ABORT_ON_ERROR"
	TemplateDTOFailurePolicyABORTONERROR string = "ABORT_ON_ERROR"

	// TemplateDTOFailurePolicyCONTINUEONERROR captures enum value "CONTINUE_ON_ERROR"
	TemplateDTOFailurePolicyCONTINUEONERROR string = "CONTINUE_ON_ERROR"

	// TemplateDTOFailurePolicyROLLBACKONERROR captures enum value "ROLLBACK_ON_ERROR"
	TemplateDTOFailurePolicyROLLBACKONERROR string = "ROLLBACK_ON_ERROR"

	// TemplateDTOFailurePolicyROLLBACKTARGETONERROR captures enum value "ROLLBACK_TARGET_ON_ERROR"
	TemplateDTOFailurePolicyROLLBACKTARGETONERROR string = "ROLLBACK_TARGET_ON_ERROR"

	// TemplateDTOFailurePolicyABORTTARGETONERROR captures enum value "ABORT_TARGET_ON_ERROR"
	TemplateDTOFailurePolicyABORTTARGETONERROR string = "ABORT_TARGET_ON_ERROR"
)

// prop value enum
func (m *TemplateDTO) validateFailurePolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateDTOTypeFailurePolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemplateDTO) validateFailurePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.FailurePolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateFailurePolicyEnum("failurePolicy", "body", m.FailurePolicy); err != nil {
		return err
	}

	return nil
}

func (m *TemplateDTO) validateRollbackTemplateParams(formats strfmt.Registry) error {

	if swag.IsZero(m.RollbackTemplateParams) { // not required
		return nil
	}

	for i := 0; i < len(m.RollbackTemplateParams); i++ {
		if swag.IsZero(m.RollbackTemplateParams[i]) { // not required
			continue
		}

		if m.RollbackTemplateParams[i] != nil {
			if err := m.RollbackTemplateParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rollbackTemplateParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TemplateDTO) validateTemplateParams(formats strfmt.Registry) error {

	if swag.IsZero(m.TemplateParams) { // not required
		return nil
	}

	for i := 0; i < len(m.TemplateParams); i++ {
		if swag.IsZero(m.TemplateParams[i]) { // not required
			continue
		}

		if m.TemplateParams[i] != nil {
			if err := m.TemplateParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templateParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTO) UnmarshalBinary(b []byte) error {
	var res TemplateDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTOContainingTemplatesItems0 template d t o containing templates items0
// swagger:model TemplateDTOContainingTemplatesItems0
type TemplateDTOContainingTemplatesItems0 struct {

	// composite
	Composite bool `json:"composite,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this template d t o containing templates items0
func (m *TemplateDTOContainingTemplatesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTOContainingTemplatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTOContainingTemplatesItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDTOContainingTemplatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTODeviceTypesItems0 template d t o device types items0
// swagger:model TemplateDTODeviceTypesItems0
type TemplateDTODeviceTypesItems0 struct {

	// product family
	ProductFamily string `json:"productFamily,omitempty"`

	// product series
	ProductSeries string `json:"productSeries,omitempty"`

	// product type
	ProductType string `json:"productType,omitempty"`
}

// Validate validates this template d t o device types items0
func (m *TemplateDTODeviceTypesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTODeviceTypesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTODeviceTypesItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDTODeviceTypesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTORollbackTemplateParamsItems0 template d t o rollback template params items0
// swagger:model TemplateDTORollbackTemplateParamsItems0
type TemplateDTORollbackTemplateParamsItems0 struct {

	// binding
	Binding string `json:"binding,omitempty"`

	// data type
	// Enum: [STRING INTEGER IPADDRESS MACADDRESS SECTIONDIVIDER]
	DataType string `json:"dataType,omitempty"`

	// default value
	DefaultValue string `json:"defaultValue,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// instruction text
	InstructionText string `json:"instructionText,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// not param
	NotParam bool `json:"notParam,omitempty"`

	// order
	Order int64 `json:"order,omitempty"`

	// param array
	ParamArray bool `json:"paramArray,omitempty"`

	// parameter name
	ParameterName string `json:"parameterName,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// range
	Range []*TemplateDTORollbackTemplateParamsItems0RangeItems0 `json:"range"`

	// required
	Required bool `json:"required,omitempty"`

	// selection
	Selection *TemplateDTORollbackTemplateParamsItems0Selection `json:"selection,omitempty"`
}

// Validate validates this template d t o rollback template params items0
func (m *TemplateDTORollbackTemplateParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var templateDTORollbackTemplateParamsItems0TypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRING","INTEGER","IPADDRESS","MACADDRESS","SECTIONDIVIDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateDTORollbackTemplateParamsItems0TypeDataTypePropEnum = append(templateDTORollbackTemplateParamsItems0TypeDataTypePropEnum, v)
	}
}

const (

	// TemplateDTORollbackTemplateParamsItems0DataTypeSTRING captures enum value "STRING"
	TemplateDTORollbackTemplateParamsItems0DataTypeSTRING string = "STRING"

	// TemplateDTORollbackTemplateParamsItems0DataTypeINTEGER captures enum value "INTEGER"
	TemplateDTORollbackTemplateParamsItems0DataTypeINTEGER string = "INTEGER"

	// TemplateDTORollbackTemplateParamsItems0DataTypeIPADDRESS captures enum value "IPADDRESS"
	TemplateDTORollbackTemplateParamsItems0DataTypeIPADDRESS string = "IPADDRESS"

	// TemplateDTORollbackTemplateParamsItems0DataTypeMACADDRESS captures enum value "MACADDRESS"
	TemplateDTORollbackTemplateParamsItems0DataTypeMACADDRESS string = "MACADDRESS"

	// TemplateDTORollbackTemplateParamsItems0DataTypeSECTIONDIVIDER captures enum value "SECTIONDIVIDER"
	TemplateDTORollbackTemplateParamsItems0DataTypeSECTIONDIVIDER string = "SECTIONDIVIDER"
)

// prop value enum
func (m *TemplateDTORollbackTemplateParamsItems0) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateDTORollbackTemplateParamsItems0TypeDataTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemplateDTORollbackTemplateParamsItems0) validateDataType(formats strfmt.Registry) error {

	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *TemplateDTORollbackTemplateParamsItems0) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	for i := 0; i < len(m.Range); i++ {
		if swag.IsZero(m.Range[i]) { // not required
			continue
		}

		if m.Range[i] != nil {
			if err := m.Range[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("range" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TemplateDTORollbackTemplateParamsItems0) validateSelection(formats strfmt.Registry) error {

	if swag.IsZero(m.Selection) { // not required
		return nil
	}

	if m.Selection != nil {
		if err := m.Selection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTORollbackTemplateParamsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTORollbackTemplateParamsItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDTORollbackTemplateParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTORollbackTemplateParamsItems0RangeItems0 template d t o rollback template params items0 range items0
// swagger:model TemplateDTORollbackTemplateParamsItems0RangeItems0
type TemplateDTORollbackTemplateParamsItems0RangeItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// max value
	MaxValue int64 `json:"maxValue,omitempty"`

	// min value
	MinValue int64 `json:"minValue,omitempty"`
}

// Validate validates this template d t o rollback template params items0 range items0
func (m *TemplateDTORollbackTemplateParamsItems0RangeItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTORollbackTemplateParamsItems0RangeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTORollbackTemplateParamsItems0RangeItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDTORollbackTemplateParamsItems0RangeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTORollbackTemplateParamsItems0Selection template d t o rollback template params items0 selection
// swagger:model TemplateDTORollbackTemplateParamsItems0Selection
type TemplateDTORollbackTemplateParamsItems0Selection struct {

	// id
	ID string `json:"id,omitempty"`

	// selection type
	// Enum: [SINGLE_SELECT MULTI_SELECT]
	SelectionType string `json:"selectionType,omitempty"`

	// selection values
	SelectionValues interface{} `json:"selectionValues,omitempty"`
}

// Validate validates this template d t o rollback template params items0 selection
func (m *TemplateDTORollbackTemplateParamsItems0Selection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelectionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var templateDTORollbackTemplateParamsItems0SelectionTypeSelectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE_SELECT","MULTI_SELECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateDTORollbackTemplateParamsItems0SelectionTypeSelectionTypePropEnum = append(templateDTORollbackTemplateParamsItems0SelectionTypeSelectionTypePropEnum, v)
	}
}

const (

	// TemplateDTORollbackTemplateParamsItems0SelectionSelectionTypeSINGLESELECT captures enum value "SINGLE_SELECT"
	TemplateDTORollbackTemplateParamsItems0SelectionSelectionTypeSINGLESELECT string = "SINGLE_SELECT"

	// TemplateDTORollbackTemplateParamsItems0SelectionSelectionTypeMULTISELECT captures enum value "MULTI_SELECT"
	TemplateDTORollbackTemplateParamsItems0SelectionSelectionTypeMULTISELECT string = "MULTI_SELECT"
)

// prop value enum
func (m *TemplateDTORollbackTemplateParamsItems0Selection) validateSelectionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateDTORollbackTemplateParamsItems0SelectionTypeSelectionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemplateDTORollbackTemplateParamsItems0Selection) validateSelectionType(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSelectionTypeEnum("selection"+"."+"selectionType", "body", m.SelectionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTORollbackTemplateParamsItems0Selection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTORollbackTemplateParamsItems0Selection) UnmarshalBinary(b []byte) error {
	var res TemplateDTORollbackTemplateParamsItems0Selection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTOTemplateParamsItems0 template d t o template params items0
// swagger:model TemplateDTOTemplateParamsItems0
type TemplateDTOTemplateParamsItems0 struct {

	// binding
	Binding string `json:"binding,omitempty"`

	// data type
	// Enum: [STRING INTEGER IPADDRESS MACADDRESS SECTIONDIVIDER]
	DataType string `json:"dataType,omitempty"`

	// default value
	DefaultValue string `json:"defaultValue,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// instruction text
	InstructionText string `json:"instructionText,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// not param
	NotParam bool `json:"notParam,omitempty"`

	// order
	Order int64 `json:"order,omitempty"`

	// param array
	ParamArray bool `json:"paramArray,omitempty"`

	// parameter name
	ParameterName string `json:"parameterName,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// range
	Range []*TemplateDTOTemplateParamsItems0RangeItems0 `json:"range"`

	// required
	Required bool `json:"required,omitempty"`

	// selection
	Selection *TemplateDTOTemplateParamsItems0Selection `json:"selection,omitempty"`
}

// Validate validates this template d t o template params items0
func (m *TemplateDTOTemplateParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var templateDTOTemplateParamsItems0TypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRING","INTEGER","IPADDRESS","MACADDRESS","SECTIONDIVIDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateDTOTemplateParamsItems0TypeDataTypePropEnum = append(templateDTOTemplateParamsItems0TypeDataTypePropEnum, v)
	}
}

const (

	// TemplateDTOTemplateParamsItems0DataTypeSTRING captures enum value "STRING"
	TemplateDTOTemplateParamsItems0DataTypeSTRING string = "STRING"

	// TemplateDTOTemplateParamsItems0DataTypeINTEGER captures enum value "INTEGER"
	TemplateDTOTemplateParamsItems0DataTypeINTEGER string = "INTEGER"

	// TemplateDTOTemplateParamsItems0DataTypeIPADDRESS captures enum value "IPADDRESS"
	TemplateDTOTemplateParamsItems0DataTypeIPADDRESS string = "IPADDRESS"

	// TemplateDTOTemplateParamsItems0DataTypeMACADDRESS captures enum value "MACADDRESS"
	TemplateDTOTemplateParamsItems0DataTypeMACADDRESS string = "MACADDRESS"

	// TemplateDTOTemplateParamsItems0DataTypeSECTIONDIVIDER captures enum value "SECTIONDIVIDER"
	TemplateDTOTemplateParamsItems0DataTypeSECTIONDIVIDER string = "SECTIONDIVIDER"
)

// prop value enum
func (m *TemplateDTOTemplateParamsItems0) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateDTOTemplateParamsItems0TypeDataTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemplateDTOTemplateParamsItems0) validateDataType(formats strfmt.Registry) error {

	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *TemplateDTOTemplateParamsItems0) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	for i := 0; i < len(m.Range); i++ {
		if swag.IsZero(m.Range[i]) { // not required
			continue
		}

		if m.Range[i] != nil {
			if err := m.Range[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("range" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TemplateDTOTemplateParamsItems0) validateSelection(formats strfmt.Registry) error {

	if swag.IsZero(m.Selection) { // not required
		return nil
	}

	if m.Selection != nil {
		if err := m.Selection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTOTemplateParamsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTOTemplateParamsItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDTOTemplateParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTOTemplateParamsItems0RangeItems0 template d t o template params items0 range items0
// swagger:model TemplateDTOTemplateParamsItems0RangeItems0
type TemplateDTOTemplateParamsItems0RangeItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// max value
	MaxValue int64 `json:"maxValue,omitempty"`

	// min value
	MinValue int64 `json:"minValue,omitempty"`
}

// Validate validates this template d t o template params items0 range items0
func (m *TemplateDTOTemplateParamsItems0RangeItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTOTemplateParamsItems0RangeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTOTemplateParamsItems0RangeItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDTOTemplateParamsItems0RangeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDTOTemplateParamsItems0Selection template d t o template params items0 selection
// swagger:model TemplateDTOTemplateParamsItems0Selection
type TemplateDTOTemplateParamsItems0Selection struct {

	// id
	ID string `json:"id,omitempty"`

	// selection type
	// Enum: [SINGLE_SELECT MULTI_SELECT]
	SelectionType string `json:"selectionType,omitempty"`

	// selection values
	SelectionValues interface{} `json:"selectionValues,omitempty"`
}

// Validate validates this template d t o template params items0 selection
func (m *TemplateDTOTemplateParamsItems0Selection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelectionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var templateDTOTemplateParamsItems0SelectionTypeSelectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE_SELECT","MULTI_SELECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateDTOTemplateParamsItems0SelectionTypeSelectionTypePropEnum = append(templateDTOTemplateParamsItems0SelectionTypeSelectionTypePropEnum, v)
	}
}

const (

	// TemplateDTOTemplateParamsItems0SelectionSelectionTypeSINGLESELECT captures enum value "SINGLE_SELECT"
	TemplateDTOTemplateParamsItems0SelectionSelectionTypeSINGLESELECT string = "SINGLE_SELECT"

	// TemplateDTOTemplateParamsItems0SelectionSelectionTypeMULTISELECT captures enum value "MULTI_SELECT"
	TemplateDTOTemplateParamsItems0SelectionSelectionTypeMULTISELECT string = "MULTI_SELECT"
)

// prop value enum
func (m *TemplateDTOTemplateParamsItems0Selection) validateSelectionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateDTOTemplateParamsItems0SelectionTypeSelectionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemplateDTOTemplateParamsItems0Selection) validateSelectionType(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSelectionTypeEnum("selection"+"."+"selectionType", "body", m.SelectionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDTOTemplateParamsItems0Selection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDTOTemplateParamsItems0Selection) UnmarshalBinary(b []byte) error {
	var res TemplateDTOTemplateParamsItems0Selection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
