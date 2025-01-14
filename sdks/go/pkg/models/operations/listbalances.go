// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ListBalancesRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ListBalancesResponse struct {
	ContentType string
	// Balances list
	ListBalancesResponse *shared.ListBalancesResponse
	StatusCode           int
	RawResponse          *http.Response
}
