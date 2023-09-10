/*
Fuse

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuse

import (
	"encoding/json"
	"fmt"
)

// AccountType The account's type. '-' means we were not able to map the upstream type.
type AccountType string

// List of AccountType
const (
	ACCOUNTTYPE_DEPOSITORY AccountType = "depository"
	ACCOUNTTYPE_CREDIT AccountType = "credit"
	ACCOUNTTYPE_LOAN AccountType = "loan"
	ACCOUNTTYPE_INVESTMENT AccountType = "investment"
	ACCOUNTTYPE_INSURANCE AccountType = "insurance"
	ACCOUNTTYPE_PROPERTY AccountType = "property"
	ACCOUNTTYPE_ANNUITY AccountType = "annuity"
	ACCOUNTTYPE_OTHER AccountType = "other"
	ACCOUNTTYPE_MINUS AccountType = "-"
)

// All allowed values of AccountType enum
var AllowedAccountTypeEnumValues = []AccountType{
	"depository",
	"credit",
	"loan",
	"investment",
	"insurance",
	"property",
	"annuity",
	"other",
	"-",
}

func (v *AccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountType(value)
	for _, existing := range AllowedAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountType", value)
}

// NewAccountTypeFromValue returns a pointer to a valid AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypeFromValue(v string) (*AccountType, error) {
	ev := AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountType: valid values are %v", v, AllowedAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountType) IsValid() bool {
	for _, existing := range AllowedAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountType value
func (v AccountType) Ptr() *AccountType {
	return &v
}

type NullableAccountType struct {
	value *AccountType
	isSet bool
}

func (v NullableAccountType) Get() *AccountType {
	return v.value
}

func (v *NullableAccountType) Set(val *AccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountType(val *AccountType) *NullableAccountType {
	return &NullableAccountType{value: val, isSet: true}
}

func (v NullableAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
