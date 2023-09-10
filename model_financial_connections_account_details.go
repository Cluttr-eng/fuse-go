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

// checks if the FinancialConnectionsAccountDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsAccountDetails{}

// FinancialConnectionsAccountDetails struct for FinancialConnectionsAccountDetails
type FinancialConnectionsAccountDetails struct {
	// Remote Id of the account, ie Plaid or Teller account id
	RemoteId string `json:"remote_id"`
	Ach *FinancialConnectionsAccountDetailsAch `json:"ach,omitempty"`
	AccountNumber *FinancialConnectionsAccountDetailsAccountNumber `json:"account_number,omitempty"`
	RemoteData interface{} `json:"remote_data"`
}

// NewFinancialConnectionsAccountDetails instantiates a new FinancialConnectionsAccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsAccountDetails(remoteId string, remoteData interface{}) *FinancialConnectionsAccountDetails {
	this := FinancialConnectionsAccountDetails{}
	this.RemoteId = remoteId
	this.RemoteData = remoteData
	return &this
}

// NewFinancialConnectionsAccountDetailsWithDefaults instantiates a new FinancialConnectionsAccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsAccountDetailsWithDefaults() *FinancialConnectionsAccountDetails {
	this := FinancialConnectionsAccountDetails{}
	return &this
}

// GetRemoteId returns the RemoteId field value
func (o *FinancialConnectionsAccountDetails) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetails) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *FinancialConnectionsAccountDetails) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetAch returns the Ach field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetails) GetAch() FinancialConnectionsAccountDetailsAch {
	if o == nil || IsNil(o.Ach) {
		var ret FinancialConnectionsAccountDetailsAch
		return ret
	}
	return *o.Ach
}

// GetAchOk returns a tuple with the Ach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetails) GetAchOk() (*FinancialConnectionsAccountDetailsAch, bool) {
	if o == nil || IsNil(o.Ach) {
		return nil, false
	}
	return o.Ach, true
}

// HasAch returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetails) HasAch() bool {
	if o != nil && !IsNil(o.Ach) {
		return true
	}

	return false
}

// SetAch gets a reference to the given FinancialConnectionsAccountDetailsAch and assigns it to the Ach field.
func (o *FinancialConnectionsAccountDetails) SetAch(v FinancialConnectionsAccountDetailsAch) {
	o.Ach = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountDetails) GetAccountNumber() FinancialConnectionsAccountDetailsAccountNumber {
	if o == nil || IsNil(o.AccountNumber) {
		var ret FinancialConnectionsAccountDetailsAccountNumber
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountDetails) GetAccountNumberOk() (*FinancialConnectionsAccountDetailsAccountNumber, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountDetails) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given FinancialConnectionsAccountDetailsAccountNumber and assigns it to the AccountNumber field.
func (o *FinancialConnectionsAccountDetails) SetAccountNumber(v FinancialConnectionsAccountDetailsAccountNumber) {
	o.AccountNumber = &v
}

// GetRemoteData returns the RemoteData field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FinancialConnectionsAccountDetails) GetRemoteData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FinancialConnectionsAccountDetails) GetRemoteDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RemoteData) {
		return nil, false
	}
	return &o.RemoteData, true
}

// SetRemoteData sets field value
func (o *FinancialConnectionsAccountDetails) SetRemoteData(v interface{}) {
	o.RemoteData = v
}

func (o FinancialConnectionsAccountDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsAccountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remote_id"] = o.RemoteId
	if !IsNil(o.Ach) {
		toSerialize["ach"] = o.Ach
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return toSerialize, nil
}

type NullableFinancialConnectionsAccountDetails struct {
	value *FinancialConnectionsAccountDetails
	isSet bool
}

func (v NullableFinancialConnectionsAccountDetails) Get() *FinancialConnectionsAccountDetails {
	return v.value
}

func (v *NullableFinancialConnectionsAccountDetails) Set(val *FinancialConnectionsAccountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsAccountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsAccountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsAccountDetails(val *FinancialConnectionsAccountDetails) *NullableFinancialConnectionsAccountDetails {
	return &NullableFinancialConnectionsAccountDetails{value: val, isSet: true}
}

func (v NullableFinancialConnectionsAccountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsAccountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


