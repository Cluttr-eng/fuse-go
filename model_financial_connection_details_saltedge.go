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

// checks if the FinancialConnectionDetailsSaltedge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsSaltedge{}

// FinancialConnectionDetailsSaltedge Data needed to query data from SaltEdge.
type FinancialConnectionDetailsSaltedge struct {
	ConnectionId string `json:"connection_id"`
}

// NewFinancialConnectionDetailsSaltedge instantiates a new FinancialConnectionDetailsSaltedge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsSaltedge(connectionId string) *FinancialConnectionDetailsSaltedge {
	this := FinancialConnectionDetailsSaltedge{}
	this.ConnectionId = connectionId
	return &this
}

// NewFinancialConnectionDetailsSaltedgeWithDefaults instantiates a new FinancialConnectionDetailsSaltedge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsSaltedgeWithDefaults() *FinancialConnectionDetailsSaltedge {
	this := FinancialConnectionDetailsSaltedge{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *FinancialConnectionDetailsSaltedge) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsSaltedge) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *FinancialConnectionDetailsSaltedge) SetConnectionId(v string) {
	o.ConnectionId = v
}

func (o FinancialConnectionDetailsSaltedge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsSaltedge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection_id"] = o.ConnectionId
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsSaltedge struct {
	value *FinancialConnectionDetailsSaltedge
	isSet bool
}

func (v NullableFinancialConnectionDetailsSaltedge) Get() *FinancialConnectionDetailsSaltedge {
	return v.value
}

func (v *NullableFinancialConnectionDetailsSaltedge) Set(val *FinancialConnectionDetailsSaltedge) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsSaltedge) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsSaltedge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsSaltedge(val *FinancialConnectionDetailsSaltedge) *NullableFinancialConnectionDetailsSaltedge {
	return &NullableFinancialConnectionDetailsSaltedge{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsSaltedge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsSaltedge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


