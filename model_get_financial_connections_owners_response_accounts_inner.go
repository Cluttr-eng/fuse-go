/*
Fuse

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuse

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetFinancialConnectionsOwnersResponseAccountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinancialConnectionsOwnersResponseAccountsInner{}

// GetFinancialConnectionsOwnersResponseAccountsInner struct for GetFinancialConnectionsOwnersResponseAccountsInner
type GetFinancialConnectionsOwnersResponseAccountsInner struct {
	// The remote account id of the account
	RemoteAccountId string `json:"remote_account_id"`
	Owners []FinancialConnectionsOwner `json:"owners"`
}

type _GetFinancialConnectionsOwnersResponseAccountsInner GetFinancialConnectionsOwnersResponseAccountsInner

// NewGetFinancialConnectionsOwnersResponseAccountsInner instantiates a new GetFinancialConnectionsOwnersResponseAccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinancialConnectionsOwnersResponseAccountsInner(remoteAccountId string, owners []FinancialConnectionsOwner) *GetFinancialConnectionsOwnersResponseAccountsInner {
	this := GetFinancialConnectionsOwnersResponseAccountsInner{}
	this.RemoteAccountId = remoteAccountId
	this.Owners = owners
	return &this
}

// NewGetFinancialConnectionsOwnersResponseAccountsInnerWithDefaults instantiates a new GetFinancialConnectionsOwnersResponseAccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinancialConnectionsOwnersResponseAccountsInnerWithDefaults() *GetFinancialConnectionsOwnersResponseAccountsInner {
	this := GetFinancialConnectionsOwnersResponseAccountsInner{}
	return &this
}

// GetRemoteAccountId returns the RemoteAccountId field value
func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetRemoteAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteAccountId
}

// GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetRemoteAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAccountId, true
}

// SetRemoteAccountId sets field value
func (o *GetFinancialConnectionsOwnersResponseAccountsInner) SetRemoteAccountId(v string) {
	o.RemoteAccountId = v
}

// GetOwners returns the Owners field value
func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetOwners() []FinancialConnectionsOwner {
	if o == nil {
		var ret []FinancialConnectionsOwner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetOwnersOk() ([]FinancialConnectionsOwner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owners, true
}

// SetOwners sets field value
func (o *GetFinancialConnectionsOwnersResponseAccountsInner) SetOwners(v []FinancialConnectionsOwner) {
	o.Owners = v
}

func (o GetFinancialConnectionsOwnersResponseAccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinancialConnectionsOwnersResponseAccountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remote_account_id"] = o.RemoteAccountId
	toSerialize["owners"] = o.Owners
	return toSerialize, nil
}

func (o *GetFinancialConnectionsOwnersResponseAccountsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"remote_account_id",
		"owners",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetFinancialConnectionsOwnersResponseAccountsInner := _GetFinancialConnectionsOwnersResponseAccountsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFinancialConnectionsOwnersResponseAccountsInner)

	if err != nil {
		return err
	}

	*o = GetFinancialConnectionsOwnersResponseAccountsInner(varGetFinancialConnectionsOwnersResponseAccountsInner)

	return err
}

type NullableGetFinancialConnectionsOwnersResponseAccountsInner struct {
	value *GetFinancialConnectionsOwnersResponseAccountsInner
	isSet bool
}

func (v NullableGetFinancialConnectionsOwnersResponseAccountsInner) Get() *GetFinancialConnectionsOwnersResponseAccountsInner {
	return v.value
}

func (v *NullableGetFinancialConnectionsOwnersResponseAccountsInner) Set(val *GetFinancialConnectionsOwnersResponseAccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinancialConnectionsOwnersResponseAccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinancialConnectionsOwnersResponseAccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinancialConnectionsOwnersResponseAccountsInner(val *GetFinancialConnectionsOwnersResponseAccountsInner) *NullableGetFinancialConnectionsOwnersResponseAccountsInner {
	return &NullableGetFinancialConnectionsOwnersResponseAccountsInner{value: val, isSet: true}
}

func (v NullableGetFinancialConnectionsOwnersResponseAccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinancialConnectionsOwnersResponseAccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


