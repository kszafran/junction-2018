// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewAuthenticationAPIParams creates a new AuthenticationAPIParams object
// with the default values initialized.
func NewAuthenticationAPIParams() *AuthenticationAPIParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &AuthenticationAPIParams{
		ContentType: contentTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationAPIParamsWithTimeout creates a new AuthenticationAPIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthenticationAPIParamsWithTimeout(timeout time.Duration) *AuthenticationAPIParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &AuthenticationAPIParams{
		ContentType: contentTypeDefault,

		timeout: timeout,
	}
}

// NewAuthenticationAPIParamsWithContext creates a new AuthenticationAPIParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthenticationAPIParamsWithContext(ctx context.Context) *AuthenticationAPIParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &AuthenticationAPIParams{
		ContentType: contentTypeDefault,

		Context: ctx,
	}
}

// NewAuthenticationAPIParamsWithHTTPClient creates a new AuthenticationAPIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthenticationAPIParamsWithHTTPClient(client *http.Client) *AuthenticationAPIParams {
	var (
		contentTypeDefault = string("application/json")
	)
	return &AuthenticationAPIParams{
		ContentType: contentTypeDefault,
		HTTPClient:  client,
	}
}

/*AuthenticationAPIParams contains all the parameters to send to the API endpoint
for the authentication API operation typically these are written to a http.Request
*/
type AuthenticationAPIParams struct {

	/*Authorization
	  Basic Auth Base64 encoding of <username>:<password>

	*/
	Authorization string
	/*ContentType
	  Request body content type

	*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the authentication API params
func (o *AuthenticationAPIParams) WithTimeout(timeout time.Duration) *AuthenticationAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication API params
func (o *AuthenticationAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication API params
func (o *AuthenticationAPIParams) WithContext(ctx context.Context) *AuthenticationAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication API params
func (o *AuthenticationAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication API params
func (o *AuthenticationAPIParams) WithHTTPClient(client *http.Client) *AuthenticationAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication API params
func (o *AuthenticationAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the authentication API params
func (o *AuthenticationAPIParams) WithAuthorization(authorization string) *AuthenticationAPIParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the authentication API params
func (o *AuthenticationAPIParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithContentType adds the contentType to the authentication API params
func (o *AuthenticationAPIParams) WithContentType(contentType string) *AuthenticationAPIParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the authentication API params
func (o *AuthenticationAPIParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
