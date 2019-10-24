// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// NewUpdateClusterParams creates a new UpdateClusterParams object
// with the default values initialized.
func NewUpdateClusterParams() *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterParamsWithTimeout creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClusterParamsWithTimeout(timeout time.Duration) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		timeout: timeout,
	}
}

// NewUpdateClusterParamsWithContext creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClusterParamsWithContext(ctx context.Context) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		Context: ctx,
	}
}

// NewUpdateClusterParamsWithHTTPClient creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClusterParamsWithHTTPClient(client *http.Client) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{
		HTTPClient: client,
	}
}

/*UpdateClusterParams contains all the parameters to send to the API endpoint
for the update cluster operation typically these are written to a http.Request
*/
type UpdateClusterParams struct {

	/*Cluster
	  Cluster that will be updated.

	*/
	Cluster *models.ClusterUpdate
	/*ClusterID
	  ID of the cluster.

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) WithTimeout(timeout time.Duration) *UpdateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster params
func (o *UpdateClusterParams) WithContext(ctx context.Context) *UpdateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster params
func (o *UpdateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) WithHTTPClient(client *http.Client) *UpdateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the update cluster params
func (o *UpdateClusterParams) WithCluster(cluster *models.ClusterUpdate) *UpdateClusterParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the update cluster params
func (o *UpdateClusterParams) SetCluster(cluster *models.ClusterUpdate) {
	o.Cluster = cluster
}

// WithClusterID adds the clusterID to the update cluster params
func (o *UpdateClusterParams) WithClusterID(clusterID string) *UpdateClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update cluster params
func (o *UpdateClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}