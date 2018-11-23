// Code generated by go-swagger; DO NOT EDIT.

package network_discovery

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

// NewUpdateNetconfCredentialsParams creates a new UpdateNetconfCredentialsParams object
// with the default values initialized.
func NewUpdateNetconfCredentialsParams() *UpdateNetconfCredentialsParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateNetconfCredentialsParams{
		ContentType: contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetconfCredentialsParamsWithTimeout creates a new UpdateNetconfCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetconfCredentialsParamsWithTimeout(timeout time.Duration) *UpdateNetconfCredentialsParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateNetconfCredentialsParams{
		ContentType: contentTypeDefault,

		timeout: timeout,
	}
}

// NewUpdateNetconfCredentialsParamsWithContext creates a new UpdateNetconfCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetconfCredentialsParamsWithContext(ctx context.Context) *UpdateNetconfCredentialsParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateNetconfCredentialsParams{
		ContentType: contentTypeDefault,

		Context: ctx,
	}
}

// NewUpdateNetconfCredentialsParamsWithHTTPClient creates a new UpdateNetconfCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetconfCredentialsParamsWithHTTPClient(client *http.Client) *UpdateNetconfCredentialsParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateNetconfCredentialsParams{
		ContentType: contentTypeDefault,
		HTTPClient:  client,
	}
}

/*UpdateNetconfCredentialsParams contains all the parameters to send to the API endpoint
for the update netconf credentials operation typically these are written to a http.Request
*/
type UpdateNetconfCredentialsParams struct {

	/*ContentType
	  Request body content type

	*/
	ContentType string
	/*Request
	  request

	*/
	Request *models.NetconfCredentialDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) WithTimeout(timeout time.Duration) *UpdateNetconfCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) WithContext(ctx context.Context) *UpdateNetconfCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) WithHTTPClient(client *http.Client) *UpdateNetconfCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) WithContentType(contentType string) *UpdateNetconfCredentialsParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithRequest adds the request to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) WithRequest(request *models.NetconfCredentialDTO) *UpdateNetconfCredentialsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update netconf credentials params
func (o *UpdateNetconfCredentialsParams) SetRequest(request *models.NetconfCredentialDTO) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetconfCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
