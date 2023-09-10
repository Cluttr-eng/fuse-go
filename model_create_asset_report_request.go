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

// checks if the CreateAssetReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetReportRequest{}

// CreateAssetReportRequest struct for CreateAssetReportRequest
type CreateAssetReportRequest struct {
	// Access fuse token corresponding to the financial account to be create the Asset Report for.
	AccessToken string `json:"access_token"`
	// The maximum integer number of days of history to include in the Asset Report
	DaysRequested float32 `json:"days_requested"`
}

// NewCreateAssetReportRequest instantiates a new CreateAssetReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetReportRequest(accessToken string, daysRequested float32) *CreateAssetReportRequest {
	this := CreateAssetReportRequest{}
	this.AccessToken = accessToken
	this.DaysRequested = daysRequested
	return &this
}

// NewCreateAssetReportRequestWithDefaults instantiates a new CreateAssetReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetReportRequestWithDefaults() *CreateAssetReportRequest {
	this := CreateAssetReportRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *CreateAssetReportRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *CreateAssetReportRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *CreateAssetReportRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetDaysRequested returns the DaysRequested field value
func (o *CreateAssetReportRequest) GetDaysRequested() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value
// and a boolean to check if the value has been set.
func (o *CreateAssetReportRequest) GetDaysRequestedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysRequested, true
}

// SetDaysRequested sets field value
func (o *CreateAssetReportRequest) SetDaysRequested(v float32) {
	o.DaysRequested = v
}

func (o CreateAssetReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["days_requested"] = o.DaysRequested
	return toSerialize, nil
}

type NullableCreateAssetReportRequest struct {
	value *CreateAssetReportRequest
	isSet bool
}

func (v NullableCreateAssetReportRequest) Get() *CreateAssetReportRequest {
	return v.value
}

func (v *NullableCreateAssetReportRequest) Set(val *CreateAssetReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetReportRequest(val *CreateAssetReportRequest) *NullableCreateAssetReportRequest {
	return &NullableCreateAssetReportRequest{value: val, isSet: true}
}

func (v NullableCreateAssetReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


