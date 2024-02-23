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

// checks if the FinancialConnectionsAccountLiability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsAccountLiability{}

// FinancialConnectionsAccountLiability struct for FinancialConnectionsAccountLiability
type FinancialConnectionsAccountLiability struct {
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
	RemoteData map[string]interface{} `json:"remote_data"`
	// The various interest rates that apply to the account. If APR data is not available, this array will be empty.
	Aprs []FinancialConnectionsAccountLiabilityAllOfAprs `json:"aprs,omitempty"`
	// The interest rate on the loan as a percentage.
	InterestRatePercentage *float32 `json:"interest_rate_percentage,omitempty"`
	// The original principal balance of the loan.
	OriginationPrincipalAmount *float32 `json:"origination_principal_amount,omitempty"`
	// The due date for the next payment. The due date is null if a payment is not expected.
	NextPaymentDueDate *string `json:"next_payment_due_date,omitempty"`
	// The date of the last payment. Dates are returned in an ISO 8601 format (YYYY-MM-DD).
	LastPaymentDate *string `json:"last_payment_date,omitempty"`
	// The amount of the last payment.
	LastPaymentAmount *float32 `json:"last_payment_amount,omitempty"`
	// The minimum payment required for an account. This can apply to any debt account.
	MinimumPaymentAmount *float32 `json:"minimum_payment_amount,omitempty"`
}

type _FinancialConnectionsAccountLiability FinancialConnectionsAccountLiability

// NewFinancialConnectionsAccountLiability instantiates a new FinancialConnectionsAccountLiability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsAccountLiability(remoteId string, currency string, fingerprint string, name string, type_ AccountType, balance FinancialConnectionsAccountCachedBalance, remoteData map[string]interface{}) *FinancialConnectionsAccountLiability {
	this := FinancialConnectionsAccountLiability{}
	this.RemoteId = remoteId
	this.Currency = currency
	this.Fingerprint = fingerprint
	this.Name = name
	this.Type = type_
	this.Balance = balance
	this.RemoteData = remoteData
	return &this
}

// NewFinancialConnectionsAccountLiabilityWithDefaults instantiates a new FinancialConnectionsAccountLiability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsAccountLiabilityWithDefaults() *FinancialConnectionsAccountLiability {
	this := FinancialConnectionsAccountLiability{}
	return &this
}

// GetRemoteId returns the RemoteId field value
func (o *FinancialConnectionsAccountLiability) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *FinancialConnectionsAccountLiability) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetCurrency returns the Currency field value
func (o *FinancialConnectionsAccountLiability) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *FinancialConnectionsAccountLiability) SetCurrency(v string) {
	o.Currency = v
}

// GetFingerprint returns the Fingerprint field value
func (o *FinancialConnectionsAccountLiability) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *FinancialConnectionsAccountLiability) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetInstitution returns the Institution field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetInstitution() FinancialConnectionsAccountInstitution {
	if o == nil || IsNil(o.Institution) {
		var ret FinancialConnectionsAccountInstitution
		return ret
	}
	return *o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetInstitutionOk() (*FinancialConnectionsAccountInstitution, bool) {
	if o == nil || IsNil(o.Institution) {
		return nil, false
	}
	return o.Institution, true
}

// HasInstitution returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasInstitution() bool {
	if o != nil && !IsNil(o.Institution) {
		return true
	}

	return false
}

// SetInstitution gets a reference to the given FinancialConnectionsAccountInstitution and assigns it to the Institution field.
func (o *FinancialConnectionsAccountLiability) SetInstitution(v FinancialConnectionsAccountInstitution) {
	o.Institution = &v
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FinancialConnectionsAccountLiability) GetMask() string {
	if o == nil || IsNil(o.Mask.Get()) {
		var ret string
		return ret
	}
	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FinancialConnectionsAccountLiability) GetMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// HasMask returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasMask() bool {
	if o != nil && o.Mask.IsSet() {
		return true
	}

	return false
}

