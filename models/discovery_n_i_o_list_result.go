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

// DiscoveryNIOListResult discovery n i o list result
// swagger:model DiscoveryNIOListResult
type DiscoveryNIOListResult struct {

	// response
	Response []*DiscoveryNIOListResultResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this discovery n i o list result
func (m *DiscoveryNIOListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscoveryNIOListResult) validateResponse(formats strfmt.Registry) error {

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
func (m *DiscoveryNIOListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryNIOListResult) UnmarshalBinary(b []byte) error {
	var res DiscoveryNIOListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DiscoveryNIOListResultResponseItems0 discovery n i o list result response items0
// swagger:model DiscoveryNIOListResultResponseItems0
type DiscoveryNIOListResultResponseItems0 struct {

	// attribute info
	AttributeInfo interface{} `json:"attributeInfo,omitempty"`

	// cdp level
	CdpLevel int64 `json:"cdpLevel,omitempty"`

	// device ids
	DeviceIds string `json:"deviceIds,omitempty"`

	// discovery condition
	DiscoveryCondition string `json:"discoveryCondition,omitempty"`

	// discovery status
	DiscoveryStatus string `json:"discoveryStatus,omitempty"`

	// discovery type
	DiscoveryType string `json:"discoveryType,omitempty"`

	// enable password list
	EnablePasswordList string `json:"enablePasswordList,omitempty"`

	// global credential Id list
	GlobalCredentialIDList []string `json:"globalCredentialIdList"`

	// http read credential
	HTTPReadCredential *DiscoveryNIOListResultResponseItems0HTTPReadCredential `json:"httpReadCredential,omitempty"`

	// http write credential
	HTTPWriteCredential *DiscoveryNIOListResultResponseItems0HTTPWriteCredential `json:"httpWriteCredential,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ip address list
	IPAddressList string `json:"ipAddressList,omitempty"`

	// ip filter list
	IPFilterList string `json:"ipFilterList,omitempty"`

	// is auto cdp
	IsAutoCdp bool `json:"isAutoCdp,omitempty"`

	// lldp level
	LldpLevel int64 `json:"lldpLevel,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// netconf port
	NetconfPort string `json:"netconfPort,omitempty"`

	// num devices
	NumDevices int64 `json:"numDevices,omitempty"`

	// parent discovery Id
	ParentDiscoveryID string `json:"parentDiscoveryId,omitempty"`

	// password list
	PasswordList string `json:"passwordList,omitempty"`

	// preferred mgmt IP method
	PreferredMgmtIPMethod string `json:"preferredMgmtIPMethod,omitempty"`

	// protocol order
	ProtocolOrder string `json:"protocolOrder,omitempty"`

	// retry count
	RetryCount int64 `json:"retryCount,omitempty"`

	// snmp auth passphrase
	SnmpAuthPassphrase string `json:"snmpAuthPassphrase,omitempty"`

	// snmp auth protocol
	SnmpAuthProtocol string `json:"snmpAuthProtocol,omitempty"`

	// snmp mode
	SnmpMode string `json:"snmpMode,omitempty"`

	// snmp priv passphrase
	SnmpPrivPassphrase string `json:"snmpPrivPassphrase,omitempty"`

	// snmp priv protocol
	SnmpPrivProtocol string `json:"snmpPrivProtocol,omitempty"`

	// snmp ro community
	SnmpRoCommunity string `json:"snmpRoCommunity,omitempty"`

	// snmp ro community desc
	SnmpRoCommunityDesc string `json:"snmpRoCommunityDesc,omitempty"`

	// snmp rw community
	SnmpRwCommunity string `json:"snmpRwCommunity,omitempty"`

	// snmp rw community desc
	SnmpRwCommunityDesc string `json:"snmpRwCommunityDesc,omitempty"`

	// snmp user name
	SnmpUserName string `json:"snmpUserName,omitempty"`

	// time out
	TimeOut int64 `json:"timeOut,omitempty"`

	// update mgmt Ip
	UpdateMgmtIP bool `json:"updateMgmtIp,omitempty"`

	// user name list
	UserNameList string `json:"userNameList,omitempty"`
}

// Validate validates this discovery n i o list result response items0
func (m *DiscoveryNIOListResultResponseItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPReadCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPWriteCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscoveryNIOListResultResponseItems0) validateHTTPReadCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPReadCredential) { // not required
		return nil
	}

	if m.HTTPReadCredential != nil {
		if err := m.HTTPReadCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpReadCredential")
			}
			return err
		}
	}

	return nil
}

