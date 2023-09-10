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

// checks if the CreateLinkTokenRequestMx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLinkTokenRequestMx{}

// CreateLinkTokenRequestMx An object specifying information about the MX configuration to use for deciding which MX supported financial institutions to display.
type CreateLinkTokenRequestMx struct {
	// Follows the same schema as MX's request a connect url(https://docs.mx.com/api#connect_request_a_url) schema.
	Config map[string]interface{} `json:"config,omitempty"`
}

// NewCreateLinkTokenRequestMx instantiates a new CreateLinkTokenRequestMx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkTokenRequestMx() *CreateLinkTokenRequestMx {
	this := CreateLinkTokenRequestMx{}
	return &this
}

// NewCreateLinkTokenRequestMxWithDefaults instantiates a new CreateLinkTokenRequestMx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkTokenRequestMxWithDefaults() *CreateLinkTokenRequestMx {
	this := CreateLinkTokenRequestMx{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateLinkTokenRequestMx) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkTokenRequestMx) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateLinkTokenRequestMx) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *CreateLinkTokenRequestMx) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o CreateLinkTokenRequestMx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLinkTokenRequestMx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableCreateLinkTokenRequestMx struct {
	value *CreateLinkTokenRequestMx
	isSet bool
}

func (v NullableCreateLinkTokenRequestMx) Get() *CreateLinkTokenRequestMx {
	return v.value
}

func (v *NullableCreateLinkTokenRequestMx) Set(val *CreateLinkTokenRequestMx) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkTokenRequestMx) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkTokenRequestMx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkTokenRequestMx(val *CreateLinkTokenRequestMx) *NullableCreateLinkTokenRequestMx {
	return &NullableCreateLinkTokenRequestMx{value: val, isSet: true}
}

func (v NullableCreateLinkTokenRequestMx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkTokenRequestMx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