// SetMask gets a reference to the given NullableString and assigns it to the Mask field.
func (o *FinancialConnectionsAccountLiability) SetMask(v string) {
	o.Mask.Set(&v)
}
// SetMaskNil sets the value for Mask to be an explicit nil
func (o *FinancialConnectionsAccountLiability) SetMaskNil() {
	o.Mask.Set(nil)
}

// UnsetMask ensures that no value is present for Mask, not even an explicit nil
func (o *FinancialConnectionsAccountLiability) UnsetMask() {
	o.Mask.Unset()
}

// GetName returns the Name field value
func (o *FinancialConnectionsAccountLiability) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FinancialConnectionsAccountLiability) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FinancialConnectionsAccountLiability) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialConnectionsAccountLiability) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FinancialConnectionsAccountLiability) GetSubtype() AccountSubtype {
	if o == nil || IsNil(o.Subtype.Get()) {
		var ret AccountSubtype
		return ret
	}
	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FinancialConnectionsAccountLiability) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// HasSubtype returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasSubtype() bool {
	if o != nil && o.Subtype.IsSet() {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given NullableAccountSubtype and assigns it to the Subtype field.
func (o *FinancialConnectionsAccountLiability) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}
// SetSubtypeNil sets the value for Subtype to be an explicit nil
func (o *FinancialConnectionsAccountLiability) SetSubtypeNil() {
	o.Subtype.Set(nil)
}

// UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
func (o *FinancialConnectionsAccountLiability) UnsetSubtype() {
	o.Subtype.Unset()
}

// GetBalance returns the Balance field value
func (o *FinancialConnectionsAccountLiability) GetBalance() FinancialConnectionsAccountCachedBalance {
	if o == nil {
		var ret FinancialConnectionsAccountCachedBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetBalanceOk() (*FinancialConnectionsAccountCachedBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *FinancialConnectionsAccountLiability) SetBalance(v FinancialConnectionsAccountCachedBalance) {
	o.Balance = v
}

// GetAdditionalBalances returns the AdditionalBalances field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetAdditionalBalances() []FinancialConnectionsAccountCachedBalance {
	if o == nil || IsNil(o.AdditionalBalances) {
		var ret []FinancialConnectionsAccountCachedBalance
		return ret
	}
	return o.AdditionalBalances
}

// GetAdditionalBalancesOk returns a tuple with the AdditionalBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetAdditionalBalancesOk() ([]FinancialConnectionsAccountCachedBalance, bool) {
	if o == nil || IsNil(o.AdditionalBalances) {
		return nil, false
	}
	return o.AdditionalBalances, true
}

// HasAdditionalBalances returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasAdditionalBalances() bool {
	if o != nil && !IsNil(o.AdditionalBalances) {
		return true
	}

	return false
}

// SetAdditionalBalances gets a reference to the given []FinancialConnectionsAccountCachedBalance and assigns it to the AdditionalBalances field.
func (o *FinancialConnectionsAccountLiability) SetAdditionalBalances(v []FinancialConnectionsAccountCachedBalance) {
	o.AdditionalBalances = v
}

// GetRemoteData returns the RemoteData field value
func (o *FinancialConnectionsAccountLiability) GetRemoteData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetRemoteDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.RemoteData, true
}

// SetRemoteData sets field value
func (o *FinancialConnectionsAccountLiability) SetRemoteData(v map[string]interface{}) {
	o.RemoteData = v
}

// GetAprs returns the Aprs field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetAprs() []FinancialConnectionsAccountLiabilityAllOfAprs {
	if o == nil || IsNil(o.Aprs) {
		var ret []FinancialConnectionsAccountLiabilityAllOfAprs
		return ret
	}
	return o.Aprs
}

// GetAprsOk returns a tuple with the Aprs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetAprsOk() ([]FinancialConnectionsAccountLiabilityAllOfAprs, bool) {
	if o == nil || IsNil(o.Aprs) {
		return nil, false
	}
	return o.Aprs, true
}

