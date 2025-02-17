/*
Auth API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: AUTH_VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authclient

import (
	"encoding/json"
)

// checks if the CreateScopeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScopeResponse{}

// CreateScopeResponse struct for CreateScopeResponse
type CreateScopeResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// NewCreateScopeResponse instantiates a new CreateScopeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScopeResponse() *CreateScopeResponse {
	this := CreateScopeResponse{}
	return &this
}

// NewCreateScopeResponseWithDefaults instantiates a new CreateScopeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScopeResponseWithDefaults() *CreateScopeResponse {
	this := CreateScopeResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateScopeResponse) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateScopeResponse) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateScopeResponse) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *CreateScopeResponse) SetData(v interface{}) {
	o.Data = v
}

func (o CreateScopeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateScopeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateScopeResponse struct {
	value *CreateScopeResponse
	isSet bool
}

func (v NullableCreateScopeResponse) Get() *CreateScopeResponse {
	return v.value
}

func (v *NullableCreateScopeResponse) Set(val *CreateScopeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScopeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScopeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScopeResponse(val *CreateScopeResponse) *NullableCreateScopeResponse {
	return &NullableCreateScopeResponse{value: val, isSet: true}
}

func (v NullableCreateScopeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateScopeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


