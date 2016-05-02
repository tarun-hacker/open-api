package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/swagger-api-client/go/models"
)

// UpdateSiteSnippetReader is a Reader for the UpdateSiteSnippet structure.
type UpdateSiteSnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateSiteSnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSiteSnippetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSiteSnippetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpdateSiteSnippetNoContent creates a UpdateSiteSnippetNoContent with default headers values
func NewUpdateSiteSnippetNoContent() *UpdateSiteSnippetNoContent {
	return &UpdateSiteSnippetNoContent{}
}

/*UpdateSiteSnippetNoContent handles this case with default header values.

No content
*/
type UpdateSiteSnippetNoContent struct {
}

func (o *UpdateSiteSnippetNoContent) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/snippets/{snippet_id}][%d] updateSiteSnippetNoContent ", 204)
}

func (o *UpdateSiteSnippetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSiteSnippetDefault creates a UpdateSiteSnippetDefault with default headers values
func NewUpdateSiteSnippetDefault(code int) *UpdateSiteSnippetDefault {
	return &UpdateSiteSnippetDefault{
		_statusCode: code,
	}
}

/*UpdateSiteSnippetDefault handles this case with default header values.

error
*/
type UpdateSiteSnippetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update site snippet default response
func (o *UpdateSiteSnippetDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSiteSnippetDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/snippets/{snippet_id}][%d] updateSiteSnippet default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSiteSnippetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
