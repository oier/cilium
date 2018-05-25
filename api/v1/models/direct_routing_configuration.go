// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DirectRoutingConfiguration Direct routing configuration of a node
// swagger:model DirectRoutingConfiguration

type DirectRoutingConfiguration struct {

	// Indicates that this node installs and uses direct routes to reach
	// endpoints on other cluster nodes.
	//
	InstallRoutes bool `json:"install-routes,omitempty"`
}

/* polymorph DirectRoutingConfiguration install-routes false */

// Validate validates this direct routing configuration
func (m *DirectRoutingConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DirectRoutingConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectRoutingConfiguration) UnmarshalBinary(b []byte) error {
	var res DirectRoutingConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
