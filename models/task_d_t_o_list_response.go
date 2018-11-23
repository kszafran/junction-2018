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

// TaskDTOListResponse task d t o list response
// swagger:model TaskDTOListResponse
type TaskDTOListResponse struct {

	// response
	Response []*TaskDTOListResponseResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this task d t o list response
func (m *TaskDTOListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDTOListResponse) validateResponse(formats strfmt.Registry) error {

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
func (m *TaskDTOListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDTOListResponse) UnmarshalBinary(b []byte) error {
	var res TaskDTOListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskDTOListResponseResponseItems0 task d t o list response response items0
// swagger:model TaskDTOListResponseResponseItems0
type TaskDTOListResponseResponseItems0 struct {

	// additional status URL
	AdditionalStatusURL string `json:"additionalStatusURL,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// error code
	ErrorCode string `json:"errorCode,omitempty"`

	// error key
	ErrorKey string `json:"errorKey,omitempty"`

	// failure reason
	FailureReason string `json:"failureReason,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// instance tenant Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"`

	// is error
	IsError bool `json:"isError,omitempty"`

	// last update
	LastUpdate string `json:"lastUpdate,omitempty"`

	// operation Id list
	OperationIDList interface{} `json:"operationIdList,omitempty"`

	// parent Id
	ParentID string `json:"parentId,omitempty"`

	// progress
	Progress string `json:"progress,omitempty"`

	// root Id
	RootID string `json:"rootId,omitempty"`

	// service type
	ServiceType string `json:"serviceType,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this task d t o list response response items0
func (m *TaskDTOListResponseResponseItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskDTOListResponseResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDTOListResponseResponseItems0) UnmarshalBinary(b []byte) error {
	var res TaskDTOListResponseResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}