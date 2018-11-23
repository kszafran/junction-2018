// Code generated by go-swagger; DO NOT EDIT.

package network_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// UpdateSNMPv3CredentialsReader is a Reader for the UpdateSNMPv3Credentials structure.
type UpdateSNMPv3CredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSNMPv3CredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSNMPv3CredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewUpdateSNMPv3CredentialsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewUpdateSNMPv3CredentialsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewUpdateSNMPv3CredentialsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewUpdateSNMPv3CredentialsPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSNMPv3CredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateSNMPv3CredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateSNMPv3CredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSNMPv3CredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateSNMPv3CredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewUpdateSNMPv3CredentialsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateSNMPv3CredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewUpdateSNMPv3CredentialsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewUpdateSNMPv3CredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewUpdateSNMPv3CredentialsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSNMPv3CredentialsOK creates a UpdateSNMPv3CredentialsOK with default headers values
func NewUpdateSNMPv3CredentialsOK() *UpdateSNMPv3CredentialsOK {
	return &UpdateSNMPv3CredentialsOK{}
}

/*UpdateSNMPv3CredentialsOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type UpdateSNMPv3CredentialsOK struct {
	Payload *models.TaskIDResult
}

func (o *UpdateSNMPv3CredentialsOK) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsOK  %+v", 200, o.Payload)
}

func (o *UpdateSNMPv3CredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskIDResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSNMPv3CredentialsCreated creates a UpdateSNMPv3CredentialsCreated with default headers values
func NewUpdateSNMPv3CredentialsCreated() *UpdateSNMPv3CredentialsCreated {
	return &UpdateSNMPv3CredentialsCreated{}
}

/*UpdateSNMPv3CredentialsCreated handles this case with default header values.

The POST/PUT request was fulfilled and a new resource has been created. Information about the resource is in the response body.
*/
type UpdateSNMPv3CredentialsCreated struct {
}

func (o *UpdateSNMPv3CredentialsCreated) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsCreated ", 201)
}

func (o *UpdateSNMPv3CredentialsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsAccepted creates a UpdateSNMPv3CredentialsAccepted with default headers values
func NewUpdateSNMPv3CredentialsAccepted() *UpdateSNMPv3CredentialsAccepted {
	return &UpdateSNMPv3CredentialsAccepted{}
}

/*UpdateSNMPv3CredentialsAccepted handles this case with default header values.

The request was accepted for processing, but the processing has not been completed.
*/
type UpdateSNMPv3CredentialsAccepted struct {
}

func (o *UpdateSNMPv3CredentialsAccepted) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsAccepted ", 202)
}

