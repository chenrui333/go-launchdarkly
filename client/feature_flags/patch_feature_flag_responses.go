// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/launchdarkly-api-client/models"
)

// PatchFeatureFlagReader is a Reader for the PatchFeatureFlag structure.
type PatchFeatureFlagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchFeatureFlagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchFeatureFlagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchFeatureFlagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchFeatureFlagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchFeatureFlagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchFeatureFlagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchFeatureFlagOK creates a PatchFeatureFlagOK with default headers values
func NewPatchFeatureFlagOK() *PatchFeatureFlagOK {
	return &PatchFeatureFlagOK{}
}

/*PatchFeatureFlagOK handles this case with default header values.

Feature flag response.
*/
type PatchFeatureFlagOK struct {
	Payload *models.FeatureFlag
}

func (o *PatchFeatureFlagOK) Error() string {
	return fmt.Sprintf("[PATCH /flags/{projectKey}/{featureFlagKey}][%d] patchFeatureFlagOK  %+v", 200, o.Payload)
}

func (o *PatchFeatureFlagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeatureFlag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchFeatureFlagBadRequest creates a PatchFeatureFlagBadRequest with default headers values
func NewPatchFeatureFlagBadRequest() *PatchFeatureFlagBadRequest {
	return &PatchFeatureFlagBadRequest{}
}

/*PatchFeatureFlagBadRequest handles this case with default header values.

Invalid request body.
*/
type PatchFeatureFlagBadRequest struct {
}

func (o *PatchFeatureFlagBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /flags/{projectKey}/{featureFlagKey}][%d] patchFeatureFlagBadRequest ", 400)
}

func (o *PatchFeatureFlagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchFeatureFlagUnauthorized creates a PatchFeatureFlagUnauthorized with default headers values
func NewPatchFeatureFlagUnauthorized() *PatchFeatureFlagUnauthorized {
	return &PatchFeatureFlagUnauthorized{}
}

/*PatchFeatureFlagUnauthorized handles this case with default header values.

Invalid access token.
*/
type PatchFeatureFlagUnauthorized struct {
}

func (o *PatchFeatureFlagUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /flags/{projectKey}/{featureFlagKey}][%d] patchFeatureFlagUnauthorized ", 401)
}

func (o *PatchFeatureFlagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchFeatureFlagNotFound creates a PatchFeatureFlagNotFound with default headers values
func NewPatchFeatureFlagNotFound() *PatchFeatureFlagNotFound {
	return &PatchFeatureFlagNotFound{}
}

/*PatchFeatureFlagNotFound handles this case with default header values.

Invalid resource specifier.
*/
type PatchFeatureFlagNotFound struct {
}

func (o *PatchFeatureFlagNotFound) Error() string {
	return fmt.Sprintf("[PATCH /flags/{projectKey}/{featureFlagKey}][%d] patchFeatureFlagNotFound ", 404)
}

func (o *PatchFeatureFlagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchFeatureFlagConflict creates a PatchFeatureFlagConflict with default headers values
func NewPatchFeatureFlagConflict() *PatchFeatureFlagConflict {
	return &PatchFeatureFlagConflict{}
}

/*PatchFeatureFlagConflict handles this case with default header values.

Status conflict.
*/
type PatchFeatureFlagConflict struct {
}

func (o *PatchFeatureFlagConflict) Error() string {
	return fmt.Sprintf("[PATCH /flags/{projectKey}/{featureFlagKey}][%d] patchFeatureFlagConflict ", 409)
}

func (o *PatchFeatureFlagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PatchFeatureFlagBody patch feature flag body
swagger:model PatchFeatureFlagBody
*/
type PatchFeatureFlagBody struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// patch
	Patch []*models.PatchOperation `json:"patch"`
}

// Validate validates this patch feature flag body
func (o *PatchFeatureFlagBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchFeatureFlagBody) validatePatch(formats strfmt.Registry) error {

	if swag.IsZero(o.Patch) { // not required
		return nil
	}

	for i := 0; i < len(o.Patch); i++ {
		if swag.IsZero(o.Patch[i]) { // not required
			continue
		}

		if o.Patch[i] != nil {
			if err := o.Patch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patchComment" + "." + "patch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchFeatureFlagBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchFeatureFlagBody) UnmarshalBinary(b []byte) error {
	var res PatchFeatureFlagBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
