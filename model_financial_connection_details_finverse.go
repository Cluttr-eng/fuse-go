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

// checks if the FinancialConnectionDetailsFinverse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsFinverse{}

// FinancialConnectionDetailsFinverse Data needed to query data from Finverse
type FinancialConnectionDetailsFinverse struct {
	// Access token for Finverse
	AccessToken string `json:"access_token"`
	// Login Identity Id for Finverse
	LoginIdentityId *string `json:"login_identity_id,omitempty"`
}

// NewFinancialConnectionDetailsFinverse instantiates a new FinancialConnectionDetailsFinverse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsFinverse(accessToken string) *FinancialConnectionDetailsFinverse {
	this := FinancialConnectionDetailsFinverse{}
	this.AccessToken = accessToken
	return &this
}

// NewFinancialConnectionDetailsFinverseWithDefaults instantiates a new FinancialConnectionDetailsFinverse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsFinverseWithDefaults() *FinancialConnectionDetailsFinverse {
	this := FinancialConnectionDetailsFinverse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *FinancialConnectionDetailsFinverse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsFinverse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *FinancialConnectionDetailsFinverse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetLoginIdentityId returns the LoginIdentityId field value if set, zero value otherwise.
func (o *FinancialConnectionDetailsFinverse) GetLoginIdentityId() string {
	if o == nil || IsNil(o.LoginIdentityId) {
		var ret string
		return ret
	}
	return *o.LoginIdentityId
}

// GetLoginIdentityIdOk returns a tuple with the LoginIdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsFinverse) GetLoginIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.LoginIdentityId) {
		return nil, false
	}
	return o.LoginIdentityId, true
}

// HasLoginIdentityId returns a boolean if a field has been set.
func (o *FinancialConnectionDetailsFinverse) HasLoginIdentityId() bool {
	if o != nil && !IsNil(o.LoginIdentityId) {
		return true
	}

	return false
}

// SetLoginIdentityId gets a reference to the given string and assigns it to the LoginIdentityId field.
func (o *FinancialConnectionDetailsFinverse) SetLoginIdentityId(v string) {
	o.LoginIdentityId = &v
}

func (o FinancialConnectionDetailsFinverse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsFinverse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	if !IsNil(o.LoginIdentityId) {
		toSerialize["login_identity_id"] = o.LoginIdentityId
	}
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsFinverse struct {
	value *FinancialConnectionDetailsFinverse
	isSet bool
}

func (v NullableFinancialConnectionDetailsFinverse) Get() *FinancialConnectionDetailsFinverse {
	return v.value
}

func (v *NullableFinancialConnectionDetailsFinverse) Set(val *FinancialConnectionDetailsFinverse) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsFinverse) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsFinverse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsFinverse(val *FinancialConnectionDetailsFinverse) *NullableFinancialConnectionDetailsFinverse {
	return &NullableFinancialConnectionDetailsFinverse{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsFinverse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsFinverse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


