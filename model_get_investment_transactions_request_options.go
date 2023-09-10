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

// checks if the GetInvestmentTransactionsRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvestmentTransactionsRequestOptions{}

// GetInvestmentTransactionsRequestOptions struct for GetInvestmentTransactionsRequestOptions
type GetInvestmentTransactionsRequestOptions struct {
	// An array of account_ids to retrieve transactions for.
	RemoteAccountIds []string `json:"remote_account_ids,omitempty"`
}

// NewGetInvestmentTransactionsRequestOptions instantiates a new GetInvestmentTransactionsRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvestmentTransactionsRequestOptions() *GetInvestmentTransactionsRequestOptions {
	this := GetInvestmentTransactionsRequestOptions{}
	return &this
}

// NewGetInvestmentTransactionsRequestOptionsWithDefaults instantiates a new GetInvestmentTransactionsRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvestmentTransactionsRequestOptionsWithDefaults() *GetInvestmentTransactionsRequestOptions {
	this := GetInvestmentTransactionsRequestOptions{}
	return &this
}

// GetRemoteAccountIds returns the RemoteAccountIds field value if set, zero value otherwise.
func (o *GetInvestmentTransactionsRequestOptions) GetRemoteAccountIds() []string {
	if o == nil || IsNil(o.RemoteAccountIds) {
		var ret []string
		return ret
	}
	return o.RemoteAccountIds
}

// GetRemoteAccountIdsOk returns a tuple with the RemoteAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvestmentTransactionsRequestOptions) GetRemoteAccountIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoteAccountIds) {
		return nil, false
	}
	return o.RemoteAccountIds, true
}

// HasRemoteAccountIds returns a boolean if a field has been set.
func (o *GetInvestmentTransactionsRequestOptions) HasRemoteAccountIds() bool {
	if o != nil && !IsNil(o.RemoteAccountIds) {
		return true
	}

	return false
}

// SetRemoteAccountIds gets a reference to the given []string and assigns it to the RemoteAccountIds field.
func (o *GetInvestmentTransactionsRequestOptions) SetRemoteAccountIds(v []string) {
	o.RemoteAccountIds = v
}

func (o GetInvestmentTransactionsRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvestmentTransactionsRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemoteAccountIds) {
		toSerialize["remote_account_ids"] = o.RemoteAccountIds
	}
	return toSerialize, nil
}

type NullableGetInvestmentTransactionsRequestOptions struct {
	value *GetInvestmentTransactionsRequestOptions
	isSet bool
}

func (v NullableGetInvestmentTransactionsRequestOptions) Get() *GetInvestmentTransactionsRequestOptions {
	return v.value
}

func (v *NullableGetInvestmentTransactionsRequestOptions) Set(val *GetInvestmentTransactionsRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvestmentTransactionsRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvestmentTransactionsRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvestmentTransactionsRequestOptions(val *GetInvestmentTransactionsRequestOptions) *NullableGetInvestmentTransactionsRequestOptions {
	return &NullableGetInvestmentTransactionsRequestOptions{value: val, isSet: true}
}

func (v NullableGetInvestmentTransactionsRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvestmentTransactionsRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

