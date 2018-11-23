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

// NetworkDeviceNIOListResult network device n i o list result
// swagger:model NetworkDeviceNIOListResult
type NetworkDeviceNIOListResult struct {

	// response
	Response []*NetworkDeviceNIOListResultResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this network device n i o list result
func (m *NetworkDeviceNIOListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkDeviceNIOListResult) validateResponse(formats strfmt.Registry) error {

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
func (m *NetworkDeviceNIOListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDeviceNIOListResult) UnmarshalBinary(b []byte) error {
	var res NetworkDeviceNIOListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkDeviceNIOListResultResponseItems0 network device n i o list result response items0
// swagger:model NetworkDeviceNIOListResultResponseItems0
type NetworkDeviceNIOListResultResponseItems0 struct {

	// anchor wlc for ap
	AnchorWlcForAp string `json:"anchorWlcForAp,omitempty"`

	// auth model Id
	AuthModelID string `json:"authModelId,omitempty"`

	// avg update frequency
	AvgUpdateFrequency int64 `json:"avgUpdateFrequency,omitempty"`

	// boot date time
	BootDateTime string `json:"bootDateTime,omitempty"`

	// cli status
	CliStatus string `json:"cliStatus,omitempty"`

	// duplicate device Id
	DuplicateDeviceID string `json:"duplicateDeviceId,omitempty"`

	// error code
	ErrorCode string `json:"errorCode,omitempty"`

	// error description
	ErrorDescription string `json:"errorDescription,omitempty"`

	// family
	Family string `json:"family,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// http status
	HTTPStatus string `json:"httpStatus,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image name
	ImageName string `json:"imageName,omitempty"`

	// ingress queue config
	IngressQueueConfig string `json:"ingressQueueConfig,omitempty"`

	// interface count
	InterfaceCount string `json:"interfaceCount,omitempty"`

	// inventory collection status
	InventoryCollectionStatus string `json:"inventoryCollectionStatus,omitempty"`

	// inventory reachability status
	InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"`

	// last updated
	LastUpdated string `json:"lastUpdated,omitempty"`

	// line card count
	LineCardCount string `json:"lineCardCount,omitempty"`

	// line card Id
	LineCardID string `json:"lineCardId,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// location name
	LocationName string `json:"locationName,omitempty"`

	// mac address
	MacAddress string `json:"macAddress,omitempty"`

	// management Ip address
	ManagementIPAddress string `json:"managementIpAddress,omitempty"`

	// memory size
	MemorySize string `json:"memorySize,omitempty"`

	// netconf status
	NetconfStatus string `json:"netconfStatus,omitempty"`

	// num updates
	NumUpdates int64 `json:"numUpdates,omitempty"`

	// ping status
	PingStatus string `json:"pingStatus,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// port range
	PortRange string `json:"portRange,omitempty"`

	// qos status
	QosStatus string `json:"qosStatus,omitempty"`

	// reachability failure reason
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"`

	// reachability status
	ReachabilityStatus string `json:"reachabilityStatus,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// role source
	RoleSource string `json:"roleSource,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

	// snmp contact
	SnmpContact string `json:"snmpContact,omitempty"`

	// snmp location
	SnmpLocation string `json:"snmpLocation,omitempty"`

	// snmp status
	SnmpStatus string `json:"snmpStatus,omitempty"`

	// software version
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// tag count
	TagCount int64 `json:"tagCount,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// up time
	UpTime string `json:"upTime,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`

	// wlc ap device status
	WlcApDeviceStatus string `json:"wlcApDeviceStatus,omitempty"`
}

// Validate validates this network device n i o list result response items0
func (m *NetworkDeviceNIOListResultResponseItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkDeviceNIOListResultResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDeviceNIOListResultResponseItems0) UnmarshalBinary(b []byte) error {
	var res NetworkDeviceNIOListResultResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
