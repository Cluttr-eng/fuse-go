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

// checks if the TransactionMerchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionMerchant{}

// TransactionMerchant struct for TransactionMerchant
type TransactionMerchant struct {
	// Merchant name
	Name *string `json:"name,omitempty"`
}

// NewTransactionMerchant instantiates a new TransactionMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMerchant() *TransactionMerchant {
	this := TransactionMerchant{}
	return &this
}

// NewTransactionMerchantWithDefaults instantiates a new TransactionMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionMerchantWithDefaults() *TransactionMerchant {
	this := TransactionMerchant{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransactionMerchant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMerchant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransactionMerchant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransactionMerchant) SetName(v string) {
	o.Name = &v
}

func (o TransactionMerchant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableTransactionMerchant struct {
	value *TransactionMerchant
	isSet bool
}

func (v NullableTransactionMerchant) Get() *TransactionMerchant {
	return v.value
}

func (v *NullableTransactionMerchant) Set(val *TransactionMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMerchant(val *TransactionMerchant) *NullableTransactionMerchant {
	return &NullableTransactionMerchant{value: val, isSet: true}
}

func (v NullableTransactionMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


