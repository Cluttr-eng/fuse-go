# FuseApiErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**Aggregator**](Aggregator.md) |  | [optional] 
**Errors** | Pointer to [**[]FuseApiAggregatorError**](FuseApiAggregatorError.md) |  | [optional] 

## Methods

### NewFuseApiErrorData

`func NewFuseApiErrorData() *FuseApiErrorData`

NewFuseApiErrorData instantiates a new FuseApiErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuseApiErrorDataWithDefaults

`func NewFuseApiErrorDataWithDefaults() *FuseApiErrorData`

NewFuseApiErrorDataWithDefaults instantiates a new FuseApiErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *FuseApiErrorData) GetAggregator() Aggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *FuseApiErrorData) GetAggregatorOk() (*Aggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *FuseApiErrorData) SetAggregator(v Aggregator)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *FuseApiErrorData) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetErrors

`func (o *FuseApiErrorData) GetErrors() []FuseApiAggregatorError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FuseApiErrorData) GetErrorsOk() (*[]FuseApiAggregatorError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FuseApiErrorData) SetErrors(v []FuseApiAggregatorError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FuseApiErrorData) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


