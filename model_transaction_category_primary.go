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

// TransactionCategoryPrimary Primary transaction category
type TransactionCategoryPrimary string

// List of TransactionCategoryPrimary
const (
	TRANSACTIONCATEGORYPRIMARY_AUTO_AND_TRANSPORT TransactionCategoryPrimary = "auto_and_transport"
	TRANSACTIONCATEGORYPRIMARY_BILLS_AND_UTILITIES TransactionCategoryPrimary = "bills_and_utilities"
	TRANSACTIONCATEGORYPRIMARY_BUSINESS_SERVICES TransactionCategoryPrimary = "business_services"
	TRANSACTIONCATEGORYPRIMARY_EDUCATION TransactionCategoryPrimary = "education"
	TRANSACTIONCATEGORYPRIMARY_ENTERTAINMENT TransactionCategoryPrimary = "entertainment"
	TRANSACTIONCATEGORYPRIMARY_FEES_AND_CHARGES TransactionCategoryPrimary = "fees_and_charges"
	TRANSACTIONCATEGORYPRIMARY_FINANCIAL TransactionCategoryPrimary = "financial"
	TRANSACTIONCATEGORYPRIMARY_FOOD_AND_DINING TransactionCategoryPrimary = "food_and_dining"
	TRANSACTIONCATEGORYPRIMARY_GIFTS_AND_DONATIONS TransactionCategoryPrimary = "gifts_and_donations"
	TRANSACTIONCATEGORYPRIMARY_HEALTH_AND_FITNESS TransactionCategoryPrimary = "health_and_fitness"
	TRANSACTIONCATEGORYPRIMARY_HOME TransactionCategoryPrimary = "home"
	TRANSACTIONCATEGORYPRIMARY_INCOME TransactionCategoryPrimary = "income"
	TRANSACTIONCATEGORYPRIMARY_INVESTMENTS TransactionCategoryPrimary = "investments"
	TRANSACTIONCATEGORYPRIMARY_KIDS TransactionCategoryPrimary = "kids"
	TRANSACTIONCATEGORYPRIMARY_PERSONAL_CARE TransactionCategoryPrimary = "personal_care"
	TRANSACTIONCATEGORYPRIMARY_PETS TransactionCategoryPrimary = "pets"
	TRANSACTIONCATEGORYPRIMARY_SHOPPING TransactionCategoryPrimary = "shopping"
	TRANSACTIONCATEGORYPRIMARY_TAXES TransactionCategoryPrimary = "taxes"
	TRANSACTIONCATEGORYPRIMARY_TRANSFER TransactionCategoryPrimary = "transfer"
	TRANSACTIONCATEGORYPRIMARY_TRAVEL TransactionCategoryPrimary = "travel"
	TRANSACTIONCATEGORYPRIMARY_UNCATEGORIZED TransactionCategoryPrimary = "uncategorized"
)

// All allowed values of TransactionCategoryPrimary enum
var AllowedTransactionCategoryPrimaryEnumValues = []TransactionCategoryPrimary{
	"auto_and_transport",
	"bills_and_utilities",
	"business_services",
	"education",
	"entertainment",
	"fees_and_charges",
	"financial",
	"food_and_dining",
	"gifts_and_donations",
	"health_and_fitness",
	"home",
	"income",
	"investments",
	"kids",
	"personal_care",
	"pets",
	"shopping",
	"taxes",
	"transfer",
	"travel",
	"uncategorized",
}

func (v *TransactionCategoryPrimary) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionCategoryPrimary(value)
	for _, existing := range AllowedTransactionCategoryPrimaryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionCategoryPrimary", value)
}

// NewTransactionCategoryPrimaryFromValue returns a pointer to a valid TransactionCategoryPrimary
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionCategoryPrimaryFromValue(v string) (*TransactionCategoryPrimary, error) {
	ev := TransactionCategoryPrimary(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionCategoryPrimary: valid values are %v", v, AllowedTransactionCategoryPrimaryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionCategoryPrimary) IsValid() bool {
	for _, existing := range AllowedTransactionCategoryPrimaryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionCategoryPrimary value
func (v TransactionCategoryPrimary) Ptr() *TransactionCategoryPrimary {
	return &v
}

type NullableTransactionCategoryPrimary struct {
	value *TransactionCategoryPrimary
	isSet bool
}

func (v NullableTransactionCategoryPrimary) Get() *TransactionCategoryPrimary {
	return v.value
}

func (v *NullableTransactionCategoryPrimary) Set(val *TransactionCategoryPrimary) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCategoryPrimary) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCategoryPrimary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCategoryPrimary(val *TransactionCategoryPrimary) *NullableTransactionCategoryPrimary {
	return &NullableTransactionCategoryPrimary{value: val, isSet: true}
}

func (v NullableTransactionCategoryPrimary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCategoryPrimary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