func (m *DiscoveryNIOListResultResponseItems0) validateHTTPWriteCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPWriteCredential) { // not required
		return nil
	}

	if m.HTTPWriteCredential != nil {
		if err := m.HTTPWriteCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpWriteCredential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveryNIOListResultResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryNIOListResultResponseItems0) UnmarshalBinary(b []byte) error {
	var res DiscoveryNIOListResultResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DiscoveryNIOListResultResponseItems0HTTPReadCredential discovery n i o list result response items0 HTTP read credential
// swagger:model DiscoveryNIOListResultResponseItems0HTTPReadCredential
type DiscoveryNIOListResultResponseItems0HTTPReadCredential struct {

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

	// password
	Password string `json:"password,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// secure
	Secure bool `json:"secure,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this discovery n i o list result response items0 HTTP read credential
func (m *DiscoveryNIOListResultResponseItems0HTTPReadCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var discoveryNIOListResultResponseItems0HttpReadCredentialTypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discoveryNIOListResultResponseItems0HttpReadCredentialTypeCredentialTypePropEnum = append(discoveryNIOListResultResponseItems0HttpReadCredentialTypeCredentialTypePropEnum, v)
	}
}

const (

	// DiscoveryNIOListResultResponseItems0HTTPReadCredentialCredentialTypeGLOBAL captures enum value "GLOBAL"
	DiscoveryNIOListResultResponseItems0HTTPReadCredentialCredentialTypeGLOBAL string = "GLOBAL"

	// DiscoveryNIOListResultResponseItems0HTTPReadCredentialCredentialTypeAPP captures enum value "APP"
	DiscoveryNIOListResultResponseItems0HTTPReadCredentialCredentialTypeAPP string = "APP"
)

// prop value enum
func (m *DiscoveryNIOListResultResponseItems0HTTPReadCredential) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, discoveryNIOListResultResponseItems0HttpReadCredentialTypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DiscoveryNIOListResultResponseItems0HTTPReadCredential) validateCredentialType(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCredentialTypeEnum("httpReadCredential"+"."+"credentialType", "body", m.CredentialType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveryNIOListResultResponseItems0HTTPReadCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryNIOListResultResponseItems0HTTPReadCredential) UnmarshalBinary(b []byte) error {
	var res DiscoveryNIOListResultResponseItems0HTTPReadCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DiscoveryNIOListResultResponseItems0HTTPWriteCredential discovery n i o list result response items0 HTTP write credential
// swagger:model DiscoveryNIOListResultResponseItems0HTTPWriteCredential
type DiscoveryNIOListResultResponseItems0HTTPWriteCredential struct {

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

	// password
	Password string `json:"password,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// secure
	Secure bool `json:"secure,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this discovery n i o list result response items0 HTTP write credential
func (m *DiscoveryNIOListResultResponseItems0HTTPWriteCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var discoveryNIOListResultResponseItems0HttpWriteCredentialTypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discoveryNIOListResultResponseItems0HttpWriteCredentialTypeCredentialTypePropEnum = append(discoveryNIOListResultResponseItems0HttpWriteCredentialTypeCredentialTypePropEnum, v)
	}
}

const (

	// DiscoveryNIOListResultResponseItems0HTTPWriteCredentialCredentialTypeGLOBAL captures enum value "GLOBAL"
	DiscoveryNIOListResultResponseItems0HTTPWriteCredentialCredentialTypeGLOBAL string = "GLOBAL"

	// DiscoveryNIOListResultResponseItems0HTTPWriteCredentialCredentialTypeAPP captures enum value "APP"
	DiscoveryNIOListResultResponseItems0HTTPWriteCredentialCredentialTypeAPP string = "APP"
)

// prop value enum
func (m *DiscoveryNIOListResultResponseItems0HTTPWriteCredential) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, discoveryNIOListResultResponseItems0HttpWriteCredentialTypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DiscoveryNIOListResultResponseItems0HTTPWriteCredential) validateCredentialType(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCredentialTypeEnum("httpWriteCredential"+"."+"credentialType", "body", m.CredentialType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveryNIOListResultResponseItems0HTTPWriteCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveryNIOListResultResponseItems0HTTPWriteCredential) UnmarshalBinary(b []byte) error {
	var res DiscoveryNIOListResultResponseItems0HTTPWriteCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}