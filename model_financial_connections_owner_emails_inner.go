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

// checks if the FinancialConnectionsOwnerEmailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsOwnerEmailsInner{}

// FinancialConnectionsOwnerEmailsInner struct for FinancialConnectionsOwnerEmailsInner
type FinancialConnectionsOwnerEmailsInner struct {
	// Email address
	Data string `json:"data"`
	// Indicating if it is the primary email
	Primary *bool `json:"primary,omitempty"`
	// Type of the email
	Type *string `json:"type,omitempty"`
}

// NewFinancialConnectionsOwnerEmailsInner instantiates a new FinancialConnectionsOwnerEmailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsOwnerEmailsInner(data string) *FinancialConnectionsOwnerEmailsInner {
	this := FinancialConnectionsOwnerEmailsInner{}
	this.Data = data
	return &this
}

// NewFinancialConnectionsOwnerEmailsInnerWithDefaults instantiates a new FinancialConnectionsOwnerEmailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsOwnerEmailsInnerWithDefaults() *FinancialConnectionsOwnerEmailsInner {
	this := FinancialConnectionsOwnerEmailsInner{}
	return &this
}

// GetData returns the Data field value
func (o *FinancialConnectionsOwnerEmailsInner) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsOwnerEmailsInner) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FinancialConnectionsOwnerEmailsInner) SetData(v string) {
	o.Data = v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *FinancialConnectionsOwnerEmailsInner) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsOwnerEmailsInner) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *FinancialConnectionsOwnerEmailsInner) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *FinancialConnectionsOwnerEmailsInner) SetPrimary(v bool) {
	o.Primary = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FinancialConnectionsOwnerEmailsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsOwnerEmailsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FinancialConnectionsOwnerEmailsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FinancialConnectionsOwnerEmailsInner) SetType(v string) {
	o.Type = &v
}

func (o FinancialConnectionsOwnerEmailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsOwnerEmailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFinancialConnectionsOwnerEmailsInner struct {
	value *FinancialConnectionsOwnerEmailsInner
	isSet bool
}

func (v NullableFinancialConnectionsOwnerEmailsInner) Get() *FinancialConnectionsOwnerEmailsInner {
	return v.value
}

func (v *NullableFinancialConnectionsOwnerEmailsInner) Set(val *FinancialConnectionsOwnerEmailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsOwnerEmailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsOwnerEmailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsOwnerEmailsInner(val *FinancialConnectionsOwnerEmailsInner) *NullableFinancialConnectionsOwnerEmailsInner {
	return &NullableFinancialConnectionsOwnerEmailsInner{value: val, isSet: true}
}

func (v NullableFinancialConnectionsOwnerEmailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsOwnerEmailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

