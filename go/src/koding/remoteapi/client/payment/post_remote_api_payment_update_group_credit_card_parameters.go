package payment

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

// NewPostRemoteAPIPaymentUpdateGroupCreditCardParams creates a new PostRemoteAPIPaymentUpdateGroupCreditCardParams object
// with the default values initialized.
func NewPostRemoteAPIPaymentUpdateGroupCreditCardParams() *PostRemoteAPIPaymentUpdateGroupCreditCardParams {
	var ()
	return &PostRemoteAPIPaymentUpdateGroupCreditCardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIPaymentUpdateGroupCreditCardParamsWithTimeout creates a new PostRemoteAPIPaymentUpdateGroupCreditCardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIPaymentUpdateGroupCreditCardParamsWithTimeout(timeout time.Duration) *PostRemoteAPIPaymentUpdateGroupCreditCardParams {
	var ()
	return &PostRemoteAPIPaymentUpdateGroupCreditCardParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIPaymentUpdateGroupCreditCardParamsWithContext creates a new PostRemoteAPIPaymentUpdateGroupCreditCardParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIPaymentUpdateGroupCreditCardParamsWithContext(ctx context.Context) *PostRemoteAPIPaymentUpdateGroupCreditCardParams {
	var ()
	return &PostRemoteAPIPaymentUpdateGroupCreditCardParams{

		Context: ctx,
	}
}

/*PostRemoteAPIPaymentUpdateGroupCreditCardParams contains all the parameters to send to the API endpoint
for the post remote API payment update group credit card operation typically these are written to a http.Request
*/
type PostRemoteAPIPaymentUpdateGroupCreditCardParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API payment update group credit card params
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) WithTimeout(timeout time.Duration) *PostRemoteAPIPaymentUpdateGroupCreditCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API payment update group credit card params
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API payment update group credit card params
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) WithContext(ctx context.Context) *PostRemoteAPIPaymentUpdateGroupCreditCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API payment update group credit card params
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API payment update group credit card params
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIPaymentUpdateGroupCreditCardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API payment update group credit card params
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIPaymentUpdateGroupCreditCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
