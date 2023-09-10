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

// checks if the AssetReportAccountsInnerHistoricalBalancesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetReportAccountsInnerHistoricalBalancesInner{}

// AssetReportAccountsInnerHistoricalBalancesInner struct for AssetReportAccountsInnerHistoricalBalancesInner
type AssetReportAccountsInnerHistoricalBalancesInner struct {
	// The date of the calculated historical balance, in an ISO 8601 format (YYYY-MM-DD)
	Date *string `json:"date,omitempty"`
	// The total amount of funds in the account, calculated from the current balance in the balance object by subtracting inflows and adding back outflows according to the posted date of each transaction.
	Current *float32 `json:"current,omitempty"`
	// The ISO-4217 currency code of the balance.
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
}

// NewAssetReportAccountsInnerHistoricalBalancesInner instantiates a new AssetReportAccountsInnerHistoricalBalancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportAccountsInnerHistoricalBalancesInner() *AssetReportAccountsInnerHistoricalBalancesInner {
	this := AssetReportAccountsInnerHistoricalBalancesInner{}
	return &this
}

// NewAssetReportAccountsInnerHistoricalBalancesInnerWithDefaults instantiates a new AssetReportAccountsInnerHistoricalBalancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportAccountsInnerHistoricalBalancesInnerWithDefaults() *AssetReportAccountsInnerHistoricalBalancesInner {
	this := AssetReportAccountsInnerHistoricalBalancesInner{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) SetDate(v string) {
	o.Date = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetCurrent() float32 {
	if o == nil || IsNil(o.Current) {
		var ret float32
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetCurrentOk() (*float32, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given float32 and assigns it to the Current field.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) SetCurrent(v float32) {
	o.Current = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetIsoCurrencyCode() string {
	if o == nil || IsNil(o.IsoCurrencyCode) {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.IsoCurrencyCode) {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) HasIsoCurrencyCode() bool {
	if o != nil && !IsNil(o.IsoCurrencyCode) {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *AssetReportAccountsInnerHistoricalBalancesInner) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

func (o AssetReportAccountsInnerHistoricalBalancesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetReportAccountsInnerHistoricalBalancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.IsoCurrencyCode) {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	return toSerialize, nil
}

type NullableAssetReportAccountsInnerHistoricalBalancesInner struct {
	value *AssetReportAccountsInnerHistoricalBalancesInner
	isSet bool
}

func (v NullableAssetReportAccountsInnerHistoricalBalancesInner) Get() *AssetReportAccountsInnerHistoricalBalancesInner {
	return v.value
}

func (v *NullableAssetReportAccountsInnerHistoricalBalancesInner) Set(val *AssetReportAccountsInnerHistoricalBalancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportAccountsInnerHistoricalBalancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportAccountsInnerHistoricalBalancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportAccountsInnerHistoricalBalancesInner(val *AssetReportAccountsInnerHistoricalBalancesInner) *NullableAssetReportAccountsInnerHistoricalBalancesInner {
	return &NullableAssetReportAccountsInnerHistoricalBalancesInner{value: val, isSet: true}
}

func (v NullableAssetReportAccountsInnerHistoricalBalancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportAccountsInnerHistoricalBalancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


