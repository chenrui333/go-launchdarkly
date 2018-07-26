// Code generated by go-swagger; DO NOT EDIT.

package audit_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/launchdarkly-api-client/models"
)

// GetAuditLogEntryReader is a Reader for the GetAuditLogEntry structure.
type GetAuditLogEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditLogEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAuditLogEntryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAuditLogEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuditLogEntryOK creates a GetAuditLogEntryOK with default headers values
func NewGetAuditLogEntryOK() *GetAuditLogEntryOK {
	return &GetAuditLogEntryOK{}
}

/*GetAuditLogEntryOK handles this case with default header values.

Audit log entry response.
*/
type GetAuditLogEntryOK struct {
	Payload *models.AuditLogEntry
}

func (o *GetAuditLogEntryOK) Error() string {
	return fmt.Sprintf("[GET /auditlog/{resourceId}][%d] getAuditLogEntryOK  %+v", 200, o.Payload)
}

func (o *GetAuditLogEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditLogEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogEntryUnauthorized creates a GetAuditLogEntryUnauthorized with default headers values
func NewGetAuditLogEntryUnauthorized() *GetAuditLogEntryUnauthorized {
	return &GetAuditLogEntryUnauthorized{}
}

/*GetAuditLogEntryUnauthorized handles this case with default header values.

Invalid access token.
*/
type GetAuditLogEntryUnauthorized struct {
}

func (o *GetAuditLogEntryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auditlog/{resourceId}][%d] getAuditLogEntryUnauthorized ", 401)
}

func (o *GetAuditLogEntryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAuditLogEntryNotFound creates a GetAuditLogEntryNotFound with default headers values
func NewGetAuditLogEntryNotFound() *GetAuditLogEntryNotFound {
	return &GetAuditLogEntryNotFound{}
}

/*GetAuditLogEntryNotFound handles this case with default header values.

Invalid resource specifier.
*/
type GetAuditLogEntryNotFound struct {
}

func (o *GetAuditLogEntryNotFound) Error() string {
	return fmt.Sprintf("[GET /auditlog/{resourceId}][%d] getAuditLogEntryNotFound ", 404)
}

func (o *GetAuditLogEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
