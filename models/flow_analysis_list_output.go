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

// FlowAnalysisListOutput flow analysis list output
// swagger:model FlowAnalysisListOutput
type FlowAnalysisListOutput struct {

	// response
	Response []*FlowAnalysisListOutputResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this flow analysis list output
func (m *FlowAnalysisListOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowAnalysisListOutput) validateResponse(formats strfmt.Registry) error {

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
func (m *FlowAnalysisListOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowAnalysisListOutput) UnmarshalBinary(b []byte) error {
	var res FlowAnalysisListOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FlowAnalysisListOutputResponseItems0 flow analysis list output response items0
// swagger:model FlowAnalysisListOutputResponseItems0
type FlowAnalysisListOutputResponseItems0 struct {

	// control path
	ControlPath bool `json:"controlPath,omitempty"`

	// create time
	CreateTime int64 `json:"createTime,omitempty"`

	// dest IP
	DestIP string `json:"destIP,omitempty"`

	// dest port
	DestPort string `json:"destPort,omitempty"`

	// failure reason
	FailureReason string `json:"failureReason,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// inclusions
	Inclusions []string `json:"inclusions"`

	// last update time
	LastUpdateTime int64 `json:"lastUpdateTime,omitempty"`

	// periodic refresh
	PeriodicRefresh bool `json:"periodicRefresh,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// source IP
	SourceIP string `json:"sourceIP,omitempty"`

	// source port
	SourcePort string `json:"sourcePort,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this flow analysis list output response items0
func (m *FlowAnalysisListOutputResponseItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlowAnalysisListOutputResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowAnalysisListOutputResponseItems0) UnmarshalBinary(b []byte) error {
	var res FlowAnalysisListOutputResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}