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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeviceConfigByIDParams creates a new GetDeviceConfigByIDParams object
// with the default values initialized.
func NewGetDeviceConfigByIDParams() *GetDeviceConfigByIDParams {
	var ()
	return &GetDeviceConfigByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceConfigByIDParamsWithTimeout creates a new GetDeviceConfigByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceConfigByIDParamsWithTimeout(timeout time.Duration) *GetDeviceConfigByIDParams {
	var ()
	return &GetDeviceConfigByIDParams{

		timeout: timeout,
	}
}

// NewGetDeviceConfigByIDParamsWithContext creates a new GetDeviceConfigByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceConfigByIDParamsWithContext(ctx context.Context) *GetDeviceConfigByIDParams {
	var ()
	return &GetDeviceConfigByIDParams{

		Context: ctx,
	}
}

// NewGetDeviceConfigByIDParamsWithHTTPClient creates a new GetDeviceConfigByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceConfigByIDParamsWithHTTPClient(client *http.Client) *GetDeviceConfigByIDParams {
	var ()
	return &GetDeviceConfigByIDParams{
		HTTPClient: client,
	}
}

/*GetDeviceConfigByIDParams contains all the parameters to send to the API endpoint
for the get device config by Id operation typically these are written to a http.Request
*/
type GetDeviceConfigByIDParams struct {

	/*NetworkDeviceID
	  networkDeviceId

	*/
	NetworkDeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device config by Id params
func (o *GetDeviceConfigByIDParams) WithTimeout(timeout time.Duration) *GetDeviceConfigByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device config by Id params
func (o *GetDeviceConfigByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device config by Id params
func (o *GetDeviceConfigByIDParams) WithContext(ctx context.Context) *GetDeviceConfigByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device config by Id params
func (o *GetDeviceConfigByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device config by Id params
func (o *GetDeviceConfigByIDParams) WithHTTPClient(client *http.Client) *GetDeviceConfigByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device config by Id params
func (o *GetDeviceConfigByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkDeviceID adds the networkDeviceID to the get device config by Id params
func (o *GetDeviceConfigByIDParams) WithNetworkDeviceID(networkDeviceID string) *GetDeviceConfigByIDParams {
	o.SetNetworkDeviceID(networkDeviceID)
	return o
}

// SetNetworkDeviceID adds the networkDeviceId to the get device config by Id params
func (o *GetDeviceConfigByIDParams) SetNetworkDeviceID(networkDeviceID string) {
	o.NetworkDeviceID = networkDeviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceConfigByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkDeviceId
	if err := r.SetPathParam("networkDeviceId", o.NetworkDeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
