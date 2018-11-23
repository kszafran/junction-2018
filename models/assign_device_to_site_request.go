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

// AssignDeviceToSiteRequest assign device to site request
// swagger:model AssignDeviceToSiteRequest
type AssignDeviceToSiteRequest struct {

	// device
	Device []*AssignDeviceToSiteRequestDeviceItems0 `json:"device"`
}

// Validate validates this assign device to site request
func (m *AssignDeviceToSiteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssignDeviceToSiteRequest) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	for i := 0; i < len(m.Device); i++ {
		if swag.IsZero(m.Device[i]) { // not required
			continue
		}

		if m.Device[i] != nil {
			if err := m.Device[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("device" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssignDeviceToSiteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignDeviceToSiteRequest) UnmarshalBinary(b []byte) error {
	var res AssignDeviceToSiteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AssignDeviceToSiteRequestDeviceItems0 assign device to site request device items0
// swagger:model AssignDeviceToSiteRequestDeviceItems0
type AssignDeviceToSiteRequestDeviceItems0 struct {

	// ip
	IP string `json:"ip,omitempty"`
}

// Validate validates this assign device to site request device items0
func (m *AssignDeviceToSiteRequestDeviceItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssignDeviceToSiteRequestDeviceItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignDeviceToSiteRequestDeviceItems0) UnmarshalBinary(b []byte) error {
	var res AssignDeviceToSiteRequestDeviceItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}