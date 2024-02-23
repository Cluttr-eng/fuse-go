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

// checks if the FinancialConnectionsAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsAccount{}

// FinancialConnectionsAccount struct for FinancialConnectionsAccount
type FinancialConnectionsAccount struct {
	// Remote Id of the account, ie Plaid or Teller account id
	RemoteId string `json:"remote_id"`
	// The ISO-4217 currency code of the account.
	Currency string `json:"currency"`
	// Uniquely identifies this account across all accounts for a single financial connection. Used for reconnection deduplication. See more information here: https://letsfuse.readme.io/docs/duplicate-accounts
	Fingerprint string `json:"fingerprint"`
	Institution *FinancialConnectionsAccountInstitution `json:"institution,omitempty"`
	// The partial account number.
	Mask NullableString `json:"mask,omitempty"`
	// The account's name, ie 'My Checking'
	Name string `json:"name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype,omitempty"`
	Balance FinancialConnectionsAccountCachedBalance `json:"balance"`
	// An array of additional balances. This may be used for investment type accounts where the user can have multiple balances across different currencies.
	AdditionalBalances []FinancialConnectionsAccountCachedBalance `json:"additional_balances,omitempty"`
	RemoteData interface{} `json:"remote_data"`
}

type _FinancialConnectionsAccount FinancialConnectionsAccount

// NewFinancialConnectionsAccount instantiates a new FinancialConnectionsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsAccount(remoteId string, currency string, fingerprint string, name string, type_ AccountType, balance FinancialConnectionsAccountCachedBalance, remoteData interface{}) *FinancialConnectionsAccount {
	this := FinancialConnectionsAccount{}
	this.RemoteId = remoteId
	this.Currency = currency
	this.Fingerprint = fingerprint
	this.Name = name
	this.Type = type_
	this.Balance = balance
	this.RemoteData = remoteData
	return &this
}

// NewFinancialConnectionsAccountWithDefaults instantiates a new FinancialConnectionsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsAccountWithDefaults() *FinancialConnectionsAccount {
	this := FinancialConnectionsAccount{}
	return &this
}

// GetRemoteId returns the RemoteId field value
func (o *FinancialConnectionsAccount) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *FinancialConnectionsAccount) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetCurrency returns the Currency field value
func (o *FinancialConnectionsAccount) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *FinancialConnectionsAccount) SetCurrency(v string) {
	o.Currency = v
}

// GetFingerprint returns the Fingerprint field value
func (o *FinancialConnectionsAccount) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *FinancialConnectionsAccount) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetInstitution returns the Institution field value if set, zero value otherwise.
func (o *FinancialConnectionsAccount) GetInstitution() FinancialConnectionsAccountInstitution {
	if o == nil || IsNil(o.Institution) {
		var ret FinancialConnectionsAccountInstitution
		return ret
	}
	return *o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetInstitutionOk() (*FinancialConnectionsAccountInstitution, bool) {
	if o == nil || IsNil(o.Institution) {
		return nil, false
	}
	return o.Institution, true
}

// HasInstitution returns a boolean if a field has been set.
func (o *FinancialConnectionsAccount) HasInstitution() bool {
	if o != nil && !IsNil(o.Institution) {
		return true
	}

	return false
}

// SetInstitution gets a reference to the given FinancialConnectionsAccountInstitution and assigns it to the Institution field.
func (o *FinancialConnectionsAccount) SetInstitution(v FinancialConnectionsAccountInstitution) {
	o.Institution = &v
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FinancialConnectionsAccount) GetMask() string {
	if o == nil || IsNil(o.Mask.Get()) {
		var ret string
		return ret
	}
	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FinancialConnectionsAccount) GetMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// HasMask returns a boolean if a field has been set.
func (o *FinancialConnectionsAccount) HasMask() bool {
	if o != nil && o.Mask.IsSet() {
		return true
	}

	return false
}

// SetMask gets a reference to the given NullableString and assigns it to the Mask field.
func (o *FinancialConnectionsAccount) SetMask(v string) {
	o.Mask.Set(&v)
}
// SetMaskNil sets the value for Mask to be an explicit nil
func (o *FinancialConnectionsAccount) SetMaskNil() {
	o.Mask.Set(nil)
}

// UnsetMask ensures that no value is present for Mask, not even an explicit nil
func (o *FinancialConnectionsAccount) UnsetMask() {
	o.Mask.Unset()
}

// GetName returns the Name field value
func (o *FinancialConnectionsAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FinancialConnectionsAccount) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FinancialConnectionsAccount) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialConnectionsAccount) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FinancialConnectionsAccount) GetSubtype() AccountSubtype {
	if o == nil || IsNil(o.Subtype.Get()) {
		var ret AccountSubtype
		return ret
	}
	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FinancialConnectionsAccount) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// HasSubtype returns a boolean if a field has been set.
