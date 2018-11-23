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

// CreateSNMPv3CredentialsReader is a Reader for the CreateSNMPv3Credentials structure.
type CreateSNMPv3CredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSNMPv3CredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSNMPv3CredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewCreateSNMPv3CredentialsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewCreateSNMPv3CredentialsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewCreateSNMPv3CredentialsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewCreateSNMPv3CredentialsPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSNMPv3CredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateSNMPv3CredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateSNMPv3CredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateSNMPv3CredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSNMPv3CredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewCreateSNMPv3CredentialsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateSNMPv3CredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewCreateSNMPv3CredentialsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateSNMPv3CredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewCreateSNMPv3CredentialsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSNMPv3CredentialsOK creates a CreateSNMPv3CredentialsOK with default headers values
func NewCreateSNMPv3CredentialsOK() *CreateSNMPv3CredentialsOK {
	return &CreateSNMPv3CredentialsOK{}
}

/*CreateSNMPv3CredentialsOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type CreateSNMPv3CredentialsOK struct {
	Payload *models.TaskIDResult
}

func (o *CreateSNMPv3CredentialsOK) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsOK  %+v", 200, o.Payload)
}

func (o *CreateSNMPv3CredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskIDResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSNMPv3CredentialsCreated creates a CreateSNMPv3CredentialsCreated with default headers values
func NewCreateSNMPv3CredentialsCreated() *CreateSNMPv3CredentialsCreated {
	return &CreateSNMPv3CredentialsCreated{}
}

/*CreateSNMPv3CredentialsCreated handles this case with default header values.

The POST/PUT request was fulfilled and a new resource has been created. Information about the resource is in the response body.
*/
type CreateSNMPv3CredentialsCreated struct {
}

func (o *CreateSNMPv3CredentialsCreated) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsCreated ", 201)
}

func (o *CreateSNMPv3CredentialsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsAccepted creates a CreateSNMPv3CredentialsAccepted with default headers values
func NewCreateSNMPv3CredentialsAccepted() *CreateSNMPv3CredentialsAccepted {
	return &CreateSNMPv3CredentialsAccepted{}
}

/*CreateSNMPv3CredentialsAccepted handles this case with default header values.

The request was accepted for processing, but the processing has not been completed.
*/
type CreateSNMPv3CredentialsAccepted struct {
}

func (o *CreateSNMPv3CredentialsAccepted) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsAccepted ", 202)
}

