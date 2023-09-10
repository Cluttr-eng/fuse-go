# GetFinancialConnectionsBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | [**[]FinancialConnectionsAccountBalance**](FinancialConnectionsAccountBalance.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialConnectionsBalanceResponse

`func NewGetFinancialConnectionsBalanceResponse(balances []FinancialConnectionsAccountBalance, requestId string, ) *GetFinancialConnectionsBalanceResponse`

NewGetFinancialConnectionsBalanceResponse instantiates a new GetFinancialConnectionsBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsBalanceResponseWithDefaults

`func NewGetFinancialConnectionsBalanceResponseWithDefaults() *GetFinancialConnectionsBalanceResponse`

NewGetFinancialConnectionsBalanceResponseWithDefaults instantiates a new GetFinancialConnectionsBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *GetFinancialConnectionsBalanceResponse) GetBalances() []FinancialConnectionsAccountBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *GetFinancialConnectionsBalanceResponse) GetBalancesOk() (*[]FinancialConnectionsAccountBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *GetFinancialConnectionsBalanceResponse) SetBalances(v []FinancialConnectionsAccountBalance)`

SetBalances sets Balances field to given value.


### GetRequestId

`func (o *GetFinancialConnectionsBalanceResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialConnectionsBalanceResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialConnectionsBalanceResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


