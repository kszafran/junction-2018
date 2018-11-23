// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// GetFunctionalCapabilityByIDReader is a Reader for the GetFunctionalCapabilityByID structure.
type GetFunctionalCapabilityByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFunctionalCapabilityByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFunctionalCapabilityByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetFunctionalCapabilityByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetFunctionalCapabilityByIDPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFunctionalCapabilityByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFunctionalCapabilityByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFunctionalCapabilityByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFunctionalCapabilityByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetFunctionalCapabilityByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetFunctionalCapabilityByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFunctionalCapabilityByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetFunctionalCapabilityByIDNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetFunctionalCapabilityByIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetFunctionalCapabilityByIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFunctionalCapabilityByIDOK creates a GetFunctionalCapabilityByIDOK with default headers values
func NewGetFunctionalCapabilityByIDOK() *GetFunctionalCapabilityByIDOK {
	return &GetFunctionalCapabilityByIDOK{}
}

/*GetFunctionalCapabilityByIDOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetFunctionalCapabilityByIDOK struct {
	Payload *models.FunctionalCapabilityResult
}

func (o *GetFunctionalCapabilityByIDOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdOK  %+v", 200, o.Payload)
}

func (o *GetFunctionalCapabilityByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionalCapabilityResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunctionalCapabilityByIDNoContent creates a GetFunctionalCapabilityByIDNoContent with default headers values
func NewGetFunctionalCapabilityByIDNoContent() *GetFunctionalCapabilityByIDNoContent {
	return &GetFunctionalCapabilityByIDNoContent{}
}

/*GetFunctionalCapabilityByIDNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type GetFunctionalCapabilityByIDNoContent struct {
}

func (o *GetFunctionalCapabilityByIDNoContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdNoContent ", 204)
}

func (o *GetFunctionalCapabilityByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDPartialContent creates a GetFunctionalCapabilityByIDPartialContent with default headers values
func NewGetFunctionalCapabilityByIDPartialContent() *GetFunctionalCapabilityByIDPartialContent {
	return &GetFunctionalCapabilityByIDPartialContent{}
}

/*GetFunctionalCapabilityByIDPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type GetFunctionalCapabilityByIDPartialContent struct {
}

func (o *GetFunctionalCapabilityByIDPartialContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdPartialContent ", 206)
}

func (o *GetFunctionalCapabilityByIDPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDBadRequest creates a GetFunctionalCapabilityByIDBadRequest with default headers values
func NewGetFunctionalCapabilityByIDBadRequest() *GetFunctionalCapabilityByIDBadRequest {
	return &GetFunctionalCapabilityByIDBadRequest{}
}

/*GetFunctionalCapabilityByIDBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetFunctionalCapabilityByIDBadRequest struct {
}

func (o *GetFunctionalCapabilityByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdBadRequest ", 400)
}

func (o *GetFunctionalCapabilityByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDUnauthorized creates a GetFunctionalCapabilityByIDUnauthorized with default headers values
func NewGetFunctionalCapabilityByIDUnauthorized() *GetFunctionalCapabilityByIDUnauthorized {
	return &GetFunctionalCapabilityByIDUnauthorized{}
}

/*GetFunctionalCapabilityByIDUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetFunctionalCapabilityByIDUnauthorized struct {
}

func (o *GetFunctionalCapabilityByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdUnauthorized ", 401)
}

func (o *GetFunctionalCapabilityByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDForbidden creates a GetFunctionalCapabilityByIDForbidden with default headers values
func NewGetFunctionalCapabilityByIDForbidden() *GetFunctionalCapabilityByIDForbidden {
	return &GetFunctionalCapabilityByIDForbidden{}
}

/*GetFunctionalCapabilityByIDForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type GetFunctionalCapabilityByIDForbidden struct {
}

func (o *GetFunctionalCapabilityByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdForbidden ", 403)
}

func (o *GetFunctionalCapabilityByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDNotFound creates a GetFunctionalCapabilityByIDNotFound with default headers values
func NewGetFunctionalCapabilityByIDNotFound() *GetFunctionalCapabilityByIDNotFound {
	return &GetFunctionalCapabilityByIDNotFound{}
}

/*GetFunctionalCapabilityByIDNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetFunctionalCapabilityByIDNotFound struct {
}

func (o *GetFunctionalCapabilityByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdNotFound ", 404)
}

func (o *GetFunctionalCapabilityByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDConflict creates a GetFunctionalCapabilityByIDConflict with default headers values
func NewGetFunctionalCapabilityByIDConflict() *GetFunctionalCapabilityByIDConflict {
	return &GetFunctionalCapabilityByIDConflict{}
}

/*GetFunctionalCapabilityByIDConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type GetFunctionalCapabilityByIDConflict struct {
}

func (o *GetFunctionalCapabilityByIDConflict) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdConflict ", 409)
}

func (o *GetFunctionalCapabilityByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDUnsupportedMediaType creates a GetFunctionalCapabilityByIDUnsupportedMediaType with default headers values
func NewGetFunctionalCapabilityByIDUnsupportedMediaType() *GetFunctionalCapabilityByIDUnsupportedMediaType {
	return &GetFunctionalCapabilityByIDUnsupportedMediaType{}
}

/*GetFunctionalCapabilityByIDUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type GetFunctionalCapabilityByIDUnsupportedMediaType struct {
}

func (o *GetFunctionalCapabilityByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdUnsupportedMediaType ", 415)
}

func (o *GetFunctionalCapabilityByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDInternalServerError creates a GetFunctionalCapabilityByIDInternalServerError with default headers values
func NewGetFunctionalCapabilityByIDInternalServerError() *GetFunctionalCapabilityByIDInternalServerError {
	return &GetFunctionalCapabilityByIDInternalServerError{}
}

/*GetFunctionalCapabilityByIDInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetFunctionalCapabilityByIDInternalServerError struct {
}

func (o *GetFunctionalCapabilityByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdInternalServerError ", 500)
}

func (o *GetFunctionalCapabilityByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDNotImplemented creates a GetFunctionalCapabilityByIDNotImplemented with default headers values
func NewGetFunctionalCapabilityByIDNotImplemented() *GetFunctionalCapabilityByIDNotImplemented {
	return &GetFunctionalCapabilityByIDNotImplemented{}
}

/*GetFunctionalCapabilityByIDNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type GetFunctionalCapabilityByIDNotImplemented struct {
}

func (o *GetFunctionalCapabilityByIDNotImplemented) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdNotImplemented ", 501)
}

func (o *GetFunctionalCapabilityByIDNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDServiceUnavailable creates a GetFunctionalCapabilityByIDServiceUnavailable with default headers values
func NewGetFunctionalCapabilityByIDServiceUnavailable() *GetFunctionalCapabilityByIDServiceUnavailable {
	return &GetFunctionalCapabilityByIDServiceUnavailable{}
}

/*GetFunctionalCapabilityByIDServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type GetFunctionalCapabilityByIDServiceUnavailable struct {
}

func (o *GetFunctionalCapabilityByIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdServiceUnavailable ", 503)
}

func (o *GetFunctionalCapabilityByIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunctionalCapabilityByIDGatewayTimeout creates a GetFunctionalCapabilityByIDGatewayTimeout with default headers values
func NewGetFunctionalCapabilityByIDGatewayTimeout() *GetFunctionalCapabilityByIDGatewayTimeout {
	return &GetFunctionalCapabilityByIDGatewayTimeout{}
}

/*GetFunctionalCapabilityByIDGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type GetFunctionalCapabilityByIDGatewayTimeout struct {
}

func (o *GetFunctionalCapabilityByIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-device/functional-capability/{id}][%d] getFunctionalCapabilityByIdGatewayTimeout ", 504)
}

func (o *GetFunctionalCapabilityByIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
