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

// GetSiteDeployReader is a Reader for the GetSiteDeploy structure.
type GetSiteDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSiteDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSiteDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSiteDeployDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetSiteDeployOK creates a GetSiteDeployOK with default headers values
func NewGetSiteDeployOK() *GetSiteDeployOK {
	return &GetSiteDeployOK{}
}

/*GetSiteDeployOK handles this case with default header values.

OK
*/
type GetSiteDeployOK struct {
	Payload *models.Deploy
}

func (o *GetSiteDeployOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/deploys/{deploy_id}][%d] getSiteDeployOK  %+v", 200, o.Payload)
}

func (o *GetSiteDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteDeployDefault creates a GetSiteDeployDefault with default headers values
func NewGetSiteDeployDefault(code int) *GetSiteDeployDefault {
	return &GetSiteDeployDefault{
		_statusCode: code,
	}
}

/*GetSiteDeployDefault handles this case with default header values.

error
*/
type GetSiteDeployDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get site deploy default response
func (o *GetSiteDeployDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteDeployDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/deploys/{deploy_id}][%d] getSiteDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *GetSiteDeployDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
