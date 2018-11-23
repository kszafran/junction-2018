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

// GetListOfDiscoveriesByDiscoveryIDReader is a Reader for the GetListOfDiscoveriesByDiscoveryID structure.
type GetListOfDiscoveriesByDiscoveryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListOfDiscoveriesByDiscoveryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetListOfDiscoveriesByDiscoveryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetListOfDiscoveriesByDiscoveryIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetListOfDiscoveriesByDiscoveryIDPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetListOfDiscoveriesByDiscoveryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetListOfDiscoveriesByDiscoveryIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetListOfDiscoveriesByDiscoveryIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetListOfDiscoveriesByDiscoveryIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetListOfDiscoveriesByDiscoveryIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetListOfDiscoveriesByDiscoveryIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetListOfDiscoveriesByDiscoveryIDNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetListOfDiscoveriesByDiscoveryIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetListOfDiscoveriesByDiscoveryIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetListOfDiscoveriesByDiscoveryIDOK creates a GetListOfDiscoveriesByDiscoveryIDOK with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDOK() *GetListOfDiscoveriesByDiscoveryIDOK {
	return &GetListOfDiscoveriesByDiscoveryIDOK{}
}

/*GetListOfDiscoveriesByDiscoveryIDOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetListOfDiscoveriesByDiscoveryIDOK struct {
	Payload *models.DiscoveryJobNIOListResult
}

func (o *GetListOfDiscoveriesByDiscoveryIDOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdOK  %+v", 200, o.Payload)
}

func (o *GetListOfDiscoveriesByDiscoveryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiscoveryJobNIOListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDNoContent creates a GetListOfDiscoveriesByDiscoveryIDNoContent with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDNoContent() *GetListOfDiscoveriesByDiscoveryIDNoContent {
	return &GetListOfDiscoveriesByDiscoveryIDNoContent{}
}

/*GetListOfDiscoveriesByDiscoveryIDNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type GetListOfDiscoveriesByDiscoveryIDNoContent struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDNoContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdNoContent ", 204)
}

func (o *GetListOfDiscoveriesByDiscoveryIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDPartialContent creates a GetListOfDiscoveriesByDiscoveryIDPartialContent with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDPartialContent() *GetListOfDiscoveriesByDiscoveryIDPartialContent {
	return &GetListOfDiscoveriesByDiscoveryIDPartialContent{}
}

/*GetListOfDiscoveriesByDiscoveryIDPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type GetListOfDiscoveriesByDiscoveryIDPartialContent struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDPartialContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdPartialContent ", 206)
}

func (o *GetListOfDiscoveriesByDiscoveryIDPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDBadRequest creates a GetListOfDiscoveriesByDiscoveryIDBadRequest with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDBadRequest() *GetListOfDiscoveriesByDiscoveryIDBadRequest {
	return &GetListOfDiscoveriesByDiscoveryIDBadRequest{}
}

/*GetListOfDiscoveriesByDiscoveryIDBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetListOfDiscoveriesByDiscoveryIDBadRequest struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdBadRequest ", 400)
}

func (o *GetListOfDiscoveriesByDiscoveryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDUnauthorized creates a GetListOfDiscoveriesByDiscoveryIDUnauthorized with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDUnauthorized() *GetListOfDiscoveriesByDiscoveryIDUnauthorized {
	return &GetListOfDiscoveriesByDiscoveryIDUnauthorized{}
}

/*GetListOfDiscoveriesByDiscoveryIDUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetListOfDiscoveriesByDiscoveryIDUnauthorized struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdUnauthorized ", 401)
}

func (o *GetListOfDiscoveriesByDiscoveryIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDForbidden creates a GetListOfDiscoveriesByDiscoveryIDForbidden with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDForbidden() *GetListOfDiscoveriesByDiscoveryIDForbidden {
	return &GetListOfDiscoveriesByDiscoveryIDForbidden{}
}

/*GetListOfDiscoveriesByDiscoveryIDForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type GetListOfDiscoveriesByDiscoveryIDForbidden struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDForbidden) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdForbidden ", 403)
}

func (o *GetListOfDiscoveriesByDiscoveryIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDNotFound creates a GetListOfDiscoveriesByDiscoveryIDNotFound with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDNotFound() *GetListOfDiscoveriesByDiscoveryIDNotFound {
	return &GetListOfDiscoveriesByDiscoveryIDNotFound{}
}

/*GetListOfDiscoveriesByDiscoveryIDNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetListOfDiscoveriesByDiscoveryIDNotFound struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdNotFound ", 404)
}

func (o *GetListOfDiscoveriesByDiscoveryIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDConflict creates a GetListOfDiscoveriesByDiscoveryIDConflict with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDConflict() *GetListOfDiscoveriesByDiscoveryIDConflict {
	return &GetListOfDiscoveriesByDiscoveryIDConflict{}
}

/*GetListOfDiscoveriesByDiscoveryIDConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type GetListOfDiscoveriesByDiscoveryIDConflict struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDConflict) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdConflict ", 409)
}

func (o *GetListOfDiscoveriesByDiscoveryIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType creates a GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType() *GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType {
	return &GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType{}
}

/*GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdUnsupportedMediaType ", 415)
}

func (o *GetListOfDiscoveriesByDiscoveryIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDInternalServerError creates a GetListOfDiscoveriesByDiscoveryIDInternalServerError with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDInternalServerError() *GetListOfDiscoveriesByDiscoveryIDInternalServerError {
	return &GetListOfDiscoveriesByDiscoveryIDInternalServerError{}
}

/*GetListOfDiscoveriesByDiscoveryIDInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetListOfDiscoveriesByDiscoveryIDInternalServerError struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdInternalServerError ", 500)
}

func (o *GetListOfDiscoveriesByDiscoveryIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDNotImplemented creates a GetListOfDiscoveriesByDiscoveryIDNotImplemented with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDNotImplemented() *GetListOfDiscoveriesByDiscoveryIDNotImplemented {
	return &GetListOfDiscoveriesByDiscoveryIDNotImplemented{}
}

/*GetListOfDiscoveriesByDiscoveryIDNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type GetListOfDiscoveriesByDiscoveryIDNotImplemented struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDNotImplemented) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdNotImplemented ", 501)
}

func (o *GetListOfDiscoveriesByDiscoveryIDNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDServiceUnavailable creates a GetListOfDiscoveriesByDiscoveryIDServiceUnavailable with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDServiceUnavailable() *GetListOfDiscoveriesByDiscoveryIDServiceUnavailable {
	return &GetListOfDiscoveriesByDiscoveryIDServiceUnavailable{}
}

/*GetListOfDiscoveriesByDiscoveryIDServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type GetListOfDiscoveriesByDiscoveryIDServiceUnavailable struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdServiceUnavailable ", 503)
}

func (o *GetListOfDiscoveriesByDiscoveryIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfDiscoveriesByDiscoveryIDGatewayTimeout creates a GetListOfDiscoveriesByDiscoveryIDGatewayTimeout with default headers values
func NewGetListOfDiscoveriesByDiscoveryIDGatewayTimeout() *GetListOfDiscoveriesByDiscoveryIDGatewayTimeout {
	return &GetListOfDiscoveriesByDiscoveryIDGatewayTimeout{}
}

/*GetListOfDiscoveriesByDiscoveryIDGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type GetListOfDiscoveriesByDiscoveryIDGatewayTimeout struct {
}

func (o *GetListOfDiscoveriesByDiscoveryIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/discovery/{id}/job][%d] getListOfDiscoveriesByDiscoveryIdGatewayTimeout ", 504)
}

func (o *GetListOfDiscoveriesByDiscoveryIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