func (o *CreateSNMPv3CredentialsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsNoContent creates a CreateSNMPv3CredentialsNoContent with default headers values
func NewCreateSNMPv3CredentialsNoContent() *CreateSNMPv3CredentialsNoContent {
	return &CreateSNMPv3CredentialsNoContent{}
}

/*CreateSNMPv3CredentialsNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type CreateSNMPv3CredentialsNoContent struct {
}

func (o *CreateSNMPv3CredentialsNoContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsNoContent ", 204)
}

func (o *CreateSNMPv3CredentialsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsPartialContent creates a CreateSNMPv3CredentialsPartialContent with default headers values
func NewCreateSNMPv3CredentialsPartialContent() *CreateSNMPv3CredentialsPartialContent {
	return &CreateSNMPv3CredentialsPartialContent{}
}

/*CreateSNMPv3CredentialsPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type CreateSNMPv3CredentialsPartialContent struct {
}

func (o *CreateSNMPv3CredentialsPartialContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsPartialContent ", 206)
}

func (o *CreateSNMPv3CredentialsPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsBadRequest creates a CreateSNMPv3CredentialsBadRequest with default headers values
func NewCreateSNMPv3CredentialsBadRequest() *CreateSNMPv3CredentialsBadRequest {
	return &CreateSNMPv3CredentialsBadRequest{}
}

/*CreateSNMPv3CredentialsBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type CreateSNMPv3CredentialsBadRequest struct {
}

func (o *CreateSNMPv3CredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsBadRequest ", 400)
}

func (o *CreateSNMPv3CredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsUnauthorized creates a CreateSNMPv3CredentialsUnauthorized with default headers values
func NewCreateSNMPv3CredentialsUnauthorized() *CreateSNMPv3CredentialsUnauthorized {
	return &CreateSNMPv3CredentialsUnauthorized{}
}

/*CreateSNMPv3CredentialsUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type CreateSNMPv3CredentialsUnauthorized struct {
}

func (o *CreateSNMPv3CredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsUnauthorized ", 401)
}

func (o *CreateSNMPv3CredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsForbidden creates a CreateSNMPv3CredentialsForbidden with default headers values
func NewCreateSNMPv3CredentialsForbidden() *CreateSNMPv3CredentialsForbidden {
	return &CreateSNMPv3CredentialsForbidden{}
}

/*CreateSNMPv3CredentialsForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type CreateSNMPv3CredentialsForbidden struct {
}

func (o *CreateSNMPv3CredentialsForbidden) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsForbidden ", 403)
}

func (o *CreateSNMPv3CredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsNotFound creates a CreateSNMPv3CredentialsNotFound with default headers values
func NewCreateSNMPv3CredentialsNotFound() *CreateSNMPv3CredentialsNotFound {
	return &CreateSNMPv3CredentialsNotFound{}
}

/*CreateSNMPv3CredentialsNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type CreateSNMPv3CredentialsNotFound struct {
}

func (o *CreateSNMPv3CredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsNotFound ", 404)
}

func (o *CreateSNMPv3CredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsConflict creates a CreateSNMPv3CredentialsConflict with default headers values
func NewCreateSNMPv3CredentialsConflict() *CreateSNMPv3CredentialsConflict {
	return &CreateSNMPv3CredentialsConflict{}
}

/*CreateSNMPv3CredentialsConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type CreateSNMPv3CredentialsConflict struct {
}

func (o *CreateSNMPv3CredentialsConflict) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsConflict ", 409)
}

func (o *CreateSNMPv3CredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsUnsupportedMediaType creates a CreateSNMPv3CredentialsUnsupportedMediaType with default headers values
func NewCreateSNMPv3CredentialsUnsupportedMediaType() *CreateSNMPv3CredentialsUnsupportedMediaType {
	return &CreateSNMPv3CredentialsUnsupportedMediaType{}
}

/*CreateSNMPv3CredentialsUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type CreateSNMPv3CredentialsUnsupportedMediaType struct {
}

func (o *CreateSNMPv3CredentialsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsUnsupportedMediaType ", 415)
}

func (o *CreateSNMPv3CredentialsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsInternalServerError creates a CreateSNMPv3CredentialsInternalServerError with default headers values
func NewCreateSNMPv3CredentialsInternalServerError() *CreateSNMPv3CredentialsInternalServerError {
	return &CreateSNMPv3CredentialsInternalServerError{}
}

/*CreateSNMPv3CredentialsInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type CreateSNMPv3CredentialsInternalServerError struct {
}

func (o *CreateSNMPv3CredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsInternalServerError ", 500)
}

func (o *CreateSNMPv3CredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsNotImplemented creates a CreateSNMPv3CredentialsNotImplemented with default headers values
func NewCreateSNMPv3CredentialsNotImplemented() *CreateSNMPv3CredentialsNotImplemented {
	return &CreateSNMPv3CredentialsNotImplemented{}
}

/*CreateSNMPv3CredentialsNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type CreateSNMPv3CredentialsNotImplemented struct {
}

func (o *CreateSNMPv3CredentialsNotImplemented) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsNotImplemented ", 501)
}

func (o *CreateSNMPv3CredentialsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsServiceUnavailable creates a CreateSNMPv3CredentialsServiceUnavailable with default headers values
func NewCreateSNMPv3CredentialsServiceUnavailable() *CreateSNMPv3CredentialsServiceUnavailable {
	return &CreateSNMPv3CredentialsServiceUnavailable{}
}

/*CreateSNMPv3CredentialsServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type CreateSNMPv3CredentialsServiceUnavailable struct {
}

func (o *CreateSNMPv3CredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsServiceUnavailable ", 503)
}

func (o *CreateSNMPv3CredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSNMPv3CredentialsGatewayTimeout creates a CreateSNMPv3CredentialsGatewayTimeout with default headers values
func NewCreateSNMPv3CredentialsGatewayTimeout() *CreateSNMPv3CredentialsGatewayTimeout {
	return &CreateSNMPv3CredentialsGatewayTimeout{}
}

/*CreateSNMPv3CredentialsGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type CreateSNMPv3CredentialsGatewayTimeout struct {
}

func (o *CreateSNMPv3CredentialsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/snmpv3][%d] createSNMPv3CredentialsGatewayTimeout ", 504)
}

func (o *CreateSNMPv3CredentialsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
