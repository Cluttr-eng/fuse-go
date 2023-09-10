/*
Fuse

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuse

import (
	"encoding/json"
)

// checks if the ExchangeFinancialConnectionsPublicTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeFinancialConnectionsPublicTokenRequest{}

// ExchangeFinancialConnectionsPublicTokenRequest struct for ExchangeFinancialConnectionsPublicTokenRequest
type ExchangeFinancialConnectionsPublicTokenRequest struct {
	// The public token created after a user connects with their financial institution
	PublicToken string `json:"public_token"`
}

// NewExchangeFinancialConnectionsPublicTokenRequest instantiates a new ExchangeFinancialConnectionsPublicTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeFinancialConnectionsPublicTokenRequest(publicToken string) *ExchangeFinancialConnectionsPublicTokenRequest {
	this := ExchangeFinancialConnectionsPublicTokenRequest{}
	this.PublicToken = publicToken
	return &this
}

// NewExchangeFinancialConnectionsPublicTokenRequestWithDefaults instantiates a new ExchangeFinancialConnectionsPublicTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeFinancialConnectionsPublicTokenRequestWithDefaults() *ExchangeFinancialConnectionsPublicTokenRequest {
	this := ExchangeFinancialConnectionsPublicTokenRequest{}
	return &this
}

// GetPublicToken returns the PublicToken field value
func (o *ExchangeFinancialConnectionsPublicTokenRequest) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *ExchangeFinancialConnectionsPublicTokenRequest) GetPublicTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *ExchangeFinancialConnectionsPublicTokenRequest) SetPublicToken(v string) {
	o.PublicToken = v
}

func (o ExchangeFinancialConnectionsPublicTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeFinancialConnectionsPublicTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["public_token"] = o.PublicToken
	return toSerialize, nil
}

type NullableExchangeFinancialConnectionsPublicTokenRequest struct {
	value *ExchangeFinancialConnectionsPublicTokenRequest
	isSet bool
}

func (v NullableExchangeFinancialConnectionsPublicTokenRequest) Get() *ExchangeFinancialConnectionsPublicTokenRequest {
	return v.value
}

func (v *NullableExchangeFinancialConnectionsPublicTokenRequest) Set(val *ExchangeFinancialConnectionsPublicTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeFinancialConnectionsPublicTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeFinancialConnectionsPublicTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeFinancialConnectionsPublicTokenRequest(val *ExchangeFinancialConnectionsPublicTokenRequest) *NullableExchangeFinancialConnectionsPublicTokenRequest {
	return &NullableExchangeFinancialConnectionsPublicTokenRequest{value: val, isSet: true}
}

func (v NullableExchangeFinancialConnectionsPublicTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeFinancialConnectionsPublicTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


