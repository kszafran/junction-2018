// Code generated by go-swagger; DO NOT EDIT.

package pn_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// ClaimDeviceReader is a Reader for the ClaimDevice structure.
type ClaimDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClaimDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClaimDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewClaimDeviceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewClaimDeviceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewClaimDeviceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewClaimDevicePartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewClaimDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewClaimDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewClaimDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewClaimDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewClaimDeviceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewClaimDeviceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewClaimDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewClaimDeviceNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewClaimDeviceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewClaimDeviceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClaimDeviceOK creates a ClaimDeviceOK with default headers values
func NewClaimDeviceOK() *ClaimDeviceOK {
	return &ClaimDeviceOK{}
}

/*ClaimDeviceOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type ClaimDeviceOK struct {
	Payload *models.ClaimDeviceResponse
}

func (o *ClaimDeviceOK) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceOK  %+v", 200, o.Payload)
}

func (o *ClaimDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClaimDeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClaimDeviceCreated creates a ClaimDeviceCreated with default headers values
func NewClaimDeviceCreated() *ClaimDeviceCreated {
	return &ClaimDeviceCreated{}
}

/*ClaimDeviceCreated handles this case with default header values.

The POST/PUT request was fulfilled and a new resource has been created. Information about the resource is in the response body.
*/
type ClaimDeviceCreated struct {
}

func (o *ClaimDeviceCreated) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceCreated ", 201)
}

func (o *ClaimDeviceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceAccepted creates a ClaimDeviceAccepted with default headers values
func NewClaimDeviceAccepted() *ClaimDeviceAccepted {
	return &ClaimDeviceAccepted{}
}

/*ClaimDeviceAccepted handles this case with default header values.

The request was accepted for processing, but the processing has not been completed.
*/
type ClaimDeviceAccepted struct {
}

func (o *ClaimDeviceAccepted) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceAccepted ", 202)
}

func (o *ClaimDeviceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceNoContent creates a ClaimDeviceNoContent with default headers values
func NewClaimDeviceNoContent() *ClaimDeviceNoContent {
	return &ClaimDeviceNoContent{}
}

/*ClaimDeviceNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type ClaimDeviceNoContent struct {
}

func (o *ClaimDeviceNoContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceNoContent ", 204)
}

func (o *ClaimDeviceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDevicePartialContent creates a ClaimDevicePartialContent with default headers values
func NewClaimDevicePartialContent() *ClaimDevicePartialContent {
	return &ClaimDevicePartialContent{}
}

/*ClaimDevicePartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type ClaimDevicePartialContent struct {
}

func (o *ClaimDevicePartialContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDevicePartialContent ", 206)
}

func (o *ClaimDevicePartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceBadRequest creates a ClaimDeviceBadRequest with default headers values
func NewClaimDeviceBadRequest() *ClaimDeviceBadRequest {
	return &ClaimDeviceBadRequest{}
}

/*ClaimDeviceBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type ClaimDeviceBadRequest struct {
}

func (o *ClaimDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceBadRequest ", 400)
}

func (o *ClaimDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceUnauthorized creates a ClaimDeviceUnauthorized with default headers values
func NewClaimDeviceUnauthorized() *ClaimDeviceUnauthorized {
	return &ClaimDeviceUnauthorized{}
}

/*ClaimDeviceUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type ClaimDeviceUnauthorized struct {
}

func (o *ClaimDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceUnauthorized ", 401)
}

func (o *ClaimDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceForbidden creates a ClaimDeviceForbidden with default headers values
func NewClaimDeviceForbidden() *ClaimDeviceForbidden {
	return &ClaimDeviceForbidden{}
}

/*ClaimDeviceForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type ClaimDeviceForbidden struct {
}

func (o *ClaimDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceForbidden ", 403)
}

func (o *ClaimDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceNotFound creates a ClaimDeviceNotFound with default headers values
func NewClaimDeviceNotFound() *ClaimDeviceNotFound {
	return &ClaimDeviceNotFound{}
}

/*ClaimDeviceNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type ClaimDeviceNotFound struct {
}

func (o *ClaimDeviceNotFound) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceNotFound ", 404)
}

func (o *ClaimDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceConflict creates a ClaimDeviceConflict with default headers values
func NewClaimDeviceConflict() *ClaimDeviceConflict {
	return &ClaimDeviceConflict{}
}

/*ClaimDeviceConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type ClaimDeviceConflict struct {
}

func (o *ClaimDeviceConflict) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceConflict ", 409)
}

func (o *ClaimDeviceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceUnsupportedMediaType creates a ClaimDeviceUnsupportedMediaType with default headers values
func NewClaimDeviceUnsupportedMediaType() *ClaimDeviceUnsupportedMediaType {
	return &ClaimDeviceUnsupportedMediaType{}
}

/*ClaimDeviceUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type ClaimDeviceUnsupportedMediaType struct {
}

func (o *ClaimDeviceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceUnsupportedMediaType ", 415)
}

func (o *ClaimDeviceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceInternalServerError creates a ClaimDeviceInternalServerError with default headers values
func NewClaimDeviceInternalServerError() *ClaimDeviceInternalServerError {
	return &ClaimDeviceInternalServerError{}
}

/*ClaimDeviceInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type ClaimDeviceInternalServerError struct {
}

func (o *ClaimDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceInternalServerError ", 500)
}

func (o *ClaimDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceNotImplemented creates a ClaimDeviceNotImplemented with default headers values
func NewClaimDeviceNotImplemented() *ClaimDeviceNotImplemented {
	return &ClaimDeviceNotImplemented{}
}

/*ClaimDeviceNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type ClaimDeviceNotImplemented struct {
}

func (o *ClaimDeviceNotImplemented) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceNotImplemented ", 501)
}

func (o *ClaimDeviceNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceServiceUnavailable creates a ClaimDeviceServiceUnavailable with default headers values
func NewClaimDeviceServiceUnavailable() *ClaimDeviceServiceUnavailable {
	return &ClaimDeviceServiceUnavailable{}
}

/*ClaimDeviceServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type ClaimDeviceServiceUnavailable struct {
}

func (o *ClaimDeviceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceServiceUnavailable ", 503)
}

func (o *ClaimDeviceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClaimDeviceGatewayTimeout creates a ClaimDeviceGatewayTimeout with default headers values
func NewClaimDeviceGatewayTimeout() *ClaimDeviceGatewayTimeout {
	return &ClaimDeviceGatewayTimeout{}
}

/*ClaimDeviceGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type ClaimDeviceGatewayTimeout struct {
}

func (o *ClaimDeviceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/onboarding/pnp-device/claim][%d] claimDeviceGatewayTimeout ", 504)
}

func (o *ClaimDeviceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
