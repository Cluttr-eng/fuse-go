# EnrichTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]TransactionToEnrich**](TransactionToEnrich.md) |  | 

## Methods

### NewEnrichTransactionsRequest

`func NewEnrichTransactionsRequest(transactions []TransactionToEnrich, ) *EnrichTransactionsRequest`

NewEnrichTransactionsRequest instantiates a new EnrichTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichTransactionsRequestWithDefaults

`func NewEnrichTransactionsRequestWithDefaults() *EnrichTransactionsRequest`

NewEnrichTransactionsRequestWithDefaults instantiates a new EnrichTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *EnrichTransactionsRequest) GetTransactions() []TransactionToEnrich`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *EnrichTransactionsRequest) GetTransactionsOk() (*[]TransactionToEnrich, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *EnrichTransactionsRequest) SetTransactions(v []TransactionToEnrich)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


