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

// checks if the FinancialConnectionDetailsAkoya type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsAkoya{}

// FinancialConnectionDetailsAkoya Data needed to query data from Akoya.
type FinancialConnectionDetailsAkoya struct {
	IdToken string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt string `json:"expires_at"`
}

// NewFinancialConnectionDetailsAkoya instantiates a new FinancialConnectionDetailsAkoya object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsAkoya(idToken string, refreshToken string, expiresAt string) *FinancialConnectionDetailsAkoya {
	this := FinancialConnectionDetailsAkoya{}
	this.IdToken = idToken
	this.RefreshToken = refreshToken
	this.ExpiresAt = expiresAt
	return &this
}

// NewFinancialConnectionDetailsAkoyaWithDefaults instantiates a new FinancialConnectionDetailsAkoya object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsAkoyaWithDefaults() *FinancialConnectionDetailsAkoya {
	this := FinancialConnectionDetailsAkoya{}
	return &this
}

// GetIdToken returns the IdToken field value
func (o *FinancialConnectionDetailsAkoya) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsAkoya) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *FinancialConnectionDetailsAkoya) SetIdToken(v string) {
	o.IdToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *FinancialConnectionDetailsAkoya) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsAkoya) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *FinancialConnectionDetailsAkoya) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *FinancialConnectionDetailsAkoya) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsAkoya) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *FinancialConnectionDetailsAkoya) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

func (o FinancialConnectionDetailsAkoya) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsAkoya) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id_token"] = o.IdToken
	toSerialize["refresh_token"] = o.RefreshToken
	toSerialize["expires_at"] = o.ExpiresAt
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsAkoya struct {
	value *FinancialConnectionDetailsAkoya
	isSet bool
}

func (v NullableFinancialConnectionDetailsAkoya) Get() *FinancialConnectionDetailsAkoya {
	return v.value
}

func (v *NullableFinancialConnectionDetailsAkoya) Set(val *FinancialConnectionDetailsAkoya) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsAkoya) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsAkoya) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsAkoya(val *FinancialConnectionDetailsAkoya) *NullableFinancialConnectionDetailsAkoya {
	return &NullableFinancialConnectionDetailsAkoya{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsAkoya) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsAkoya) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

