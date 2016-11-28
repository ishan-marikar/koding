package j_proposed_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJProposedDomainBindMachineIDParams creates a new PostRemoteAPIJProposedDomainBindMachineIDParams object
// with the default values initialized.
func NewPostRemoteAPIJProposedDomainBindMachineIDParams() *PostRemoteAPIJProposedDomainBindMachineIDParams {
	var ()
	return &PostRemoteAPIJProposedDomainBindMachineIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProposedDomainBindMachineIDParamsWithTimeout creates a new PostRemoteAPIJProposedDomainBindMachineIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProposedDomainBindMachineIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProposedDomainBindMachineIDParams {
	var ()
	return &PostRemoteAPIJProposedDomainBindMachineIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProposedDomainBindMachineIDParamsWithContext creates a new PostRemoteAPIJProposedDomainBindMachineIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProposedDomainBindMachineIDParamsWithContext(ctx context.Context) *PostRemoteAPIJProposedDomainBindMachineIDParams {
	var ()
	return &PostRemoteAPIJProposedDomainBindMachineIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProposedDomainBindMachineIDParams contains all the parameters to send to the API endpoint
for the post remote API j proposed domain bind machine ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJProposedDomainBindMachineIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j proposed domain bind machine ID params
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProposedDomainBindMachineIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j proposed domain bind machine ID params
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j proposed domain bind machine ID params
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) WithContext(ctx context.Context) *PostRemoteAPIJProposedDomainBindMachineIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j proposed domain bind machine ID params
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j proposed domain bind machine ID params
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) WithID(id string) *PostRemoteAPIJProposedDomainBindMachineIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j proposed domain bind machine ID params
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProposedDomainBindMachineIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
