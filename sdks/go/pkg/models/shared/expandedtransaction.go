// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type ExpandedTransaction struct {
	Metadata          map[string]string            `json:"metadata"`
	PostCommitVolumes map[string]map[string]Volume `json:"postCommitVolumes"`
	Postings          []Posting                    `json:"postings"`
	PreCommitVolumes  map[string]map[string]Volume `json:"preCommitVolumes"`
	Reference         *string                      `json:"reference,omitempty"`
	Timestamp         time.Time                    `json:"timestamp"`
	Txid              int64                        `json:"txid"`
}