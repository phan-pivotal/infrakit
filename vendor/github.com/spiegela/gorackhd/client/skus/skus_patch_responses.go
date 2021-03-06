package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// SkusPatchReader is a Reader for the SkusPatch structure.
type SkusPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SkusPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSkusPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSkusPatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSkusPatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSkusPatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSkusPatchOK creates a SkusPatchOK with default headers values
func NewSkusPatchOK() *SkusPatchOK {
	return &SkusPatchOK{}
}

/*SkusPatchOK handles this case with default header values.

Successfully modified the specified SKU
*/
type SkusPatchOK struct {
	Payload SkusPatchOKBody
}

func (o *SkusPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] skusPatchOK  %+v", 200, o.Payload)
}

func (o *SkusPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusPatchNotFound creates a SkusPatchNotFound with default headers values
func NewSkusPatchNotFound() *SkusPatchNotFound {
	return &SkusPatchNotFound{}
}

/*SkusPatchNotFound handles this case with default header values.

The SKU with the specified identifier was not found
*/
type SkusPatchNotFound struct {
	Payload *models.Error
}

func (o *SkusPatchNotFound) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] skusPatchNotFound  %+v", 404, o.Payload)
}

func (o *SkusPatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusPatchInternalServerError creates a SkusPatchInternalServerError with default headers values
func NewSkusPatchInternalServerError() *SkusPatchInternalServerError {
	return &SkusPatchInternalServerError{}
}

/*SkusPatchInternalServerError handles this case with default header values.

SKU patch failed
*/
type SkusPatchInternalServerError struct {
	Payload *models.Error
}

func (o *SkusPatchInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] skusPatchInternalServerError  %+v", 500, o.Payload)
}

func (o *SkusPatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusPatchDefault creates a SkusPatchDefault with default headers values
func NewSkusPatchDefault(code int) *SkusPatchDefault {
	return &SkusPatchDefault{
		_statusCode: code,
	}
}

/*SkusPatchDefault handles this case with default header values.

Unexpected error
*/
type SkusPatchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the skus patch default response
func (o *SkusPatchDefault) Code() int {
	return o._statusCode
}

func (o *SkusPatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] skusPatch default  %+v", o._statusCode, o.Payload)
}

func (o *SkusPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SkusPatchOKBody skus patch o k body
swagger:model SkusPatchOKBody
*/
type SkusPatchOKBody interface{}
