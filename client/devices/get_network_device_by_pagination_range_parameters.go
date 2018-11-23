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

// NewGetNetworkDeviceByPaginationRangeParams creates a new GetNetworkDeviceByPaginationRangeParams object
// with the default values initialized.
func NewGetNetworkDeviceByPaginationRangeParams() *GetNetworkDeviceByPaginationRangeParams {
	var ()
	return &GetNetworkDeviceByPaginationRangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkDeviceByPaginationRangeParamsWithTimeout creates a new GetNetworkDeviceByPaginationRangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkDeviceByPaginationRangeParamsWithTimeout(timeout time.Duration) *GetNetworkDeviceByPaginationRangeParams {
	var ()
	return &GetNetworkDeviceByPaginationRangeParams{

		timeout: timeout,
	}
}

// NewGetNetworkDeviceByPaginationRangeParamsWithContext creates a new GetNetworkDeviceByPaginationRangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkDeviceByPaginationRangeParamsWithContext(ctx context.Context) *GetNetworkDeviceByPaginationRangeParams {
	var ()
	return &GetNetworkDeviceByPaginationRangeParams{

		Context: ctx,
	}
}

// NewGetNetworkDeviceByPaginationRangeParamsWithHTTPClient creates a new GetNetworkDeviceByPaginationRangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkDeviceByPaginationRangeParamsWithHTTPClient(client *http.Client) *GetNetworkDeviceByPaginationRangeParams {
	var ()
	return &GetNetworkDeviceByPaginationRangeParams{
		HTTPClient: client,
	}
}

/*GetNetworkDeviceByPaginationRangeParams contains all the parameters to send to the API endpoint
for the get network device by pagination range operation typically these are written to a http.Request
*/
type GetNetworkDeviceByPaginationRangeParams struct {

	/*RecordsToReturn
	  Number of records to return

	*/
	RecordsToReturn int64
	/*StartIndex
	  Start index

	*/
	StartIndex int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) WithTimeout(timeout time.Duration) *GetNetworkDeviceByPaginationRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) WithContext(ctx context.Context) *GetNetworkDeviceByPaginationRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) WithHTTPClient(client *http.Client) *GetNetworkDeviceByPaginationRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecordsToReturn adds the recordsToReturn to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) WithRecordsToReturn(recordsToReturn int64) *GetNetworkDeviceByPaginationRangeParams {
	o.SetRecordsToReturn(recordsToReturn)
	return o
}

// SetRecordsToReturn adds the recordsToReturn to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) SetRecordsToReturn(recordsToReturn int64) {
	o.RecordsToReturn = recordsToReturn
}

// WithStartIndex adds the startIndex to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) WithStartIndex(startIndex int64) *GetNetworkDeviceByPaginationRangeParams {
	o.SetStartIndex(startIndex)
	return o
}

// SetStartIndex adds the startIndex to the get network device by pagination range params
func (o *GetNetworkDeviceByPaginationRangeParams) SetStartIndex(startIndex int64) {
	o.StartIndex = startIndex
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkDeviceByPaginationRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recordsToReturn
	if err := r.SetPathParam("recordsToReturn", swag.FormatInt64(o.RecordsToReturn)); err != nil {
		return err
	}

	// path param startIndex
	if err := r.SetPathParam("startIndex", swag.FormatInt64(o.StartIndex)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
