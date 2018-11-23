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

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 206:
		result := NewDeleteProjectPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewDeleteProjectUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewDeleteProjectNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteProjectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewDeleteProjectGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/*DeleteProjectOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type DeleteProjectOK struct {
	Payload *models.TaskIDResult
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskIDResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectNoContent creates a DeleteProjectNoContent with default headers values
func NewDeleteProjectNoContent() *DeleteProjectNoContent {
	return &DeleteProjectNoContent{}
}

/*DeleteProjectNoContent handles this case with default header values.

The request was successful, however no content was returned.
*/
type DeleteProjectNoContent struct {
}

func (o *DeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectNoContent ", 204)
}

func (o *DeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectPartialContent creates a DeleteProjectPartialContent with default headers values
func NewDeleteProjectPartialContent() *DeleteProjectPartialContent {
	return &DeleteProjectPartialContent{}
}

/*DeleteProjectPartialContent handles this case with default header values.

The GET request included a Range Header, and the server responded with the partial content matching the range.
*/
type DeleteProjectPartialContent struct {
}

func (o *DeleteProjectPartialContent) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectPartialContent ", 206)
}

func (o *DeleteProjectPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectBadRequest creates a DeleteProjectBadRequest with default headers values
func NewDeleteProjectBadRequest() *DeleteProjectBadRequest {
	return &DeleteProjectBadRequest{}
}

/*DeleteProjectBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type DeleteProjectBadRequest struct {
}

func (o *DeleteProjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectBadRequest ", 400)
}

func (o *DeleteProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUnauthorized creates a DeleteProjectUnauthorized with default headers values
func NewDeleteProjectUnauthorized() *DeleteProjectUnauthorized {
	return &DeleteProjectUnauthorized{}
}

/*DeleteProjectUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type DeleteProjectUnauthorized struct {
}

func (o *DeleteProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectUnauthorized ", 401)
}

func (o *DeleteProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*DeleteProjectForbidden handles this case with default header values.

The server recognizes the authentication credentials, but the client is not authorized to perform this request.
*/
type DeleteProjectForbidden struct {
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectForbidden ", 403)
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/*DeleteProjectNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type DeleteProjectNotFound struct {
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectNotFound ", 404)
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectConflict creates a DeleteProjectConflict with default headers values
func NewDeleteProjectConflict() *DeleteProjectConflict {
	return &DeleteProjectConflict{}
}

/*DeleteProjectConflict handles this case with default header values.

The target resource is in a conflicted state (for example, an edit conflict where a resource is being edited by multiple users). Retrying the request later might succeed.
*/
type DeleteProjectConflict struct {
}

func (o *DeleteProjectConflict) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectConflict ", 409)
}

func (o *DeleteProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUnsupportedMediaType creates a DeleteProjectUnsupportedMediaType with default headers values
func NewDeleteProjectUnsupportedMediaType() *DeleteProjectUnsupportedMediaType {
	return &DeleteProjectUnsupportedMediaType{}
}

/*DeleteProjectUnsupportedMediaType handles this case with default header values.

The client sent a request body in a format that the server does not support (for example, XML to a server that only accepts JSON).
*/
type DeleteProjectUnsupportedMediaType struct {
}

func (o *DeleteProjectUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectUnsupportedMediaType ", 415)
}

func (o *DeleteProjectUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectInternalServerError creates a DeleteProjectInternalServerError with default headers values
func NewDeleteProjectInternalServerError() *DeleteProjectInternalServerError {
	return &DeleteProjectInternalServerError{}
}

/*DeleteProjectInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type DeleteProjectInternalServerError struct {
}

func (o *DeleteProjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectInternalServerError ", 500)
}

func (o *DeleteProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectNotImplemented creates a DeleteProjectNotImplemented with default headers values
func NewDeleteProjectNotImplemented() *DeleteProjectNotImplemented {
	return &DeleteProjectNotImplemented{}
}

/*DeleteProjectNotImplemented handles this case with default header values.

The server has not implemented the functionality required to fulfill the request.
*/
type DeleteProjectNotImplemented struct {
}

func (o *DeleteProjectNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectNotImplemented ", 501)
}

func (o *DeleteProjectNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectServiceUnavailable creates a DeleteProjectServiceUnavailable with default headers values
func NewDeleteProjectServiceUnavailable() *DeleteProjectServiceUnavailable {
	return &DeleteProjectServiceUnavailable{}
}

/*DeleteProjectServiceUnavailable handles this case with default header values.

The server is (temporarily) unavailable.
*/
type DeleteProjectServiceUnavailable struct {
}

func (o *DeleteProjectServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectServiceUnavailable ", 503)
}

func (o *DeleteProjectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectGatewayTimeout creates a DeleteProjectGatewayTimeout with default headers values
func NewDeleteProjectGatewayTimeout() *DeleteProjectGatewayTimeout {
	return &DeleteProjectGatewayTimeout{}
}

/*DeleteProjectGatewayTimeout handles this case with default header values.

The server did not respond inside time restrictions and timed-out.
*/
type DeleteProjectGatewayTimeout struct {
}

func (o *DeleteProjectGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /dna/intent/api/v1/template-programmer/project/{projectId}][%d] deleteProjectGatewayTimeout ", 504)
}

func (o *DeleteProjectGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
