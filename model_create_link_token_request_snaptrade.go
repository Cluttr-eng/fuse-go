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

// checks if the CreateLinkTokenRequestSnaptrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLinkTokenRequestSnaptrade{}

// CreateLinkTokenRequestSnaptrade An object specifying information about the SnapTrade configuration to use when creating a link token. 
type CreateLinkTokenRequestSnaptrade struct {
	Config *CreateLinkTokenRequestSnaptradeConfig `json:"config,omitempty"`
}

// NewCreateLinkTokenRequestSnaptrade instantiates a new CreateLinkTokenRequestSnaptrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkTokenRequestSnaptrade() *CreateLinkTokenRequestSnaptrade {
	this := CreateLinkTokenRequestSnaptrade{}
	return &this
}

// NewCreateLinkTokenRequestSnaptradeWithDefaults instantiates a new CreateLinkTokenRequestSnaptrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkTokenRequestSnaptradeWithDefaults() *CreateLinkTokenRequestSnaptrade {
	this := CreateLinkTokenRequestSnaptrade{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateLinkTokenRequestSnaptrade) GetConfig() CreateLinkTokenRequestSnaptradeConfig {
	if o == nil || IsNil(o.Config) {
		var ret CreateLinkTokenRequestSnaptradeConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkTokenRequestSnaptrade) GetConfigOk() (*CreateLinkTokenRequestSnaptradeConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateLinkTokenRequestSnaptrade) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given CreateLinkTokenRequestSnaptradeConfig and assigns it to the Config field.
func (o *CreateLinkTokenRequestSnaptrade) SetConfig(v CreateLinkTokenRequestSnaptradeConfig) {
	o.Config = &v
}

func (o CreateLinkTokenRequestSnaptrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLinkTokenRequestSnaptrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableCreateLinkTokenRequestSnaptrade struct {
	value *CreateLinkTokenRequestSnaptrade
	isSet bool
}

func (v NullableCreateLinkTokenRequestSnaptrade) Get() *CreateLinkTokenRequestSnaptrade {
	return v.value
}

func (v *NullableCreateLinkTokenRequestSnaptrade) Set(val *CreateLinkTokenRequestSnaptrade) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkTokenRequestSnaptrade) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkTokenRequestSnaptrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkTokenRequestSnaptrade(val *CreateLinkTokenRequestSnaptrade) *NullableCreateLinkTokenRequestSnaptrade {
	return &NullableCreateLinkTokenRequestSnaptrade{value: val, isSet: true}
}

func (v NullableCreateLinkTokenRequestSnaptrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkTokenRequestSnaptrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


