// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceStopDaemonPostReader is a Reader for the StorageServiceStopDaemonPost structure.
type StorageServiceStopDaemonPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceStopDaemonPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceStopDaemonPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceStopDaemonPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceStopDaemonPostOK creates a StorageServiceStopDaemonPostOK with default headers values
func NewStorageServiceStopDaemonPostOK() *StorageServiceStopDaemonPostOK {
	return &StorageServiceStopDaemonPostOK{}
}

/*StorageServiceStopDaemonPostOK handles this case with default header values.

StorageServiceStopDaemonPostOK storage service stop daemon post o k
*/
type StorageServiceStopDaemonPostOK struct {
}

func (o *StorageServiceStopDaemonPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceStopDaemonPostDefault creates a StorageServiceStopDaemonPostDefault with default headers values
func NewStorageServiceStopDaemonPostDefault(code int) *StorageServiceStopDaemonPostDefault {
	return &StorageServiceStopDaemonPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceStopDaemonPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceStopDaemonPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service stop daemon post default response
func (o *StorageServiceStopDaemonPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceStopDaemonPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceStopDaemonPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceStopDaemonPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}