# GetFinanceScoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinanceScore** | [**FinanceScore**](FinanceScore.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinanceScoreResponse

`func NewGetFinanceScoreResponse(financeScore FinanceScore, requestId string, ) *GetFinanceScoreResponse`

NewGetFinanceScoreResponse instantiates a new GetFinanceScoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceScoreResponseWithDefaults

`func NewGetFinanceScoreResponseWithDefaults() *GetFinanceScoreResponse`

NewGetFinanceScoreResponseWithDefaults instantiates a new GetFinanceScoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinanceScore

`func (o *GetFinanceScoreResponse) GetFinanceScore() FinanceScore`

GetFinanceScore returns the FinanceScore field if non-nil, zero value otherwise.

### GetFinanceScoreOk

`func (o *GetFinanceScoreResponse) GetFinanceScoreOk() (*FinanceScore, bool)`

GetFinanceScoreOk returns a tuple with the FinanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceScore

`func (o *GetFinanceScoreResponse) SetFinanceScore(v FinanceScore)`

SetFinanceScore sets FinanceScore field to given value.


### GetRequestId

`func (o *GetFinanceScoreResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinanceScoreResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinanceScoreResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


