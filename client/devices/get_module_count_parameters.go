// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetModuleCountParams creates a new GetModuleCountParams object
// with the default values initialized.
func NewGetModuleCountParams() *GetModuleCountParams {
	var ()
	return &GetModuleCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetModuleCountParamsWithTimeout creates a new GetModuleCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetModuleCountParamsWithTimeout(timeout time.Duration) *GetModuleCountParams {
	var ()
	return &GetModuleCountParams{

		timeout: timeout,
	}
}

// NewGetModuleCountParamsWithContext creates a new GetModuleCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetModuleCountParamsWithContext(ctx context.Context) *GetModuleCountParams {
	var ()
	return &GetModuleCountParams{

		Context: ctx,
	}
}

// NewGetModuleCountParamsWithHTTPClient creates a new GetModuleCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetModuleCountParamsWithHTTPClient(client *http.Client) *GetModuleCountParams {
	var ()
	return &GetModuleCountParams{
		HTTPClient: client,
	}
}

/*GetModuleCountParams contains all the parameters to send to the API endpoint
for the get module count operation typically these are written to a http.Request
*/
type GetModuleCountParams struct {

	/*DeviceID
	  deviceId

	*/
	DeviceID string
	/*NameList
	  nameList

	*/
	NameList []string
	/*OperationalStateCodeList
	  operationalStateCodeList

	*/
	OperationalStateCodeList []string
	/*PartNumberList
	  partNumberList

	*/
	PartNumberList []string
	/*VendorEquipmentTypeList
	  vendorEquipmentTypeList

	*/
	VendorEquipmentTypeList []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get module count params
func (o *GetModuleCountParams) WithTimeout(timeout time.Duration) *GetModuleCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get module count params
func (o *GetModuleCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get module count params
func (o *GetModuleCountParams) WithContext(ctx context.Context) *GetModuleCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get module count params
func (o *GetModuleCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get module count params
func (o *GetModuleCountParams) WithHTTPClient(client *http.Client) *GetModuleCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get module count params
func (o *GetModuleCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get module count params
func (o *GetModuleCountParams) WithDeviceID(deviceID string) *GetModuleCountParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get module count params
func (o *GetModuleCountParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNameList adds the nameList to the get module count params
func (o *GetModuleCountParams) WithNameList(nameList []string) *GetModuleCountParams {
	o.SetNameList(nameList)
	return o
}

// SetNameList adds the nameList to the get module count params
func (o *GetModuleCountParams) SetNameList(nameList []string) {
	o.NameList = nameList
}

// WithOperationalStateCodeList adds the operationalStateCodeList to the get module count params
func (o *GetModuleCountParams) WithOperationalStateCodeList(operationalStateCodeList []string) *GetModuleCountParams {
	o.SetOperationalStateCodeList(operationalStateCodeList)
	return o
}

// SetOperationalStateCodeList adds the operationalStateCodeList to the get module count params
func (o *GetModuleCountParams) SetOperationalStateCodeList(operationalStateCodeList []string) {
	o.OperationalStateCodeList = operationalStateCodeList
}

// WithPartNumberList adds the partNumberList to the get module count params
func (o *GetModuleCountParams) WithPartNumberList(partNumberList []string) *GetModuleCountParams {
	o.SetPartNumberList(partNumberList)
	return o
}

// SetPartNumberList adds the partNumberList to the get module count params
func (o *GetModuleCountParams) SetPartNumberList(partNumberList []string) {
	o.PartNumberList = partNumberList
}

// WithVendorEquipmentTypeList adds the vendorEquipmentTypeList to the get module count params
func (o *GetModuleCountParams) WithVendorEquipmentTypeList(vendorEquipmentTypeList []string) *GetModuleCountParams {
	o.SetVendorEquipmentTypeList(vendorEquipmentTypeList)
	return o
}

// SetVendorEquipmentTypeList adds the vendorEquipmentTypeList to the get module count params
func (o *GetModuleCountParams) SetVendorEquipmentTypeList(vendorEquipmentTypeList []string) {
	o.VendorEquipmentTypeList = vendorEquipmentTypeList
}

// WriteToRequest writes these params to a swagger request
func (o *GetModuleCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param deviceId
	qrDeviceID := o.DeviceID
	qDeviceID := qrDeviceID
	if qDeviceID != "" {
		if err := r.SetQueryParam("deviceId", qDeviceID); err != nil {
			return err
		}
	}

	valuesNameList := o.NameList

	joinedNameList := swag.JoinByFormat(valuesNameList, "")
	// query array param nameList
	if err := r.SetQueryParam("nameList", joinedNameList...); err != nil {
		return err
	}

	valuesOperationalStateCodeList := o.OperationalStateCodeList

	joinedOperationalStateCodeList := swag.JoinByFormat(valuesOperationalStateCodeList, "")
	// query array param operationalStateCodeList
	if err := r.SetQueryParam("operationalStateCodeList", joinedOperationalStateCodeList...); err != nil {
		return err
	}

	valuesPartNumberList := o.PartNumberList

	joinedPartNumberList := swag.JoinByFormat(valuesPartNumberList, "")
	// query array param partNumberList
	if err := r.SetQueryParam("partNumberList", joinedPartNumberList...); err != nil {
		return err
	}

	valuesVendorEquipmentTypeList := o.VendorEquipmentTypeList

	joinedVendorEquipmentTypeList := swag.JoinByFormat(valuesVendorEquipmentTypeList, "")
	// query array param vendorEquipmentTypeList
	if err := r.SetQueryParam("vendorEquipmentTypeList", joinedVendorEquipmentTypeList...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
