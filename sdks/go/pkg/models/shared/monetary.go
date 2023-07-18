// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"math/big"
)

type Monetary struct {
	// The amount of the monetary value.
	Amount *big.Int `json:"amount"`
	// The asset of the monetary value.
	Asset string `json:"asset"`
}
