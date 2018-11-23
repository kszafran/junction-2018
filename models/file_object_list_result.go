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

// FileObjectListResult file object list result
// swagger:model FileObjectListResult
type FileObjectListResult struct {

	// response
	Response []*FileObjectListResultResponseItems0 `json:"response"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this file object list result
func (m *FileObjectListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileObjectListResult) validateResponse(formats strfmt.Registry) error {

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
func (m *FileObjectListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileObjectListResult) UnmarshalBinary(b []byte) error {
	var res FileObjectListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileObjectListResultResponseItems0 file object list result response items0
// swagger:model FileObjectListResultResponseItems0
type FileObjectListResultResponseItems0 struct {

	// attribute info
	AttributeInfo interface{} `json:"attributeInfo,omitempty"`

	// download path
	DownloadPath string `json:"downloadPath,omitempty"`

	// encrypted
	Encrypted bool `json:"encrypted,omitempty"`

	// file format
	FileFormat string `json:"fileFormat,omitempty"`

	// file size
	FileSize string `json:"fileSize,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// md5 checksum
	Md5Checksum string `json:"md5Checksum,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// name space
	NameSpace string `json:"nameSpace,omitempty"`

	// sftp server list
	SftpServerList []interface{} `json:"sftpServerList"`

	// sha1 checksum
	Sha1Checksum string `json:"sha1Checksum,omitempty"`

	// task Id
	TaskID interface{} `json:"taskId,omitempty"`
}

// Validate validates this file object list result response items0
func (m *FileObjectListResultResponseItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileObjectListResultResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileObjectListResultResponseItems0) UnmarshalBinary(b []byte) error {
	var res FileObjectListResultResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
