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

// InventoryRequest inventory request
// swagger:model InventoryRequest
type InventoryRequest struct {

	// cdp level
	CdpLevel int64 `json:"cdpLevel,omitempty"`

	// discovery type
	DiscoveryType string `json:"discoveryType,omitempty"`

	// enable password list
	EnablePasswordList []string `json:"enablePasswordList"`

	// global credential Id list
	GlobalCredentialIDList []string `json:"globalCredentialIdList"`

	// http read credential
	HTTPReadCredential *InventoryRequestHTTPReadCredential `json:"httpReadCredential,omitempty"`

	// http write credential
	HTTPWriteCredential *InventoryRequestHTTPWriteCredential `json:"httpWriteCredential,omitempty"`

	// ip address list
	IPAddressList string `json:"ipAddressList,omitempty"`

	// ip filter list
	IPFilterList []string `json:"ipFilterList"`

	// lldp level
	LldpLevel int64 `json:"lldpLevel,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// netconf port
	NetconfPort string `json:"netconfPort,omitempty"`

	// no add new device
	NoAddNewDevice bool `json:"noAddNewDevice,omitempty"`

	// parent discovery Id
	ParentDiscoveryID string `json:"parentDiscoveryId,omitempty"`

	// password list
	PasswordList []string `json:"passwordList"`

	// preferred mgmt IP method
	PreferredMgmtIPMethod string `json:"preferredMgmtIPMethod,omitempty"`

	// protocol order
	ProtocolOrder string `json:"protocolOrder,omitempty"`

	// re discovery
	ReDiscovery bool `json:"reDiscovery,omitempty"`

	// retry
	Retry int64 `json:"retry,omitempty"`

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

	// snmp r o community
	SnmpROCommunity string `json:"snmpROCommunity,omitempty"`

	// snmp r o community desc
	SnmpROCommunityDesc string `json:"snmpROCommunityDesc,omitempty"`

	// snmp r w community
	SnmpRWCommunity string `json:"snmpRWCommunity,omitempty"`

	// snmp r w community desc
	SnmpRWCommunityDesc string `json:"snmpRWCommunityDesc,omitempty"`

	// snmp user name
	SnmpUserName string `json:"snmpUserName,omitempty"`

	// snmp version
	SnmpVersion string `json:"snmpVersion,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`

	// update mgmt Ip
	UpdateMgmtIP bool `json:"updateMgmtIp,omitempty"`

	// user name list
	UserNameList []string `json:"userNameList"`
}

// Validate validates this inventory request
func (m *InventoryRequest) Validate(formats strfmt.Registry) error {
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

func (m *InventoryRequest) validateHTTPReadCredential(formats strfmt.Registry) error {

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

func (m *InventoryRequest) validateHTTPWriteCredential(formats strfmt.Registry) error {

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
func (m *InventoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryRequest) UnmarshalBinary(b []byte) error {
	var res InventoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryRequestHTTPReadCredential inventory request HTTP read credential
// swagger:model InventoryRequestHTTPReadCredential
type InventoryRequestHTTPReadCredential struct {

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

// Validate validates this inventory request HTTP read credential
func (m *InventoryRequestHTTPReadCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var inventoryRequestHttpReadCredentialTypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryRequestHttpReadCredentialTypeCredentialTypePropEnum = append(inventoryRequestHttpReadCredentialTypeCredentialTypePropEnum, v)
	}
}

const (

	// InventoryRequestHTTPReadCredentialCredentialTypeGLOBAL captures enum value "GLOBAL"
	InventoryRequestHTTPReadCredentialCredentialTypeGLOBAL string = "GLOBAL"

	// InventoryRequestHTTPReadCredentialCredentialTypeAPP captures enum value "APP"
	InventoryRequestHTTPReadCredentialCredentialTypeAPP string = "APP"
)

// prop value enum
func (m *InventoryRequestHTTPReadCredential) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, inventoryRequestHttpReadCredentialTypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InventoryRequestHTTPReadCredential) validateCredentialType(formats strfmt.Registry) error {

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
func (m *InventoryRequestHTTPReadCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryRequestHTTPReadCredential) UnmarshalBinary(b []byte) error {
	var res InventoryRequestHTTPReadCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryRequestHTTPWriteCredential inventory request HTTP write credential
// swagger:model InventoryRequestHTTPWriteCredential
type InventoryRequestHTTPWriteCredential struct {

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

// Validate validates this inventory request HTTP write credential
func (m *InventoryRequestHTTPWriteCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var inventoryRequestHttpWriteCredentialTypeCredentialTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GLOBAL","APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryRequestHttpWriteCredentialTypeCredentialTypePropEnum = append(inventoryRequestHttpWriteCredentialTypeCredentialTypePropEnum, v)
	}
}

const (

	// InventoryRequestHTTPWriteCredentialCredentialTypeGLOBAL captures enum value "GLOBAL"
	InventoryRequestHTTPWriteCredentialCredentialTypeGLOBAL string = "GLOBAL"

	// InventoryRequestHTTPWriteCredentialCredentialTypeAPP captures enum value "APP"
	InventoryRequestHTTPWriteCredentialCredentialTypeAPP string = "APP"
)

// prop value enum
func (m *InventoryRequestHTTPWriteCredential) validateCredentialTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, inventoryRequestHttpWriteCredentialTypeCredentialTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InventoryRequestHTTPWriteCredential) validateCredentialType(formats strfmt.Registry) error {

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
func (m *InventoryRequestHTTPWriteCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryRequestHTTPWriteCredential) UnmarshalBinary(b []byte) error {
	var res InventoryRequestHTTPWriteCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
