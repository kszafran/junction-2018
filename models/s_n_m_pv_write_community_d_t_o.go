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

// SNMPvWriteCommunityDTO s n m pv write community d t o
// swagger:model SNMPvWriteCommunityDTO
type SNMPvWriteCommunityDTO []*SNMPvWriteCommunityDTOItems0

// Validate validates this s n m pv write community d t o
func (m SNMPvWriteCommunityDTO) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// SNMPvWriteCommunityDTOItems0 s n m pv write community d t o items0
// swagger:model SNMPvWriteCommunityDTOItems0
type SNMPvWriteCommunityDTOItems0 struct {

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

	// write community
	WriteCommunity string `json:"writeCommunity,omitempty"`
}

// Validate validates this s n m pv write community d t o items0
func (m *SNMPvWriteCommunityDTOItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sNMPvWriteCommunityDTOItems0TypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sNMPvWriteCommunityDTOItems0TypeCredentialTypePropEnum = append(sNMPvWriteCommunityDTOItems0TypeCredentialTypePropEnum, v)
	}
}

const (

	// SNMPvWriteCommunityDTOItems0CredentialTypeGLOBAL captures enum value "GLOBAL"
	SNMPvWriteCommunityDTOItems0CredentialTypeGLOBAL string = "GLOBAL"

	// SNMPvWriteCommunityDTOItems0CredentialTypeAPP captures enum value "APP"
	SNMPvWriteCommunityDTOItems0CredentialTypeAPP string = "APP"
)

// prop value enum
func (m *SNMPvWriteCommunityDTOItems0) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sNMPvWriteCommunityDTOItems0TypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SNMPvWriteCommunityDTOItems0) validateCredentialType(formats strfmt.Registry) error {

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
func (m *SNMPvWriteCommunityDTOItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPvWriteCommunityDTOItems0) UnmarshalBinary(b []byte) error {
	var res SNMPvWriteCommunityDTOItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
