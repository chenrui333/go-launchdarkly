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

// GetCustomRolesReader is a Reader for the GetCustomRoles structure.
type GetCustomRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCustomRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomRolesOK creates a GetCustomRolesOK with default headers values
func NewGetCustomRolesOK() *GetCustomRolesOK {
	return &GetCustomRolesOK{}
}

/*GetCustomRolesOK handles this case with default header values.

Custom roles response.
*/
type GetCustomRolesOK struct {
	Payload *models.CustomRoles
}

func (o *GetCustomRolesOK) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getCustomRolesOK  %+v", 200, o.Payload)
}

func (o *GetCustomRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomRoles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomRolesUnauthorized creates a GetCustomRolesUnauthorized with default headers values
func NewGetCustomRolesUnauthorized() *GetCustomRolesUnauthorized {
	return &GetCustomRolesUnauthorized{}
}

/*GetCustomRolesUnauthorized handles this case with default header values.

Invalid access token.
*/
type GetCustomRolesUnauthorized struct {
}

func (o *GetCustomRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getCustomRolesUnauthorized ", 401)
}

func (o *GetCustomRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}