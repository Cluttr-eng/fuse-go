# FinancialConnectionsInvestmentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | The remote ID of the Investment transaction | 
**RemoteAccountId** | **string** | Remote Account Id of the transaction, ie Plaid Account Id | 
**AccountName** | Pointer to **string** | The name of the account associated with the investment transaction | [optional] 
**Amount** | **float32** | The amount of the investment transaction, in cents. The format of this value is a double. | 
**Description** | **string** | A description of the investment transaction | 
**Fees** | **float32** | The fees associated with the investment transaction, in cents. The format of this value is a double. | 
**Currency** | [**Currency**](Currency.md) |  | 
**Date** | **time.Time** | The date and time of the investment transaction | 
**Type** | **string** | The type of the investment transaction (e.g., &#39;buy&#39;, &#39;sell&#39;, &#39;dividend&#39;) | 
**Subtype** | Pointer to [**FinancialConnectionsInvestmentTransactionSubtype**](FinancialConnectionsInvestmentTransactionSubtype.md) |  | [optional] 
**Quantity** | **float32** | The number of units of the security involved in this transaction | 
**Price** | **float32** | The price of the security involved in this transaction, in cents. The format of this value is a double. | 
**Security** | [**FinancialConnectionsInvestmentSecurity**](FinancialConnectionsInvestmentSecurity.md) |  | 

## Methods

### NewFinancialConnectionsInvestmentTransaction

`func NewFinancialConnectionsInvestmentTransaction(remoteId string, remoteAccountId string, amount float32, description string, fees float32, currency Currency, date time.Time, type_ string, quantity float32, price float32, security FinancialConnectionsInvestmentSecurity, ) *FinancialConnectionsInvestmentTransaction`

NewFinancialConnectionsInvestmentTransaction instantiates a new FinancialConnectionsInvestmentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsInvestmentTransactionWithDefaults

`func NewFinancialConnectionsInvestmentTransactionWithDefaults() *FinancialConnectionsInvestmentTransaction`

NewFinancialConnectionsInvestmentTransactionWithDefaults instantiates a new FinancialConnectionsInvestmentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *FinancialConnectionsInvestmentTransaction) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FinancialConnectionsInvestmentTransaction) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FinancialConnectionsInvestmentTransaction) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetRemoteAccountId

`func (o *FinancialConnectionsInvestmentTransaction) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *FinancialConnectionsInvestmentTransaction) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *FinancialConnectionsInvestmentTransaction) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetAccountName

`func (o *FinancialConnectionsInvestmentTransaction) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *FinancialConnectionsInvestmentTransaction) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *FinancialConnectionsInvestmentTransaction) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *FinancialConnectionsInvestmentTransaction) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAmount

`func (o *FinancialConnectionsInvestmentTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialConnectionsInvestmentTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialConnectionsInvestmentTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *FinancialConnectionsInvestmentTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FinancialConnectionsInvestmentTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FinancialConnectionsInvestmentTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFees

`func (o *FinancialConnectionsInvestmentTransaction) GetFees() float32`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *FinancialConnectionsInvestmentTransaction) GetFeesOk() (*float32, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *FinancialConnectionsInvestmentTransaction) SetFees(v float32)`

SetFees sets Fees field to given value.


### GetCurrency

`func (o *FinancialConnectionsInvestmentTransaction) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialConnectionsInvestmentTransaction) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialConnectionsInvestmentTransaction) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetDate

`func (o *FinancialConnectionsInvestmentTransaction) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FinancialConnectionsInvestmentTransaction) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FinancialConnectionsInvestmentTransaction) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetType

`func (o *FinancialConnectionsInvestmentTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialConnectionsInvestmentTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialConnectionsInvestmentTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetSubtype

`func (o *FinancialConnectionsInvestmentTransaction) GetSubtype() FinancialConnectionsInvestmentTransactionSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FinancialConnectionsInvestmentTransaction) GetSubtypeOk() (*FinancialConnectionsInvestmentTransactionSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FinancialConnectionsInvestmentTransaction) SetSubtype(v FinancialConnectionsInvestmentTransactionSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *FinancialConnectionsInvestmentTransaction) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetQuantity

`func (o *FinancialConnectionsInvestmentTransaction) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *FinancialConnectionsInvestmentTransaction) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *FinancialConnectionsInvestmentTransaction) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *FinancialConnectionsInvestmentTransaction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FinancialConnectionsInvestmentTransaction) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FinancialConnectionsInvestmentTransaction) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetSecurity

`func (o *FinancialConnectionsInvestmentTransaction) GetSecurity() FinancialConnectionsInvestmentSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FinancialConnectionsInvestmentTransaction) GetSecurityOk() (*FinancialConnectionsInvestmentSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FinancialConnectionsInvestmentTransaction) SetSecurity(v FinancialConnectionsInvestmentSecurity)`

SetSecurity sets Security field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


