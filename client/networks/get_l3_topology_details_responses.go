// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// GetL3TopologyDetailsReader is a Reader for the GetL3TopologyDetails structure.
type GetL3TopologyDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetL3TopologyDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetL3TopologyDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetL3TopologyDetailsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetL3TopologyDetailsPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetL3TopologyDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetL3TopologyDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetL3TopologyDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetL3TopologyDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetL3TopologyDetailsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetL3TopologyDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetL3TopologyDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetL3TopologyDetailsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetL3TopologyDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetL3TopologyDetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetL3TopologyDetailsOK creates a GetL3TopologyDetailsOK with default headers values
func NewGetL3TopologyDetailsOK() *GetL3TopologyDetailsOK {
	return &GetL3TopologyDetailsOK{}
}

/*GetL3TopologyDetailsOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetL3TopologyDetailsOK struct {
	Payload *models.TopologyResult
}

func (o *GetL3TopologyDetailsOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsOK  %+v", 200, o.Payload)
}

func (o *GetL3TopologyDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopologyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetL3TopologyDetailsNoContent creates a GetL3TopologyDetailsNoContent with default headers values
func NewGetL3TopologyDetailsNoContent() *GetL3TopologyDetailsNoContent {
	return &GetL3TopologyDetailsNoContent{}
}

/*GetL3TopologyDetailsNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type GetL3TopologyDetailsNoContent struct {
}

func (o *GetL3TopologyDetailsNoContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsNoContent ", 204)
}

func (o *GetL3TopologyDetailsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsPartialContent creates a GetL3TopologyDetailsPartialContent with default headers values
func NewGetL3TopologyDetailsPartialContent() *GetL3TopologyDetailsPartialContent {
	return &GetL3TopologyDetailsPartialContent{}
}

/*GetL3TopologyDetailsPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type GetL3TopologyDetailsPartialContent struct {
}

func (o *GetL3TopologyDetailsPartialContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsPartialContent ", 206)
}

func (o *GetL3TopologyDetailsPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsBadRequest creates a GetL3TopologyDetailsBadRequest with default headers values
func NewGetL3TopologyDetailsBadRequest() *GetL3TopologyDetailsBadRequest {
	return &GetL3TopologyDetailsBadRequest{}
}

/*GetL3TopologyDetailsBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetL3TopologyDetailsBadRequest struct {
}

func (o *GetL3TopologyDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsBadRequest ", 400)
}

func (o *GetL3TopologyDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsUnauthorized creates a GetL3TopologyDetailsUnauthorized with default headers values
func NewGetL3TopologyDetailsUnauthorized() *GetL3TopologyDetailsUnauthorized {
	return &GetL3TopologyDetailsUnauthorized{}
}

/*GetL3TopologyDetailsUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetL3TopologyDetailsUnauthorized struct {
}

func (o *GetL3TopologyDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsUnauthorized ", 401)
}

func (o *GetL3TopologyDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsForbidden creates a GetL3TopologyDetailsForbidden with default headers values
func NewGetL3TopologyDetailsForbidden() *GetL3TopologyDetailsForbidden {
	return &GetL3TopologyDetailsForbidden{}
}

/*GetL3TopologyDetailsForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type GetL3TopologyDetailsForbidden struct {
}

func (o *GetL3TopologyDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsForbidden ", 403)
}

func (o *GetL3TopologyDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsNotFound creates a GetL3TopologyDetailsNotFound with default headers values
func NewGetL3TopologyDetailsNotFound() *GetL3TopologyDetailsNotFound {
	return &GetL3TopologyDetailsNotFound{}
}

/*GetL3TopologyDetailsNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetL3TopologyDetailsNotFound struct {
}

func (o *GetL3TopologyDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsNotFound ", 404)
}

func (o *GetL3TopologyDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsConflict creates a GetL3TopologyDetailsConflict with default headers values
func NewGetL3TopologyDetailsConflict() *GetL3TopologyDetailsConflict {
	return &GetL3TopologyDetailsConflict{}
}

/*GetL3TopologyDetailsConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type GetL3TopologyDetailsConflict struct {
}

func (o *GetL3TopologyDetailsConflict) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsConflict ", 409)
}

func (o *GetL3TopologyDetailsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsUnsupportedMediaType creates a GetL3TopologyDetailsUnsupportedMediaType with default headers values
func NewGetL3TopologyDetailsUnsupportedMediaType() *GetL3TopologyDetailsUnsupportedMediaType {
	return &GetL3TopologyDetailsUnsupportedMediaType{}
}

/*GetL3TopologyDetailsUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type GetL3TopologyDetailsUnsupportedMediaType struct {
}

func (o *GetL3TopologyDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsUnsupportedMediaType ", 415)
}

func (o *GetL3TopologyDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsInternalServerError creates a GetL3TopologyDetailsInternalServerError with default headers values
func NewGetL3TopologyDetailsInternalServerError() *GetL3TopologyDetailsInternalServerError {
	return &GetL3TopologyDetailsInternalServerError{}
}

/*GetL3TopologyDetailsInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetL3TopologyDetailsInternalServerError struct {
}

func (o *GetL3TopologyDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsInternalServerError ", 500)
}

func (o *GetL3TopologyDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsNotImplemented creates a GetL3TopologyDetailsNotImplemented with default headers values
func NewGetL3TopologyDetailsNotImplemented() *GetL3TopologyDetailsNotImplemented {
	return &GetL3TopologyDetailsNotImplemented{}
}

/*GetL3TopologyDetailsNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type GetL3TopologyDetailsNotImplemented struct {
}

func (o *GetL3TopologyDetailsNotImplemented) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsNotImplemented ", 501)
}

func (o *GetL3TopologyDetailsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsServiceUnavailable creates a GetL3TopologyDetailsServiceUnavailable with default headers values
func NewGetL3TopologyDetailsServiceUnavailable() *GetL3TopologyDetailsServiceUnavailable {
	return &GetL3TopologyDetailsServiceUnavailable{}
}

/*GetL3TopologyDetailsServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type GetL3TopologyDetailsServiceUnavailable struct {
}

func (o *GetL3TopologyDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsServiceUnavailable ", 503)
}

func (o *GetL3TopologyDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetL3TopologyDetailsGatewayTimeout creates a GetL3TopologyDetailsGatewayTimeout with default headers values
func NewGetL3TopologyDetailsGatewayTimeout() *GetL3TopologyDetailsGatewayTimeout {
	return &GetL3TopologyDetailsGatewayTimeout{}
}

/*GetL3TopologyDetailsGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type GetL3TopologyDetailsGatewayTimeout struct {
}

func (o *GetL3TopologyDetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/topology/l3/{topologyType}][%d] getL3TopologyDetailsGatewayTimeout ", 504)
}

func (o *GetL3TopologyDetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
