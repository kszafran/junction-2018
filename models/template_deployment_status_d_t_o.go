// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TemplateDeploymentStatusDTO template deployment status d t o
// swagger:model TemplateDeploymentStatusDTO
type TemplateDeploymentStatusDTO struct {

	// deployment Id
	DeploymentID string `json:"deploymentId,omitempty"`

	// deployment name
	DeploymentName string `json:"deploymentName,omitempty"`

	// devices
	Devices []*TemplateDeploymentStatusDTODevicesItems0 `json:"devices"`

	// duration
	Duration string `json:"duration,omitempty"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// template name
	TemplateName string `json:"templateName,omitempty"`

	// template version
	TemplateVersion string `json:"templateVersion,omitempty"`
}

// Validate validates this template deployment status d t o
func (m *TemplateDeploymentStatusDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateDeploymentStatusDTO) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDeploymentStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDeploymentStatusDTO) UnmarshalBinary(b []byte) error {
	var res TemplateDeploymentStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TemplateDeploymentStatusDTODevicesItems0 template deployment status d t o devices items0
// swagger:model TemplateDeploymentStatusDTODevicesItems0
type TemplateDeploymentStatusDTODevicesItems0 struct {

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this template deployment status d t o devices items0
func (m *TemplateDeploymentStatusDTODevicesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateDeploymentStatusDTODevicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateDeploymentStatusDTODevicesItems0) UnmarshalBinary(b []byte) error {
	var res TemplateDeploymentStatusDTODevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
