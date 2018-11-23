// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// GetListOfFilesReader is a Reader for the GetListOfFiles structure.
type GetListOfFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListOfFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetListOfFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetListOfFilesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewGetListOfFilesPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetListOfFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetListOfFilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetListOfFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetListOfFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetListOfFilesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGetListOfFilesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetListOfFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetListOfFilesNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetListOfFilesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetListOfFilesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetListOfFilesOK creates a GetListOfFilesOK with default headers values
func NewGetListOfFilesOK() *GetListOfFilesOK {
	return &GetListOfFilesOK{}
}

/*GetListOfFilesOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetListOfFilesOK struct {
	Payload *models.FileObjectListResult
}

func (o *GetListOfFilesOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesOK  %+v", 200, o.Payload)
}

func (o *GetListOfFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileObjectListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListOfFilesNoContent creates a GetListOfFilesNoContent with default headers values
func NewGetListOfFilesNoContent() *GetListOfFilesNoContent {
	return &GetListOfFilesNoContent{}
}

/*GetListOfFilesNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type GetListOfFilesNoContent struct {
}

func (o *GetListOfFilesNoContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesNoContent ", 204)
}

func (o *GetListOfFilesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesPartialContent creates a GetListOfFilesPartialContent with default headers values
func NewGetListOfFilesPartialContent() *GetListOfFilesPartialContent {
	return &GetListOfFilesPartialContent{}
}

/*GetListOfFilesPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type GetListOfFilesPartialContent struct {
}

func (o *GetListOfFilesPartialContent) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesPartialContent ", 206)
}

func (o *GetListOfFilesPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesBadRequest creates a GetListOfFilesBadRequest with default headers values
func NewGetListOfFilesBadRequest() *GetListOfFilesBadRequest {
	return &GetListOfFilesBadRequest{}
}

/*GetListOfFilesBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetListOfFilesBadRequest struct {
}

func (o *GetListOfFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesBadRequest ", 400)
}

func (o *GetListOfFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesUnauthorized creates a GetListOfFilesUnauthorized with default headers values
func NewGetListOfFilesUnauthorized() *GetListOfFilesUnauthorized {
	return &GetListOfFilesUnauthorized{}
}

/*GetListOfFilesUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetListOfFilesUnauthorized struct {
}

func (o *GetListOfFilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesUnauthorized ", 401)
}

func (o *GetListOfFilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesForbidden creates a GetListOfFilesForbidden with default headers values
func NewGetListOfFilesForbidden() *GetListOfFilesForbidden {
	return &GetListOfFilesForbidden{}
}

/*GetListOfFilesForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type GetListOfFilesForbidden struct {
}

func (o *GetListOfFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesForbidden ", 403)
}

func (o *GetListOfFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesNotFound creates a GetListOfFilesNotFound with default headers values
func NewGetListOfFilesNotFound() *GetListOfFilesNotFound {
	return &GetListOfFilesNotFound{}
}

/*GetListOfFilesNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetListOfFilesNotFound struct {
}

func (o *GetListOfFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesNotFound ", 404)
}

func (o *GetListOfFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesConflict creates a GetListOfFilesConflict with default headers values
func NewGetListOfFilesConflict() *GetListOfFilesConflict {
	return &GetListOfFilesConflict{}
}

/*GetListOfFilesConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type GetListOfFilesConflict struct {
}

func (o *GetListOfFilesConflict) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesConflict ", 409)
}

func (o *GetListOfFilesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesUnsupportedMediaType creates a GetListOfFilesUnsupportedMediaType with default headers values
func NewGetListOfFilesUnsupportedMediaType() *GetListOfFilesUnsupportedMediaType {
	return &GetListOfFilesUnsupportedMediaType{}
}

/*GetListOfFilesUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type GetListOfFilesUnsupportedMediaType struct {
}

func (o *GetListOfFilesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesUnsupportedMediaType ", 415)
}

func (o *GetListOfFilesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesInternalServerError creates a GetListOfFilesInternalServerError with default headers values
func NewGetListOfFilesInternalServerError() *GetListOfFilesInternalServerError {
	return &GetListOfFilesInternalServerError{}
}

/*GetListOfFilesInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetListOfFilesInternalServerError struct {
}

func (o *GetListOfFilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesInternalServerError ", 500)
}

func (o *GetListOfFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesNotImplemented creates a GetListOfFilesNotImplemented with default headers values
func NewGetListOfFilesNotImplemented() *GetListOfFilesNotImplemented {
	return &GetListOfFilesNotImplemented{}
}

/*GetListOfFilesNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type GetListOfFilesNotImplemented struct {
}

func (o *GetListOfFilesNotImplemented) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesNotImplemented ", 501)
}

func (o *GetListOfFilesNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesServiceUnavailable creates a GetListOfFilesServiceUnavailable with default headers values
func NewGetListOfFilesServiceUnavailable() *GetListOfFilesServiceUnavailable {
	return &GetListOfFilesServiceUnavailable{}
}

/*GetListOfFilesServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type GetListOfFilesServiceUnavailable struct {
}

func (o *GetListOfFilesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesServiceUnavailable ", 503)
}

func (o *GetListOfFilesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListOfFilesGatewayTimeout creates a GetListOfFilesGatewayTimeout with default headers values
func NewGetListOfFilesGatewayTimeout() *GetListOfFilesGatewayTimeout {
	return &GetListOfFilesGatewayTimeout{}
}

/*GetListOfFilesGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type GetListOfFilesGatewayTimeout struct {
}

func (o *GetListOfFilesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/file/namespace/{nameSpace}][%d] getListOfFilesGatewayTimeout ", 504)
}

func (o *GetListOfFilesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
