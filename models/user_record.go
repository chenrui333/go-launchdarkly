// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserRecord user record
// swagger:model UserRecord
type UserRecord struct {

	// avatar
	Avatar string `json:"avatar,omitempty"`

	// environment Id
	EnvironmentID string `json:"environmentId,omitempty"`

	// last ping
	LastPing int64 `json:"lastPing,omitempty"`

	// owner Id
	OwnerID ID `json:"ownerId,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this user record
func (m *UserRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRecord) validateOwnerID(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerID) { // not required
		return nil
	}

	if err := m.OwnerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ownerId")
		}
		return err
	}

	return nil
}

func (m *UserRecord) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRecord) UnmarshalBinary(b []byte) error {
	var res UserRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
