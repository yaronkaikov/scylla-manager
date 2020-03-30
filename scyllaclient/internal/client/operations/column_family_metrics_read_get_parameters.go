// Code generated by go-swagger; DO NOT EDIT.

package operations

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
)

// NewColumnFamilyMetricsReadGetParams creates a new ColumnFamilyMetricsReadGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadGetParams() *ColumnFamilyMetricsReadGetParams {

	return &ColumnFamilyMetricsReadGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadGetParamsWithTimeout creates a new ColumnFamilyMetricsReadGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadGetParams {

	return &ColumnFamilyMetricsReadGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadGetParamsWithContext creates a new ColumnFamilyMetricsReadGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadGetParams {

	return &ColumnFamilyMetricsReadGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadGetParams {

	return &ColumnFamilyMetricsReadGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsReadGetParams contains all the parameters to send to the API endpoint
for the column family metrics read get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read get params
func (o *ColumnFamilyMetricsReadGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read get params
func (o *ColumnFamilyMetricsReadGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read get params
func (o *ColumnFamilyMetricsReadGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read get params
func (o *ColumnFamilyMetricsReadGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read get params
func (o *ColumnFamilyMetricsReadGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read get params
func (o *ColumnFamilyMetricsReadGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}