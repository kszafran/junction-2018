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

// InventoryDeviceInfo inventory device info
// swagger:model InventoryDeviceInfo
type InventoryDeviceInfo struct {

	// cli transport
	CliTransport string `json:"cliTransport,omitempty"`

	// compute device
	ComputeDevice bool `json:"computeDevice,omitempty"`

	// enable password
	EnablePassword string `json:"enablePassword,omitempty"`

	// extended discovery info
	ExtendedDiscoveryInfo string `json:"extendedDiscoveryInfo,omitempty"`

	// http password
	HTTPPassword string `json:"httpPassword,omitempty"`

	// http port
	HTTPPort string `json:"httpPort,omitempty"`

	// http secure
	HTTPSecure bool `json:"httpSecure,omitempty"`

	// http user name
	HTTPUserName string `json:"httpUserName,omitempty"`

	// ip address
	IPAddress []string `json:"ipAddress"`

	// meraki org Id
	MerakiOrgID []string `json:"merakiOrgId"`

	// netconf port
	NetconfPort string `json:"netconfPort,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

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

	// snmp r w community
	SnmpRWCommunity string `json:"snmpRWCommunity,omitempty"`

	// snmp retry
	SnmpRetry int64 `json:"snmpRetry,omitempty"`

	// snmp timeout
	SnmpTimeout int64 `json:"snmpTimeout,omitempty"`

	// snmp user name
	SnmpUserName string `json:"snmpUserName,omitempty"`

	// snmp version
	SnmpVersion string `json:"snmpVersion,omitempty"`

	// type
	// Enum: [COMPUTE_DEVICE MERAKI_DASHBOARD NETWORK_DEVICE NODATACHANGE]
	Type string `json:"type,omitempty"`

	// update mgmt ipaddress list
	UpdateMgmtIpaddressList []*InventoryDeviceInfoUpdateMgmtIpaddressListItems0 `json:"updateMgmtIPaddressList"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this inventory device info
func (m *InventoryDeviceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateMgmtIpaddressList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var inventoryDeviceInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COMPUTE_DEVICE","MERAKI_DASHBOARD","NETWORK_DEVICE","NODATACHANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryDeviceInfoTypeTypePropEnum = append(inventoryDeviceInfoTypeTypePropEnum, v)
	}
}

const (

	// InventoryDeviceInfoTypeCOMPUTEDEVICE captures enum value "COMPUTE_DEVICE"
	InventoryDeviceInfoTypeCOMPUTEDEVICE string = "COMPUTE_DEVICE"

	// InventoryDeviceInfoTypeMERAKIDASHBOARD captures enum value "MERAKI_DASHBOARD"
	InventoryDeviceInfoTypeMERAKIDASHBOARD string = "MERAKI_DASHBOARD"

	// InventoryDeviceInfoTypeNETWORKDEVICE captures enum value "NETWORK_DEVICE"
	InventoryDeviceInfoTypeNETWORKDEVICE string = "NETWORK_DEVICE"

	// InventoryDeviceInfoTypeNODATACHANGE captures enum value "NODATACHANGE"
	InventoryDeviceInfoTypeNODATACHANGE string = "NODATACHANGE"
)

// prop value enum
func (m *InventoryDeviceInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, inventoryDeviceInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InventoryDeviceInfo) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InventoryDeviceInfo) validateUpdateMgmtIpaddressList(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateMgmtIpaddressList) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateMgmtIpaddressList); i++ {
		if swag.IsZero(m.UpdateMgmtIpaddressList[i]) { // not required
			continue
		}

		if m.UpdateMgmtIpaddressList[i] != nil {
			if err := m.UpdateMgmtIpaddressList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateMgmtIPaddressList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryDeviceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryDeviceInfo) UnmarshalBinary(b []byte) error {
	var res InventoryDeviceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InventoryDeviceInfoUpdateMgmtIpaddressListItems0 inventory device info update mgmt ipaddress list items0
// swagger:model InventoryDeviceInfoUpdateMgmtIpaddressListItems0
type InventoryDeviceInfoUpdateMgmtIpaddressListItems0 struct {

	// exist mgmt Ip address
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"`

	// new mgmt Ip address
	NewMgmtIPAddress string `json:"newMgmtIpAddress,omitempty"`
}

// Validate validates this inventory device info update mgmt ipaddress list items0
func (m *InventoryDeviceInfoUpdateMgmtIpaddressListItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryDeviceInfoUpdateMgmtIpaddressListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryDeviceInfoUpdateMgmtIpaddressListItems0) UnmarshalBinary(b []byte) error {
	var res InventoryDeviceInfoUpdateMgmtIpaddressListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
