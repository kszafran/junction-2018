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

// SAVAMapping s a v a mapping
// swagger:model SAVAMapping
type SAVAMapping struct {

	// auto sync period
	AutoSyncPeriod int64 `json:"autoSyncPeriod,omitempty"`

	// cco user
	CcoUser string `json:"ccoUser,omitempty"`

	// expiry
	Expiry int64 `json:"expiry,omitempty"`

	// last sync
	LastSync int64 `json:"lastSync,omitempty"`

	// profile
	Profile *SAVAMappingProfile `json:"profile,omitempty"`

	// smart account Id
	SmartAccountID string `json:"smartAccountId,omitempty"`

	// sync result
	SyncResult *SAVAMappingSyncResult `json:"syncResult,omitempty"`

	// sync result str
	SyncResultStr string `json:"syncResultStr,omitempty"`

	// sync start time
	SyncStartTime int64 `json:"syncStartTime,omitempty"`

	// sync status
	// Enum: [NOT_SYNCED SYNCING SUCCESS FAILURE]
	SyncStatus string `json:"syncStatus,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// virtual account Id
	VirtualAccountID string `json:"virtualAccountId,omitempty"`
}

// Validate validates this s a v a mapping
func (m *SAVAMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAVAMapping) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *SAVAMapping) validateSyncResult(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncResult) { // not required
		return nil
	}

	if m.SyncResult != nil {
		if err := m.SyncResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncResult")
			}
			return err
		}
	}

	return nil
}

var sAVAMappingTypeSyncStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_SYNCED","SYNCING","SUCCESS","FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sAVAMappingTypeSyncStatusPropEnum = append(sAVAMappingTypeSyncStatusPropEnum, v)
	}
}

const (

	// SAVAMappingSyncStatusNOTSYNCED captures enum value "NOT_SYNCED"
	SAVAMappingSyncStatusNOTSYNCED string = "NOT_SYNCED"

	// SAVAMappingSyncStatusSYNCING captures enum value "SYNCING"
	SAVAMappingSyncStatusSYNCING string = "SYNCING"

	// SAVAMappingSyncStatusSUCCESS captures enum value "SUCCESS"
	SAVAMappingSyncStatusSUCCESS string = "SUCCESS"

	// SAVAMappingSyncStatusFAILURE captures enum value "FAILURE"
	SAVAMappingSyncStatusFAILURE string = "FAILURE"
)

// prop value enum
func (m *SAVAMapping) validateSyncStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sAVAMappingTypeSyncStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SAVAMapping) validateSyncStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSyncStatusEnum("syncStatus", "body", m.SyncStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SAVAMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAVAMapping) UnmarshalBinary(b []byte) error {
	var res SAVAMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SAVAMappingProfile s a v a mapping profile
// swagger:model SAVAMappingProfile
type SAVAMappingProfile struct {

	// address fqdn
	AddressFqdn string `json:"addressFqdn,omitempty"`

	// address Ip v4
	AddressIPV4 string `json:"addressIpV4,omitempty"`

	// cert
	Cert string `json:"cert,omitempty"`

	// make default
	MakeDefault bool `json:"makeDefault,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// profile Id
	ProfileID string `json:"profileId,omitempty"`

	// proxy
	Proxy bool `json:"proxy,omitempty"`
}

// Validate validates this s a v a mapping profile
func (m *SAVAMappingProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SAVAMappingProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAVAMappingProfile) UnmarshalBinary(b []byte) error {
	var res SAVAMappingProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SAVAMappingSyncResult s a v a mapping sync result
// swagger:model SAVAMappingSyncResult
type SAVAMappingSyncResult struct {

	// sync list
	SyncList []*SAVAMappingSyncResultSyncListItems0 `json:"syncList"`

	// sync msg
	SyncMsg string `json:"syncMsg,omitempty"`
}

// Validate validates this s a v a mapping sync result
func (m *SAVAMappingSyncResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAVAMappingSyncResult) validateSyncList(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncList) { // not required
		return nil
	}

	for i := 0; i < len(m.SyncList); i++ {
		if swag.IsZero(m.SyncList[i]) { // not required
			continue
		}

		if m.SyncList[i] != nil {
			if err := m.SyncList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("syncResult" + "." + "syncList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SAVAMappingSyncResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAVAMappingSyncResult) UnmarshalBinary(b []byte) error {
	var res SAVAMappingSyncResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SAVAMappingSyncResultSyncListItems0 s a v a mapping sync result sync list items0
// swagger:model SAVAMappingSyncResultSyncListItems0
type SAVAMappingSyncResultSyncListItems0 struct {

	// device sn list
	DeviceSnList []string `json:"deviceSnList"`

	// sync type
	// Enum: [Add Update Delete MismatchError]
	SyncType string `json:"syncType,omitempty"`
}

// Validate validates this s a v a mapping sync result sync list items0
func (m *SAVAMappingSyncResultSyncListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sAVAMappingSyncResultSyncListItems0TypeSyncTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Add","Update","Delete","MismatchError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sAVAMappingSyncResultSyncListItems0TypeSyncTypePropEnum = append(sAVAMappingSyncResultSyncListItems0TypeSyncTypePropEnum, v)
	}
}

const (

	// SAVAMappingSyncResultSyncListItems0SyncTypeAdd captures enum value "Add"
	SAVAMappingSyncResultSyncListItems0SyncTypeAdd string = "Add"

	// SAVAMappingSyncResultSyncListItems0SyncTypeUpdate captures enum value "Update"
	SAVAMappingSyncResultSyncListItems0SyncTypeUpdate string = "Update"

	// SAVAMappingSyncResultSyncListItems0SyncTypeDelete captures enum value "Delete"
	SAVAMappingSyncResultSyncListItems0SyncTypeDelete string = "Delete"

	// SAVAMappingSyncResultSyncListItems0SyncTypeMismatchError captures enum value "MismatchError"
	SAVAMappingSyncResultSyncListItems0SyncTypeMismatchError string = "MismatchError"
)

// prop value enum
func (m *SAVAMappingSyncResultSyncListItems0) validateSyncTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sAVAMappingSyncResultSyncListItems0TypeSyncTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SAVAMappingSyncResultSyncListItems0) validateSyncType(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSyncTypeEnum("syncType", "body", m.SyncType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SAVAMappingSyncResultSyncListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAVAMappingSyncResultSyncListItems0) UnmarshalBinary(b []byte) error {
	var res SAVAMappingSyncResultSyncListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
