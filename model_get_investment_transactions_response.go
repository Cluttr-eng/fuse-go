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

// checks if the GetInvestmentTransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvestmentTransactionsResponse{}

// GetInvestmentTransactionsResponse struct for GetInvestmentTransactionsResponse
type GetInvestmentTransactionsResponse struct {
	Accounts []FinancialConnectionsAccount `json:"accounts"`
	InvestmentTransactions []FinancialConnectionsInvestmentTransaction `json:"investment_transactions"`
	// The total number of transactions within the specified date range.
	TotalTransactions *float32 `json:"total_transactions,omitempty"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

// NewGetInvestmentTransactionsResponse instantiates a new GetInvestmentTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvestmentTransactionsResponse(accounts []FinancialConnectionsAccount, investmentTransactions []FinancialConnectionsInvestmentTransaction, requestId string) *GetInvestmentTransactionsResponse {
	this := GetInvestmentTransactionsResponse{}
	this.Accounts = accounts
	this.InvestmentTransactions = investmentTransactions
	this.RequestId = requestId
	return &this
}

// NewGetInvestmentTransactionsResponseWithDefaults instantiates a new GetInvestmentTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvestmentTransactionsResponseWithDefaults() *GetInvestmentTransactionsResponse {
	this := GetInvestmentTransactionsResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *GetInvestmentTransactionsResponse) GetAccounts() []FinancialConnectionsAccount {
	if o == nil {
		var ret []FinancialConnectionsAccount
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *GetInvestmentTransactionsResponse) GetAccountsOk() ([]FinancialConnectionsAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *GetInvestmentTransactionsResponse) SetAccounts(v []FinancialConnectionsAccount) {
	o.Accounts = v
}

// GetInvestmentTransactions returns the InvestmentTransactions field value
func (o *GetInvestmentTransactionsResponse) GetInvestmentTransactions() []FinancialConnectionsInvestmentTransaction {
	if o == nil {
		var ret []FinancialConnectionsInvestmentTransaction
		return ret
	}

	return o.InvestmentTransactions
}

// GetInvestmentTransactionsOk returns a tuple with the InvestmentTransactions field value
// and a boolean to check if the value has been set.
func (o *GetInvestmentTransactionsResponse) GetInvestmentTransactionsOk() ([]FinancialConnectionsInvestmentTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvestmentTransactions, true
}

// SetInvestmentTransactions sets field value
func (o *GetInvestmentTransactionsResponse) SetInvestmentTransactions(v []FinancialConnectionsInvestmentTransaction) {
	o.InvestmentTransactions = v
}

// GetTotalTransactions returns the TotalTransactions field value if set, zero value otherwise.
func (o *GetInvestmentTransactionsResponse) GetTotalTransactions() float32 {
	if o == nil || IsNil(o.TotalTransactions) {
		var ret float32
		return ret
	}
	return *o.TotalTransactions
}

// GetTotalTransactionsOk returns a tuple with the TotalTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvestmentTransactionsResponse) GetTotalTransactionsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTransactions) {
		return nil, false
	}
	return o.TotalTransactions, true
}

// HasTotalTransactions returns a boolean if a field has been set.
func (o *GetInvestmentTransactionsResponse) HasTotalTransactions() bool {
	if o != nil && !IsNil(o.TotalTransactions) {
		return true
	}

	return false
}

// SetTotalTransactions gets a reference to the given float32 and assigns it to the TotalTransactions field.
func (o *GetInvestmentTransactionsResponse) SetTotalTransactions(v float32) {
	o.TotalTransactions = &v
}

// GetRequestId returns the RequestId field value
func (o *GetInvestmentTransactionsResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetInvestmentTransactionsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetInvestmentTransactionsResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o GetInvestmentTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvestmentTransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	toSerialize["investment_transactions"] = o.InvestmentTransactions
	if !IsNil(o.TotalTransactions) {
		toSerialize["total_transactions"] = o.TotalTransactions
	}
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

type NullableGetInvestmentTransactionsResponse struct {
	value *GetInvestmentTransactionsResponse
	isSet bool
}

func (v NullableGetInvestmentTransactionsResponse) Get() *GetInvestmentTransactionsResponse {
	return v.value
}

func (v *NullableGetInvestmentTransactionsResponse) Set(val *GetInvestmentTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvestmentTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvestmentTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvestmentTransactionsResponse(val *GetInvestmentTransactionsResponse) *NullableGetInvestmentTransactionsResponse {
	return &NullableGetInvestmentTransactionsResponse{value: val, isSet: true}
}

func (v NullableGetInvestmentTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvestmentTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

