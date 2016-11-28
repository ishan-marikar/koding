package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountUnlinkOauthIDReader is a Reader for the PostRemoteAPIJAccountUnlinkOauthID structure.
type PostRemoteAPIJAccountUnlinkOauthIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountUnlinkOauthIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountUnlinkOauthIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountUnlinkOauthIDOK creates a PostRemoteAPIJAccountUnlinkOauthIDOK with default headers values
func NewPostRemoteAPIJAccountUnlinkOauthIDOK() *PostRemoteAPIJAccountUnlinkOauthIDOK {
	return &PostRemoteAPIJAccountUnlinkOauthIDOK{}
}

/*PostRemoteAPIJAccountUnlinkOauthIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountUnlinkOauthIDOK struct {
	Payload *models.JAccount
}

func (o *PostRemoteAPIJAccountUnlinkOauthIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.unlinkOauth/{id}][%d] postRemoteApiJAccountUnlinkOauthIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountUnlinkOauthIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
