// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/launchdarkly-api-client/models"
)

// GetFeatureFlagsReader is a Reader for the GetFeatureFlags structure.
type GetFeatureFlagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeatureFlagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFeatureFlagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetFeatureFlagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeatureFlagsOK creates a GetFeatureFlagsOK with default headers values
func NewGetFeatureFlagsOK() *GetFeatureFlagsOK {
	return &GetFeatureFlagsOK{}
}

/*GetFeatureFlagsOK handles this case with default header values.

Flags response.
*/
type GetFeatureFlagsOK struct {
	Payload *models.FeatureFlags
}

func (o *GetFeatureFlagsOK) Error() string {
	return fmt.Sprintf("[GET /flags/{projectKey}][%d] getFeatureFlagsOK  %+v", 200, o.Payload)
}

func (o *GetFeatureFlagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeatureFlags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeatureFlagsUnauthorized creates a GetFeatureFlagsUnauthorized with default headers values
func NewGetFeatureFlagsUnauthorized() *GetFeatureFlagsUnauthorized {
	return &GetFeatureFlagsUnauthorized{}
}

/*GetFeatureFlagsUnauthorized handles this case with default header values.

Invalid access token.
*/
type GetFeatureFlagsUnauthorized struct {
}

func (o *GetFeatureFlagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flags/{projectKey}][%d] getFeatureFlagsUnauthorized ", 401)
}

func (o *GetFeatureFlagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}