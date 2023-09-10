# ConsumerRiskReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CustomizationId** | **string** |  | 
**SpendLimit** | **float32** | The total limit for the current timeframe, in cents. | 
**CurrentSpend** | **float32** | The total current spend in the current timeframe, in cents, without factoring in pending payments. | 
**PendingPaymentsAmount** | **float32** | The total unpaid amount, in cents, from all timeframes. | 
**IsoCurrencyCode** | **string** | The ISO-4217 currency code of the transaction | 
**LastUpdated** | **string** | The datetime of when the consumer risk report was most recently updated, in ISO-8601 format. | 
**FinanceScore** | [**FinanceScore**](FinanceScore.md) |  | 
**PredictedBalance** | **float32** | Predicted balance for the timeframe. | 

## Methods

### NewConsumerRiskReport

`func NewConsumerRiskReport(id string, customizationId string, spendLimit float32, currentSpend float32, pendingPaymentsAmount float32, isoCurrencyCode string, lastUpdated string, financeScore FinanceScore, predictedBalance float32, ) *ConsumerRiskReport`

NewConsumerRiskReport instantiates a new ConsumerRiskReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRiskReportWithDefaults

`func NewConsumerRiskReportWithDefaults() *ConsumerRiskReport`

NewConsumerRiskReportWithDefaults instantiates a new ConsumerRiskReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsumerRiskReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerRiskReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerRiskReport) SetId(v string)`

SetId sets Id field to given value.


### GetCustomizationId

`func (o *ConsumerRiskReport) GetCustomizationId() string`

GetCustomizationId returns the CustomizationId field if non-nil, zero value otherwise.

### GetCustomizationIdOk

`func (o *ConsumerRiskReport) GetCustomizationIdOk() (*string, bool)`

GetCustomizationIdOk returns a tuple with the CustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationId

`func (o *ConsumerRiskReport) SetCustomizationId(v string)`

SetCustomizationId sets CustomizationId field to given value.


### GetSpendLimit

`func (o *ConsumerRiskReport) GetSpendLimit() float32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *ConsumerRiskReport) GetSpendLimitOk() (*float32, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *ConsumerRiskReport) SetSpendLimit(v float32)`

SetSpendLimit sets SpendLimit field to given value.


### GetCurrentSpend

`func (o *ConsumerRiskReport) GetCurrentSpend() float32`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *ConsumerRiskReport) GetCurrentSpendOk() (*float32, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *ConsumerRiskReport) SetCurrentSpend(v float32)`

SetCurrentSpend sets CurrentSpend field to given value.


### GetPendingPaymentsAmount

`func (o *ConsumerRiskReport) GetPendingPaymentsAmount() float32`

GetPendingPaymentsAmount returns the PendingPaymentsAmount field if non-nil, zero value otherwise.

### GetPendingPaymentsAmountOk

`func (o *ConsumerRiskReport) GetPendingPaymentsAmountOk() (*float32, bool)`

GetPendingPaymentsAmountOk returns a tuple with the PendingPaymentsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPaymentsAmount

`func (o *ConsumerRiskReport) SetPendingPaymentsAmount(v float32)`

SetPendingPaymentsAmount sets PendingPaymentsAmount field to given value.


### GetIsoCurrencyCode

`func (o *ConsumerRiskReport) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *ConsumerRiskReport) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *ConsumerRiskReport) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetLastUpdated

`func (o *ConsumerRiskReport) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConsumerRiskReport) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConsumerRiskReport) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetFinanceScore

`func (o *ConsumerRiskReport) GetFinanceScore() FinanceScore`

GetFinanceScore returns the FinanceScore field if non-nil, zero value otherwise.

### GetFinanceScoreOk

`func (o *ConsumerRiskReport) GetFinanceScoreOk() (*FinanceScore, bool)`

GetFinanceScoreOk returns a tuple with the FinanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceScore

`func (o *ConsumerRiskReport) SetFinanceScore(v FinanceScore)`

SetFinanceScore sets FinanceScore field to given value.


### GetPredictedBalance

`func (o *ConsumerRiskReport) GetPredictedBalance() float32`

GetPredictedBalance returns the PredictedBalance field if non-nil, zero value otherwise.

### GetPredictedBalanceOk

`func (o *ConsumerRiskReport) GetPredictedBalanceOk() (*float32, bool)`

GetPredictedBalanceOk returns a tuple with the PredictedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedBalance

`func (o *ConsumerRiskReport) SetPredictedBalance(v float32)`

SetPredictedBalance sets PredictedBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


