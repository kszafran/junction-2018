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

// GetDiscoveriesByRangeReader is a Reader for the GetDiscoveriesByRange structure.
type GetDiscoveriesByRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoveriesByRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDiscoveriesByRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetDiscoveriesByRangeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetDiscoveriesByRangePartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDiscoveriesByRangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetDiscoveriesByRangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDiscoveriesByRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDiscoveriesByRangeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetDiscoveriesByRangeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetDiscoveriesByRangeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDiscoveriesByRangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetDiscoveriesByRangeNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetDiscoveriesByRangeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetDiscoveriesByRangeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDiscoveriesByRangeOK creates a GetDiscoveriesByRangeOK with default headers values
func NewGetDiscoveriesByRangeOK() *GetDiscoveriesByRangeOK {
	return &GetDiscoveriesByRangeOK{}
}

/*GetDiscoveriesByRangeOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetDiscoveriesByRangeOK struct {
	Payload *models.DiscoveryNIOListResult
}

func (o *GetDiscoveriesByRangeOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeOK  %+v", 200, o.Payload)
}

func (o *GetDiscoveriesByRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiscoveryNIOListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveriesByRangeNoContent creates a GetDiscoveriesByRangeNoContent with default headers values
func NewGetDiscoveriesByRangeNoContent() *GetDiscoveriesByRangeNoContent {
	return &GetDiscoveriesByRangeNoContent{}
}

/*GetDiscoveriesByRangeNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type GetDiscoveriesByRangeNoContent struct {
}

func (o *GetDiscoveriesByRangeNoContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeNoContent ", 204)
}

func (o *GetDiscoveriesByRangeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangePartialContent creates a GetDiscoveriesByRangePartialContent with default headers values
func NewGetDiscoveriesByRangePartialContent() *GetDiscoveriesByRangePartialContent {
	return &GetDiscoveriesByRangePartialContent{}
}

/*GetDiscoveriesByRangePartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type GetDiscoveriesByRangePartialContent struct {
}

func (o *GetDiscoveriesByRangePartialContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangePartialContent ", 206)
}

func (o *GetDiscoveriesByRangePartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeBadRequest creates a GetDiscoveriesByRangeBadRequest with default headers values
func NewGetDiscoveriesByRangeBadRequest() *GetDiscoveriesByRangeBadRequest {
	return &GetDiscoveriesByRangeBadRequest{}
}

/*GetDiscoveriesByRangeBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetDiscoveriesByRangeBadRequest struct {
}

func (o *GetDiscoveriesByRangeBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeBadRequest ", 400)
}

func (o *GetDiscoveriesByRangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeUnauthorized creates a GetDiscoveriesByRangeUnauthorized with default headers values
func NewGetDiscoveriesByRangeUnauthorized() *GetDiscoveriesByRangeUnauthorized {
	return &GetDiscoveriesByRangeUnauthorized{}
}

/*GetDiscoveriesByRangeUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetDiscoveriesByRangeUnauthorized struct {
}

func (o *GetDiscoveriesByRangeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeUnauthorized ", 401)
}

func (o *GetDiscoveriesByRangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeForbidden creates a GetDiscoveriesByRangeForbidden with default headers values
func NewGetDiscoveriesByRangeForbidden() *GetDiscoveriesByRangeForbidden {
	return &GetDiscoveriesByRangeForbidden{}
}

/*GetDiscoveriesByRangeForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type GetDiscoveriesByRangeForbidden struct {
}

func (o *GetDiscoveriesByRangeForbidden) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeForbidden ", 403)
}

func (o *GetDiscoveriesByRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeNotFound creates a GetDiscoveriesByRangeNotFound with default headers values
func NewGetDiscoveriesByRangeNotFound() *GetDiscoveriesByRangeNotFound {
	return &GetDiscoveriesByRangeNotFound{}
}

/*GetDiscoveriesByRangeNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetDiscoveriesByRangeNotFound struct {
}

func (o *GetDiscoveriesByRangeNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeNotFound ", 404)
}

func (o *GetDiscoveriesByRangeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeConflict creates a GetDiscoveriesByRangeConflict with default headers values
func NewGetDiscoveriesByRangeConflict() *GetDiscoveriesByRangeConflict {
	return &GetDiscoveriesByRangeConflict{}
}

/*GetDiscoveriesByRangeConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type GetDiscoveriesByRangeConflict struct {
}

func (o *GetDiscoveriesByRangeConflict) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeConflict ", 409)
}

func (o *GetDiscoveriesByRangeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeUnsupportedMediaType creates a GetDiscoveriesByRangeUnsupportedMediaType with default headers values
func NewGetDiscoveriesByRangeUnsupportedMediaType() *GetDiscoveriesByRangeUnsupportedMediaType {
	return &GetDiscoveriesByRangeUnsupportedMediaType{}
}

/*GetDiscoveriesByRangeUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type GetDiscoveriesByRangeUnsupportedMediaType struct {
}

func (o *GetDiscoveriesByRangeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeUnsupportedMediaType ", 415)
}

func (o *GetDiscoveriesByRangeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeInternalServerError creates a GetDiscoveriesByRangeInternalServerError with default headers values
func NewGetDiscoveriesByRangeInternalServerError() *GetDiscoveriesByRangeInternalServerError {
	return &GetDiscoveriesByRangeInternalServerError{}
}

/*GetDiscoveriesByRangeInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetDiscoveriesByRangeInternalServerError struct {
}

func (o *GetDiscoveriesByRangeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeInternalServerError ", 500)
}

func (o *GetDiscoveriesByRangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeNotImplemented creates a GetDiscoveriesByRangeNotImplemented with default headers values
func NewGetDiscoveriesByRangeNotImplemented() *GetDiscoveriesByRangeNotImplemented {
	return &GetDiscoveriesByRangeNotImplemented{}
}

/*GetDiscoveriesByRangeNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type GetDiscoveriesByRangeNotImplemented struct {
}

func (o *GetDiscoveriesByRangeNotImplemented) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeNotImplemented ", 501)
}

func (o *GetDiscoveriesByRangeNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeServiceUnavailable creates a GetDiscoveriesByRangeServiceUnavailable with default headers values
func NewGetDiscoveriesByRangeServiceUnavailable() *GetDiscoveriesByRangeServiceUnavailable {
	return &GetDiscoveriesByRangeServiceUnavailable{}
}

/*GetDiscoveriesByRangeServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type GetDiscoveriesByRangeServiceUnavailable struct {
}

func (o *GetDiscoveriesByRangeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeServiceUnavailable ", 503)
}

func (o *GetDiscoveriesByRangeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscoveriesByRangeGatewayTimeout creates a GetDiscoveriesByRangeGatewayTimeout with default headers values
func NewGetDiscoveriesByRangeGatewayTimeout() *GetDiscoveriesByRangeGatewayTimeout {
	return &GetDiscoveriesByRangeGatewayTimeout{}
}

/*GetDiscoveriesByRangeGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type GetDiscoveriesByRangeGatewayTimeout struct {
}

func (o *GetDiscoveriesByRangeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}][%d] getDiscoveriesByRangeGatewayTimeout ", 504)
}

func (o *GetDiscoveriesByRangeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
