// Code generated by go-swagger; DO NOT EDIT.

package pn_p

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

	models "github.com/kszafran/junction-2018/models"
)

// NewUpdatePnPServerProfileParams creates a new UpdatePnPServerProfileParams object
// with the default values initialized.
func NewUpdatePnPServerProfileParams() *UpdatePnPServerProfileParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdatePnPServerProfileParams{
		ContentType: contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePnPServerProfileParamsWithTimeout creates a new UpdatePnPServerProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePnPServerProfileParamsWithTimeout(timeout time.Duration) *UpdatePnPServerProfileParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdatePnPServerProfileParams{
		ContentType: contentTypeDefault,

		timeout: timeout,
	}
}

// NewUpdatePnPServerProfileParamsWithContext creates a new UpdatePnPServerProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePnPServerProfileParamsWithContext(ctx context.Context) *UpdatePnPServerProfileParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdatePnPServerProfileParams{
		ContentType: contentTypeDefault,

		Context: ctx,
	}
}

// NewUpdatePnPServerProfileParamsWithHTTPClient creates a new UpdatePnPServerProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePnPServerProfileParamsWithHTTPClient(client *http.Client) *UpdatePnPServerProfileParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdatePnPServerProfileParams{
		ContentType: contentTypeDefault,
		HTTPClient:  client,
	}
}

/*UpdatePnPServerProfileParams contains all the parameters to send to the API endpoint
for the update pn p server profile operation typically these are written to a http.Request
*/
type UpdatePnPServerProfileParams struct {

	/*ContentType
	  Request body content type

	*/
	ContentType string
	/*Request
	  request

	*/
	Request *models.SAVAMapping

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) WithTimeout(timeout time.Duration) *UpdatePnPServerProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) WithContext(ctx context.Context) *UpdatePnPServerProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) WithHTTPClient(client *http.Client) *UpdatePnPServerProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) WithContentType(contentType string) *UpdatePnPServerProfileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithRequest adds the request to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) WithRequest(request *models.SAVAMapping) *UpdatePnPServerProfileParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update pn p server profile params
func (o *UpdatePnPServerProfileParams) SetRequest(request *models.SAVAMapping) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePnPServerProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
