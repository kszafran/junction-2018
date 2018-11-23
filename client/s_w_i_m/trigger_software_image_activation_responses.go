// Code generated by go-swagger; DO NOT EDIT.

package s_w_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// TriggerSoftwareImageActivationReader is a Reader for the TriggerSoftwareImageActivation structure.
type TriggerSoftwareImageActivationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerSoftwareImageActivationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTriggerSoftwareImageActivationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewTriggerSoftwareImageActivationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewTriggerSoftwareImageActivationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewTriggerSoftwareImageActivationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewTriggerSoftwareImageActivationPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTriggerSoftwareImageActivationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTriggerSoftwareImageActivationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTriggerSoftwareImageActivationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTriggerSoftwareImageActivationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewTriggerSoftwareImageActivationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewTriggerSoftwareImageActivationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTriggerSoftwareImageActivationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewTriggerSoftwareImageActivationNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTriggerSoftwareImageActivationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewTriggerSoftwareImageActivationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggerSoftwareImageActivationOK creates a TriggerSoftwareImageActivationOK with default headers values
func NewTriggerSoftwareImageActivationOK() *TriggerSoftwareImageActivationOK {
	return &TriggerSoftwareImageActivationOK{}
}

/*TriggerSoftwareImageActivationOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type TriggerSoftwareImageActivationOK struct {
	Payload *models.TaskIDResult
}

func (o *TriggerSoftwareImageActivationOK) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationOK  %+v", 200, o.Payload)
}

func (o *TriggerSoftwareImageActivationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskIDResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerSoftwareImageActivationCreated creates a TriggerSoftwareImageActivationCreated with default headers values
func NewTriggerSoftwareImageActivationCreated() *TriggerSoftwareImageActivationCreated {
	return &TriggerSoftwareImageActivationCreated{}
}

/*TriggerSoftwareImageActivationCreated handles this case with default header values.

The POST/PUT request was fulfilled and a new resource has been created. Information about the resource is in the response body.
*/
type TriggerSoftwareImageActivationCreated struct {
}

func (o *TriggerSoftwareImageActivationCreated) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationCreated ", 201)
}

func (o *TriggerSoftwareImageActivationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationAccepted creates a TriggerSoftwareImageActivationAccepted with default headers values
func NewTriggerSoftwareImageActivationAccepted() *TriggerSoftwareImageActivationAccepted {
	return &TriggerSoftwareImageActivationAccepted{}
}

/*TriggerSoftwareImageActivationAccepted handles this case with default header values.

The request was accepted for processing, but the processing has not been completed.
*/
type TriggerSoftwareImageActivationAccepted struct {
}

func (o *TriggerSoftwareImageActivationAccepted) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationAccepted ", 202)
}

func (o *TriggerSoftwareImageActivationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationNoContent creates a TriggerSoftwareImageActivationNoContent with default headers values
func NewTriggerSoftwareImageActivationNoContent() *TriggerSoftwareImageActivationNoContent {
	return &TriggerSoftwareImageActivationNoContent{}
}

/*TriggerSoftwareImageActivationNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type TriggerSoftwareImageActivationNoContent struct {
}

func (o *TriggerSoftwareImageActivationNoContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationNoContent ", 204)
}

func (o *TriggerSoftwareImageActivationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationPartialContent creates a TriggerSoftwareImageActivationPartialContent with default headers values
func NewTriggerSoftwareImageActivationPartialContent() *TriggerSoftwareImageActivationPartialContent {
	return &TriggerSoftwareImageActivationPartialContent{}
}

/*TriggerSoftwareImageActivationPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type TriggerSoftwareImageActivationPartialContent struct {
}

func (o *TriggerSoftwareImageActivationPartialContent) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationPartialContent ", 206)
}

func (o *TriggerSoftwareImageActivationPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationBadRequest creates a TriggerSoftwareImageActivationBadRequest with default headers values
func NewTriggerSoftwareImageActivationBadRequest() *TriggerSoftwareImageActivationBadRequest {
	return &TriggerSoftwareImageActivationBadRequest{}
}

/*TriggerSoftwareImageActivationBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type TriggerSoftwareImageActivationBadRequest struct {
}

func (o *TriggerSoftwareImageActivationBadRequest) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationBadRequest ", 400)
}

func (o *TriggerSoftwareImageActivationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationUnauthorized creates a TriggerSoftwareImageActivationUnauthorized with default headers values
func NewTriggerSoftwareImageActivationUnauthorized() *TriggerSoftwareImageActivationUnauthorized {
	return &TriggerSoftwareImageActivationUnauthorized{}
}

/*TriggerSoftwareImageActivationUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type TriggerSoftwareImageActivationUnauthorized struct {
}

func (o *TriggerSoftwareImageActivationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationUnauthorized ", 401)
}

func (o *TriggerSoftwareImageActivationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationForbidden creates a TriggerSoftwareImageActivationForbidden with default headers values
func NewTriggerSoftwareImageActivationForbidden() *TriggerSoftwareImageActivationForbidden {
	return &TriggerSoftwareImageActivationForbidden{}
}

/*TriggerSoftwareImageActivationForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type TriggerSoftwareImageActivationForbidden struct {
}

func (o *TriggerSoftwareImageActivationForbidden) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationForbidden ", 403)
}

func (o *TriggerSoftwareImageActivationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationNotFound creates a TriggerSoftwareImageActivationNotFound with default headers values
func NewTriggerSoftwareImageActivationNotFound() *TriggerSoftwareImageActivationNotFound {
	return &TriggerSoftwareImageActivationNotFound{}
}

/*TriggerSoftwareImageActivationNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type TriggerSoftwareImageActivationNotFound struct {
}

func (o *TriggerSoftwareImageActivationNotFound) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationNotFound ", 404)
}

func (o *TriggerSoftwareImageActivationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationConflict creates a TriggerSoftwareImageActivationConflict with default headers values
func NewTriggerSoftwareImageActivationConflict() *TriggerSoftwareImageActivationConflict {
	return &TriggerSoftwareImageActivationConflict{}
}

/*TriggerSoftwareImageActivationConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type TriggerSoftwareImageActivationConflict struct {
}

func (o *TriggerSoftwareImageActivationConflict) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationConflict ", 409)
}

func (o *TriggerSoftwareImageActivationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationUnsupportedMediaType creates a TriggerSoftwareImageActivationUnsupportedMediaType with default headers values
func NewTriggerSoftwareImageActivationUnsupportedMediaType() *TriggerSoftwareImageActivationUnsupportedMediaType {
	return &TriggerSoftwareImageActivationUnsupportedMediaType{}
}

/*TriggerSoftwareImageActivationUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type TriggerSoftwareImageActivationUnsupportedMediaType struct {
}

func (o *TriggerSoftwareImageActivationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationUnsupportedMediaType ", 415)
}

func (o *TriggerSoftwareImageActivationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationInternalServerError creates a TriggerSoftwareImageActivationInternalServerError with default headers values
func NewTriggerSoftwareImageActivationInternalServerError() *TriggerSoftwareImageActivationInternalServerError {
	return &TriggerSoftwareImageActivationInternalServerError{}
}

/*TriggerSoftwareImageActivationInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type TriggerSoftwareImageActivationInternalServerError struct {
}

func (o *TriggerSoftwareImageActivationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationInternalServerError ", 500)
}

func (o *TriggerSoftwareImageActivationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationNotImplemented creates a TriggerSoftwareImageActivationNotImplemented with default headers values
func NewTriggerSoftwareImageActivationNotImplemented() *TriggerSoftwareImageActivationNotImplemented {
	return &TriggerSoftwareImageActivationNotImplemented{}
}

/*TriggerSoftwareImageActivationNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type TriggerSoftwareImageActivationNotImplemented struct {
}

func (o *TriggerSoftwareImageActivationNotImplemented) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationNotImplemented ", 501)
}

func (o *TriggerSoftwareImageActivationNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationServiceUnavailable creates a TriggerSoftwareImageActivationServiceUnavailable with default headers values
func NewTriggerSoftwareImageActivationServiceUnavailable() *TriggerSoftwareImageActivationServiceUnavailable {
	return &TriggerSoftwareImageActivationServiceUnavailable{}
}

/*TriggerSoftwareImageActivationServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type TriggerSoftwareImageActivationServiceUnavailable struct {
}

func (o *TriggerSoftwareImageActivationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationServiceUnavailable ", 503)
}

func (o *TriggerSoftwareImageActivationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerSoftwareImageActivationGatewayTimeout creates a TriggerSoftwareImageActivationGatewayTimeout with default headers values
func NewTriggerSoftwareImageActivationGatewayTimeout() *TriggerSoftwareImageActivationGatewayTimeout {
	return &TriggerSoftwareImageActivationGatewayTimeout{}
}

/*TriggerSoftwareImageActivationGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type TriggerSoftwareImageActivationGatewayTimeout struct {
}

func (o *TriggerSoftwareImageActivationGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/image/activation/device][%d] triggerSoftwareImageActivationGatewayTimeout ", 504)
}

func (o *TriggerSoftwareImageActivationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
