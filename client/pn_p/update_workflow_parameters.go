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

// NewUpdateWorkflowParams creates a new UpdateWorkflowParams object
// with the default values initialized.
func NewUpdateWorkflowParams() *UpdateWorkflowParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateWorkflowParams{
		ContentType: contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWorkflowParamsWithTimeout creates a new UpdateWorkflowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWorkflowParamsWithTimeout(timeout time.Duration) *UpdateWorkflowParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateWorkflowParams{
		ContentType: contentTypeDefault,

		timeout: timeout,
	}
}

// NewUpdateWorkflowParamsWithContext creates a new UpdateWorkflowParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWorkflowParamsWithContext(ctx context.Context) *UpdateWorkflowParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateWorkflowParams{
		ContentType: contentTypeDefault,

		Context: ctx,
	}
}

// NewUpdateWorkflowParamsWithHTTPClient creates a new UpdateWorkflowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWorkflowParamsWithHTTPClient(client *http.Client) *UpdateWorkflowParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &UpdateWorkflowParams{
		ContentType: contentTypeDefault,
		HTTPClient:  client,
	}
}

/*UpdateWorkflowParams contains all the parameters to send to the API endpoint
for the update workflow operation typically these are written to a http.Request
*/
type UpdateWorkflowParams struct {

	/*ContentType
	  Request body content type

	*/
	ContentType string
	/*ID
	  id

	*/
	ID string
	/*Request
	  request

	*/
	Request *models.Workflow

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update workflow params
func (o *UpdateWorkflowParams) WithTimeout(timeout time.Duration) *UpdateWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update workflow params
func (o *UpdateWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update workflow params
func (o *UpdateWorkflowParams) WithContext(ctx context.Context) *UpdateWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update workflow params
func (o *UpdateWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update workflow params
func (o *UpdateWorkflowParams) WithHTTPClient(client *http.Client) *UpdateWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update workflow params
func (o *UpdateWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the update workflow params
func (o *UpdateWorkflowParams) WithContentType(contentType string) *UpdateWorkflowParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the update workflow params
func (o *UpdateWorkflowParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the update workflow params
func (o *UpdateWorkflowParams) WithID(id string) *UpdateWorkflowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update workflow params
func (o *UpdateWorkflowParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the update workflow params
func (o *UpdateWorkflowParams) WithRequest(request *models.Workflow) *UpdateWorkflowParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update workflow params
func (o *UpdateWorkflowParams) SetRequest(request *models.Workflow) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