// HasAprs returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasAprs() bool {
	if o != nil && !IsNil(o.Aprs) {
		return true
	}

	return false
}

// SetAprs gets a reference to the given []FinancialConnectionsAccountLiabilityAllOfAprs and assigns it to the Aprs field.
func (o *FinancialConnectionsAccountLiability) SetAprs(v []FinancialConnectionsAccountLiabilityAllOfAprs) {
	o.Aprs = v
}

// GetInterestRatePercentage returns the InterestRatePercentage field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetInterestRatePercentage() float32 {
	if o == nil || IsNil(o.InterestRatePercentage) {
		var ret float32
		return ret
	}
	return *o.InterestRatePercentage
}

// GetInterestRatePercentageOk returns a tuple with the InterestRatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetInterestRatePercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.InterestRatePercentage) {
		return nil, false
	}
	return o.InterestRatePercentage, true
}

// HasInterestRatePercentage returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasInterestRatePercentage() bool {
	if o != nil && !IsNil(o.InterestRatePercentage) {
		return true
	}

	return false
}

// SetInterestRatePercentage gets a reference to the given float32 and assigns it to the InterestRatePercentage field.
func (o *FinancialConnectionsAccountLiability) SetInterestRatePercentage(v float32) {
	o.InterestRatePercentage = &v
}

// GetOriginationPrincipalAmount returns the OriginationPrincipalAmount field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetOriginationPrincipalAmount() float32 {
	if o == nil || IsNil(o.OriginationPrincipalAmount) {
		var ret float32
		return ret
	}
	return *o.OriginationPrincipalAmount
}

// GetOriginationPrincipalAmountOk returns a tuple with the OriginationPrincipalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetOriginationPrincipalAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginationPrincipalAmount) {
		return nil, false
	}
	return o.OriginationPrincipalAmount, true
}

// HasOriginationPrincipalAmount returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasOriginationPrincipalAmount() bool {
	if o != nil && !IsNil(o.OriginationPrincipalAmount) {
		return true
	}

	return false
}

// SetOriginationPrincipalAmount gets a reference to the given float32 and assigns it to the OriginationPrincipalAmount field.
func (o *FinancialConnectionsAccountLiability) SetOriginationPrincipalAmount(v float32) {
	o.OriginationPrincipalAmount = &v
}

// GetNextPaymentDueDate returns the NextPaymentDueDate field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetNextPaymentDueDate() string {
	if o == nil || IsNil(o.NextPaymentDueDate) {
		var ret string
		return ret
	}
	return *o.NextPaymentDueDate
}

// GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetNextPaymentDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextPaymentDueDate) {
		return nil, false
	}
	return o.NextPaymentDueDate, true
}

// HasNextPaymentDueDate returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasNextPaymentDueDate() bool {
	if o != nil && !IsNil(o.NextPaymentDueDate) {
		return true
	}

	return false
}

// SetNextPaymentDueDate gets a reference to the given string and assigns it to the NextPaymentDueDate field.
func (o *FinancialConnectionsAccountLiability) SetNextPaymentDueDate(v string) {
	o.NextPaymentDueDate = &v
}

// GetLastPaymentDate returns the LastPaymentDate field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetLastPaymentDate() string {
	if o == nil || IsNil(o.LastPaymentDate) {
		var ret string
		return ret
	}
	return *o.LastPaymentDate
}

// GetLastPaymentDateOk returns a tuple with the LastPaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetLastPaymentDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastPaymentDate) {
		return nil, false
	}
	return o.LastPaymentDate, true
}

// HasLastPaymentDate returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasLastPaymentDate() bool {
	if o != nil && !IsNil(o.LastPaymentDate) {
		return true
	}

	return false
}

// SetLastPaymentDate gets a reference to the given string and assigns it to the LastPaymentDate field.
func (o *FinancialConnectionsAccountLiability) SetLastPaymentDate(v string) {
	o.LastPaymentDate = &v
}

