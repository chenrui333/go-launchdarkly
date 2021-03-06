// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostEnvironmentReader is a Reader for the PostEnvironment structure.
type PostEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostEnvironmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostEnvironmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostEnvironmentCreated creates a PostEnvironmentCreated with default headers values
func NewPostEnvironmentCreated() *PostEnvironmentCreated {
	return &PostEnvironmentCreated{}
}

/*PostEnvironmentCreated handles this case with default header values.

Resource created.
*/
type PostEnvironmentCreated struct {
}

func (o *PostEnvironmentCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{projectKey}/environments][%d] postEnvironmentCreated ", 201)
}

func (o *PostEnvironmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEnvironmentBadRequest creates a PostEnvironmentBadRequest with default headers values
func NewPostEnvironmentBadRequest() *PostEnvironmentBadRequest {
	return &PostEnvironmentBadRequest{}
}

/*PostEnvironmentBadRequest handles this case with default header values.

Invalid request body.
*/
type PostEnvironmentBadRequest struct {
}

func (o *PostEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{projectKey}/environments][%d] postEnvironmentBadRequest ", 400)
}

func (o *PostEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEnvironmentUnauthorized creates a PostEnvironmentUnauthorized with default headers values
func NewPostEnvironmentUnauthorized() *PostEnvironmentUnauthorized {
	return &PostEnvironmentUnauthorized{}
}

/*PostEnvironmentUnauthorized handles this case with default header values.

Invalid access token.
*/
type PostEnvironmentUnauthorized struct {
}

func (o *PostEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{projectKey}/environments][%d] postEnvironmentUnauthorized ", 401)
}

func (o *PostEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEnvironmentConflict creates a PostEnvironmentConflict with default headers values
func NewPostEnvironmentConflict() *PostEnvironmentConflict {
	return &PostEnvironmentConflict{}
}

/*PostEnvironmentConflict handles this case with default header values.

Status conflict.
*/
type PostEnvironmentConflict struct {
}

func (o *PostEnvironmentConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{projectKey}/environments][%d] postEnvironmentConflict ", 409)
}

func (o *PostEnvironmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
