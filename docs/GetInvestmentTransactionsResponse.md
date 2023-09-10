# GetInvestmentTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]FinancialConnectionsAccount**](FinancialConnectionsAccount.md) |  | 
**InvestmentTransactions** | [**[]FinancialConnectionsInvestmentTransaction**](FinancialConnectionsInvestmentTransaction.md) |  | 
**TotalTransactions** | Pointer to **float32** | The total number of transactions within the specified date range. | [optional] 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetInvestmentTransactionsResponse

`func NewGetInvestmentTransactionsResponse(accounts []FinancialConnectionsAccount, investmentTransactions []FinancialConnectionsInvestmentTransaction, requestId string, ) *GetInvestmentTransactionsResponse`

NewGetInvestmentTransactionsResponse instantiates a new GetInvestmentTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvestmentTransactionsResponseWithDefaults

`func NewGetInvestmentTransactionsResponseWithDefaults() *GetInvestmentTransactionsResponse`

NewGetInvestmentTransactionsResponseWithDefaults instantiates a new GetInvestmentTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetInvestmentTransactionsResponse) GetAccounts() []FinancialConnectionsAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetInvestmentTransactionsResponse) GetAccountsOk() (*[]FinancialConnectionsAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetInvestmentTransactionsResponse) SetAccounts(v []FinancialConnectionsAccount)`

SetAccounts sets Accounts field to given value.


### GetInvestmentTransactions

`func (o *GetInvestmentTransactionsResponse) GetInvestmentTransactions() []FinancialConnectionsInvestmentTransaction`

GetInvestmentTransactions returns the InvestmentTransactions field if non-nil, zero value otherwise.

### GetInvestmentTransactionsOk

`func (o *GetInvestmentTransactionsResponse) GetInvestmentTransactionsOk() (*[]FinancialConnectionsInvestmentTransaction, bool)`

GetInvestmentTransactionsOk returns a tuple with the InvestmentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentTransactions

`func (o *GetInvestmentTransactionsResponse) SetInvestmentTransactions(v []FinancialConnectionsInvestmentTransaction)`

SetInvestmentTransactions sets InvestmentTransactions field to given value.


### GetTotalTransactions

`func (o *GetInvestmentTransactionsResponse) GetTotalTransactions() float32`

GetTotalTransactions returns the TotalTransactions field if non-nil, zero value otherwise.

### GetTotalTransactionsOk

`func (o *GetInvestmentTransactionsResponse) GetTotalTransactionsOk() (*float32, bool)`

GetTotalTransactionsOk returns a tuple with the TotalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactions

`func (o *GetInvestmentTransactionsResponse) SetTotalTransactions(v float32)`

SetTotalTransactions sets TotalTransactions field to given value.

### HasTotalTransactions

`func (o *GetInvestmentTransactionsResponse) HasTotalTransactions() bool`

HasTotalTransactions returns a boolean if a field has been set.

### GetRequestId

`func (o *GetInvestmentTransactionsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetInvestmentTransactionsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetInvestmentTransactionsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


