// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type SendEventRequestBody struct {
	Name string `json:"name"`
}

type SendEventRequest struct {
	RequestBody *SendEventRequestBody `request:"mediaType=application/json"`
	// The instance id
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceID"`
}

type SendEventResponse struct {
	ContentType string
	// General error
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
