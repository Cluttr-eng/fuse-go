# GetConsumerRiskReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerRiskReport** | [**ConsumerRiskReport**](ConsumerRiskReport.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetConsumerRiskReportResponse

`func NewGetConsumerRiskReportResponse(consumerRiskReport ConsumerRiskReport, requestId string, ) *GetConsumerRiskReportResponse`

NewGetConsumerRiskReportResponse instantiates a new GetConsumerRiskReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsumerRiskReportResponseWithDefaults

`func NewGetConsumerRiskReportResponseWithDefaults() *GetConsumerRiskReportResponse`

NewGetConsumerRiskReportResponseWithDefaults instantiates a new GetConsumerRiskReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerRiskReport

`func (o *GetConsumerRiskReportResponse) GetConsumerRiskReport() ConsumerRiskReport`

GetConsumerRiskReport returns the ConsumerRiskReport field if non-nil, zero value otherwise.

### GetConsumerRiskReportOk

`func (o *GetConsumerRiskReportResponse) GetConsumerRiskReportOk() (*ConsumerRiskReport, bool)`

GetConsumerRiskReportOk returns a tuple with the ConsumerRiskReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRiskReport

`func (o *GetConsumerRiskReportResponse) SetConsumerRiskReport(v ConsumerRiskReport)`

SetConsumerRiskReport sets ConsumerRiskReport field to given value.


### GetRequestId

`func (o *GetConsumerRiskReportResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetConsumerRiskReportResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetConsumerRiskReportResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


