# UpdateConsumerRiskReportCustomizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to [**ConsumerRiskReportTimeFrame**](ConsumerRiskReportTimeFrame.md) |  | [optional] 
**MinLimit** | Pointer to **float32** | The minimum allowed limit, in cents. | [optional] 
**MaxLimit** | Pointer to **float32** | The maximum allowed limit, in cents. | [optional] 
**RiskTolerance** | Pointer to **float32** | This parameter indicates the risk tolerance associated with spend limits. A high risk tolerance allow for higher limits, increasing both potential gains and losses. A Lower risk tolerance enforces strict limits, reducing the potential for loss but also limiting transaction volume for reliable users. | [optional] 

## Methods

### NewUpdateConsumerRiskReportCustomizationRequest

`func NewUpdateConsumerRiskReportCustomizationRequest() *UpdateConsumerRiskReportCustomizationRequest`

NewUpdateConsumerRiskReportCustomizationRequest instantiates a new UpdateConsumerRiskReportCustomizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConsumerRiskReportCustomizationRequestWithDefaults

`func NewUpdateConsumerRiskReportCustomizationRequestWithDefaults() *UpdateConsumerRiskReportCustomizationRequest`

NewUpdateConsumerRiskReportCustomizationRequestWithDefaults instantiates a new UpdateConsumerRiskReportCustomizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetTimeframe() ConsumerRiskReportTimeFrame`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetTimeframeOk() (*ConsumerRiskReportTimeFrame, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *UpdateConsumerRiskReportCustomizationRequest) SetTimeframe(v ConsumerRiskReportTimeFrame)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *UpdateConsumerRiskReportCustomizationRequest) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetMinLimit

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetMinLimit() float32`

GetMinLimit returns the MinLimit field if non-nil, zero value otherwise.

### GetMinLimitOk

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetMinLimitOk() (*float32, bool)`

GetMinLimitOk returns a tuple with the MinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLimit

`func (o *UpdateConsumerRiskReportCustomizationRequest) SetMinLimit(v float32)`

SetMinLimit sets MinLimit field to given value.

### HasMinLimit

`func (o *UpdateConsumerRiskReportCustomizationRequest) HasMinLimit() bool`

HasMinLimit returns a boolean if a field has been set.

### GetMaxLimit

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetMaxLimit() float32`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetMaxLimitOk() (*float32, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *UpdateConsumerRiskReportCustomizationRequest) SetMaxLimit(v float32)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *UpdateConsumerRiskReportCustomizationRequest) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.

### GetRiskTolerance

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetRiskTolerance() float32`

GetRiskTolerance returns the RiskTolerance field if non-nil, zero value otherwise.

### GetRiskToleranceOk

`func (o *UpdateConsumerRiskReportCustomizationRequest) GetRiskToleranceOk() (*float32, bool)`

GetRiskToleranceOk returns a tuple with the RiskTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTolerance

`func (o *UpdateConsumerRiskReportCustomizationRequest) SetRiskTolerance(v float32)`

SetRiskTolerance sets RiskTolerance field to given value.

### HasRiskTolerance

`func (o *UpdateConsumerRiskReportCustomizationRequest) HasRiskTolerance() bool`

HasRiskTolerance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


