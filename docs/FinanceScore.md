# FinanceScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | The likelihood of a credit default. A higher score implies lower risk. The finance score and all finance score metrics are values between 0 and 1. This value is calculated by the weighted sum of the metrics below. | 
**SavingsScore** | **float32** | This quantifies a user&#39;s monthly savings habits. A lower score represents minimal savings, while a higher score indicates a user who consistently sets aside a substantial portion of their income. | 
**ExpenseStabilityScore** | **float32** | This assesses the consistency of a user&#39;s monthly spending. A lower score indicates variable monthly expenditure, while a higher score represents consistent spending habits. | 
**ActivityAgeScore** | **float32** | This measures the regularity of a user&#39;s financial activity over a period of time. A lower score suggests limited activity, while a higher score is indicative of regular daily transactions over a long period of time. | 
**IncomeScore** | **float32** | This evaluates the stability of a user&#39;s income. A lower score suggests inconsistent or low income, while a higher score represents high, consistent income. | 
**LoanPaymentsScore** | **float32** | This evaluates a user&#39;s loan repayment behaviour. A lower score is assigned to those without loan payments, while a higher score denotes consistent loan payments, such as a mortgage. | 
**RepaymentsScore** | **float32** | This quantifies a user&#39;s ability to repay debts. A lower score corresponds to missed payments, while a higher score signifies consistent debt repayment. | 

## Methods

### NewFinanceScore

`func NewFinanceScore(value float32, savingsScore float32, expenseStabilityScore float32, activityAgeScore float32, incomeScore float32, loanPaymentsScore float32, repaymentsScore float32, ) *FinanceScore`

NewFinanceScore instantiates a new FinanceScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinanceScoreWithDefaults

`func NewFinanceScoreWithDefaults() *FinanceScore`

NewFinanceScoreWithDefaults instantiates a new FinanceScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FinanceScore) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FinanceScore) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FinanceScore) SetValue(v float32)`

SetValue sets Value field to given value.


### GetSavingsScore

`func (o *FinanceScore) GetSavingsScore() float32`

GetSavingsScore returns the SavingsScore field if non-nil, zero value otherwise.

### GetSavingsScoreOk

`func (o *FinanceScore) GetSavingsScoreOk() (*float32, bool)`

GetSavingsScoreOk returns a tuple with the SavingsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingsScore

`func (o *FinanceScore) SetSavingsScore(v float32)`

SetSavingsScore sets SavingsScore field to given value.


### GetExpenseStabilityScore

`func (o *FinanceScore) GetExpenseStabilityScore() float32`

GetExpenseStabilityScore returns the ExpenseStabilityScore field if non-nil, zero value otherwise.

### GetExpenseStabilityScoreOk

`func (o *FinanceScore) GetExpenseStabilityScoreOk() (*float32, bool)`

GetExpenseStabilityScoreOk returns a tuple with the ExpenseStabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseStabilityScore

`func (o *FinanceScore) SetExpenseStabilityScore(v float32)`

SetExpenseStabilityScore sets ExpenseStabilityScore field to given value.


### GetActivityAgeScore

`func (o *FinanceScore) GetActivityAgeScore() float32`

GetActivityAgeScore returns the ActivityAgeScore field if non-nil, zero value otherwise.

### GetActivityAgeScoreOk

`func (o *FinanceScore) GetActivityAgeScoreOk() (*float32, bool)`

GetActivityAgeScoreOk returns a tuple with the ActivityAgeScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityAgeScore

`func (o *FinanceScore) SetActivityAgeScore(v float32)`

SetActivityAgeScore sets ActivityAgeScore field to given value.


### GetIncomeScore

`func (o *FinanceScore) GetIncomeScore() float32`

GetIncomeScore returns the IncomeScore field if non-nil, zero value otherwise.

### GetIncomeScoreOk

`func (o *FinanceScore) GetIncomeScoreOk() (*float32, bool)`

GetIncomeScoreOk returns a tuple with the IncomeScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeScore

`func (o *FinanceScore) SetIncomeScore(v float32)`

SetIncomeScore sets IncomeScore field to given value.


### GetLoanPaymentsScore

`func (o *FinanceScore) GetLoanPaymentsScore() float32`

GetLoanPaymentsScore returns the LoanPaymentsScore field if non-nil, zero value otherwise.

### GetLoanPaymentsScoreOk

`func (o *FinanceScore) GetLoanPaymentsScoreOk() (*float32, bool)`

GetLoanPaymentsScoreOk returns a tuple with the LoanPaymentsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanPaymentsScore

`func (o *FinanceScore) SetLoanPaymentsScore(v float32)`

SetLoanPaymentsScore sets LoanPaymentsScore field to given value.


### GetRepaymentsScore

`func (o *FinanceScore) GetRepaymentsScore() float32`

GetRepaymentsScore returns the RepaymentsScore field if non-nil, zero value otherwise.

### GetRepaymentsScoreOk

`func (o *FinanceScore) GetRepaymentsScoreOk() (*float32, bool)`

GetRepaymentsScoreOk returns a tuple with the RepaymentsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentsScore

`func (o *FinanceScore) SetRepaymentsScore(v float32)`

SetRepaymentsScore sets RepaymentsScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


