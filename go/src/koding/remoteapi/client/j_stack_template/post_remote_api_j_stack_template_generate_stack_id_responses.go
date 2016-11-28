package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJStackTemplateGenerateStackIDReader is a Reader for the PostRemoteAPIJStackTemplateGenerateStackID structure.
type PostRemoteAPIJStackTemplateGenerateStackIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJStackTemplateGenerateStackIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJStackTemplateGenerateStackIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJStackTemplateGenerateStackIDOK creates a PostRemoteAPIJStackTemplateGenerateStackIDOK with default headers values
func NewPostRemoteAPIJStackTemplateGenerateStackIDOK() *PostRemoteAPIJStackTemplateGenerateStackIDOK {
	return &PostRemoteAPIJStackTemplateGenerateStackIDOK{}
}

/*PostRemoteAPIJStackTemplateGenerateStackIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJStackTemplateGenerateStackIDOK struct {
	Payload *models.JStackTemplate
}

func (o *PostRemoteAPIJStackTemplateGenerateStackIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.generateStack/{id}][%d] postRemoteApiJStackTemplateGenerateStackIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJStackTemplateGenerateStackIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JStackTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
