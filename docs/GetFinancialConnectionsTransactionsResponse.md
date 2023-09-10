# GetFinancialConnectionsTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]Transaction**](Transaction.md) |  | 
**TotalTransactions** | **float32** | The total number of transactions available within the date range specified. If total_transactions is larger than the size of the transactions array, more transactions are available and can be fetched | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialConnectionsTransactionsResponse

`func NewGetFinancialConnectionsTransactionsResponse(transactions []Transaction, totalTransactions float32, requestId string, ) *GetFinancialConnectionsTransactionsResponse`

NewGetFinancialConnectionsTransactionsResponse instantiates a new GetFinancialConnectionsTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsTransactionsResponseWithDefaults

`func NewGetFinancialConnectionsTransactionsResponseWithDefaults() *GetFinancialConnectionsTransactionsResponse`

NewGetFinancialConnectionsTransactionsResponseWithDefaults instantiates a new GetFinancialConnectionsTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *GetFinancialConnectionsTransactionsResponse) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetFinancialConnectionsTransactionsResponse) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetFinancialConnectionsTransactionsResponse) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.


### GetTotalTransactions

`func (o *GetFinancialConnectionsTransactionsResponse) GetTotalTransactions() float32`

GetTotalTransactions returns the TotalTransactions field if non-nil, zero value otherwise.

### GetTotalTransactionsOk

`func (o *GetFinancialConnectionsTransactionsResponse) GetTotalTransactionsOk() (*float32, bool)`

GetTotalTransactionsOk returns a tuple with the TotalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactions

`func (o *GetFinancialConnectionsTransactionsResponse) SetTotalTransactions(v float32)`

SetTotalTransactions sets TotalTransactions field to given value.


### GetRequestId

`func (o *GetFinancialConnectionsTransactionsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialConnectionsTransactionsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialConnectionsTransactionsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


