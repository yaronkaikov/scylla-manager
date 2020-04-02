// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/scyllaclient/internal/agent/models"
)

// SyncCopyReader is a Reader for the SyncCopy structure.
type SyncCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncCopyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSyncCopyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncCopyOK creates a SyncCopyOK with default headers values
func NewSyncCopyOK() *SyncCopyOK {
	return &SyncCopyOK{}
}

/*SyncCopyOK handles this case with default header values.

Job ID
*/
type SyncCopyOK struct {
	Payload *models.Jobid
	JobID   int64
}

func (o *SyncCopyOK) GetPayload() *models.Jobid {
	return o.Payload
}

func (o *SyncCopyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewSyncCopyDefault creates a SyncCopyDefault with default headers values
func NewSyncCopyDefault(code int) *SyncCopyDefault {
	return &SyncCopyDefault{
		_statusCode: code,
	}
}

/*SyncCopyDefault handles this case with default header values.

Server error
*/
type SyncCopyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the sync copy default response
func (o *SyncCopyDefault) Code() int {
	return o._statusCode
}

func (o *SyncCopyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SyncCopyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *SyncCopyDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

/*SyncCopyBody sync copy body
swagger:model SyncCopyBody
*/
type SyncCopyBody struct {

	// A remote name string eg. drive: for the destination
	DstFs string `json:"dstFs,omitempty"`

	// A remote name string eg. drive: for the source
	SrcFs string `json:"srcFs,omitempty"`
}

// Validate validates this sync copy body
func (o *SyncCopyBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SyncCopyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SyncCopyBody) UnmarshalBinary(b []byte) error {
	var res SyncCopyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}