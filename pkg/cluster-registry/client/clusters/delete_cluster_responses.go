// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// DeleteClusterReader is a Reader for the DeleteCluster structure.
type DeleteClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteClusterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClusterNoContent creates a DeleteClusterNoContent with default headers values
func NewDeleteClusterNoContent() *DeleteClusterNoContent {
	return &DeleteClusterNoContent{}
}

/*DeleteClusterNoContent handles this case with default header values.

Cluster deleted
*/
type DeleteClusterNoContent struct {
}

func (o *DeleteClusterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}][%d] deleteClusterNoContent ", 204)
}

func (o *DeleteClusterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterBadRequest creates a DeleteClusterBadRequest with default headers values
func NewDeleteClusterBadRequest() *DeleteClusterBadRequest {
	return &DeleteClusterBadRequest{}
}

/*DeleteClusterBadRequest handles this case with default header values.

Invalid request
*/
type DeleteClusterBadRequest struct {
	Payload *models.Error
}

func (o *DeleteClusterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}][%d] deleteClusterBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterUnauthorized creates a DeleteClusterUnauthorized with default headers values
func NewDeleteClusterUnauthorized() *DeleteClusterUnauthorized {
	return &DeleteClusterUnauthorized{}
}

/*DeleteClusterUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteClusterUnauthorized struct {
}

func (o *DeleteClusterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}][%d] deleteClusterUnauthorized ", 401)
}

func (o *DeleteClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterForbidden creates a DeleteClusterForbidden with default headers values
func NewDeleteClusterForbidden() *DeleteClusterForbidden {
	return &DeleteClusterForbidden{}
}

/*DeleteClusterForbidden handles this case with default header values.

Forbidden
*/
type DeleteClusterForbidden struct {
	Payload *models.Error
}

func (o *DeleteClusterForbidden) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}][%d] deleteClusterForbidden  %+v", 403, o.Payload)
}

func (o *DeleteClusterForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterNotFound creates a DeleteClusterNotFound with default headers values
func NewDeleteClusterNotFound() *DeleteClusterNotFound {
	return &DeleteClusterNotFound{}
}

/*DeleteClusterNotFound handles this case with default header values.

Cluster not found
*/
type DeleteClusterNotFound struct {
}

func (o *DeleteClusterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}][%d] deleteClusterNotFound ", 404)
}

func (o *DeleteClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterInternalServerError creates a DeleteClusterInternalServerError with default headers values
func NewDeleteClusterInternalServerError() *DeleteClusterInternalServerError {
	return &DeleteClusterInternalServerError{}
}

/*DeleteClusterInternalServerError handles this case with default header values.

Unexpected error
*/
type DeleteClusterInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteClusterInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}][%d] deleteClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
