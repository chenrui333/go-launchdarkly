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

// PatchCustomRoleReader is a Reader for the PatchCustomRole structure.
type PatchCustomRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCustomRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchCustomRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchCustomRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchCustomRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchCustomRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchCustomRoleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchCustomRoleOK creates a PatchCustomRoleOK with default headers values
func NewPatchCustomRoleOK() *PatchCustomRoleOK {
	return &PatchCustomRoleOK{}
}

/*PatchCustomRoleOK handles this case with default header values.

Custom role response.
*/
type PatchCustomRoleOK struct {
	Payload *models.CustomRole
}

func (o *PatchCustomRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /roles/{customRoleKey}][%d] patchCustomRoleOK  %+v", 200, o.Payload)
}

func (o *PatchCustomRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCustomRoleBadRequest creates a PatchCustomRoleBadRequest with default headers values
func NewPatchCustomRoleBadRequest() *PatchCustomRoleBadRequest {
	return &PatchCustomRoleBadRequest{}
}

/*PatchCustomRoleBadRequest handles this case with default header values.

Invalid request body.
*/
type PatchCustomRoleBadRequest struct {
}

func (o *PatchCustomRoleBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /roles/{customRoleKey}][%d] patchCustomRoleBadRequest ", 400)
}

func (o *PatchCustomRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCustomRoleUnauthorized creates a PatchCustomRoleUnauthorized with default headers values
func NewPatchCustomRoleUnauthorized() *PatchCustomRoleUnauthorized {
	return &PatchCustomRoleUnauthorized{}
}

/*PatchCustomRoleUnauthorized handles this case with default header values.

Invalid access token.
*/
type PatchCustomRoleUnauthorized struct {
}

func (o *PatchCustomRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /roles/{customRoleKey}][%d] patchCustomRoleUnauthorized ", 401)
}

func (o *PatchCustomRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCustomRoleNotFound creates a PatchCustomRoleNotFound with default headers values
func NewPatchCustomRoleNotFound() *PatchCustomRoleNotFound {
	return &PatchCustomRoleNotFound{}
}

/*PatchCustomRoleNotFound handles this case with default header values.

Invalid resource specifier.
*/
type PatchCustomRoleNotFound struct {
}

func (o *PatchCustomRoleNotFound) Error() string {
	return fmt.Sprintf("[PATCH /roles/{customRoleKey}][%d] patchCustomRoleNotFound ", 404)
}

func (o *PatchCustomRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCustomRoleConflict creates a PatchCustomRoleConflict with default headers values
func NewPatchCustomRoleConflict() *PatchCustomRoleConflict {
	return &PatchCustomRoleConflict{}
}

/*PatchCustomRoleConflict handles this case with default header values.

Status conflict.
*/
type PatchCustomRoleConflict struct {
}

func (o *PatchCustomRoleConflict) Error() string {
	return fmt.Sprintf("[PATCH /roles/{customRoleKey}][%d] patchCustomRoleConflict ", 409)
}

func (o *PatchCustomRoleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
