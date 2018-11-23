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

// CreateCLICredentialsReader is a Reader for the CreateCLICredentials structure.
type CreateCLICredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCLICredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCLICredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewCreateCLICredentialsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewCreateCLICredentialsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewCreateCLICredentialsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewCreateCLICredentialsPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateCLICredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateCLICredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateCLICredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateCLICredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateCLICredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewCreateCLICredentialsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateCLICredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewCreateCLICredentialsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateCLICredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewCreateCLICredentialsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCLICredentialsOK creates a CreateCLICredentialsOK with default headers values
func NewCreateCLICredentialsOK() *CreateCLICredentialsOK {
	return &CreateCLICredentialsOK{}
}

/*CreateCLICredentialsOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type CreateCLICredentialsOK struct {
	Payload *models.TaskIDResult
}

func (o *CreateCLICredentialsOK) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsOK  %+v", 200, o.Payload)
}

func (o *CreateCLICredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskIDResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCLICredentialsCreated creates a CreateCLICredentialsCreated with default headers values
func NewCreateCLICredentialsCreated() *CreateCLICredentialsCreated {
	return &CreateCLICredentialsCreated{}
}

/*CreateCLICredentialsCreated handles this case with default header values.

The POST/PUT request was fulfilled and a new resource has been created. Information about the resource is in the response body.
*/
type CreateCLICredentialsCreated struct {
}

func (o *CreateCLICredentialsCreated) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsCreated ", 201)
}

func (o *CreateCLICredentialsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsAccepted creates a CreateCLICredentialsAccepted with default headers values
func NewCreateCLICredentialsAccepted() *CreateCLICredentialsAccepted {
	return &CreateCLICredentialsAccepted{}
}

/*CreateCLICredentialsAccepted handles this case with default header values.

The request was accepted for processing, but the processing has not been completed.
*/
type CreateCLICredentialsAccepted struct {
}

func (o *CreateCLICredentialsAccepted) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsAccepted ", 202)
}

func (o *CreateCLICredentialsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsNoContent creates a CreateCLICredentialsNoContent with default headers values
func NewCreateCLICredentialsNoContent() *CreateCLICredentialsNoContent {
	return &CreateCLICredentialsNoContent{}
}

/*CreateCLICredentialsNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type CreateCLICredentialsNoContent struct {
}

func (o *CreateCLICredentialsNoContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsNoContent ", 204)
}

func (o *CreateCLICredentialsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsPartialContent creates a CreateCLICredentialsPartialContent with default headers values
func NewCreateCLICredentialsPartialContent() *CreateCLICredentialsPartialContent {
	return &CreateCLICredentialsPartialContent{}
}

/*CreateCLICredentialsPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type CreateCLICredentialsPartialContent struct {
}

func (o *CreateCLICredentialsPartialContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsPartialContent ", 206)
}

func (o *CreateCLICredentialsPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsBadRequest creates a CreateCLICredentialsBadRequest with default headers values
func NewCreateCLICredentialsBadRequest() *CreateCLICredentialsBadRequest {
	return &CreateCLICredentialsBadRequest{}
}

/*CreateCLICredentialsBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type CreateCLICredentialsBadRequest struct {
}

func (o *CreateCLICredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsBadRequest ", 400)
}

func (o *CreateCLICredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsUnauthorized creates a CreateCLICredentialsUnauthorized with default headers values
func NewCreateCLICredentialsUnauthorized() *CreateCLICredentialsUnauthorized {
	return &CreateCLICredentialsUnauthorized{}
}

/*CreateCLICredentialsUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type CreateCLICredentialsUnauthorized struct {
}

func (o *CreateCLICredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsUnauthorized ", 401)
}

func (o *CreateCLICredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsForbidden creates a CreateCLICredentialsForbidden with default headers values
func NewCreateCLICredentialsForbidden() *CreateCLICredentialsForbidden {
	return &CreateCLICredentialsForbidden{}
}

/*CreateCLICredentialsForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type CreateCLICredentialsForbidden struct {
}

func (o *CreateCLICredentialsForbidden) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsForbidden ", 403)
}

func (o *CreateCLICredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsNotFound creates a CreateCLICredentialsNotFound with default headers values
func NewCreateCLICredentialsNotFound() *CreateCLICredentialsNotFound {
	return &CreateCLICredentialsNotFound{}
}

/*CreateCLICredentialsNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type CreateCLICredentialsNotFound struct {
}

func (o *CreateCLICredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsNotFound ", 404)
}

func (o *CreateCLICredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsConflict creates a CreateCLICredentialsConflict with default headers values
func NewCreateCLICredentialsConflict() *CreateCLICredentialsConflict {
	return &CreateCLICredentialsConflict{}
}

/*CreateCLICredentialsConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type CreateCLICredentialsConflict struct {
}

func (o *CreateCLICredentialsConflict) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsConflict ", 409)
}

func (o *CreateCLICredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsUnsupportedMediaType creates a CreateCLICredentialsUnsupportedMediaType with default headers values
func NewCreateCLICredentialsUnsupportedMediaType() *CreateCLICredentialsUnsupportedMediaType {
	return &CreateCLICredentialsUnsupportedMediaType{}
}

/*CreateCLICredentialsUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type CreateCLICredentialsUnsupportedMediaType struct {
}

func (o *CreateCLICredentialsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsUnsupportedMediaType ", 415)
}

func (o *CreateCLICredentialsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsInternalServerError creates a CreateCLICredentialsInternalServerError with default headers values
func NewCreateCLICredentialsInternalServerError() *CreateCLICredentialsInternalServerError {
	return &CreateCLICredentialsInternalServerError{}
}

/*CreateCLICredentialsInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type CreateCLICredentialsInternalServerError struct {
}

func (o *CreateCLICredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsInternalServerError ", 500)
}

func (o *CreateCLICredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsNotImplemented creates a CreateCLICredentialsNotImplemented with default headers values
func NewCreateCLICredentialsNotImplemented() *CreateCLICredentialsNotImplemented {
	return &CreateCLICredentialsNotImplemented{}
}

/*CreateCLICredentialsNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type CreateCLICredentialsNotImplemented struct {
}

func (o *CreateCLICredentialsNotImplemented) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsNotImplemented ", 501)
}

func (o *CreateCLICredentialsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsServiceUnavailable creates a CreateCLICredentialsServiceUnavailable with default headers values
func NewCreateCLICredentialsServiceUnavailable() *CreateCLICredentialsServiceUnavailable {
	return &CreateCLICredentialsServiceUnavailable{}
}

/*CreateCLICredentialsServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type CreateCLICredentialsServiceUnavailable struct {
}

func (o *CreateCLICredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsServiceUnavailable ", 503)
}

func (o *CreateCLICredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCLICredentialsGatewayTimeout creates a CreateCLICredentialsGatewayTimeout with default headers values
func NewCreateCLICredentialsGatewayTimeout() *CreateCLICredentialsGatewayTimeout {
	return &CreateCLICredentialsGatewayTimeout{}
}

/*CreateCLICredentialsGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type CreateCLICredentialsGatewayTimeout struct {
}

func (o *CreateCLICredentialsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/global-credential/cli][%d] createCLICredentialsGatewayTimeout ", 504)
}

func (o *CreateCLICredentialsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
