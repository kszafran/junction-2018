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

// Settings settings
// swagger:model Settings
type Settings struct {

	// id
	ID string `json:"_id,omitempty"`

	// aaa credentials
	AaaCredentials *SettingsAaaCredentials `json:"aaaCredentials,omitempty"`

	// accept eula
	AcceptEula bool `json:"acceptEula,omitempty"`

	// default profile
	DefaultProfile *SettingsDefaultProfile `json:"defaultProfile,omitempty"`

	// sava mapping list
	SavaMappingList []*SettingsSavaMappingListItems0 `json:"savaMappingList"`

	// task time outs
	TaskTimeOuts *SettingsTaskTimeOuts `json:"taskTimeOuts,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this settings
func (m *Settings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAaaCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSavaMappingList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskTimeOuts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Settings) validateAaaCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.AaaCredentials) { // not required
		return nil
	}

	if m.AaaCredentials != nil {
		if err := m.AaaCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aaaCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *Settings) validateDefaultProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultProfile) { // not required
		return nil
	}

	if m.DefaultProfile != nil {
		if err := m.DefaultProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultProfile")
			}
			return err
		}
	}

	return nil
}

func (m *Settings) validateSavaMappingList(formats strfmt.Registry) error {

	if swag.IsZero(m.SavaMappingList) { // not required
		return nil
	}

	for i := 0; i < len(m.SavaMappingList); i++ {
		if swag.IsZero(m.SavaMappingList[i]) { // not required
			continue
		}

		if m.SavaMappingList[i] != nil {
			if err := m.SavaMappingList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("savaMappingList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Settings) validateTaskTimeOuts(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskTimeOuts) { // not required
		return nil
	}

	if m.TaskTimeOuts != nil {
		if err := m.TaskTimeOuts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskTimeOuts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Settings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Settings) UnmarshalBinary(b []byte) error {
	var res Settings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsAaaCredentials settings aaa credentials
// swagger:model SettingsAaaCredentials
type SettingsAaaCredentials struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this settings aaa credentials
func (m *SettingsAaaCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsAaaCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsAaaCredentials) UnmarshalBinary(b []byte) error {
	var res SettingsAaaCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsDefaultProfile settings default profile
// swagger:model SettingsDefaultProfile
type SettingsDefaultProfile struct {

	// cert
	Cert string `json:"cert,omitempty"`

	// fqdn addresses
	FqdnAddresses []string `json:"fqdnAddresses"`

	// ip addresses
	IPAddresses []string `json:"ipAddresses"`

	// port
	Port int64 `json:"port,omitempty"`

	// proxy
	Proxy bool `json:"proxy,omitempty"`
}

// Validate validates this settings default profile
func (m *SettingsDefaultProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsDefaultProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsDefaultProfile) UnmarshalBinary(b []byte) error {
	var res SettingsDefaultProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsSavaMappingListItems0 settings sava mapping list items0
// swagger:model SettingsSavaMappingListItems0
type SettingsSavaMappingListItems0 struct {

	// auto sync period
	AutoSyncPeriod int64 `json:"autoSyncPeriod,omitempty"`

	// cco user
	CcoUser string `json:"ccoUser,omitempty"`

	// expiry
	Expiry int64 `json:"expiry,omitempty"`

	// last sync
	LastSync int64 `json:"lastSync,omitempty"`

	// profile
	Profile *SettingsSavaMappingListItems0Profile `json:"profile,omitempty"`

	// smart account Id
	SmartAccountID string `json:"smartAccountId,omitempty"`

	// sync result
	SyncResult *SettingsSavaMappingListItems0SyncResult `json:"syncResult,omitempty"`

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

// Validate validates this settings sava mapping list items0
func (m *SettingsSavaMappingListItems0) Validate(formats strfmt.Registry) error {
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

func (m *SettingsSavaMappingListItems0) validateProfile(formats strfmt.Registry) error {

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

func (m *SettingsSavaMappingListItems0) validateSyncResult(formats strfmt.Registry) error {

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

var settingsSavaMappingListItems0TypeSyncStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_SYNCED","SYNCING","SUCCESS","FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settingsSavaMappingListItems0TypeSyncStatusPropEnum = append(settingsSavaMappingListItems0TypeSyncStatusPropEnum, v)
	}
}

const (

	// SettingsSavaMappingListItems0SyncStatusNOTSYNCED captures enum value "NOT_SYNCED"
	SettingsSavaMappingListItems0SyncStatusNOTSYNCED string = "NOT_SYNCED"

	// SettingsSavaMappingListItems0SyncStatusSYNCING captures enum value "SYNCING"
	SettingsSavaMappingListItems0SyncStatusSYNCING string = "SYNCING"

	// SettingsSavaMappingListItems0SyncStatusSUCCESS captures enum value "SUCCESS"
	SettingsSavaMappingListItems0SyncStatusSUCCESS string = "SUCCESS"

	// SettingsSavaMappingListItems0SyncStatusFAILURE captures enum value "FAILURE"
	SettingsSavaMappingListItems0SyncStatusFAILURE string = "FAILURE"
)

// prop value enum
func (m *SettingsSavaMappingListItems0) validateSyncStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, settingsSavaMappingListItems0TypeSyncStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SettingsSavaMappingListItems0) validateSyncStatus(formats strfmt.Registry) error {

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
func (m *SettingsSavaMappingListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsSavaMappingListItems0) UnmarshalBinary(b []byte) error {
	var res SettingsSavaMappingListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsSavaMappingListItems0Profile settings sava mapping list items0 profile
// swagger:model SettingsSavaMappingListItems0Profile
type SettingsSavaMappingListItems0Profile struct {

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

// Validate validates this settings sava mapping list items0 profile
func (m *SettingsSavaMappingListItems0Profile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsSavaMappingListItems0Profile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsSavaMappingListItems0Profile) UnmarshalBinary(b []byte) error {
	var res SettingsSavaMappingListItems0Profile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsSavaMappingListItems0SyncResult settings sava mapping list items0 sync result
// swagger:model SettingsSavaMappingListItems0SyncResult
type SettingsSavaMappingListItems0SyncResult struct {

	// sync list
	SyncList []*SettingsSavaMappingListItems0SyncResultSyncListItems0 `json:"syncList"`

	// sync msg
	SyncMsg string `json:"syncMsg,omitempty"`
}

// Validate validates this settings sava mapping list items0 sync result
func (m *SettingsSavaMappingListItems0SyncResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsSavaMappingListItems0SyncResult) validateSyncList(formats strfmt.Registry) error {

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
func (m *SettingsSavaMappingListItems0SyncResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsSavaMappingListItems0SyncResult) UnmarshalBinary(b []byte) error {
	var res SettingsSavaMappingListItems0SyncResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsSavaMappingListItems0SyncResultSyncListItems0 settings sava mapping list items0 sync result sync list items0
// swagger:model SettingsSavaMappingListItems0SyncResultSyncListItems0
type SettingsSavaMappingListItems0SyncResultSyncListItems0 struct {

	// device sn list
	DeviceSnList []string `json:"deviceSnList"`

	// sync type
	// Enum: [Add Update Delete MismatchError]
	SyncType string `json:"syncType,omitempty"`
}

// Validate validates this settings sava mapping list items0 sync result sync list items0
func (m *SettingsSavaMappingListItems0SyncResultSyncListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var settingsSavaMappingListItems0SyncResultSyncListItems0TypeSyncTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Add","Update","Delete","MismatchError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settingsSavaMappingListItems0SyncResultSyncListItems0TypeSyncTypePropEnum = append(settingsSavaMappingListItems0SyncResultSyncListItems0TypeSyncTypePropEnum, v)
	}
}

const (

	// SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeAdd captures enum value "Add"
	SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeAdd string = "Add"

	// SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeUpdate captures enum value "Update"
	SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeUpdate string = "Update"

	// SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeDelete captures enum value "Delete"
	SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeDelete string = "Delete"

	// SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeMismatchError captures enum value "MismatchError"
	SettingsSavaMappingListItems0SyncResultSyncListItems0SyncTypeMismatchError string = "MismatchError"
)

// prop value enum
func (m *SettingsSavaMappingListItems0SyncResultSyncListItems0) validateSyncTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, settingsSavaMappingListItems0SyncResultSyncListItems0TypeSyncTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SettingsSavaMappingListItems0SyncResultSyncListItems0) validateSyncType(formats strfmt.Registry) error {

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
func (m *SettingsSavaMappingListItems0SyncResultSyncListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsSavaMappingListItems0SyncResultSyncListItems0) UnmarshalBinary(b []byte) error {
	var res SettingsSavaMappingListItems0SyncResultSyncListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SettingsTaskTimeOuts settings task time outs
// swagger:model SettingsTaskTimeOuts
type SettingsTaskTimeOuts struct {

	// config time out
	ConfigTimeOut int64 `json:"configTimeOut,omitempty"`

	// general time out
	GeneralTimeOut int64 `json:"generalTimeOut,omitempty"`

	// image download time out
	ImageDownloadTimeOut int64 `json:"imageDownloadTimeOut,omitempty"`
}

// Validate validates this settings task time outs
func (m *SettingsTaskTimeOuts) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsTaskTimeOuts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsTaskTimeOuts) UnmarshalBinary(b []byte) error {
	var res SettingsTaskTimeOuts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
