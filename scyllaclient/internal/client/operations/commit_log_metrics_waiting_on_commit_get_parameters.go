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

// NewCommitLogMetricsWaitingOnCommitGetParams creates a new CommitLogMetricsWaitingOnCommitGetParams object
// with the default values initialized.
func NewCommitLogMetricsWaitingOnCommitGetParams() *CommitLogMetricsWaitingOnCommitGetParams {

	return &CommitLogMetricsWaitingOnCommitGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCommitLogMetricsWaitingOnCommitGetParamsWithTimeout creates a new CommitLogMetricsWaitingOnCommitGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCommitLogMetricsWaitingOnCommitGetParamsWithTimeout(timeout time.Duration) *CommitLogMetricsWaitingOnCommitGetParams {

	return &CommitLogMetricsWaitingOnCommitGetParams{

		timeout: timeout,
	}
}

// NewCommitLogMetricsWaitingOnCommitGetParamsWithContext creates a new CommitLogMetricsWaitingOnCommitGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCommitLogMetricsWaitingOnCommitGetParamsWithContext(ctx context.Context) *CommitLogMetricsWaitingOnCommitGetParams {

	return &CommitLogMetricsWaitingOnCommitGetParams{

		Context: ctx,
	}
}

// NewCommitLogMetricsWaitingOnCommitGetParamsWithHTTPClient creates a new CommitLogMetricsWaitingOnCommitGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCommitLogMetricsWaitingOnCommitGetParamsWithHTTPClient(client *http.Client) *CommitLogMetricsWaitingOnCommitGetParams {

	return &CommitLogMetricsWaitingOnCommitGetParams{
		HTTPClient: client,
	}
}

/*CommitLogMetricsWaitingOnCommitGetParams contains all the parameters to send to the API endpoint
for the commit log metrics waiting on commit get operation typically these are written to a http.Request
*/
type CommitLogMetricsWaitingOnCommitGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the commit log metrics waiting on commit get params
func (o *CommitLogMetricsWaitingOnCommitGetParams) WithTimeout(timeout time.Duration) *CommitLogMetricsWaitingOnCommitGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commit log metrics waiting on commit get params
func (o *CommitLogMetricsWaitingOnCommitGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commit log metrics waiting on commit get params
func (o *CommitLogMetricsWaitingOnCommitGetParams) WithContext(ctx context.Context) *CommitLogMetricsWaitingOnCommitGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commit log metrics waiting on commit get params
func (o *CommitLogMetricsWaitingOnCommitGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commit log metrics waiting on commit get params
func (o *CommitLogMetricsWaitingOnCommitGetParams) WithHTTPClient(client *http.Client) *CommitLogMetricsWaitingOnCommitGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commit log metrics waiting on commit get params
func (o *CommitLogMetricsWaitingOnCommitGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CommitLogMetricsWaitingOnCommitGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}