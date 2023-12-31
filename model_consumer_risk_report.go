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

// checks if the ConsumerRiskReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumerRiskReport{}

// ConsumerRiskReport struct for ConsumerRiskReport
type ConsumerRiskReport struct {
	Id string `json:"id"`
	CustomizationId string `json:"customization_id"`
	// The total limit for the current timeframe, in cents.
	SpendLimit float32 `json:"spend_limit"`
	// The total current spend in the current timeframe, in cents, without factoring in pending payments.
	CurrentSpend float32 `json:"current_spend"`
	// The total unpaid amount, in cents, from all timeframes.
	PendingPaymentsAmount float32 `json:"pending_payments_amount"`
	// The ISO-4217 currency code of the transaction
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The datetime of when the consumer risk report was most recently updated, in ISO-8601 format.
	LastUpdated string `json:"last_updated"`
	FinanceScore FinanceScore `json:"finance_score"`
	// Predicted balance for the timeframe.
	PredictedBalance float32 `json:"predicted_balance"`
}

// NewConsumerRiskReport instantiates a new ConsumerRiskReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerRiskReport(id string, customizationId string, spendLimit float32, currentSpend float32, pendingPaymentsAmount float32, isoCurrencyCode string, lastUpdated string, financeScore FinanceScore, predictedBalance float32) *ConsumerRiskReport {
	this := ConsumerRiskReport{}
	this.Id = id
	this.CustomizationId = customizationId
	this.SpendLimit = spendLimit
	this.CurrentSpend = currentSpend
	this.PendingPaymentsAmount = pendingPaymentsAmount
	this.IsoCurrencyCode = isoCurrencyCode
	this.LastUpdated = lastUpdated
	this.FinanceScore = financeScore
	this.PredictedBalance = predictedBalance
	return &this
}

// NewConsumerRiskReportWithDefaults instantiates a new ConsumerRiskReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerRiskReportWithDefaults() *ConsumerRiskReport {
	this := ConsumerRiskReport{}
	return &this
}

// GetId returns the Id field value
func (o *ConsumerRiskReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsumerRiskReport) SetId(v string) {
	o.Id = v
}

// GetCustomizationId returns the CustomizationId field value
func (o *ConsumerRiskReport) GetCustomizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomizationId
}

// GetCustomizationIdOk returns a tuple with the CustomizationId field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetCustomizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomizationId, true
}

// SetCustomizationId sets field value
func (o *ConsumerRiskReport) SetCustomizationId(v string) {
	o.CustomizationId = v
}

// GetSpendLimit returns the SpendLimit field value
func (o *ConsumerRiskReport) GetSpendLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpendLimit
}

// GetSpendLimitOk returns a tuple with the SpendLimit field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetSpendLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpendLimit, true
}

// SetSpendLimit sets field value
func (o *ConsumerRiskReport) SetSpendLimit(v float32) {
	o.SpendLimit = v
}

// GetCurrentSpend returns the CurrentSpend field value
func (o *ConsumerRiskReport) GetCurrentSpend() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentSpend
}

// GetCurrentSpendOk returns a tuple with the CurrentSpend field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetCurrentSpendOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentSpend, true
}

// SetCurrentSpend sets field value
func (o *ConsumerRiskReport) SetCurrentSpend(v float32) {
	o.CurrentSpend = v
}

// GetPendingPaymentsAmount returns the PendingPaymentsAmount field value
func (o *ConsumerRiskReport) GetPendingPaymentsAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingPaymentsAmount
}

// GetPendingPaymentsAmountOk returns a tuple with the PendingPaymentsAmount field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetPendingPaymentsAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingPaymentsAmount, true
}

// SetPendingPaymentsAmount sets field value
func (o *ConsumerRiskReport) SetPendingPaymentsAmount(v float32) {
	o.PendingPaymentsAmount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *ConsumerRiskReport) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *ConsumerRiskReport) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *ConsumerRiskReport) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *ConsumerRiskReport) SetLastUpdated(v string) {
	o.LastUpdated = v
}

// GetFinanceScore returns the FinanceScore field value
func (o *ConsumerRiskReport) GetFinanceScore() FinanceScore {
	if o == nil {
		var ret FinanceScore
		return ret
	}

	return o.FinanceScore
}

// GetFinanceScoreOk returns a tuple with the FinanceScore field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetFinanceScoreOk() (*FinanceScore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinanceScore, true
}

// SetFinanceScore sets field value
func (o *ConsumerRiskReport) SetFinanceScore(v FinanceScore) {
	o.FinanceScore = v
}

// GetPredictedBalance returns the PredictedBalance field value
func (o *ConsumerRiskReport) GetPredictedBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PredictedBalance
}

// GetPredictedBalanceOk returns a tuple with the PredictedBalance field value
// and a boolean to check if the value has been set.
func (o *ConsumerRiskReport) GetPredictedBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PredictedBalance, true
}

// SetPredictedBalance sets field value
func (o *ConsumerRiskReport) SetPredictedBalance(v float32) {
	o.PredictedBalance = v
}

func (o ConsumerRiskReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumerRiskReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customization_id"] = o.CustomizationId
	toSerialize["spend_limit"] = o.SpendLimit
	toSerialize["current_spend"] = o.CurrentSpend
	toSerialize["pending_payments_amount"] = o.PendingPaymentsAmount
	toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	toSerialize["last_updated"] = o.LastUpdated
	toSerialize["finance_score"] = o.FinanceScore
	toSerialize["predicted_balance"] = o.PredictedBalance
	return toSerialize, nil
}

type NullableConsumerRiskReport struct {
	value *ConsumerRiskReport
	isSet bool
}

func (v NullableConsumerRiskReport) Get() *ConsumerRiskReport {
	return v.value
}

func (v *NullableConsumerRiskReport) Set(val *ConsumerRiskReport) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerRiskReport) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerRiskReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerRiskReport(val *ConsumerRiskReport) *NullableConsumerRiskReport {
	return &NullableConsumerRiskReport{value: val, isSet: true}
}

func (v NullableConsumerRiskReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerRiskReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


