// Code generated by go-swagger; DO NOT EDIT.

package site_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kszafran/junction-2018/models"
)

// ProvisionNFVReader is a Reader for the ProvisionNFV structure.
type ProvisionNFVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionNFVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProvisionNFVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProvisionNFVBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewProvisionNFVUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProvisionNFVNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewProvisionNFVInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProvisionNFVOK creates a ProvisionNFVOK with default headers values
func NewProvisionNFVOK() *ProvisionNFVOK {
	return &ProvisionNFVOK{}
}

/*ProvisionNFVOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type ProvisionNFVOK struct {
	Payload *models.ProvisionNFVResponse
}

func (o *ProvisionNFVOK) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/provision-nfv][%d] provisionNFVOK  %+v", 200, o.Payload)
}

func (o *ProvisionNFVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvisionNFVResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionNFVBadRequest creates a ProvisionNFVBadRequest with default headers values
func NewProvisionNFVBadRequest() *ProvisionNFVBadRequest {
	return &ProvisionNFVBadRequest{}
}

/*ProvisionNFVBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type ProvisionNFVBadRequest struct {
}

func (o *ProvisionNFVBadRequest) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/provision-nfv][%d] provisionNFVBadRequest ", 400)
}

func (o *ProvisionNFVBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProvisionNFVUnauthorized creates a ProvisionNFVUnauthorized with default headers values
func NewProvisionNFVUnauthorized() *ProvisionNFVUnauthorized {
	return &ProvisionNFVUnauthorized{}
}

/*ProvisionNFVUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type ProvisionNFVUnauthorized struct {
}

func (o *ProvisionNFVUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/provision-nfv][%d] provisionNFVUnauthorized ", 401)
}

func (o *ProvisionNFVUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProvisionNFVNotFound creates a ProvisionNFVNotFound with default headers values
func NewProvisionNFVNotFound() *ProvisionNFVNotFound {
	return &ProvisionNFVNotFound{}
}

/*ProvisionNFVNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type ProvisionNFVNotFound struct {
}

func (o *ProvisionNFVNotFound) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/provision-nfv][%d] provisionNFVNotFound ", 404)
}

func (o *ProvisionNFVNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProvisionNFVInternalServerError creates a ProvisionNFVInternalServerError with default headers values
func NewProvisionNFVInternalServerError() *ProvisionNFVInternalServerError {
	return &ProvisionNFVInternalServerError{}
}

/*ProvisionNFVInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type ProvisionNFVInternalServerError struct {
}

func (o *ProvisionNFVInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dna/intent/api/v1/provision-nfv][%d] provisionNFVInternalServerError ", 500)
}

func (o *ProvisionNFVInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
