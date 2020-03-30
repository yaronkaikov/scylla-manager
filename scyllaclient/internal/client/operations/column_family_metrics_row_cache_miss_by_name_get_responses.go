// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsRowCacheMissByNameGetReader is a Reader for the ColumnFamilyMetricsRowCacheMissByNameGet structure.
type ColumnFamilyMetricsRowCacheMissByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRowCacheMissByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsRowCacheMissByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsRowCacheMissByNameGetOK creates a ColumnFamilyMetricsRowCacheMissByNameGetOK with default headers values
func NewColumnFamilyMetricsRowCacheMissByNameGetOK() *ColumnFamilyMetricsRowCacheMissByNameGetOK {
	return &ColumnFamilyMetricsRowCacheMissByNameGetOK{}
}

/*ColumnFamilyMetricsRowCacheMissByNameGetOK handles this case with default header values.

ColumnFamilyMetricsRowCacheMissByNameGetOK column family metrics row cache miss by name get o k
*/
type ColumnFamilyMetricsRowCacheMissByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/row_cache_miss/{name}][%d] columnFamilyMetricsRowCacheMissByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}