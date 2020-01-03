// Code generated by go-swagger; DO NOT EDIT.

package v_p_cs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteVpcRouteReader is a Reader for the DeleteVpcRoute structure.
type DeleteVpcRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVpcRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVpcRouteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVpcRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVpcRouteNoContent creates a DeleteVpcRouteNoContent with default headers values
func NewDeleteVpcRouteNoContent() *DeleteVpcRouteNoContent {
	return &DeleteVpcRouteNoContent{}
}

/*DeleteVpcRouteNoContent handles this case with default header values.

The route was deleted successfully.
*/
type DeleteVpcRouteNoContent struct {
}

func (o *DeleteVpcRouteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{vpc_id}/routes/{id}][%d] deleteVpcRouteNoContent ", 204)
}

func (o *DeleteVpcRouteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVpcRouteNotFound creates a DeleteVpcRouteNotFound with default headers values
func NewDeleteVpcRouteNotFound() *DeleteVpcRouteNotFound {
	return &DeleteVpcRouteNotFound{}
}

/*DeleteVpcRouteNotFound handles this case with default header values.

A route with the specified identifier could not be found.
*/
type DeleteVpcRouteNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpcRouteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{vpc_id}/routes/{id}][%d] deleteVpcRouteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVpcRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}