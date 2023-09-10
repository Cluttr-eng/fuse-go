# EvalConsumerRiskReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowSize** | **float32** | The size of the window for training | 
**TimeFrame** | [**ConsumerRiskReportTimeFrame**](ConsumerRiskReportTimeFrame.md) |  | 
**Events** | [**[]EvalConsumerRiskReportRequestEventsInner**](EvalConsumerRiskReportRequestEventsInner.md) |  | 

## Methods

### NewEvalConsumerRiskReportRequest

`func NewEvalConsumerRiskReportRequest(windowSize float32, timeFrame ConsumerRiskReportTimeFrame, events []EvalConsumerRiskReportRequestEventsInner, ) *EvalConsumerRiskReportRequest`

NewEvalConsumerRiskReportRequest instantiates a new EvalConsumerRiskReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalConsumerRiskReportRequestWithDefaults

`func NewEvalConsumerRiskReportRequestWithDefaults() *EvalConsumerRiskReportRequest`

NewEvalConsumerRiskReportRequestWithDefaults instantiates a new EvalConsumerRiskReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowSize

`func (o *EvalConsumerRiskReportRequest) GetWindowSize() float32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *EvalConsumerRiskReportRequest) GetWindowSizeOk() (*float32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *EvalConsumerRiskReportRequest) SetWindowSize(v float32)`

SetWindowSize sets WindowSize field to given value.


### GetTimeFrame

`func (o *EvalConsumerRiskReportRequest) GetTimeFrame() ConsumerRiskReportTimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *EvalConsumerRiskReportRequest) GetTimeFrameOk() (*ConsumerRiskReportTimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *EvalConsumerRiskReportRequest) SetTimeFrame(v ConsumerRiskReportTimeFrame)`

SetTimeFrame sets TimeFrame field to given value.


### GetEvents

`func (o *EvalConsumerRiskReportRequest) GetEvents() []EvalConsumerRiskReportRequestEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EvalConsumerRiskReportRequest) GetEventsOk() (*[]EvalConsumerRiskReportRequestEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EvalConsumerRiskReportRequest) SetEvents(v []EvalConsumerRiskReportRequestEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