// GetLastPaymentAmount returns the LastPaymentAmount field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetLastPaymentAmount() float32 {
	if o == nil || IsNil(o.LastPaymentAmount) {
		var ret float32
		return ret
	}
	return *o.LastPaymentAmount
}

// GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetLastPaymentAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.LastPaymentAmount) {
		return nil, false
	}
	return o.LastPaymentAmount, true
}

// HasLastPaymentAmount returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasLastPaymentAmount() bool {
	if o != nil && !IsNil(o.LastPaymentAmount) {
		return true
	}

	return false
}

// SetLastPaymentAmount gets a reference to the given float32 and assigns it to the LastPaymentAmount field.
func (o *FinancialConnectionsAccountLiability) SetLastPaymentAmount(v float32) {
	o.LastPaymentAmount = &v
}

// GetMinimumPaymentAmount returns the MinimumPaymentAmount field value if set, zero value otherwise.
func (o *FinancialConnectionsAccountLiability) GetMinimumPaymentAmount() float32 {
	if o == nil || IsNil(o.MinimumPaymentAmount) {
		var ret float32
		return ret
	}
	return *o.MinimumPaymentAmount
}

// GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsAccountLiability) GetMinimumPaymentAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimumPaymentAmount) {
		return nil, false
	}
	return o.MinimumPaymentAmount, true
}

// HasMinimumPaymentAmount returns a boolean if a field has been set.
func (o *FinancialConnectionsAccountLiability) HasMinimumPaymentAmount() bool {
	if o != nil && !IsNil(o.MinimumPaymentAmount) {
		return true
	}

	return false
}

// SetMinimumPaymentAmount gets a reference to the given float32 and assigns it to the MinimumPaymentAmount field.
func (o *FinancialConnectionsAccountLiability) SetMinimumPaymentAmount(v float32) {
	o.MinimumPaymentAmount = &v
}

func (o FinancialConnectionsAccountLiability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsAccountLiability) ToMap() (map[string]interface{}, error) {
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
	toSerialize["remote_data"] = o.RemoteData
	if !IsNil(o.Aprs) {
		toSerialize["aprs"] = o.Aprs
	}
	if !IsNil(o.InterestRatePercentage) {
		toSerialize["interest_rate_percentage"] = o.InterestRatePercentage
	}
	if !IsNil(o.OriginationPrincipalAmount) {
		toSerialize["origination_principal_amount"] = o.OriginationPrincipalAmount
	}
	if !IsNil(o.NextPaymentDueDate) {
		toSerialize["next_payment_due_date"] = o.NextPaymentDueDate
	}
	if !IsNil(o.LastPaymentDate) {
		toSerialize["last_payment_date"] = o.LastPaymentDate
	}
	if !IsNil(o.LastPaymentAmount) {
		toSerialize["last_payment_amount"] = o.LastPaymentAmount
	}
	if !IsNil(o.MinimumPaymentAmount) {
		toSerialize["minimum_payment_amount"] = o.MinimumPaymentAmount
	}
	return toSerialize, nil
}

func (o *FinancialConnectionsAccountLiability) UnmarshalJSON(data []byte) (err error) {
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

	varFinancialConnectionsAccountLiability := _FinancialConnectionsAccountLiability{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialConnectionsAccountLiability)

	if err != nil {
		return err
	}

	*o = FinancialConnectionsAccountLiability(varFinancialConnectionsAccountLiability)

	return err
}

type NullableFinancialConnectionsAccountLiability struct {
	value *FinancialConnectionsAccountLiability
	isSet bool
}

func (v NullableFinancialConnectionsAccountLiability) Get() *FinancialConnectionsAccountLiability {
	return v.value
}

func (v *NullableFinancialConnectionsAccountLiability) Set(val *FinancialConnectionsAccountLiability) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsAccountLiability) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsAccountLiability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsAccountLiability(val *FinancialConnectionsAccountLiability) *NullableFinancialConnectionsAccountLiability {
	return &NullableFinancialConnectionsAccountLiability{value: val, isSet: true}
}

func (v NullableFinancialConnectionsAccountLiability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsAccountLiability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


