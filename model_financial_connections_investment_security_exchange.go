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

// checks if the FinancialConnectionsInvestmentSecurityExchange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsInvestmentSecurityExchange{}

// FinancialConnectionsInvestmentSecurityExchange struct for FinancialConnectionsInvestmentSecurityExchange
type FinancialConnectionsInvestmentSecurityExchange struct {
	// The Market Identifier Code (MIC) associated with the specific financial market or exchange where the security is traded.
	MicCode *string `json:"mic_code,omitempty"`
	// The suffix of the security's ticker symbol.
	Suffix *string `json:"suffix,omitempty"`
}

// NewFinancialConnectionsInvestmentSecurityExchange instantiates a new FinancialConnectionsInvestmentSecurityExchange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsInvestmentSecurityExchange() *FinancialConnectionsInvestmentSecurityExchange {
	this := FinancialConnectionsInvestmentSecurityExchange{}
	return &this
}

// NewFinancialConnectionsInvestmentSecurityExchangeWithDefaults instantiates a new FinancialConnectionsInvestmentSecurityExchange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsInvestmentSecurityExchangeWithDefaults() *FinancialConnectionsInvestmentSecurityExchange {
	this := FinancialConnectionsInvestmentSecurityExchange{}
	return &this
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurityExchange) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurityExchange) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurityExchange) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *FinancialConnectionsInvestmentSecurityExchange) SetMicCode(v string) {
	o.MicCode = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurityExchange) GetSuffix() string {
	if o == nil || IsNil(o.Suffix) {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurityExchange) GetSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Suffix) {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurityExchange) HasSuffix() bool {
	if o != nil && !IsNil(o.Suffix) {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *FinancialConnectionsInvestmentSecurityExchange) SetSuffix(v string) {
	o.Suffix = &v
}

func (o FinancialConnectionsInvestmentSecurityExchange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsInvestmentSecurityExchange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Suffix) {
		toSerialize["suffix"] = o.Suffix
	}
	return toSerialize, nil
}

type NullableFinancialConnectionsInvestmentSecurityExchange struct {
	value *FinancialConnectionsInvestmentSecurityExchange
	isSet bool
}

func (v NullableFinancialConnectionsInvestmentSecurityExchange) Get() *FinancialConnectionsInvestmentSecurityExchange {
	return v.value
}

func (v *NullableFinancialConnectionsInvestmentSecurityExchange) Set(val *FinancialConnectionsInvestmentSecurityExchange) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsInvestmentSecurityExchange) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsInvestmentSecurityExchange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsInvestmentSecurityExchange(val *FinancialConnectionsInvestmentSecurityExchange) *NullableFinancialConnectionsInvestmentSecurityExchange {
	return &NullableFinancialConnectionsInvestmentSecurityExchange{value: val, isSet: true}
}

func (v NullableFinancialConnectionsInvestmentSecurityExchange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsInvestmentSecurityExchange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

