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

// checks if the CreateLinkTokenRequestPlaid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLinkTokenRequestPlaid{}

// CreateLinkTokenRequestPlaid An object specifying information about the Plaid configuration to use when creating a link token. 
type CreateLinkTokenRequestPlaid struct {
	// Follows the same schema as Plaid's Link Token Create Schema(https://plaid.com/docs/api/tokens/#linktokencreate). 'products', 'client_id', 'secret', 'client_user_id', 'client_name', 'webhook', 'institution_data' and 'country_codes' (only US and Canada is supported right now) will be set by Fuse and override any values you set.
	Config map[string]interface{} `json:"config,omitempty"`
}

// NewCreateLinkTokenRequestPlaid instantiates a new CreateLinkTokenRequestPlaid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkTokenRequestPlaid() *CreateLinkTokenRequestPlaid {
	this := CreateLinkTokenRequestPlaid{}
	return &this
}

// NewCreateLinkTokenRequestPlaidWithDefaults instantiates a new CreateLinkTokenRequestPlaid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkTokenRequestPlaidWithDefaults() *CreateLinkTokenRequestPlaid {
	this := CreateLinkTokenRequestPlaid{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateLinkTokenRequestPlaid) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkTokenRequestPlaid) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateLinkTokenRequestPlaid) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *CreateLinkTokenRequestPlaid) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o CreateLinkTokenRequestPlaid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLinkTokenRequestPlaid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableCreateLinkTokenRequestPlaid struct {
	value *CreateLinkTokenRequestPlaid
	isSet bool
}

func (v NullableCreateLinkTokenRequestPlaid) Get() *CreateLinkTokenRequestPlaid {
	return v.value
}

func (v *NullableCreateLinkTokenRequestPlaid) Set(val *CreateLinkTokenRequestPlaid) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkTokenRequestPlaid) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkTokenRequestPlaid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkTokenRequestPlaid(val *CreateLinkTokenRequestPlaid) *NullableCreateLinkTokenRequestPlaid {
	return &NullableCreateLinkTokenRequestPlaid{value: val, isSet: true}
}

func (v NullableCreateLinkTokenRequestPlaid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkTokenRequestPlaid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


