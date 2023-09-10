# GetFinancialConnectionsOwnersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]GetFinancialConnectionsOwnersResponseAccountsInner**](GetFinancialConnectionsOwnersResponseAccountsInner.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialConnectionsOwnersResponse

`func NewGetFinancialConnectionsOwnersResponse(accounts []GetFinancialConnectionsOwnersResponseAccountsInner, requestId string, ) *GetFinancialConnectionsOwnersResponse`

NewGetFinancialConnectionsOwnersResponse instantiates a new GetFinancialConnectionsOwnersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsOwnersResponseWithDefaults

`func NewGetFinancialConnectionsOwnersResponseWithDefaults() *GetFinancialConnectionsOwnersResponse`

NewGetFinancialConnectionsOwnersResponseWithDefaults instantiates a new GetFinancialConnectionsOwnersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetFinancialConnectionsOwnersResponse) GetAccounts() []GetFinancialConnectionsOwnersResponseAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetFinancialConnectionsOwnersResponse) GetAccountsOk() (*[]GetFinancialConnectionsOwnersResponseAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetFinancialConnectionsOwnersResponse) SetAccounts(v []GetFinancialConnectionsOwnersResponseAccountsInner)`

SetAccounts sets Accounts field to given value.


### GetRequestId

`func (o *GetFinancialConnectionsOwnersResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialConnectionsOwnersResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialConnectionsOwnersResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


