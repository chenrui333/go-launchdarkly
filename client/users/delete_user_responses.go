// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/*DeleteUserNoContent handles this case with default header values.

Action completed successfully.
*/
type DeleteUserNoContent struct {
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{projectKey}/{environmentKey}/{userKey}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUnauthorized creates a DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {
	return &DeleteUserUnauthorized{}
}

/*DeleteUserUnauthorized handles this case with default header values.

Invalid access token.
*/
type DeleteUserUnauthorized struct {
}

func (o *DeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{projectKey}/{environmentKey}/{userKey}][%d] deleteUserUnauthorized ", 401)
}

func (o *DeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/*DeleteUserNotFound handles this case with default header values.

Invalid resource specifier.
*/
type DeleteUserNotFound struct {
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{projectKey}/{environmentKey}/{userKey}][%d] deleteUserNotFound ", 404)
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
