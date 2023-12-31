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

// checks if the FinancialConnectionsAccountDetailsAccountNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsAccountDetailsAccountNumber{}

// FinancialConnectionsAccountDetailsAccountNumber struct for FinancialConnectionsAccountDetailsAccountNumber
type FinancialConnectionsAccountDetailsAccountNumber struct {
	// Unique identifier representing the account, typically referred to as the account number.
	Number string `json:"number"`
	// A six-digit number used by banks in the United Kingdom and Ireland to identify the branch where the account is held.
	SortCode *string `json:"sort_code,omitempty"`
	// International Bank Account Number (IBAN) is an internationally agreed system of identifying bank accounts across national borders to facilitate the communication and processing of cross border transactions.
	Iban *string `json:"iban,omitempty"`
	// SWIFT/BIC code is an international identifier used to distinctively recognize a particular bank during financial transactions, such as money transfers.
	SwiftBic *string `json:"swift_bic,omitempty"`
	// Bank-State-Branch (BSB) code is a six-digit numerical code used to identify an individual branch of a financial institution in Australia.
	Bsb *string `json:"bsb,omitempty"`
	// Bank Identifier Code (BIC) is an international standard identifier used by banks and financial institutions globally to carry out transactions.
	Bic *string `json:"bic,omitempty"`
}

// NewFinancialConnectionsAccountDetailsAccountNumber instantiates a new FinancialConnectionsAccountDetailsAccountNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsAccountDetailsAccountNumber(number string) *FinancialConnectionsAccountDetailsAccountNumber {
	this := FinancialConnectionsAccountDetailsAccountNumber{}
	this.Number = number
	return &this
}

// NewFinancialConnectionsAccountDetailsAccountNumberWithDefaults instantiates a new FinancialConnectionsAccountDetailsAccountNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsAccountDetailsAccountNumberWithDefaults() *FinancialConnectionsAccountDetailsAccountNumber {
	this := FinancialConnectionsAccountDetailsAccountNumber{}
	return &this
}

// GetNumber returns the Number field value
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *FinancialConnectionsAccountDetailsAccountNumber) SetNumber(v string) {
	o.Number = v
}

// GetSortCode returns the SortCode field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSortCode() string {
	if o == nil || IsNil(o.SortCode) {
		var ret string
		return ret
	}
	return *o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SortCode) {
		return nil, false
	}
	return o.SortCode, true
}

// HasSortCode returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) HasSortCode() bool {
	if o != nil && !IsNil(o.SortCode) {
		return true
	}

	return false
}

// SetSortCode gets a reference to the given string and assigns it to the SortCode field.
func (o *FinancialConnectionsAccountDetailsAccountNumber) SetSortCode(v string) {
	o.SortCode = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetIban() string {
	if o == nil || IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetIbanOk() (*string, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) HasIban() bool {
	if o != nil && !IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *FinancialConnectionsAccountDetailsAccountNumber) SetIban(v string) {
	o.Iban = &v
}

// GetSwiftBic returns the SwiftBic field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSwiftBic() string {
	if o == nil || IsNil(o.SwiftBic) {
		var ret string
		return ret
	}
	return *o.SwiftBic
}

// GetSwiftBicOk returns a tuple with the SwiftBic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSwiftBicOk() (*string, bool) {
	if o == nil || IsNil(o.SwiftBic) {
		return nil, false
	}
	return o.SwiftBic, true
}

// HasSwiftBic returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) HasSwiftBic() bool {
	if o != nil && !IsNil(o.SwiftBic) {
		return true
	}

	return false
}

// SetSwiftBic gets a reference to the given string and assigns it to the SwiftBic field.
func (o *FinancialConnectionsAccountDetailsAccountNumber) SetSwiftBic(v string) {
	o.SwiftBic = &v
}

// GetBsb returns the Bsb field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBsb() string {
	if o == nil || IsNil(o.Bsb) {
		var ret string
		return ret
	}
	return *o.Bsb
}

// GetBsbOk returns a tuple with the Bsb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBsbOk() (*string, bool) {
	if o == nil || IsNil(o.Bsb) {
		return nil, false
	}
	return o.Bsb, true
}

// HasBsb returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) HasBsb() bool {
	if o != nil && !IsNil(o.Bsb) {
		return true
	}

	return false
}

// SetBsb gets a reference to the given string and assigns it to the Bsb field.
func (o *FinancialConnectionsAccountDetailsAccountNumber) SetBsb(v string) {
	o.Bsb = &v
}

// GetBic returns the Bic field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBic() string {
	if o == nil || IsNil(o.Bic) {
		var ret string
		return ret
	}
	return *o.Bic
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBicOk() (*string, bool) {
	if o == nil || IsNil(o.Bic) {
		return nil, false
	}
	return o.Bic, true
}

// HasBic returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetailsAccountNumber) HasBic() bool {
	if o != nil && !IsNil(o.Bic) {
		return true
	}

	return false
}

// SetBic gets a reference to the given string and assigns it to the Bic field.
func (o *FinancialConnectionsAccountDetailsAccountNumber) SetBic(v string) {
	o.Bic = &v
}

func (o FinancialConnectionsAccountDetailsAccountNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsAccountDetailsAccountNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	if !IsNil(o.SortCode) {
		toSerialize["sort_code"] = o.SortCode
	}
	if !IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !IsNil(o.SwiftBic) {
		toSerialize["swift_bic"] = o.SwiftBic
	}
	if !IsNil(o.Bsb) {
		toSerialize["bsb"] = o.Bsb
	}
	if !IsNil(o.Bic) {
		toSerialize["bic"] = o.Bic
	}
	return toSerialize, nil
}

type NullableFinancialConnectionsAccountDetailsAccountNumber struct {
	value *FinancialConnectionsAccountDetailsAccountNumber
	isSet bool
}

func (v NullableFinancialConnectionsAccountDetailsAccountNumber) Get() *FinancialConnectionsAccountDetailsAccountNumber {
	return v.value
}

func (v *NullableFinancialConnectionsAccountDetailsAccountNumber) Set(val *FinancialConnectionsAccountDetailsAccountNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsAccountDetailsAccountNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsAccountDetailsAccountNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsAccountDetailsAccountNumber(val *FinancialConnectionsAccountDetailsAccountNumber) *NullableFinancialConnectionsAccountDetailsAccountNumber {
	return &NullableFinancialConnectionsAccountDetailsAccountNumber{value: val, isSet: true}
}

func (v NullableFinancialConnectionsAccountDetailsAccountNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsAccountDetailsAccountNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


