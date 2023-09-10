# ConsumerRiskReportCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Timeframe** | [**ConsumerRiskReportTimeFrame**](ConsumerRiskReportTimeFrame.md) |  | 
**MinLimit** | **int32** | The minimum allowed limit, in cents. | 
**MaxLimit** | **int32** | The maximum allowed limit, in cents. | 
**RiskTolerance** | Pointer to **int32** | This parameter indicates the risk tolerance associated with spend limits. A high risk tolerance allow for higher limits, increasing both potential gains and losses. A Lower risk tolerance enforces strict limits, reducing the potential for loss but also limiting transaction volume for reliable users. | [optional] 

## Methods

### NewConsumerRiskReportCustomization

`func NewConsumerRiskReportCustomization(id string, timeframe ConsumerRiskReportTimeFrame, minLimit int32, maxLimit int32, ) *ConsumerRiskReportCustomization`

NewConsumerRiskReportCustomization instantiates a new ConsumerRiskReportCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRiskReportCustomizationWithDefaults

`func NewConsumerRiskReportCustomizationWithDefaults() *ConsumerRiskReportCustomization`

NewConsumerRiskReportCustomizationWithDefaults instantiates a new ConsumerRiskReportCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsumerRiskReportCustomization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerRiskReportCustomization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerRiskReportCustomization) SetId(v string)`

SetId sets Id field to given value.


### GetTimeframe

`func (o *ConsumerRiskReportCustomization) GetTimeframe() ConsumerRiskReportTimeFrame`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ConsumerRiskReportCustomization) GetTimeframeOk() (*ConsumerRiskReportTimeFrame, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ConsumerRiskReportCustomization) SetTimeframe(v ConsumerRiskReportTimeFrame)`

SetTimeframe sets Timeframe field to given value.


### GetMinLimit

`func (o *ConsumerRiskReportCustomization) GetMinLimit() int32`

GetMinLimit returns the MinLimit field if non-nil, zero value otherwise.

### GetMinLimitOk

`func (o *ConsumerRiskReportCustomization) GetMinLimitOk() (*int32, bool)`

GetMinLimitOk returns a tuple with the MinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLimit

`func (o *ConsumerRiskReportCustomization) SetMinLimit(v int32)`

SetMinLimit sets MinLimit field to given value.


### GetMaxLimit

`func (o *ConsumerRiskReportCustomization) GetMaxLimit() int32`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *ConsumerRiskReportCustomization) GetMaxLimitOk() (*int32, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *ConsumerRiskReportCustomization) SetMaxLimit(v int32)`

SetMaxLimit sets MaxLimit field to given value.


### GetRiskTolerance

`func (o *ConsumerRiskReportCustomization) GetRiskTolerance() int32`

GetRiskTolerance returns the RiskTolerance field if non-nil, zero value otherwise.

### GetRiskToleranceOk

`func (o *ConsumerRiskReportCustomization) GetRiskToleranceOk() (*int32, bool)`

GetRiskToleranceOk returns a tuple with the RiskTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTolerance

`func (o *ConsumerRiskReportCustomization) SetRiskTolerance(v int32)`

SetRiskTolerance sets RiskTolerance field to given value.

### HasRiskTolerance

`func (o *ConsumerRiskReportCustomization) HasRiskTolerance() bool`

HasRiskTolerance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


