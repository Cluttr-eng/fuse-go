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

// checks if the FinancialConnectionDetailsTruelayer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsTruelayer{}

// FinancialConnectionDetailsTruelayer Data needed to query data from TrueLayer
type FinancialConnectionDetailsTruelayer struct {
	// Access token for TrueLayer
	AccessToken string `json:"access_token"`
}

// NewFinancialConnectionDetailsTruelayer instantiates a new FinancialConnectionDetailsTruelayer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsTruelayer(accessToken string) *FinancialConnectionDetailsTruelayer {
	this := FinancialConnectionDetailsTruelayer{}
	this.AccessToken = accessToken
	return &this
}

// NewFinancialConnectionDetailsTruelayerWithDefaults instantiates a new FinancialConnectionDetailsTruelayer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsTruelayerWithDefaults() *FinancialConnectionDetailsTruelayer {
	this := FinancialConnectionDetailsTruelayer{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *FinancialConnectionDetailsTruelayer) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsTruelayer) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *FinancialConnectionDetailsTruelayer) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o FinancialConnectionDetailsTruelayer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsTruelayer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsTruelayer struct {
	value *FinancialConnectionDetailsTruelayer
	isSet bool
}

func (v NullableFinancialConnectionDetailsTruelayer) Get() *FinancialConnectionDetailsTruelayer {
	return v.value
}

func (v *NullableFinancialConnectionDetailsTruelayer) Set(val *FinancialConnectionDetailsTruelayer) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsTruelayer) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsTruelayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsTruelayer(val *FinancialConnectionDetailsTruelayer) *NullableFinancialConnectionDetailsTruelayer {
	return &NullableFinancialConnectionDetailsTruelayer{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsTruelayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsTruelayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


