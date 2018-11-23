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

// GetOverallNetworkHealthReader is a Reader for the GetOverallNetworkHealth structure.
type GetOverallNetworkHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOverallNetworkHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOverallNetworkHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOverallNetworkHealthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOverallNetworkHealthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOverallNetworkHealthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOverallNetworkHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOverallNetworkHealthOK creates a GetOverallNetworkHealthOK with default headers values
func NewGetOverallNetworkHealthOK() *GetOverallNetworkHealthOK {
	return &GetOverallNetworkHealthOK{}
}

/*GetOverallNetworkHealthOK handles this case with default header values.

The request was successful. The result is contained in the response body.
*/
type GetOverallNetworkHealthOK struct {
	Payload *models.GetOverallNetworkHealthResponse
}

func (o *GetOverallNetworkHealthOK) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-health][%d] getOverallNetworkHealthOK  %+v", 200, o.Payload)
}

func (o *GetOverallNetworkHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOverallNetworkHealthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOverallNetworkHealthBadRequest creates a GetOverallNetworkHealthBadRequest with default headers values
func NewGetOverallNetworkHealthBadRequest() *GetOverallNetworkHealthBadRequest {
	return &GetOverallNetworkHealthBadRequest{}
}

/*GetOverallNetworkHealthBadRequest handles this case with default header values.

The client made a request that the server could not understand (for example, the request syntax is incorrect).
*/
type GetOverallNetworkHealthBadRequest struct {
}

func (o *GetOverallNetworkHealthBadRequest) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-health][%d] getOverallNetworkHealthBadRequest ", 400)
}

func (o *GetOverallNetworkHealthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOverallNetworkHealthUnauthorized creates a GetOverallNetworkHealthUnauthorized with default headers values
func NewGetOverallNetworkHealthUnauthorized() *GetOverallNetworkHealthUnauthorized {
	return &GetOverallNetworkHealthUnauthorized{}
}

/*GetOverallNetworkHealthUnauthorized handles this case with default header values.

The client's authentication credentials included with the request are missing or invalid.
*/
type GetOverallNetworkHealthUnauthorized struct {
}

func (o *GetOverallNetworkHealthUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-health][%d] getOverallNetworkHealthUnauthorized ", 401)
}

func (o *GetOverallNetworkHealthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOverallNetworkHealthNotFound creates a GetOverallNetworkHealthNotFound with default headers values
func NewGetOverallNetworkHealthNotFound() *GetOverallNetworkHealthNotFound {
	return &GetOverallNetworkHealthNotFound{}
}

/*GetOverallNetworkHealthNotFound handles this case with default header values.

The client made a request for a resource that does not exist.
*/
type GetOverallNetworkHealthNotFound struct {
}

func (o *GetOverallNetworkHealthNotFound) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-health][%d] getOverallNetworkHealthNotFound ", 404)
}

func (o *GetOverallNetworkHealthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOverallNetworkHealthInternalServerError creates a GetOverallNetworkHealthInternalServerError with default headers values
func NewGetOverallNetworkHealthInternalServerError() *GetOverallNetworkHealthInternalServerError {
	return &GetOverallNetworkHealthInternalServerError{}
}

/*GetOverallNetworkHealthInternalServerError handles this case with default header values.

The server could not fulfill the request.
*/
type GetOverallNetworkHealthInternalServerError struct {
}

func (o *GetOverallNetworkHealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dna/intent/api/v1/network-health][%d] getOverallNetworkHealthInternalServerError ", 500)
}

func (o *GetOverallNetworkHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
