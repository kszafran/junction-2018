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

// NetconfCredentialDTO netconf credential d t o
// swagger:model NetconfCredentialDTO
type NetconfCredentialDTO struct {

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

	// netconf port
	NetconfPort string `json:"netconfPort,omitempty"`
}

// Validate validates this netconf credential d t o
func (m *NetconfCredentialDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var netconfCredentialDTOTypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		netconfCredentialDTOTypeCredentialTypePropEnum = append(netconfCredentialDTOTypeCredentialTypePropEnum, v)
	}
}

const (

	// NetconfCredentialDTOCredentialTypeGLOBAL captures enum value "GLOBAL"
	NetconfCredentialDTOCredentialTypeGLOBAL string = "GLOBAL"

	// NetconfCredentialDTOCredentialTypeAPP captures enum value "APP"
	NetconfCredentialDTOCredentialTypeAPP string = "APP"
)

// prop value enum
func (m *NetconfCredentialDTO) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, netconfCredentialDTOTypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetconfCredentialDTO) validateCredentialType(formats strfmt.Registry) error {

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
func (m *NetconfCredentialDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetconfCredentialDTO) UnmarshalBinary(b []byte) error {
	var res NetconfCredentialDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
