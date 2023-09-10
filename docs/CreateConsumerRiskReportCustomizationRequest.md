# CreateConsumerRiskReportCustomizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | [**ConsumerRiskReportTimeFrame**](ConsumerRiskReportTimeFrame.md) |  | 
**MinLimit** | **float32** | The minimum allowed limit, in cents. | 
**MaxLimit** | **float32** | The maximum allowed limit, in cents. | 
**RiskTolerance** | **float32** | This parameter indicates the risk tolerance associated with spend limits. A high risk tolerance allow for higher limits, increasing both potential gains and losses. A Lower risk tolerance enforces strict limits, reducing the potential for loss but also limiting transaction volume for reliable users. | 

## Methods

### NewCreateConsumerRiskReportCustomizationRequest

`func NewCreateConsumerRiskReportCustomizationRequest(timeframe ConsumerRiskReportTimeFrame, minLimit float32, maxLimit float32, riskTolerance float32, ) *CreateConsumerRiskReportCustomizationRequest`

NewCreateConsumerRiskReportCustomizationRequest instantiates a new CreateConsumerRiskReportCustomizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConsumerRiskReportCustomizationRequestWithDefaults

`func NewCreateConsumerRiskReportCustomizationRequestWithDefaults() *CreateConsumerRiskReportCustomizationRequest`

NewCreateConsumerRiskReportCustomizationRequestWithDefaults instantiates a new CreateConsumerRiskReportCustomizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *CreateConsumerRiskReportCustomizationRequest) GetTimeframe() ConsumerRiskReportTimeFrame`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *CreateConsumerRiskReportCustomizationRequest) GetTimeframeOk() (*ConsumerRiskReportTimeFrame, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *CreateConsumerRiskReportCustomizationRequest) SetTimeframe(v ConsumerRiskReportTimeFrame)`

SetTimeframe sets Timeframe field to given value.


### GetMinLimit

`func (o *CreateConsumerRiskReportCustomizationRequest) GetMinLimit() float32`

GetMinLimit returns the MinLimit field if non-nil, zero value otherwise.

### GetMinLimitOk

`func (o *CreateConsumerRiskReportCustomizationRequest) GetMinLimitOk() (*float32, bool)`

GetMinLimitOk returns a tuple with the MinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLimit

`func (o *CreateConsumerRiskReportCustomizationRequest) SetMinLimit(v float32)`

SetMinLimit sets MinLimit field to given value.


### GetMaxLimit

`func (o *CreateConsumerRiskReportCustomizationRequest) GetMaxLimit() float32`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *CreateConsumerRiskReportCustomizationRequest) GetMaxLimitOk() (*float32, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *CreateConsumerRiskReportCustomizationRequest) SetMaxLimit(v float32)`

SetMaxLimit sets MaxLimit field to given value.


### GetRiskTolerance

`func (o *CreateConsumerRiskReportCustomizationRequest) GetRiskTolerance() float32`

GetRiskTolerance returns the RiskTolerance field if non-nil, zero value otherwise.

### GetRiskToleranceOk

`func (o *CreateConsumerRiskReportCustomizationRequest) GetRiskToleranceOk() (*float32, bool)`

GetRiskToleranceOk returns a tuple with the RiskTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTolerance

`func (o *CreateConsumerRiskReportCustomizationRequest) SetRiskTolerance(v float32)`

SetRiskTolerance sets RiskTolerance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


