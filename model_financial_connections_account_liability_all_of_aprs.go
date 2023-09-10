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

// checks if the FinancialConnectionsAccountLiabilityAllOfAprs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsAccountLiabilityAllOfAprs{}

// FinancialConnectionsAccountLiabilityAllOfAprs struct for FinancialConnectionsAccountLiabilityAllOfAprs
type FinancialConnectionsAccountLiabilityAllOfAprs struct {
	// Annual Percentage Rate applied.
	AprPercentage *float32 `json:"apr_percentage,omitempty"`
	// The type of balance to which the APR applies.
	AprType *string `json:"apr_type,omitempty"`
}

// NewFinancialConnectionsAccountLiabilityAllOfAprs instantiates a new FinancialConnectionsAccountLiabilityAllOfAprs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsAccountLiabilityAllOfAprs() *FinancialConnectionsAccountLiabilityAllOfAprs {
	this := FinancialConnectionsAccountLiabilityAllOfAprs{}
	return &this
}

// NewFinancialConnectionsAccountLiabilityAllOfAprsWithDefaults instantiates a new FinancialConnectionsAccountLiabilityAllOfAprs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsAccountLiabilityAllOfAprsWithDefaults() *FinancialConnectionsAccountLiabilityAllOfAprs {
	this := FinancialConnectionsAccountLiabilityAllOfAprs{}
	return &this
}

// GetAprPercentage returns the AprPercentage field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) GetAprPercentage() float32 {
	if o == nil || IsNil(o.AprPercentage) {
		var ret float32
		return ret
	}
	return *o.AprPercentage
}

// GetAprPercentageOk returns a tuple with the AprPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) GetAprPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.AprPercentage) {
		return nil, false
	}
	return o.AprPercentage, true
}

// HasAprPercentage returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) HasAprPercentage() bool {
	if o != nil && !IsNil(o.AprPercentage) {
		return true
	}

	return false
}

// SetAprPercentage gets a reference to the given float32 and assigns it to the AprPercentage field.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) SetAprPercentage(v float32) {
	o.AprPercentage = &v
}

// GetAprType returns the AprType field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) GetAprType() string {
	if o == nil || IsNil(o.AprType) {
		var ret string
		return ret
	}
	return *o.AprType
}

// GetAprTypeOk returns a tuple with the AprType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) GetAprTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AprType) {
		return nil, false
	}
	return o.AprType, true
}

// HasAprType returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) HasAprType() bool {
	if o != nil && !IsNil(o.AprType) {
		return true
	}

	return false
}

// SetAprType gets a reference to the given string and assigns it to the AprType field.
func (o *FinancialConnectionsAccountLiabilityAllOfAprs) SetAprType(v string) {
	o.AprType = &v
}

func (o FinancialConnectionsAccountLiabilityAllOfAprs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsAccountLiabilityAllOfAprs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AprPercentage) {
		toSerialize["apr_percentage"] = o.AprPercentage
	}
	if !IsNil(o.AprType) {
		toSerialize["apr_type"] = o.AprType
	}
	return toSerialize, nil
}

type NullableFinancialConnectionsAccountLiabilityAllOfAprs struct {
	value *FinancialConnectionsAccountLiabilityAllOfAprs
	isSet bool
}

func (v NullableFinancialConnectionsAccountLiabilityAllOfAprs) Get() *FinancialConnectionsAccountLiabilityAllOfAprs {
	return v.value
}

func (v *NullableFinancialConnectionsAccountLiabilityAllOfAprs) Set(val *FinancialConnectionsAccountLiabilityAllOfAprs) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsAccountLiabilityAllOfAprs) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsAccountLiabilityAllOfAprs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsAccountLiabilityAllOfAprs(val *FinancialConnectionsAccountLiabilityAllOfAprs) *NullableFinancialConnectionsAccountLiabilityAllOfAprs {
	return &NullableFinancialConnectionsAccountLiabilityAllOfAprs{value: val, isSet: true}
}

func (v NullableFinancialConnectionsAccountLiabilityAllOfAprs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsAccountLiabilityAllOfAprs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

