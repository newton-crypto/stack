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
)

// checks if the GetTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransactionResponse{}

// GetTransactionResponse struct for GetTransactionResponse
type GetTransactionResponse struct {
	Data ExpandedTransaction `json:"data"`
}

// NewGetTransactionResponse instantiates a new GetTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionResponse(data ExpandedTransaction) *GetTransactionResponse {
	this := GetTransactionResponse{}
	this.Data = data
	return &this
}

// NewGetTransactionResponseWithDefaults instantiates a new GetTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionResponseWithDefaults() *GetTransactionResponse {
	this := GetTransactionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetTransactionResponse) GetData() ExpandedTransaction {
	if o == nil {
		var ret ExpandedTransaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetDataOk() (*ExpandedTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetTransactionResponse) SetData(v ExpandedTransaction) {
	o.Data = v
}

func (o GetTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetTransactionResponse struct {
	value *GetTransactionResponse
	isSet bool
}

func (v NullableGetTransactionResponse) Get() *GetTransactionResponse {
	return v.value
}

func (v *NullableGetTransactionResponse) Set(val *GetTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionResponse(val *GetTransactionResponse) *NullableGetTransactionResponse {
	return &NullableGetTransactionResponse{value: val, isSet: true}
}

func (v NullableGetTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

