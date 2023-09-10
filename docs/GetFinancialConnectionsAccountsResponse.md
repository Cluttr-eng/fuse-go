# GetFinancialConnectionsAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]FinancialConnectionsAccount**](FinancialConnectionsAccount.md) |  | 
**FinancialConnection** | [**FinancialConnectionData**](FinancialConnectionData.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialConnectionsAccountsResponse

`func NewGetFinancialConnectionsAccountsResponse(accounts []FinancialConnectionsAccount, financialConnection FinancialConnectionData, requestId string, ) *GetFinancialConnectionsAccountsResponse`

NewGetFinancialConnectionsAccountsResponse instantiates a new GetFinancialConnectionsAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsAccountsResponseWithDefaults

`func NewGetFinancialConnectionsAccountsResponseWithDefaults() *GetFinancialConnectionsAccountsResponse`

NewGetFinancialConnectionsAccountsResponseWithDefaults instantiates a new GetFinancialConnectionsAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetFinancialConnectionsAccountsResponse) GetAccounts() []FinancialConnectionsAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetFinancialConnectionsAccountsResponse) GetAccountsOk() (*[]FinancialConnectionsAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetFinancialConnectionsAccountsResponse) SetAccounts(v []FinancialConnectionsAccount)`

SetAccounts sets Accounts field to given value.


### GetFinancialConnection

`func (o *GetFinancialConnectionsAccountsResponse) GetFinancialConnection() FinancialConnectionData`

GetFinancialConnection returns the FinancialConnection field if non-nil, zero value otherwise.

### GetFinancialConnectionOk

`func (o *GetFinancialConnectionsAccountsResponse) GetFinancialConnectionOk() (*FinancialConnectionData, bool)`

GetFinancialConnectionOk returns a tuple with the FinancialConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnection

`func (o *GetFinancialConnectionsAccountsResponse) SetFinancialConnection(v FinancialConnectionData)`

SetFinancialConnection sets FinancialConnection field to given value.


### GetRequestId

`func (o *GetFinancialConnectionsAccountsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialConnectionsAccountsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialConnectionsAccountsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


