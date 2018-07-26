// Code generated by go-swagger; DO NOT EDIT.

package custom_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/launchdarkly-api-client/models"
)

// GetCustomRoleReader is a Reader for the GetCustomRole structure.
type GetCustomRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCustomRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCustomRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomRoleOK creates a GetCustomRoleOK with default headers values
func NewGetCustomRoleOK() *GetCustomRoleOK {
	return &GetCustomRoleOK{}
}

/*GetCustomRoleOK handles this case with default header values.

Custom role response.
*/
type GetCustomRoleOK struct {
	Payload *models.CustomRole
}

func (o *GetCustomRoleOK) Error() string {
	return fmt.Sprintf("[GET /roles/{customRoleKey}][%d] getCustomRoleOK  %+v", 200, o.Payload)
}

func (o *GetCustomRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomRoleBadRequest creates a GetCustomRoleBadRequest with default headers values
func NewGetCustomRoleBadRequest() *GetCustomRoleBadRequest {
	return &GetCustomRoleBadRequest{}
}

/*GetCustomRoleBadRequest handles this case with default header values.

Invalid request body.
*/
type GetCustomRoleBadRequest struct {
}

func (o *GetCustomRoleBadRequest) Error() string {
	return fmt.Sprintf("[GET /roles/{customRoleKey}][%d] getCustomRoleBadRequest ", 400)
}

func (o *GetCustomRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomRoleUnauthorized creates a GetCustomRoleUnauthorized with default headers values
func NewGetCustomRoleUnauthorized() *GetCustomRoleUnauthorized {
	return &GetCustomRoleUnauthorized{}
}

/*GetCustomRoleUnauthorized handles this case with default header values.

Invalid access token.
*/
type GetCustomRoleUnauthorized struct {
}

func (o *GetCustomRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /roles/{customRoleKey}][%d] getCustomRoleUnauthorized ", 401)
}

func (o *GetCustomRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
