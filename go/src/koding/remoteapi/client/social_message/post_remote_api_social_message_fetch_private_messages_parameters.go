package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPISocialMessageFetchPrivateMessagesParams creates a new PostRemoteAPISocialMessageFetchPrivateMessagesParams object
// with the default values initialized.
func NewPostRemoteAPISocialMessageFetchPrivateMessagesParams() *PostRemoteAPISocialMessageFetchPrivateMessagesParams {
	var ()
	return &PostRemoteAPISocialMessageFetchPrivateMessagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialMessageFetchPrivateMessagesParamsWithTimeout creates a new PostRemoteAPISocialMessageFetchPrivateMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialMessageFetchPrivateMessagesParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageFetchPrivateMessagesParams {
	var ()
	return &PostRemoteAPISocialMessageFetchPrivateMessagesParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialMessageFetchPrivateMessagesParamsWithContext creates a new PostRemoteAPISocialMessageFetchPrivateMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialMessageFetchPrivateMessagesParamsWithContext(ctx context.Context) *PostRemoteAPISocialMessageFetchPrivateMessagesParams {
	var ()
	return &PostRemoteAPISocialMessageFetchPrivateMessagesParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialMessageFetchPrivateMessagesParams contains all the parameters to send to the API endpoint
for the post remote API social message fetch private messages operation typically these are written to a http.Request
*/
type PostRemoteAPISocialMessageFetchPrivateMessagesParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API social message fetch private messages params
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageFetchPrivateMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social message fetch private messages params
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social message fetch private messages params
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) WithContext(ctx context.Context) *PostRemoteAPISocialMessageFetchPrivateMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social message fetch private messages params
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social message fetch private messages params
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialMessageFetchPrivateMessagesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social message fetch private messages params
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialMessageFetchPrivateMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
