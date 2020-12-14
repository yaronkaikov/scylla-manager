// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// MessagingServiceMessagesRepliedGetReader is a Reader for the MessagingServiceMessagesRepliedGet structure.
type MessagingServiceMessagesRepliedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MessagingServiceMessagesRepliedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMessagingServiceMessagesRepliedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMessagingServiceMessagesRepliedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMessagingServiceMessagesRepliedGetOK creates a MessagingServiceMessagesRepliedGetOK with default headers values
func NewMessagingServiceMessagesRepliedGetOK() *MessagingServiceMessagesRepliedGetOK {
	return &MessagingServiceMessagesRepliedGetOK{}
}

/*MessagingServiceMessagesRepliedGetOK handles this case with default header values.

MessagingServiceMessagesRepliedGetOK messaging service messages replied get o k
*/
type MessagingServiceMessagesRepliedGetOK struct {
	Payload []*models.MessageCounter
}

func (o *MessagingServiceMessagesRepliedGetOK) GetPayload() []*models.MessageCounter {
	return o.Payload
}

func (o *MessagingServiceMessagesRepliedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMessagingServiceMessagesRepliedGetDefault creates a MessagingServiceMessagesRepliedGetDefault with default headers values
func NewMessagingServiceMessagesRepliedGetDefault(code int) *MessagingServiceMessagesRepliedGetDefault {
	return &MessagingServiceMessagesRepliedGetDefault{
		_statusCode: code,
	}
}

/*MessagingServiceMessagesRepliedGetDefault handles this case with default header values.

internal server error
*/
type MessagingServiceMessagesRepliedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the messaging service messages replied get default response
func (o *MessagingServiceMessagesRepliedGetDefault) Code() int {
	return o._statusCode
}

func (o *MessagingServiceMessagesRepliedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *MessagingServiceMessagesRepliedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *MessagingServiceMessagesRepliedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}