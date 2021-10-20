// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Deleteemployeeofficial deleteemployeeofficial
//
// swagger:model deleteemployeeofficial
type Deleteemployeeofficial struct {

	// jobtype
	// Example: Intern,Permanent,Resign,Terminate
	// Required: true
	Jobtype *string `json:"jobtype"`
}

// Validate validates this deleteemployeeofficial
func (m *Deleteemployeeofficial) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobtype(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deleteemployeeofficial) validateJobtype(formats strfmt.Registry) error {

	if err := validate.Required("jobtype", "body", m.Jobtype); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deleteemployeeofficial based on context it is used
func (m *Deleteemployeeofficial) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Deleteemployeeofficial) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deleteemployeeofficial) UnmarshalBinary(b []byte) error {
	var res Deleteemployeeofficial
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
