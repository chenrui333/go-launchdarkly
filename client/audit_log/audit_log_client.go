// Code generated by go-swagger; DO NOT EDIT.

package audit_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new audit log API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for audit log API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAuditLogEntries gets a list of all audit log entries the query parameters allow you to restrict the returned results by date ranges resource specifiers or a full text search query
*/
func (a *Client) GetAuditLogEntries(params *GetAuditLogEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAuditLogEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditLogEntriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditLogEntries",
		Method:             "GET",
		PathPattern:        "/auditlog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuditLogEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditLogEntriesOK), nil

}

/*
GetAuditLogEntry uses this endpoint to fetch a single audit log entry by its resouce ID
*/
func (a *Client) GetAuditLogEntry(params *GetAuditLogEntryParams, authInfo runtime.ClientAuthInfoWriter) (*GetAuditLogEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditLogEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuditLogEntry",
		Method:             "GET",
		PathPattern:        "/auditlog/{resourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuditLogEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAuditLogEntryOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}