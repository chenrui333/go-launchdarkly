// Code generated by go-swagger; DO NOT EDIT.

package user_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/launchdarkly-api-client/models"
)

// GetUserFlagSettingReader is a Reader for the GetUserFlagSetting structure.
type GetUserFlagSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserFlagSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserFlagSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUserFlagSettingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserFlagSettingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserFlagSettingOK creates a GetUserFlagSettingOK with default headers values
func NewGetUserFlagSettingOK() *GetUserFlagSettingOK {
	return &GetUserFlagSettingOK{}
}

/*GetUserFlagSettingOK handles this case with default header values.

User flag setting response.
*/
type GetUserFlagSettingOK struct {
	Payload *models.UserFlagSetting
}

func (o *GetUserFlagSettingOK) Error() string {
	return fmt.Sprintf("[GET /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey}][%d] getUserFlagSettingOK  %+v", 200, o.Payload)
}

func (o *GetUserFlagSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserFlagSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFlagSettingUnauthorized creates a GetUserFlagSettingUnauthorized with default headers values
func NewGetUserFlagSettingUnauthorized() *GetUserFlagSettingUnauthorized {
	return &GetUserFlagSettingUnauthorized{}
}

/*GetUserFlagSettingUnauthorized handles this case with default header values.

Invalid access token.
*/
type GetUserFlagSettingUnauthorized struct {
}

func (o *GetUserFlagSettingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey}][%d] getUserFlagSettingUnauthorized ", 401)
}

func (o *GetUserFlagSettingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserFlagSettingNotFound creates a GetUserFlagSettingNotFound with default headers values
func NewGetUserFlagSettingNotFound() *GetUserFlagSettingNotFound {
	return &GetUserFlagSettingNotFound{}
}

/*GetUserFlagSettingNotFound handles this case with default header values.

Invalid resource specifier.
*/
type GetUserFlagSettingNotFound struct {
}

func (o *GetUserFlagSettingNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey}][%d] getUserFlagSettingNotFound ", 404)
}

func (o *GetUserFlagSettingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
