# FuseApiWarningData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**Aggregator**](Aggregator.md) |  | [optional] 
**Warnings** | Pointer to [**[]FuseApiWarningDataWarningsInner**](FuseApiWarningDataWarningsInner.md) |  | [optional] 

## Methods

### NewFuseApiWarningData

`func NewFuseApiWarningData() *FuseApiWarningData`

NewFuseApiWarningData instantiates a new FuseApiWarningData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuseApiWarningDataWithDefaults

`func NewFuseApiWarningDataWithDefaults() *FuseApiWarningData`

NewFuseApiWarningDataWithDefaults instantiates a new FuseApiWarningData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *FuseApiWarningData) GetAggregator() Aggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *FuseApiWarningData) GetAggregatorOk() (*Aggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *FuseApiWarningData) SetAggregator(v Aggregator)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *FuseApiWarningData) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetWarnings

`func (o *FuseApiWarningData) GetWarnings() []FuseApiWarningDataWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *FuseApiWarningData) GetWarningsOk() (*[]FuseApiWarningDataWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *FuseApiWarningData) SetWarnings(v []FuseApiWarningDataWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *FuseApiWarningData) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


