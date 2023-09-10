# DeleteFinancialConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinancialConnectionId** | **string** | Id of the deleted financial connection | 
**AccessToken** | **string** | Access token of the deleted financial connection | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewDeleteFinancialConnectionResponse

`func NewDeleteFinancialConnectionResponse(financialConnectionId string, accessToken string, requestId string, ) *DeleteFinancialConnectionResponse`

NewDeleteFinancialConnectionResponse instantiates a new DeleteFinancialConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFinancialConnectionResponseWithDefaults

`func NewDeleteFinancialConnectionResponseWithDefaults() *DeleteFinancialConnectionResponse`

NewDeleteFinancialConnectionResponseWithDefaults instantiates a new DeleteFinancialConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinancialConnectionId

`func (o *DeleteFinancialConnectionResponse) GetFinancialConnectionId() string`

GetFinancialConnectionId returns the FinancialConnectionId field if non-nil, zero value otherwise.

### GetFinancialConnectionIdOk

`func (o *DeleteFinancialConnectionResponse) GetFinancialConnectionIdOk() (*string, bool)`

GetFinancialConnectionIdOk returns a tuple with the FinancialConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnectionId

`func (o *DeleteFinancialConnectionResponse) SetFinancialConnectionId(v string)`

SetFinancialConnectionId sets FinancialConnectionId field to given value.


### GetAccessToken

`func (o *DeleteFinancialConnectionResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DeleteFinancialConnectionResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DeleteFinancialConnectionResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRequestId

`func (o *DeleteFinancialConnectionResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeleteFinancialConnectionResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeleteFinancialConnectionResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


