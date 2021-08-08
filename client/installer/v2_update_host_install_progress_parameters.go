// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewV2UpdateHostInstallProgressParams creates a new V2UpdateHostInstallProgressParams object
// with the default values initialized.
func NewV2UpdateHostInstallProgressParams() *V2UpdateHostInstallProgressParams {
	var ()
	return &V2UpdateHostInstallProgressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2UpdateHostInstallProgressParamsWithTimeout creates a new V2UpdateHostInstallProgressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2UpdateHostInstallProgressParamsWithTimeout(timeout time.Duration) *V2UpdateHostInstallProgressParams {
	var ()
	return &V2UpdateHostInstallProgressParams{

		timeout: timeout,
	}
}

// NewV2UpdateHostInstallProgressParamsWithContext creates a new V2UpdateHostInstallProgressParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2UpdateHostInstallProgressParamsWithContext(ctx context.Context) *V2UpdateHostInstallProgressParams {
	var ()
	return &V2UpdateHostInstallProgressParams{

		Context: ctx,
	}
}

// NewV2UpdateHostInstallProgressParamsWithHTTPClient creates a new V2UpdateHostInstallProgressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2UpdateHostInstallProgressParamsWithHTTPClient(client *http.Client) *V2UpdateHostInstallProgressParams {
	var ()
	return &V2UpdateHostInstallProgressParams{
		HTTPClient: client,
	}
}

/*V2UpdateHostInstallProgressParams contains all the parameters to send to the API endpoint
for the v2 update host install progress operation typically these are written to a http.Request
*/
type V2UpdateHostInstallProgressParams struct {

	/*DiscoveryAgentVersion
	  The software version of the discovery agent that is updating progress.

	*/
	DiscoveryAgentVersion *string
	/*HostProgress
	  New progress value.

	*/
	HostProgress *models.HostProgress
	/*HostID
	  The ID of the host to update.

	*/
	HostID strfmt.UUID
	/*InfraEnvID
	  The InfraEnv of the host being updated.

	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithTimeout(timeout time.Duration) *V2UpdateHostInstallProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithContext(ctx context.Context) *V2UpdateHostInstallProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithHTTPClient(client *http.Client) *V2UpdateHostInstallProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiscoveryAgentVersion adds the discoveryAgentVersion to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithDiscoveryAgentVersion(discoveryAgentVersion *string) *V2UpdateHostInstallProgressParams {
	o.SetDiscoveryAgentVersion(discoveryAgentVersion)
	return o
}

// SetDiscoveryAgentVersion adds the discoveryAgentVersion to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetDiscoveryAgentVersion(discoveryAgentVersion *string) {
	o.DiscoveryAgentVersion = discoveryAgentVersion
}

// WithHostProgress adds the hostProgress to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithHostProgress(hostProgress *models.HostProgress) *V2UpdateHostInstallProgressParams {
	o.SetHostProgress(hostProgress)
	return o
}

// SetHostProgress adds the hostProgress to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetHostProgress(hostProgress *models.HostProgress) {
	o.HostProgress = hostProgress
}

// WithHostID adds the hostID to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithHostID(hostID strfmt.UUID) *V2UpdateHostInstallProgressParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithInfraEnvID adds the infraEnvID to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) WithInfraEnvID(infraEnvID strfmt.UUID) *V2UpdateHostInstallProgressParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the v2 update host install progress params
func (o *V2UpdateHostInstallProgressParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *V2UpdateHostInstallProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DiscoveryAgentVersion != nil {

		// header param discovery_agent_version
		if err := r.SetHeaderParam("discovery_agent_version", *o.DiscoveryAgentVersion); err != nil {
			return err
		}

	}

	if o.HostProgress != nil {
		if err := r.SetBodyParam(o.HostProgress); err != nil {
			return err
		}
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
