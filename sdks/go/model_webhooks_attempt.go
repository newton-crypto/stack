/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"time"
)

// WebhooksAttempt struct for WebhooksAttempt
type WebhooksAttempt struct {
	Id             *string         `json:"id,omitempty"`
	WebhookID      *string         `json:"webhookID,omitempty"`
	CreatedAt      *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time      `json:"updatedAt,omitempty"`
	Config         *WebhooksConfig `json:"config,omitempty"`
	Payload        *string         `json:"payload,omitempty"`
	StatusCode     *int32          `json:"statusCode,omitempty"`
	RetryAttempt   *int32          `json:"retryAttempt,omitempty"`
	Status         *string         `json:"status,omitempty"`
	NextRetryAfter *time.Time      `json:"nextRetryAfter,omitempty"`
}

// NewWebhooksAttempt instantiates a new WebhooksAttempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksAttempt() *WebhooksAttempt {
	this := WebhooksAttempt{}
	return &this
}

// NewWebhooksAttemptWithDefaults instantiates a new WebhooksAttempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksAttemptWithDefaults() *WebhooksAttempt {
	this := WebhooksAttempt{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhooksAttempt) SetId(v string) {
	o.Id = &v
}

// GetWebhookID returns the WebhookID field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetWebhookID() string {
	if o == nil || isNil(o.WebhookID) {
		var ret string
		return ret
	}
	return *o.WebhookID
}

// GetWebhookIDOk returns a tuple with the WebhookID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetWebhookIDOk() (*string, bool) {
	if o == nil || isNil(o.WebhookID) {
		return nil, false
	}
	return o.WebhookID, true
}

// HasWebhookID returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasWebhookID() bool {
	if o != nil && !isNil(o.WebhookID) {
		return true
	}

	return false
}

// SetWebhookID gets a reference to the given string and assigns it to the WebhookID field.
func (o *WebhooksAttempt) SetWebhookID(v string) {
	o.WebhookID = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WebhooksAttempt) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WebhooksAttempt) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetConfig() WebhooksConfig {
	if o == nil || isNil(o.Config) {
		var ret WebhooksConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetConfigOk() (*WebhooksConfig, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given WebhooksConfig and assigns it to the Config field.
func (o *WebhooksAttempt) SetConfig(v WebhooksConfig) {
	o.Config = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetPayload() string {
	if o == nil || isNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetPayloadOk() (*string, bool) {
	if o == nil || isNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *WebhooksAttempt) SetPayload(v string) {
	o.Payload = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetStatusCode() int32 {
	if o == nil || isNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetStatusCodeOk() (*int32, bool) {
	if o == nil || isNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasStatusCode() bool {
	if o != nil && !isNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *WebhooksAttempt) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRetryAttempt returns the RetryAttempt field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetRetryAttempt() int32 {
	if o == nil || isNil(o.RetryAttempt) {
		var ret int32
		return ret
	}
	return *o.RetryAttempt
}

// GetRetryAttemptOk returns a tuple with the RetryAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetRetryAttemptOk() (*int32, bool) {
	if o == nil || isNil(o.RetryAttempt) {
		return nil, false
	}
	return o.RetryAttempt, true
}

// HasRetryAttempt returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasRetryAttempt() bool {
	if o != nil && !isNil(o.RetryAttempt) {
		return true
	}

	return false
}

// SetRetryAttempt gets a reference to the given int32 and assigns it to the RetryAttempt field.
func (o *WebhooksAttempt) SetRetryAttempt(v int32) {
	o.RetryAttempt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebhooksAttempt) SetStatus(v string) {
	o.Status = &v
}

// GetNextRetryAfter returns the NextRetryAfter field value if set, zero value otherwise.
func (o *WebhooksAttempt) GetNextRetryAfter() time.Time {
	if o == nil || isNil(o.NextRetryAfter) {
		var ret time.Time
		return ret
	}
	return *o.NextRetryAfter
}

// GetNextRetryAfterOk returns a tuple with the NextRetryAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAttempt) GetNextRetryAfterOk() (*time.Time, bool) {
	if o == nil || isNil(o.NextRetryAfter) {
		return nil, false
	}
	return o.NextRetryAfter, true
}

// HasNextRetryAfter returns a boolean if a field has been set.
func (o *WebhooksAttempt) HasNextRetryAfter() bool {
	if o != nil && !isNil(o.NextRetryAfter) {
		return true
	}

	return false
}

// SetNextRetryAfter gets a reference to the given time.Time and assigns it to the NextRetryAfter field.
func (o *WebhooksAttempt) SetNextRetryAfter(v time.Time) {
	o.NextRetryAfter = &v
}

func (o WebhooksAttempt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.WebhookID) {
		toSerialize["webhookID"] = o.WebhookID
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !isNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !isNil(o.RetryAttempt) {
		toSerialize["retryAttempt"] = o.RetryAttempt
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.NextRetryAfter) {
		toSerialize["nextRetryAfter"] = o.NextRetryAfter
	}
	return json.Marshal(toSerialize)
}

type NullableWebhooksAttempt struct {
	value *WebhooksAttempt
	isSet bool
}

func (v NullableWebhooksAttempt) Get() *WebhooksAttempt {
	return v.value
}

func (v *NullableWebhooksAttempt) Set(val *WebhooksAttempt) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksAttempt) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksAttempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksAttempt(val *WebhooksAttempt) *NullableWebhooksAttempt {
	return &NullableWebhooksAttempt{value: val, isSet: true}
}

func (v NullableWebhooksAttempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksAttempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}