func (o *UpdateSNMPv3CredentialsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsNoContent creates a UpdateSNMPv3CredentialsNoContent with default headers values
func NewUpdateSNMPv3CredentialsNoContent() *UpdateSNMPv3CredentialsNoContent {
	return &UpdateSNMPv3CredentialsNoContent{}
}

/*UpdateSNMPv3CredentialsNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type UpdateSNMPv3CredentialsNoContent struct {
}

func (o *UpdateSNMPv3CredentialsNoContent) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsNoContent ", 204)
}

func (o *UpdateSNMPv3CredentialsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsPartialContent creates a UpdateSNMPv3CredentialsPartialContent with default headers values
func NewUpdateSNMPv3CredentialsPartialContent() *UpdateSNMPv3CredentialsPartialContent {
	return &UpdateSNMPv3CredentialsPartialContent{}
}

/*UpdateSNMPv3CredentialsPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type UpdateSNMPv3CredentialsPartialContent struct {
}

func (o *UpdateSNMPv3CredentialsPartialContent) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsPartialContent ", 206)
}

func (o *UpdateSNMPv3CredentialsPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsBadRequest creates a UpdateSNMPv3CredentialsBadRequest with default headers values
func NewUpdateSNMPv3CredentialsBadRequest() *UpdateSNMPv3CredentialsBadRequest {
	return &UpdateSNMPv3CredentialsBadRequest{}
}

/*UpdateSNMPv3CredentialsBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type UpdateSNMPv3CredentialsBadRequest struct {
}

func (o *UpdateSNMPv3CredentialsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsBadRequest ", 400)
}

func (o *UpdateSNMPv3CredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsUnauthorized creates a UpdateSNMPv3CredentialsUnauthorized with default headers values
func NewUpdateSNMPv3CredentialsUnauthorized() *UpdateSNMPv3CredentialsUnauthorized {
	return &UpdateSNMPv3CredentialsUnauthorized{}
}

/*UpdateSNMPv3CredentialsUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type UpdateSNMPv3CredentialsUnauthorized struct {
}

func (o *UpdateSNMPv3CredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsUnauthorized ", 401)
}

func (o *UpdateSNMPv3CredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsForbidden creates a UpdateSNMPv3CredentialsForbidden with default headers values
func NewUpdateSNMPv3CredentialsForbidden() *UpdateSNMPv3CredentialsForbidden {
	return &UpdateSNMPv3CredentialsForbidden{}
}

/*UpdateSNMPv3CredentialsForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type UpdateSNMPv3CredentialsForbidden struct {
}

func (o *UpdateSNMPv3CredentialsForbidden) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsForbidden ", 403)
}

func (o *UpdateSNMPv3CredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsNotFound creates a UpdateSNMPv3CredentialsNotFound with default headers values
func NewUpdateSNMPv3CredentialsNotFound() *UpdateSNMPv3CredentialsNotFound {
	return &UpdateSNMPv3CredentialsNotFound{}
}

/*UpdateSNMPv3CredentialsNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type UpdateSNMPv3CredentialsNotFound struct {
}

func (o *UpdateSNMPv3CredentialsNotFound) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsNotFound ", 404)
}

func (o *UpdateSNMPv3CredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsConflict creates a UpdateSNMPv3CredentialsConflict with default headers values
func NewUpdateSNMPv3CredentialsConflict() *UpdateSNMPv3CredentialsConflict {
	return &UpdateSNMPv3CredentialsConflict{}
}

/*UpdateSNMPv3CredentialsConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type UpdateSNMPv3CredentialsConflict struct {
}

func (o *UpdateSNMPv3CredentialsConflict) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsConflict ", 409)
}

func (o *UpdateSNMPv3CredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsUnsupportedMediaType creates a UpdateSNMPv3CredentialsUnsupportedMediaType with default headers values
func NewUpdateSNMPv3CredentialsUnsupportedMediaType() *UpdateSNMPv3CredentialsUnsupportedMediaType {
	return &UpdateSNMPv3CredentialsUnsupportedMediaType{}
}

/*UpdateSNMPv3CredentialsUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type UpdateSNMPv3CredentialsUnsupportedMediaType struct {
}

func (o *UpdateSNMPv3CredentialsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsUnsupportedMediaType ", 415)
}

func (o *UpdateSNMPv3CredentialsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsInternalServerError creates a UpdateSNMPv3CredentialsInternalServerError with default headers values
func NewUpdateSNMPv3CredentialsInternalServerError() *UpdateSNMPv3CredentialsInternalServerError {
	return &UpdateSNMPv3CredentialsInternalServerError{}
}

/*UpdateSNMPv3CredentialsInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type UpdateSNMPv3CredentialsInternalServerError struct {
}

func (o *UpdateSNMPv3CredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsInternalServerError ", 500)
}

func (o *UpdateSNMPv3CredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsNotImplemented creates a UpdateSNMPv3CredentialsNotImplemented with default headers values
func NewUpdateSNMPv3CredentialsNotImplemented() *UpdateSNMPv3CredentialsNotImplemented {
	return &UpdateSNMPv3CredentialsNotImplemented{}
}

/*UpdateSNMPv3CredentialsNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type UpdateSNMPv3CredentialsNotImplemented struct {
}

func (o *UpdateSNMPv3CredentialsNotImplemented) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsNotImplemented ", 501)
}

func (o *UpdateSNMPv3CredentialsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsServiceUnavailable creates a UpdateSNMPv3CredentialsServiceUnavailable with default headers values
func NewUpdateSNMPv3CredentialsServiceUnavailable() *UpdateSNMPv3CredentialsServiceUnavailable {
	return &UpdateSNMPv3CredentialsServiceUnavailable{}
}

/*UpdateSNMPv3CredentialsServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type UpdateSNMPv3CredentialsServiceUnavailable struct {
}

func (o *UpdateSNMPv3CredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsServiceUnavailable ", 503)
}

func (o *UpdateSNMPv3CredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSNMPv3CredentialsGatewayTimeout creates a UpdateSNMPv3CredentialsGatewayTimeout with default headers values
func NewUpdateSNMPv3CredentialsGatewayTimeout() *UpdateSNMPv3CredentialsGatewayTimeout {
	return &UpdateSNMPv3CredentialsGatewayTimeout{}
}

/*UpdateSNMPv3CredentialsGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type UpdateSNMPv3CredentialsGatewayTimeout struct {
}

func (o *UpdateSNMPv3CredentialsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /dna/intent/api/v1/global-credential/snmpv3][%d] updateSNMPv3CredentialsGatewayTimeout ", 504)
}

func (o *UpdateSNMPv3CredentialsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
