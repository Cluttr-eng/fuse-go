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

// checks if the FinancialConnectionDetailsTeller type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsTeller{}

// FinancialConnectionDetailsTeller Data needed to query data from Teller
type FinancialConnectionDetailsTeller struct {
	// Access token for Teller
	AccessToken string `json:"access_token"`
	// Enrollment ID associated with the access token in Teller
	EnrollmentId string `json:"enrollment_id"`
}

// NewFinancialConnectionDetailsTeller instantiates a new FinancialConnectionDetailsTeller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsTeller(accessToken string, enrollmentId string) *FinancialConnectionDetailsTeller {
	this := FinancialConnectionDetailsTeller{}
	this.AccessToken = accessToken
	this.EnrollmentId = enrollmentId
	return &this
}

// NewFinancialConnectionDetailsTellerWithDefaults instantiates a new FinancialConnectionDetailsTeller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsTellerWithDefaults() *FinancialConnectionDetailsTeller {
	this := FinancialConnectionDetailsTeller{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *FinancialConnectionDetailsTeller) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsTeller) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *FinancialConnectionDetailsTeller) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetEnrollmentId returns the EnrollmentId field value
func (o *FinancialConnectionDetailsTeller) GetEnrollmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnrollmentId
}

// GetEnrollmentIdOk returns a tuple with the EnrollmentId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsTeller) GetEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentId, true
}

// SetEnrollmentId sets field value
func (o *FinancialConnectionDetailsTeller) SetEnrollmentId(v string) {
	o.EnrollmentId = v
}

func (o FinancialConnectionDetailsTeller) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsTeller) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["enrollment_id"] = o.EnrollmentId
	return toSerialize, nil
}

type NullableFinancialConnectionDetailsTeller struct {
	value *FinancialConnectionDetailsTeller
	isSet bool
}

func (v NullableFinancialConnectionDetailsTeller) Get() *FinancialConnectionDetailsTeller {
	return v.value
}

func (v *NullableFinancialConnectionDetailsTeller) Set(val *FinancialConnectionDetailsTeller) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsTeller) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsTeller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsTeller(val *FinancialConnectionDetailsTeller) *NullableFinancialConnectionDetailsTeller {
	return &NullableFinancialConnectionDetailsTeller{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsTeller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsTeller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


