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

// UpdatePnPGlobalSettingsResponse update pn p global settings response
// swagger:model UpdatePnPGlobalSettingsResponse
type UpdatePnPGlobalSettingsResponse struct {

	// id
	ID string `json:"_id,omitempty"`

	// aaa credentials
	AaaCredentials *UpdatePnPGlobalSettingsResponseAaaCredentials `json:"aaaCredentials,omitempty"`

	// accept eula
	AcceptEula bool `json:"acceptEula,omitempty"`

	// default profile
	DefaultProfile *UpdatePnPGlobalSettingsResponseDefaultProfile `json:"defaultProfile,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// sava mapping list
	SavaMappingList []*UpdatePnPGlobalSettingsResponseSavaMappingListItems0 `json:"savaMappingList"`

	// task time outs
	TaskTimeOuts *UpdatePnPGlobalSettingsResponseTaskTimeOuts `json:"taskTimeOuts,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`

	// version
	Version float64 `json:"version,omitempty"`
}

// Validate validates this update pn p global settings response
func (m *UpdatePnPGlobalSettingsResponse) Validate(formats strfmt.Registry) error {
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

func (m *UpdatePnPGlobalSettingsResponse) validateAaaCredentials(formats strfmt.Registry) error {

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

func (m *UpdatePnPGlobalSettingsResponse) validateDefaultProfile(formats strfmt.Registry) error {

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

func (m *UpdatePnPGlobalSettingsResponse) validateSavaMappingList(formats strfmt.Registry) error {

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

func (m *UpdatePnPGlobalSettingsResponse) validateTaskTimeOuts(formats strfmt.Registry) error {

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
func (m *UpdatePnPGlobalSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponse) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseAaaCredentials update pn p global settings response aaa credentials
// swagger:model UpdatePnPGlobalSettingsResponseAaaCredentials
type UpdatePnPGlobalSettingsResponseAaaCredentials struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this update pn p global settings response aaa credentials
func (m *UpdatePnPGlobalSettingsResponseAaaCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseAaaCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseAaaCredentials) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseAaaCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseDefaultProfile update pn p global settings response default profile
// swagger:model UpdatePnPGlobalSettingsResponseDefaultProfile
type UpdatePnPGlobalSettingsResponseDefaultProfile struct {

	// cert
	Cert string `json:"cert,omitempty"`

	// fqdn addresses
	FqdnAddresses []string `json:"fqdnAddresses"`

	// ip addresses
	IPAddresses []string `json:"ipAddresses"`

	// port
	Port float64 `json:"port,omitempty"`

	// proxy
	Proxy bool `json:"proxy,omitempty"`
}

// Validate validates this update pn p global settings response default profile
func (m *UpdatePnPGlobalSettingsResponseDefaultProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseDefaultProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseDefaultProfile) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseDefaultProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseSavaMappingListItems0 update pn p global settings response sava mapping list items0
// swagger:model UpdatePnPGlobalSettingsResponseSavaMappingListItems0
type UpdatePnPGlobalSettingsResponseSavaMappingListItems0 struct {

	// auto sync period
	AutoSyncPeriod float64 `json:"autoSyncPeriod,omitempty"`

	// cco user
	CcoUser string `json:"ccoUser,omitempty"`

	// expiry
	Expiry float64 `json:"expiry,omitempty"`

	// last sync
	LastSync float64 `json:"lastSync,omitempty"`

	// profile
	Profile *UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile `json:"profile,omitempty"`

	// smart account Id
	SmartAccountID string `json:"smartAccountId,omitempty"`

	// sync result
	SyncResult *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult `json:"syncResult,omitempty"`

	// sync result str
	SyncResultStr string `json:"syncResultStr,omitempty"`

	// sync start time
	SyncStartTime float64 `json:"syncStartTime,omitempty"`

	// sync status
	SyncStatus string `json:"syncStatus,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// virtual account Id
	VirtualAccountID string `json:"virtualAccountId,omitempty"`
}

// Validate validates this update pn p global settings response sava mapping list items0
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0) validateProfile(formats strfmt.Registry) error {

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

func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0) validateSyncResult(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseSavaMappingListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile update pn p global settings response sava mapping list items0 profile
// swagger:model UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile
type UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile struct {

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
	Port float64 `json:"port,omitempty"`

	// profile Id
	ProfileID string `json:"profileId,omitempty"`

	// proxy
	Proxy bool `json:"proxy,omitempty"`
}

// Validate validates this update pn p global settings response sava mapping list items0 profile
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseSavaMappingListItems0Profile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult update pn p global settings response sava mapping list items0 sync result
// swagger:model UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult
type UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult struct {

	// sync list
	SyncList []*UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0 `json:"syncList"`

	// sync msg
	SyncMsg string `json:"syncMsg,omitempty"`
}

// Validate validates this update pn p global settings response sava mapping list items0 sync result
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult) validateSyncList(formats strfmt.Registry) error {

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
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0 update pn p global settings response sava mapping list items0 sync result sync list items0
// swagger:model UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0
type UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0 struct {

	// device sn list
	DeviceSnList []string `json:"deviceSnList"`

	// sync type
	SyncType string `json:"syncType,omitempty"`
}

// Validate validates this update pn p global settings response sava mapping list items0 sync result sync list items0
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseSavaMappingListItems0SyncResultSyncListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdatePnPGlobalSettingsResponseTaskTimeOuts update pn p global settings response task time outs
// swagger:model UpdatePnPGlobalSettingsResponseTaskTimeOuts
type UpdatePnPGlobalSettingsResponseTaskTimeOuts struct {

	// config time out
	ConfigTimeOut float64 `json:"configTimeOut,omitempty"`

	// general time out
	GeneralTimeOut float64 `json:"generalTimeOut,omitempty"`

	// image download time out
	ImageDownloadTimeOut float64 `json:"imageDownloadTimeOut,omitempty"`
}

// Validate validates this update pn p global settings response task time outs
func (m *UpdatePnPGlobalSettingsResponseTaskTimeOuts) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseTaskTimeOuts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePnPGlobalSettingsResponseTaskTimeOuts) UnmarshalBinary(b []byte) error {
	var res UpdatePnPGlobalSettingsResponseTaskTimeOuts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