func (o *FinancialConnectionsAccount) HasSubtype() bool {
	if o != nil && o.Subtype.IsSet() {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given NullableAccountSubtype and assigns it to the Subtype field.
func (o *FinancialConnectionsAccount) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}
// SetSubtypeNil sets the value for Subtype to be an explicit nil
func (o *FinancialConnectionsAccount) SetSubtypeNil() {
	o.Subtype.Set(nil)
}

// UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
func (o *FinancialConnectionsAccount) UnsetSubtype() {
	o.Subtype.Unset()
}

// GetBalance returns the Balance field value
func (o *FinancialConnectionsAccount) GetBalance() FinancialConnectionsAccountCachedBalance {
	if o == nil {
		var ret FinancialConnectionsAccountCachedBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetBalanceOk() (*FinancialConnectionsAccountCachedBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *FinancialConnectionsAccount) SetBalance(v FinancialConnectionsAccountCachedBalance) {
	o.Balance = v
}

// GetAdditionalBalances returns the AdditionalBalances field value if set, zero value otherwise.
func (o *FinancialConnectionsAccount) GetAdditionalBalances() []FinancialConnectionsAccountCachedBalance {
	if o == nil || IsNil(o.AdditionalBalances) {
		var ret []FinancialConnectionsAccountCachedBalance
		return ret
	}
	return o.AdditionalBalances
}

// GetAdditionalBalancesOk returns a tuple with the AdditionalBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccount) GetAdditionalBalancesOk() ([]FinancialConnectionsAccountCachedBalance, bool) {
	if o == nil || IsNil(o.AdditionalBalances) {
		return nil, false
	}
	return o.AdditionalBalances, true
}

// HasAdditionalBalances returns a boolean if a field has been set.
func (o *FinancialConnectionsAccount) HasAdditionalBalances() bool {
	if o != nil && !IsNil(o.AdditionalBalances) {
		return true
	}

	return false
}

// SetAdditionalBalances gets a reference to the given []FinancialConnectionsAccountCachedBalance and assigns it to the AdditionalBalances field.
func (o *FinancialConnectionsAccount) SetAdditionalBalances(v []FinancialConnectionsAccountCachedBalance) {
	o.AdditionalBalances = v
}

// GetRemoteData returns the RemoteData field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FinancialConnectionsAccount) GetRemoteData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FinancialConnectionsAccount) GetRemoteDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RemoteData) {
		return nil, false
	}
	return &o.RemoteData, true
}

// SetRemoteData sets field value
func (o *FinancialConnectionsAccount) SetRemoteData(v interface{}) {
	o.RemoteData = v
}

func (o FinancialConnectionsAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remote_id"] = o.RemoteId
	toSerialize["currency"] = o.Currency
	toSerialize["fingerprint"] = o.Fingerprint
	if !IsNil(o.Institution) {
		toSerialize["institution"] = o.Institution
	}
	if o.Mask.IsSet() {
		toSerialize["mask"] = o.Mask.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Subtype.IsSet() {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	toSerialize["balance"] = o.Balance
	if !IsNil(o.AdditionalBalances) {
		toSerialize["additional_balances"] = o.AdditionalBalances
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return toSerialize, nil
}

func (o *FinancialConnectionsAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"remote_id",
		"currency",
		"fingerprint",
		"name",
		"type",
		"balance",
		"remote_data",
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

	varFinancialConnectionsAccount := _FinancialConnectionsAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialConnectionsAccount)

	if err != nil {
		return err
	}

	*o = FinancialConnectionsAccount(varFinancialConnectionsAccount)

	return err
}

type NullableFinancialConnectionsAccount struct {
	value *FinancialConnectionsAccount
	isSet bool
}

func (v NullableFinancialConnectionsAccount) Get() *FinancialConnectionsAccount {
	return v.value
}

func (v *NullableFinancialConnectionsAccount) Set(val *FinancialConnectionsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsAccount(val *FinancialConnectionsAccount) *NullableFinancialConnectionsAccount {
	return &NullableFinancialConnectionsAccount{value: val, isSet: true}
}

func (v NullableFinancialConnectionsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


