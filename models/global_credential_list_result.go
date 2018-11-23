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

// GlobalCredentialListResult global credential list result
// swagger:model GlobalCredentialListResult
type GlobalCredentialListResult struct {

	// response
	Response []*GlobalCredentialListResultResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this global credential list result
func (m *GlobalCredentialListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalCredentialListResult) validateResponse(formats strfmt.Registry) error {

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
func (m *GlobalCredentialListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalCredentialListResult) UnmarshalBinary(b []byte) error {
	var res GlobalCredentialListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GlobalCredentialListResultResponseItems0 global credential list result response items0
// swagger:model GlobalCredentialListResultResponseItems0
type GlobalCredentialListResultResponseItems0 struct {

	// comments
	Comments string `json:"comments,omitempty"`

	// credential type
	// Enum: [GLOBAL APP]
	CredentialType string `json:"credentialType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// instance tenant Id
	InstanceTenantID string `json:"instanceTenantId,omitempty"`

	// instance Uuid
	InstanceUUID string `json:"instanceUuid,omitempty"`
}

// Validate validates this global credential list result response items0
func (m *GlobalCredentialListResultResponseItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var globalCredentialListResultResponseItems0TypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalCredentialListResultResponseItems0TypeCredentialTypePropEnum = append(globalCredentialListResultResponseItems0TypeCredentialTypePropEnum, v)
	}
}

const (

	// GlobalCredentialListResultResponseItems0CredentialTypeGLOBAL captures enum value "GLOBAL"
	GlobalCredentialListResultResponseItems0CredentialTypeGLOBAL string = "GLOBAL"

	// GlobalCredentialListResultResponseItems0CredentialTypeAPP captures enum value "APP"
	GlobalCredentialListResultResponseItems0CredentialTypeAPP string = "APP"
)

// prop value enum
func (m *GlobalCredentialListResultResponseItems0) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalCredentialListResultResponseItems0TypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GlobalCredentialListResultResponseItems0) validateCredentialType(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCredentialTypeEnum("credentialType", "body", m.CredentialType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalCredentialListResultResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalCredentialListResultResponseItems0) UnmarshalBinary(b []byte) error {
	var res GlobalCredentialListResultResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}