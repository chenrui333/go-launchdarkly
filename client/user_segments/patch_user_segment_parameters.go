// Code generated by go-swagger; DO NOT EDIT.

package user_segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/go-launchdarkly/models"
)

// NewPatchUserSegmentParams creates a new PatchUserSegmentParams object
// with the default values initialized.
func NewPatchUserSegmentParams() *PatchUserSegmentParams {
	var ()
	return &PatchUserSegmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUserSegmentParamsWithTimeout creates a new PatchUserSegmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUserSegmentParamsWithTimeout(timeout time.Duration) *PatchUserSegmentParams {
	var ()
	return &PatchUserSegmentParams{

		timeout: timeout,
	}
}

// NewPatchUserSegmentParamsWithContext creates a new PatchUserSegmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUserSegmentParamsWithContext(ctx context.Context) *PatchUserSegmentParams {
	var ()
	return &PatchUserSegmentParams{

		Context: ctx,
	}
}

// NewPatchUserSegmentParamsWithHTTPClient creates a new PatchUserSegmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUserSegmentParamsWithHTTPClient(client *http.Client) *PatchUserSegmentParams {
	var ()
	return &PatchUserSegmentParams{
		HTTPClient: client,
	}
}

/*PatchUserSegmentParams contains all the parameters to send to the API endpoint
for the patch user segment operation typically these are written to a http.Request
*/
type PatchUserSegmentParams struct {

	/*PatchOnly
	  Requires a JSON Patch representation of the desired changes to the project. 'http://jsonpatch.com/' Feature flag patches also support JSON Merge Patch format. 'https://tools.ietf.org/html/rfc7386' The addition of comments is also supported.

	*/
	PatchOnly []*models.PatchOperation
	/*EnvironmentKey
	  The environment key, used to tie together flag configuration and users under one environment so they can be managed together.

	*/
	EnvironmentKey string
	/*ProjectKey
	  The project key, used to tie the flags together under one project so they can be managed together.

	*/
	ProjectKey string
	/*UserSegmentKey
	  The user segment's key. The key identifies the user segment in your code.

	*/
	UserSegmentKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch user segment params
func (o *PatchUserSegmentParams) WithTimeout(timeout time.Duration) *PatchUserSegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch user segment params
func (o *PatchUserSegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch user segment params
func (o *PatchUserSegmentParams) WithContext(ctx context.Context) *PatchUserSegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch user segment params
func (o *PatchUserSegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch user segment params
func (o *PatchUserSegmentParams) WithHTTPClient(client *http.Client) *PatchUserSegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch user segment params
func (o *PatchUserSegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchOnly adds the patchOnly to the patch user segment params
func (o *PatchUserSegmentParams) WithPatchOnly(patchOnly []*models.PatchOperation) *PatchUserSegmentParams {
	o.SetPatchOnly(patchOnly)
	return o
}

// SetPatchOnly adds the patchOnly to the patch user segment params
func (o *PatchUserSegmentParams) SetPatchOnly(patchOnly []*models.PatchOperation) {
	o.PatchOnly = patchOnly
}

// WithEnvironmentKey adds the environmentKey to the patch user segment params
func (o *PatchUserSegmentParams) WithEnvironmentKey(environmentKey string) *PatchUserSegmentParams {
	o.SetEnvironmentKey(environmentKey)
	return o
}

// SetEnvironmentKey adds the environmentKey to the patch user segment params
func (o *PatchUserSegmentParams) SetEnvironmentKey(environmentKey string) {
	o.EnvironmentKey = environmentKey
}

// WithProjectKey adds the projectKey to the patch user segment params
func (o *PatchUserSegmentParams) WithProjectKey(projectKey string) *PatchUserSegmentParams {
	o.SetProjectKey(projectKey)
	return o
}

// SetProjectKey adds the projectKey to the patch user segment params
func (o *PatchUserSegmentParams) SetProjectKey(projectKey string) {
	o.ProjectKey = projectKey
}

// WithUserSegmentKey adds the userSegmentKey to the patch user segment params
func (o *PatchUserSegmentParams) WithUserSegmentKey(userSegmentKey string) *PatchUserSegmentParams {
	o.SetUserSegmentKey(userSegmentKey)
	return o
}

// SetUserSegmentKey adds the userSegmentKey to the patch user segment params
func (o *PatchUserSegmentParams) SetUserSegmentKey(userSegmentKey string) {
	o.UserSegmentKey = userSegmentKey
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUserSegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchOnly != nil {
		if err := r.SetBodyParam(o.PatchOnly); err != nil {
			return err
		}
	}

	// path param environmentKey
	if err := r.SetPathParam("environmentKey", o.EnvironmentKey); err != nil {
		return err
	}

	// path param projectKey
	if err := r.SetPathParam("projectKey", o.ProjectKey); err != nil {
		return err
	}

	// path param userSegmentKey
	if err := r.SetPathParam("userSegmentKey", o.UserSegmentKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
