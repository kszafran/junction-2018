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

// GetOverallNetworkHealthResponse get overall network health response
// swagger:model GetOverallNetworkHealthResponse
type GetOverallNetworkHealthResponse struct {

	// response
	Response *GetOverallNetworkHealthResponseResponse `json:"response,omitempty"`
}

// Validate validates this get overall network health response
func (m *GetOverallNetworkHealthResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetOverallNetworkHealthResponse) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetOverallNetworkHealthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOverallNetworkHealthResponse) UnmarshalBinary(b []byte) error {
	var res GetOverallNetworkHealthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetOverallNetworkHealthResponseResponse get overall network health response response
// swagger:model GetOverallNetworkHealthResponseResponse
type GetOverallNetworkHealthResponseResponse struct {

	// health distirubution
	HealthDistirubution []*GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0 `json:"healthDistirubution"`

	// latest health score
	LatestHealthScore int `json:"latestHealthScore,omitempty"`

	// latest measured by entity
	LatestMeasuredByEntity string `json:"latestMeasuredByEntity,omitempty"`

	// measured by
	MeasuredBy string `json:"measuredBy,omitempty"`

	// monitored devices
	MonitoredDevices int `json:"monitoredDevices,omitempty"`

	// un monitored devices
	UnMonitoredDevices int `json:"unMonitoredDevices,omitempty"`
}

// Validate validates this get overall network health response response
func (m *GetOverallNetworkHealthResponseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthDistirubution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetOverallNetworkHealthResponseResponse) validateHealthDistirubution(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthDistirubution) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthDistirubution); i++ {
		if swag.IsZero(m.HealthDistirubution[i]) { // not required
			continue
		}

		if m.HealthDistirubution[i] != nil {
			if err := m.HealthDistirubution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("response" + "." + "healthDistirubution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetOverallNetworkHealthResponseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOverallNetworkHealthResponseResponse) UnmarshalBinary(b []byte) error {
	var res GetOverallNetworkHealthResponseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0 get overall network health response response health distirubution items0
// swagger:model GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0
type GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0 struct {

	// bad percentage
	BadPercentage float64 `json:"badPercentage,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// fair percentage
	FairPercentage float64 `json:"fairPercentage,omitempty"`

	// good percentage
	GoodPercentage float64 `json:"goodPercentage,omitempty"`

	// health score
	HealthScore float64 `json:"healthScore,omitempty"`

	// total count
	TotalCount int `json:"totalCount,omitempty"`
}

// Validate validates this get overall network health response response health distirubution items0
func (m *GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0) UnmarshalBinary(b []byte) error {
	var res GetOverallNetworkHealthResponseResponseHealthDistirubutionItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}