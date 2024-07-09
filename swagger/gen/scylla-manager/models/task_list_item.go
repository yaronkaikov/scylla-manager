// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskListItem task list item
//
// swagger:model TaskListItem
type TaskListItem struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// error count
	ErrorCount int64 `json:"error_count,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// last error
	// Format: date-time
	LastError *strfmt.DateTime `json:"last_error,omitempty"`

	// last success
	// Format: date-time
	LastSuccess *strfmt.DateTime `json:"last_success,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// next activation
	// Format: date-time
	NextActivation *strfmt.DateTime `json:"next_activation,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`

	// retry
	Retry int64 `json:"retry,omitempty"`

	// schedule
	Schedule *Schedule `json:"schedule,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// success count
	SuccessCount int64 `json:"success_count,omitempty"`

	// suspended
	Suspended bool `json:"suspended,omitempty"`

	// This field is DEPRECATED. Use labels instead.
	Tags []string `json:"tags"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this task list item
func (m *TaskListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextActivation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskListItem) validateLastError(formats strfmt.Registry) error {

	if swag.IsZero(m.LastError) { // not required
		return nil
	}

	if err := validate.FormatOf("last_error", "body", "date-time", m.LastError.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskListItem) validateLastSuccess(formats strfmt.Registry) error {

	if swag.IsZero(m.LastSuccess) { // not required
		return nil
	}

	if err := validate.FormatOf("last_success", "body", "date-time", m.LastSuccess.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskListItem) validateNextActivation(formats strfmt.Registry) error {

	if swag.IsZero(m.NextActivation) { // not required
		return nil
	}

	if err := validate.FormatOf("next_activation", "body", "date-time", m.NextActivation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskListItem) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskListItem) UnmarshalBinary(b []byte) error {
	var res TaskListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
