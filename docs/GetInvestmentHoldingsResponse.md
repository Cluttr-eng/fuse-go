# GetInvestmentHoldingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]FinancialConnectionsAccount**](FinancialConnectionsAccount.md) |  | 
**Holdings** | [**[]FinancialConnectionsHolding**](FinancialConnectionsHolding.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetInvestmentHoldingsResponse

`func NewGetInvestmentHoldingsResponse(accounts []FinancialConnectionsAccount, holdings []FinancialConnectionsHolding, requestId string, ) *GetInvestmentHoldingsResponse`

NewGetInvestmentHoldingsResponse instantiates a new GetInvestmentHoldingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvestmentHoldingsResponseWithDefaults

`func NewGetInvestmentHoldingsResponseWithDefaults() *GetInvestmentHoldingsResponse`

NewGetInvestmentHoldingsResponseWithDefaults instantiates a new GetInvestmentHoldingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetInvestmentHoldingsResponse) GetAccounts() []FinancialConnectionsAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetInvestmentHoldingsResponse) GetAccountsOk() (*[]FinancialConnectionsAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetInvestmentHoldingsResponse) SetAccounts(v []FinancialConnectionsAccount)`

SetAccounts sets Accounts field to given value.


### GetHoldings

`func (o *GetInvestmentHoldingsResponse) GetHoldings() []FinancialConnectionsHolding`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *GetInvestmentHoldingsResponse) GetHoldingsOk() (*[]FinancialConnectionsHolding, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *GetInvestmentHoldingsResponse) SetHoldings(v []FinancialConnectionsHolding)`

SetHoldings sets Holdings field to given value.


### GetRequestId

`func (o *GetInvestmentHoldingsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetInvestmentHoldingsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetInvestmentHoldingsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


