// Code generated by go-swagger; DO NOT EDIT.

package custom_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new custom roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCustomRole deletes a custom role by key
*/
func (a *Client) DeleteCustomRole(params *DeleteCustomRoleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCustomRoleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCustomRole",
		Method:             "DELETE",
		PathPattern:        "/roles/{customRoleKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCustomRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCustomRoleNoContent), nil

}

/*
GetCustomRole gets one custom role by key
*/
func (a *Client) GetCustomRole(params *GetCustomRoleParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomRole",
		Method:             "GET",
		PathPattern:        "/roles/{customRoleKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomRoleOK), nil

}

/*
GetCustomRoles returns a complete list of custom roles
*/
func (a *Client) GetCustomRoles(params *GetCustomRolesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomRoles",
		Method:             "GET",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomRolesOK), nil

}

/*
PatchCustomRole modifies a custom role by key
*/
func (a *Client) PatchCustomRole(params *PatchCustomRoleParams, authInfo runtime.ClientAuthInfoWriter) (*PatchCustomRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCustomRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCustomRole",
		Method:             "PATCH",
		PathPattern:        "/roles/{customRoleKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCustomRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCustomRoleOK), nil

}

/*
PostCustomRole creates a new custom role
*/
func (a *Client) PostCustomRole(params *PostCustomRoleParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomRoleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCustomRole",
		Method:             "POST",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCustomRoleCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
