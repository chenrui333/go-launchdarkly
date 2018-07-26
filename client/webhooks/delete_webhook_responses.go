// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteWebhookReader is a Reader for the DeleteWebhook structure.
type DeleteWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteWebhookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWebhookNoContent creates a DeleteWebhookNoContent with default headers values
func NewDeleteWebhookNoContent() *DeleteWebhookNoContent {
	return &DeleteWebhookNoContent{}
}

/*DeleteWebhookNoContent handles this case with default header values.

Action completed successfully.
*/
type DeleteWebhookNoContent struct {
}

func (o *DeleteWebhookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{resourceId}][%d] deleteWebhookNoContent ", 204)
}

func (o *DeleteWebhookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebhookUnauthorized creates a DeleteWebhookUnauthorized with default headers values
func NewDeleteWebhookUnauthorized() *DeleteWebhookUnauthorized {
	return &DeleteWebhookUnauthorized{}
}

/*DeleteWebhookUnauthorized handles this case with default header values.

Invalid access token.
*/
type DeleteWebhookUnauthorized struct {
}

func (o *DeleteWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{resourceId}][%d] deleteWebhookUnauthorized ", 401)
}

func (o *DeleteWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebhookNotFound creates a DeleteWebhookNotFound with default headers values
func NewDeleteWebhookNotFound() *DeleteWebhookNotFound {
	return &DeleteWebhookNotFound{}
}

/*DeleteWebhookNotFound handles this case with default header values.

Invalid resource specifier.
*/
type DeleteWebhookNotFound struct {
}

func (o *DeleteWebhookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{resourceId}][%d] deleteWebhookNotFound ", 404)
}

func (o *DeleteWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
