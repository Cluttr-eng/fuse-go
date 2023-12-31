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

// checks if the GetLiabilitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLiabilitiesRequest{}

// GetLiabilitiesRequest struct for GetLiabilitiesRequest
type GetLiabilitiesRequest struct {
	// Access token for authentication
	AccessToken string `json:"access_token"`
}

// NewGetLiabilitiesRequest instantiates a new GetLiabilitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLiabilitiesRequest(accessToken string) *GetLiabilitiesRequest {
	this := GetLiabilitiesRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewGetLiabilitiesRequestWithDefaults instantiates a new GetLiabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLiabilitiesRequestWithDefaults() *GetLiabilitiesRequest {
	this := GetLiabilitiesRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *GetLiabilitiesRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *GetLiabilitiesRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *GetLiabilitiesRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o GetLiabilitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLiabilitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	return toSerialize, nil
}

type NullableGetLiabilitiesRequest struct {
	value *GetLiabilitiesRequest
	isSet bool
}

func (v NullableGetLiabilitiesRequest) Get() *GetLiabilitiesRequest {
	return v.value
}

func (v *NullableGetLiabilitiesRequest) Set(val *GetLiabilitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLiabilitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLiabilitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLiabilitiesRequest(val *GetLiabilitiesRequest) *NullableGetLiabilitiesRequest {
	return &NullableGetLiabilitiesRequest{value: val, isSet: true}
}

func (v NullableGetLiabilitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLiabilitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


