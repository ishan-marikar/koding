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

// NewPostRemoteAPISocialMessageFetchPrivateMessageCountParams creates a new PostRemoteAPISocialMessageFetchPrivateMessageCountParams object
// with the default values initialized.
func NewPostRemoteAPISocialMessageFetchPrivateMessageCountParams() *PostRemoteAPISocialMessageFetchPrivateMessageCountParams {
	var ()
	return &PostRemoteAPISocialMessageFetchPrivateMessageCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialMessageFetchPrivateMessageCountParamsWithTimeout creates a new PostRemoteAPISocialMessageFetchPrivateMessageCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialMessageFetchPrivateMessageCountParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageFetchPrivateMessageCountParams {
	var ()
	return &PostRemoteAPISocialMessageFetchPrivateMessageCountParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialMessageFetchPrivateMessageCountParamsWithContext creates a new PostRemoteAPISocialMessageFetchPrivateMessageCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialMessageFetchPrivateMessageCountParamsWithContext(ctx context.Context) *PostRemoteAPISocialMessageFetchPrivateMessageCountParams {
	var ()
	return &PostRemoteAPISocialMessageFetchPrivateMessageCountParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialMessageFetchPrivateMessageCountParams contains all the parameters to send to the API endpoint
for the post remote API social message fetch private message count operation typically these are written to a http.Request
*/
type PostRemoteAPISocialMessageFetchPrivateMessageCountParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API social message fetch private message count params
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageFetchPrivateMessageCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social message fetch private message count params
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social message fetch private message count params
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) WithContext(ctx context.Context) *PostRemoteAPISocialMessageFetchPrivateMessageCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social message fetch private message count params
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social message fetch private message count params
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialMessageFetchPrivateMessageCountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social message fetch private message count params
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialMessageFetchPrivateMessageCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
