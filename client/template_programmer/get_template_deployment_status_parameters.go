// Code generated by go-swagger; DO NOT EDIT.

package template_programmer

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

// NewGetTemplateDeploymentStatusParams creates a new GetTemplateDeploymentStatusParams object
// with the default values initialized.
func NewGetTemplateDeploymentStatusParams() *GetTemplateDeploymentStatusParams {
	var ()
	return &GetTemplateDeploymentStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplateDeploymentStatusParamsWithTimeout creates a new GetTemplateDeploymentStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTemplateDeploymentStatusParamsWithTimeout(timeout time.Duration) *GetTemplateDeploymentStatusParams {
	var ()
	return &GetTemplateDeploymentStatusParams{

		timeout: timeout,
	}
}

// NewGetTemplateDeploymentStatusParamsWithContext creates a new GetTemplateDeploymentStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTemplateDeploymentStatusParamsWithContext(ctx context.Context) *GetTemplateDeploymentStatusParams {
	var ()
	return &GetTemplateDeploymentStatusParams{

		Context: ctx,
	}
}

// NewGetTemplateDeploymentStatusParamsWithHTTPClient creates a new GetTemplateDeploymentStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTemplateDeploymentStatusParamsWithHTTPClient(client *http.Client) *GetTemplateDeploymentStatusParams {
	var ()
	return &GetTemplateDeploymentStatusParams{
		HTTPClient: client,
	}
}

/*GetTemplateDeploymentStatusParams contains all the parameters to send to the API endpoint
for the get template deployment status operation typically these are written to a http.Request
*/
type GetTemplateDeploymentStatusParams struct {

	/*DeploymentID
	  deploymentId

	*/
	DeploymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) WithTimeout(timeout time.Duration) *GetTemplateDeploymentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) WithContext(ctx context.Context) *GetTemplateDeploymentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) WithHTTPClient(client *http.Client) *GetTemplateDeploymentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) WithDeploymentID(deploymentID string) *GetTemplateDeploymentStatusParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get template deployment status params
func (o *GetTemplateDeploymentStatusParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplateDeploymentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploymentId
	if err := r.SetPathParam("deploymentId", o.DeploymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
