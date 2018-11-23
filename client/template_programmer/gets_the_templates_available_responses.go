// Code generated by go-swagger; DO NOT EDIT.

package template_programmer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// GetsTheTemplatesAvailableReader is a Reader for the GetsTheTemplatesAvailable structure.
type GetsTheTemplatesAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetsTheTemplatesAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetsTheTemplatesAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetsTheTemplatesAvailableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetsTheTemplatesAvailablePartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetsTheTemplatesAvailableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetsTheTemplatesAvailableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetsTheTemplatesAvailableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetsTheTemplatesAvailableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetsTheTemplatesAvailableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetsTheTemplatesAvailableUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetsTheTemplatesAvailableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetsTheTemplatesAvailableNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetsTheTemplatesAvailableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetsTheTemplatesAvailableGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetsTheTemplatesAvailableOK creates a GetsTheTemplatesAvailableOK with default headers values
func NewGetsTheTemplatesAvailableOK() *GetsTheTemplatesAvailableOK {
	return &GetsTheTemplatesAvailableOK{}
}

/*GetsTheTemplatesAvailableOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetsTheTemplatesAvailableOK struct {
	Payload models.CollectionTemplateInfo
}

func (o *GetsTheTemplatesAvailableOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableOK  %+v", 200, o.Payload)
}

func (o *GetsTheTemplatesAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetsTheTemplatesAvailableNoContent creates a GetsTheTemplatesAvailableNoContent with default headers values
func NewGetsTheTemplatesAvailableNoContent() *GetsTheTemplatesAvailableNoContent {
	return &GetsTheTemplatesAvailableNoContent{}
}

/*GetsTheTemplatesAvailableNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type GetsTheTemplatesAvailableNoContent struct {
}

func (o *GetsTheTemplatesAvailableNoContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableNoContent ", 204)
}

func (o *GetsTheTemplatesAvailableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailablePartialContent creates a GetsTheTemplatesAvailablePartialContent with default headers values
func NewGetsTheTemplatesAvailablePartialContent() *GetsTheTemplatesAvailablePartialContent {
	return &GetsTheTemplatesAvailablePartialContent{}
}

/*GetsTheTemplatesAvailablePartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type GetsTheTemplatesAvailablePartialContent struct {
}

func (o *GetsTheTemplatesAvailablePartialContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailablePartialContent ", 206)
}

func (o *GetsTheTemplatesAvailablePartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableBadRequest creates a GetsTheTemplatesAvailableBadRequest with default headers values
func NewGetsTheTemplatesAvailableBadRequest() *GetsTheTemplatesAvailableBadRequest {
	return &GetsTheTemplatesAvailableBadRequest{}
}

/*GetsTheTemplatesAvailableBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetsTheTemplatesAvailableBadRequest struct {
}

func (o *GetsTheTemplatesAvailableBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableBadRequest ", 400)
}

func (o *GetsTheTemplatesAvailableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableUnauthorized creates a GetsTheTemplatesAvailableUnauthorized with default headers values
func NewGetsTheTemplatesAvailableUnauthorized() *GetsTheTemplatesAvailableUnauthorized {
	return &GetsTheTemplatesAvailableUnauthorized{}
}

/*GetsTheTemplatesAvailableUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetsTheTemplatesAvailableUnauthorized struct {
}

func (o *GetsTheTemplatesAvailableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableUnauthorized ", 401)
}

func (o *GetsTheTemplatesAvailableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableForbidden creates a GetsTheTemplatesAvailableForbidden with default headers values
func NewGetsTheTemplatesAvailableForbidden() *GetsTheTemplatesAvailableForbidden {
	return &GetsTheTemplatesAvailableForbidden{}
}

/*GetsTheTemplatesAvailableForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type GetsTheTemplatesAvailableForbidden struct {
}

func (o *GetsTheTemplatesAvailableForbidden) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableForbidden ", 403)
}

func (o *GetsTheTemplatesAvailableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableNotFound creates a GetsTheTemplatesAvailableNotFound with default headers values
func NewGetsTheTemplatesAvailableNotFound() *GetsTheTemplatesAvailableNotFound {
	return &GetsTheTemplatesAvailableNotFound{}
}

/*GetsTheTemplatesAvailableNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetsTheTemplatesAvailableNotFound struct {
}

func (o *GetsTheTemplatesAvailableNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableNotFound ", 404)
}

func (o *GetsTheTemplatesAvailableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableConflict creates a GetsTheTemplatesAvailableConflict with default headers values
func NewGetsTheTemplatesAvailableConflict() *GetsTheTemplatesAvailableConflict {
	return &GetsTheTemplatesAvailableConflict{}
}

/*GetsTheTemplatesAvailableConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type GetsTheTemplatesAvailableConflict struct {
}

func (o *GetsTheTemplatesAvailableConflict) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableConflict ", 409)
}

func (o *GetsTheTemplatesAvailableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableUnsupportedMediaType creates a GetsTheTemplatesAvailableUnsupportedMediaType with default headers values
func NewGetsTheTemplatesAvailableUnsupportedMediaType() *GetsTheTemplatesAvailableUnsupportedMediaType {
	return &GetsTheTemplatesAvailableUnsupportedMediaType{}
}

/*GetsTheTemplatesAvailableUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type GetsTheTemplatesAvailableUnsupportedMediaType struct {
}

func (o *GetsTheTemplatesAvailableUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableUnsupportedMediaType ", 415)
}

func (o *GetsTheTemplatesAvailableUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableInternalServerError creates a GetsTheTemplatesAvailableInternalServerError with default headers values
func NewGetsTheTemplatesAvailableInternalServerError() *GetsTheTemplatesAvailableInternalServerError {
	return &GetsTheTemplatesAvailableInternalServerError{}
}

/*GetsTheTemplatesAvailableInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetsTheTemplatesAvailableInternalServerError struct {
}

func (o *GetsTheTemplatesAvailableInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableInternalServerError ", 500)
}

func (o *GetsTheTemplatesAvailableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableNotImplemented creates a GetsTheTemplatesAvailableNotImplemented with default headers values
func NewGetsTheTemplatesAvailableNotImplemented() *GetsTheTemplatesAvailableNotImplemented {
	return &GetsTheTemplatesAvailableNotImplemented{}
}

/*GetsTheTemplatesAvailableNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type GetsTheTemplatesAvailableNotImplemented struct {
}

func (o *GetsTheTemplatesAvailableNotImplemented) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableNotImplemented ", 501)
}

func (o *GetsTheTemplatesAvailableNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableServiceUnavailable creates a GetsTheTemplatesAvailableServiceUnavailable with default headers values
func NewGetsTheTemplatesAvailableServiceUnavailable() *GetsTheTemplatesAvailableServiceUnavailable {
	return &GetsTheTemplatesAvailableServiceUnavailable{}
}

/*GetsTheTemplatesAvailableServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type GetsTheTemplatesAvailableServiceUnavailable struct {
}

func (o *GetsTheTemplatesAvailableServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableServiceUnavailable ", 503)
}

func (o *GetsTheTemplatesAvailableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetsTheTemplatesAvailableGatewayTimeout creates a GetsTheTemplatesAvailableGatewayTimeout with default headers values
func NewGetsTheTemplatesAvailableGatewayTimeout() *GetsTheTemplatesAvailableGatewayTimeout {
	return &GetsTheTemplatesAvailableGatewayTimeout{}
}

/*GetsTheTemplatesAvailableGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type GetsTheTemplatesAvailableGatewayTimeout struct {
}

func (o *GetsTheTemplatesAvailableGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/template-programmer/template][%d] getsTheTemplatesAvailableGatewayTimeout ", 504)
}

func (o *GetsTheTemplatesAvailableGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
