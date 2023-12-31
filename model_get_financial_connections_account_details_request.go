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

// checks if the GetFinancialConnectionsAccountDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinancialConnectionsAccountDetailsRequest{}

// GetFinancialConnectionsAccountDetailsRequest struct for GetFinancialConnectionsAccountDetailsRequest
type GetFinancialConnectionsAccountDetailsRequest struct {
	// Access token for authentication
	AccessToken string `json:"access_token"`
}

// NewGetFinancialConnectionsAccountDetailsRequest instantiates a new GetFinancialConnectionsAccountDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinancialConnectionsAccountDetailsRequest(accessToken string) *GetFinancialConnectionsAccountDetailsRequest {
	this := GetFinancialConnectionsAccountDetailsRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewGetFinancialConnectionsAccountDetailsRequestWithDefaults instantiates a new GetFinancialConnectionsAccountDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinancialConnectionsAccountDetailsRequestWithDefaults() *GetFinancialConnectionsAccountDetailsRequest {
	this := GetFinancialConnectionsAccountDetailsRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *GetFinancialConnectionsAccountDetailsRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountDetailsRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *GetFinancialConnectionsAccountDetailsRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o GetFinancialConnectionsAccountDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinancialConnectionsAccountDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	return toSerialize, nil
}

type NullableGetFinancialConnectionsAccountDetailsRequest struct {
	value *GetFinancialConnectionsAccountDetailsRequest
	isSet bool
}

func (v NullableGetFinancialConnectionsAccountDetailsRequest) Get() *GetFinancialConnectionsAccountDetailsRequest {
	return v.value
}

func (v *NullableGetFinancialConnectionsAccountDetailsRequest) Set(val *GetFinancialConnectionsAccountDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinancialConnectionsAccountDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinancialConnectionsAccountDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinancialConnectionsAccountDetailsRequest(val *GetFinancialConnectionsAccountDetailsRequest) *NullableGetFinancialConnectionsAccountDetailsRequest {
	return &NullableGetFinancialConnectionsAccountDetailsRequest{value: val, isSet: true}
}

func (v NullableGetFinancialConnectionsAccountDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinancialConnectionsAccountDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


