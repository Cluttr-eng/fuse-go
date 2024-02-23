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

// checks if the FinancialConnectionDetailsSophtron type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsSophtron{}

// FinancialConnectionDetailsSophtron Data needed to query data from Sophtron.
type FinancialConnectionDetailsSophtron struct {
	UserInstitutionId string `json:"user_institution_id"`
}

type _FinancialConnectionDetailsSophtron FinancialConnectionDetailsSophtron

// NewFinancialConnectionDetailsSophtron instantiates a new FinancialConnectionDetailsSophtron object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsSophtron(userInstitutionId string) *FinancialConnectionDetailsSophtron {
	this := FinancialConnectionDetailsSophtron{}
	this.UserInstitutionId = userInstitutionId
	return &this
}

// NewFinancialConnectionDetailsSophtronWithDefaults instantiates a new FinancialConnectionDetailsSophtron object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsSophtronWithDefaults() *FinancialConnectionDetailsSophtron {
	this := FinancialConnectionDetailsSophtron{}
	return &this
}

// GetUserInstitutionId returns the UserInstitutionId field value
func (o *FinancialConnectionDetailsSophtron) GetUserInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserInstitutionId
}

// GetUserInstitutionIdOk returns a tuple with the UserInstitutionId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsSophtron) GetUserInstitutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInstitutionId, true
}

// SetUserInstitutionId sets field value
func (o *FinancialConnectionDetailsSophtron) SetUserInstitutionId(v string) {
	o.UserInstitutionId = v
}

func (o FinancialConnectionDetailsSophtron) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsSophtron) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_institution_id"] = o.UserInstitutionId
	return toSerialize, nil
}

func (o *FinancialConnectionDetailsSophtron) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_institution_id",
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

	varFinancialConnectionDetailsSophtron := _FinancialConnectionDetailsSophtron{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialConnectionDetailsSophtron)

	if err != nil {
		return err
	}

	*o = FinancialConnectionDetailsSophtron(varFinancialConnectionDetailsSophtron)

	return err
}

type NullableFinancialConnectionDetailsSophtron struct {
	value *FinancialConnectionDetailsSophtron
	isSet bool
}

func (v NullableFinancialConnectionDetailsSophtron) Get() *FinancialConnectionDetailsSophtron {
	return v.value
}

func (v *NullableFinancialConnectionDetailsSophtron) Set(val *FinancialConnectionDetailsSophtron) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsSophtron) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsSophtron) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsSophtron(val *FinancialConnectionDetailsSophtron) *NullableFinancialConnectionDetailsSophtron {
	return &NullableFinancialConnectionDetailsSophtron{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsSophtron) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsSophtron) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

