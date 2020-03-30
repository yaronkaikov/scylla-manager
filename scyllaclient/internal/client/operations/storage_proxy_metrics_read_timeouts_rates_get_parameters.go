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

// NewStorageProxyMetricsReadTimeoutsRatesGetParams creates a new StorageProxyMetricsReadTimeoutsRatesGetParams object
// with the default values initialized.
func NewStorageProxyMetricsReadTimeoutsRatesGetParams() *StorageProxyMetricsReadTimeoutsRatesGetParams {

	return &StorageProxyMetricsReadTimeoutsRatesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsReadTimeoutsRatesGetParamsWithTimeout creates a new StorageProxyMetricsReadTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsReadTimeoutsRatesGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsReadTimeoutsRatesGetParams {

	return &StorageProxyMetricsReadTimeoutsRatesGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsReadTimeoutsRatesGetParamsWithContext creates a new StorageProxyMetricsReadTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsReadTimeoutsRatesGetParamsWithContext(ctx context.Context) *StorageProxyMetricsReadTimeoutsRatesGetParams {

	return &StorageProxyMetricsReadTimeoutsRatesGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsReadTimeoutsRatesGetParamsWithHTTPClient creates a new StorageProxyMetricsReadTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsReadTimeoutsRatesGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsReadTimeoutsRatesGetParams {

	return &StorageProxyMetricsReadTimeoutsRatesGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsReadTimeoutsRatesGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics read timeouts rates get operation typically these are written to a http.Request
*/
type StorageProxyMetricsReadTimeoutsRatesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics read timeouts rates get params
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsReadTimeoutsRatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics read timeouts rates get params
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics read timeouts rates get params
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) WithContext(ctx context.Context) *StorageProxyMetricsReadTimeoutsRatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics read timeouts rates get params
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics read timeouts rates get params
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsReadTimeoutsRatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics read timeouts rates get params
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsReadTimeoutsRatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}