# GetFinancialConnectionsAccountStatementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access token for authentication | 
**RemoteAccountId** | **string** | The remote account id to retrieve the statement for. | 
**Date** | Pointer to **string** | The year and month for the account statement to be retrieved in YYYY-MM. | [optional] 

## Methods

### NewGetFinancialConnectionsAccountStatementRequest

`func NewGetFinancialConnectionsAccountStatementRequest(accessToken string, remoteAccountId string, ) *GetFinancialConnectionsAccountStatementRequest`

NewGetFinancialConnectionsAccountStatementRequest instantiates a new GetFinancialConnectionsAccountStatementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsAccountStatementRequestWithDefaults

`func NewGetFinancialConnectionsAccountStatementRequestWithDefaults() *GetFinancialConnectionsAccountStatementRequest`

NewGetFinancialConnectionsAccountStatementRequestWithDefaults instantiates a new GetFinancialConnectionsAccountStatementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetFinancialConnectionsAccountStatementRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetFinancialConnectionsAccountStatementRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetFinancialConnectionsAccountStatementRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRemoteAccountId

`func (o *GetFinancialConnectionsAccountStatementRequest) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *GetFinancialConnectionsAccountStatementRequest) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *GetFinancialConnectionsAccountStatementRequest) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetDate

`func (o *GetFinancialConnectionsAccountStatementRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetFinancialConnectionsAccountStatementRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetFinancialConnectionsAccountStatementRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetFinancialConnectionsAccountStatementRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


