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

// checks if the FinancialConnectionDetailsBasiq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsBasiq{}

// FinancialConnectionDetailsBasiq Data needed to query data from Basiq
type FinancialConnectionDetailsBasiq struct {
	// The identifier of the user for Basiq.
	UserId string `json:"user_id"`
	// The identifier of the connection for Basiq.
	ConnectionId string `json:"connection_id"`
}

// NewFinancialConnectionDetailsBasiq instantiates a new FinancialConnectionDetailsBasiq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsBasiq(userId string, connectionId string) *FinancialConnectionDetailsBasiq {
	this := FinancialConnectionDetailsBasiq{}
	this.UserId = userId
	this.ConnectionId = connectionId
	return &this
}

// NewFinancialConnectionDetailsBasiqWithDefaults instantiates a new FinancialConnectionDetailsBasiq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsBasiqWithDefaults() *FinancialConnectionDetailsBasiq {
	this := FinancialConnectionDetailsBasiq{}
	return &this
}

// GetUserId returns the UserId field value
func (o *FinancialConnectionDetailsBasiq) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsBasiq) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FinancialConnectionDetailsBasiq) SetUserId(v string) {
	o.UserId = v
}

// GetConnectionId returns the ConnectionId field value
func (o *FinancialConnectionDetailsBasiq) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsBasiq) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *FinancialConnectionDetailsBasiq) SetConnectionId(v string) {
	o.ConnectionId = v
}

func (o FinancialConnectionDetailsBasiq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsBasiq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["connection_id"] = o.ConnectionId
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsBasiq struct {
	value *FinancialConnectionDetailsBasiq
	isSet bool
}

func (v NullableFinancialConnectionDetailsBasiq) Get() *FinancialConnectionDetailsBasiq {
	return v.value
}

func (v *NullableFinancialConnectionDetailsBasiq) Set(val *FinancialConnectionDetailsBasiq) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsBasiq) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsBasiq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsBasiq(val *FinancialConnectionDetailsBasiq) *NullableFinancialConnectionDetailsBasiq {
	return &NullableFinancialConnectionDetailsBasiq{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsBasiq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsBasiq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


