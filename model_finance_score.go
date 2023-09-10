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

// checks if the FinanceScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinanceScore{}

// FinanceScore struct for FinanceScore
type FinanceScore struct {
	// The likelihood of a credit default. A higher score implies lower risk. The finance score and all finance score metrics are values between 0 and 1. This value is calculated by the weighted sum of the metrics below.
	Value float32 `json:"value"`
	// This quantifies a user's monthly savings habits. A lower score represents minimal savings, while a higher score indicates a user who consistently sets aside a substantial portion of their income.
	SavingsScore float32 `json:"savings_score"`
	// This assesses the consistency of a user's monthly spending. A lower score indicates variable monthly expenditure, while a higher score represents consistent spending habits.
	ExpenseStabilityScore float32 `json:"expense_stability_score"`
	// This measures the regularity of a user's financial activity over a period of time. A lower score suggests limited activity, while a higher score is indicative of regular daily transactions over a long period of time.
	ActivityAgeScore float32 `json:"activity_age_score"`
	// This evaluates the stability of a user's income. A lower score suggests inconsistent or low income, while a higher score represents high, consistent income.
	IncomeScore float32 `json:"income_score"`
	// This evaluates a user's loan repayment behaviour. A lower score is assigned to those without loan payments, while a higher score denotes consistent loan payments, such as a mortgage.
	LoanPaymentsScore float32 `json:"loan_payments_score"`
	// This quantifies a user's ability to repay debts. A lower score corresponds to missed payments, while a higher score signifies consistent debt repayment.
	RepaymentsScore float32 `json:"repayments_score"`
}

// NewFinanceScore instantiates a new FinanceScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinanceScore(value float32, savingsScore float32, expenseStabilityScore float32, activityAgeScore float32, incomeScore float32, loanPaymentsScore float32, repaymentsScore float32) *FinanceScore {
	this := FinanceScore{}
	this.Value = value
	this.SavingsScore = savingsScore
	this.ExpenseStabilityScore = expenseStabilityScore
	this.ActivityAgeScore = activityAgeScore
	this.IncomeScore = incomeScore
	this.LoanPaymentsScore = loanPaymentsScore
	this.RepaymentsScore = repaymentsScore
	return &this
}

// NewFinanceScoreWithDefaults instantiates a new FinanceScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinanceScoreWithDefaults() *FinanceScore {
	this := FinanceScore{}
	return &this
}

// GetValue returns the Value field value
func (o *FinanceScore) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FinanceScore) SetValue(v float32) {
	o.Value = v
}

// GetSavingsScore returns the SavingsScore field value
func (o *FinanceScore) GetSavingsScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SavingsScore
}

// GetSavingsScoreOk returns a tuple with the SavingsScore field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetSavingsScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavingsScore, true
}

// SetSavingsScore sets field value
func (o *FinanceScore) SetSavingsScore(v float32) {
	o.SavingsScore = v
}

// GetExpenseStabilityScore returns the ExpenseStabilityScore field value
func (o *FinanceScore) GetExpenseStabilityScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpenseStabilityScore
}

// GetExpenseStabilityScoreOk returns a tuple with the ExpenseStabilityScore field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetExpenseStabilityScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpenseStabilityScore, true
}

// SetExpenseStabilityScore sets field value
func (o *FinanceScore) SetExpenseStabilityScore(v float32) {
	o.ExpenseStabilityScore = v
}

// GetActivityAgeScore returns the ActivityAgeScore field value
func (o *FinanceScore) GetActivityAgeScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ActivityAgeScore
}

// GetActivityAgeScoreOk returns a tuple with the ActivityAgeScore field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetActivityAgeScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityAgeScore, true
}

// SetActivityAgeScore sets field value
func (o *FinanceScore) SetActivityAgeScore(v float32) {
	o.ActivityAgeScore = v
}

// GetIncomeScore returns the IncomeScore field value
func (o *FinanceScore) GetIncomeScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IncomeScore
}

// GetIncomeScoreOk returns a tuple with the IncomeScore field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetIncomeScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncomeScore, true
}

// SetIncomeScore sets field value
func (o *FinanceScore) SetIncomeScore(v float32) {
	o.IncomeScore = v
}

// GetLoanPaymentsScore returns the LoanPaymentsScore field value
func (o *FinanceScore) GetLoanPaymentsScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LoanPaymentsScore
}

// GetLoanPaymentsScoreOk returns a tuple with the LoanPaymentsScore field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetLoanPaymentsScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoanPaymentsScore, true
}

// SetLoanPaymentsScore sets field value
func (o *FinanceScore) SetLoanPaymentsScore(v float32) {
	o.LoanPaymentsScore = v
}

// GetRepaymentsScore returns the RepaymentsScore field value
func (o *FinanceScore) GetRepaymentsScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RepaymentsScore
}

// GetRepaymentsScoreOk returns a tuple with the RepaymentsScore field value
// and a boolean to check if the value has been set.
func (o *FinanceScore) GetRepaymentsScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepaymentsScore, true
}

// SetRepaymentsScore sets field value
func (o *FinanceScore) SetRepaymentsScore(v float32) {
	o.RepaymentsScore = v
}

func (o FinanceScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinanceScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["savings_score"] = o.SavingsScore
	toSerialize["expense_stability_score"] = o.ExpenseStabilityScore
	toSerialize["activity_age_score"] = o.ActivityAgeScore
	toSerialize["income_score"] = o.IncomeScore
	toSerialize["loan_payments_score"] = o.LoanPaymentsScore
	toSerialize["repayments_score"] = o.RepaymentsScore
	return toSerialize, nil
}

type NullableFinanceScore struct {
	value *FinanceScore
	isSet bool
}

func (v NullableFinanceScore) Get() *FinanceScore {
	return v.value
}

func (v *NullableFinanceScore) Set(val *FinanceScore) {
	v.value = val
	v.isSet = true
}

func (v NullableFinanceScore) IsSet() bool {
	return v.isSet
}

func (v *NullableFinanceScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinanceScore(val *FinanceScore) *NullableFinanceScore {
	return &NullableFinanceScore{value: val, isSet: true}
}

func (v NullableFinanceScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinanceScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


