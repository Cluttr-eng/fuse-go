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

// checks if the FinancialConnectionDetailsMono type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsMono{}

// FinancialConnectionDetailsMono Data needed to query data from Mono
type FinancialConnectionDetailsMono struct {
	// Account Id for Mono
	AccountId string `json:"account_id"`
}

// NewFinancialConnectionDetailsMono instantiates a new FinancialConnectionDetailsMono object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsMono(accountId string) *FinancialConnectionDetailsMono {
	this := FinancialConnectionDetailsMono{}
	this.AccountId = accountId
	return &this
}

// NewFinancialConnectionDetailsMonoWithDefaults instantiates a new FinancialConnectionDetailsMono object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsMonoWithDefaults() *FinancialConnectionDetailsMono {
	this := FinancialConnectionDetailsMono{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *FinancialConnectionDetailsMono) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsMono) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *FinancialConnectionDetailsMono) SetAccountId(v string) {
	o.AccountId = v
}

func (o FinancialConnectionDetailsMono) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsMono) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsMono struct {
	value *FinancialConnectionDetailsMono
	isSet bool
}

func (v NullableFinancialConnectionDetailsMono) Get() *FinancialConnectionDetailsMono {
	return v.value
}

func (v *NullableFinancialConnectionDetailsMono) Set(val *FinancialConnectionDetailsMono) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsMono) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsMono) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsMono(val *FinancialConnectionDetailsMono) *NullableFinancialConnectionDetailsMono {
	return &NullableFinancialConnectionDetailsMono{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsMono) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsMono) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


