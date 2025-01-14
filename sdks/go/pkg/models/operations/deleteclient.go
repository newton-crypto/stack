// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteClientRequest struct {
	// Client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type DeleteClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
