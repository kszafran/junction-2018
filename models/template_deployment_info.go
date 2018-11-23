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

// TemplateDeploymentInfo template deployment info
// swagger:model TemplateDeploymentInfo
type TemplateDeploymentInfo struct {

	// force push template
	ForcePushTemplate bool `json:"forcePushTemplate,omitempty"`

	// is composite
	IsComposite bool `json:"isComposite,omitempty"`

	// main template Id
	MainTemplateID string `json:"mainTemplateId,omitempty"`

	// member template deployment info
	MemberTemplateDeploymentInfo []string `json:"memberTemplateDeploymentInfo"`

	// target info
	TargetInfo []*TemplateDeploymentInfoTargetInfoItems0 `json:"targetInfo"`

	// template Id
	TemplateID string `json:"templateId,omitempty"`
}

// Validate validates this template deployment info
func (m *TemplateDeploymentInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateDeploymentInfo) validateTargetInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetInfo); i++ {
		if swag.IsZero(m.TargetInfo[i]) { // not required
			continue
		}

		if m.TargetInfo[i] != nil {
			if err := m.TargetInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDeploymentInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDeploymentInfo) UnmarshalBinary(b []byte) error {
	var res TemplateDeploymentInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDeploymentInfoTargetInfoItems0 template deployment info target info items0
// swagger:model TemplateDeploymentInfoTargetInfoItems0
type TemplateDeploymentInfoTargetInfoItems0 struct {

	// host name
	HostName string `json:"hostName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// params
	Params interface{} `json:"params,omitempty"`

	// type
	// Enum: [MANAGED_DEVICE_IP MANAGED_DEVICE_UUID PRE_PROVISIONED_SERIAL PRE_PROVISIONED_MAC DEFAULT MANAGED_DEVICE_HOSTNAME]
	Type string `json:"type,omitempty"`
}

// Validate validates this template deployment info target info items0
func (m *TemplateDeploymentInfoTargetInfoItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var templateDeploymentInfoTargetInfoItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANAGED_DEVICE_IP","MANAGED_DEVICE_UUID","PRE_PROVISIONED_SERIAL","PRE_PROVISIONED_MAC","DEFAULT","MANAGED_DEVICE_HOSTNAME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateDeploymentInfoTargetInfoItems0TypeTypePropEnum = append(templateDeploymentInfoTargetInfoItems0TypeTypePropEnum, v)
	}
}

const (

	// TemplateDeploymentInfoTargetInfoItems0TypeMANAGEDDEVICEIP captures enum value "MANAGED_DEVICE_IP"
	TemplateDeploymentInfoTargetInfoItems0TypeMANAGEDDEVICEIP string = "MANAGED_DEVICE_IP"

	// TemplateDeploymentInfoTargetInfoItems0TypeMANAGEDDEVICEUUID captures enum value "MANAGED_DEVICE_UUID"
	TemplateDeploymentInfoTargetInfoItems0TypeMANAGEDDEVICEUUID string = "MANAGED_DEVICE_UUID"

	// TemplateDeploymentInfoTargetInfoItems0TypePREPROVISIONEDSERIAL captures enum value "PRE_PROVISIONED_SERIAL"
	TemplateDeploymentInfoTargetInfoItems0TypePREPROVISIONEDSERIAL string = "PRE_PROVISIONED_SERIAL"

	// TemplateDeploymentInfoTargetInfoItems0TypePREPROVISIONEDMAC captures enum value "PRE_PROVISIONED_MAC"
	TemplateDeploymentInfoTargetInfoItems0TypePREPROVISIONEDMAC string = "PRE_PROVISIONED_MAC"

	// TemplateDeploymentInfoTargetInfoItems0TypeDEFAULT captures enum value "DEFAULT"
	TemplateDeploymentInfoTargetInfoItems0TypeDEFAULT string = "DEFAULT"

	// TemplateDeploymentInfoTargetInfoItems0TypeMANAGEDDEVICEHOSTNAME captures enum value "MANAGED_DEVICE_HOSTNAME"
	TemplateDeploymentInfoTargetInfoItems0TypeMANAGEDDEVICEHOSTNAME string = "MANAGED_DEVICE_HOSTNAME"
)

// prop value enum
func (m *TemplateDeploymentInfoTargetInfoItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, templateDeploymentInfoTargetInfoItems0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemplateDeploymentInfoTargetInfoItems0) validateType(formats strfmt.Registry) error {

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
func (m *TemplateDeploymentInfoTargetInfoItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDeploymentInfoTargetInfoItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDeploymentInfoTargetInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}