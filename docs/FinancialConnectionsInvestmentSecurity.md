# FinancialConnectionsInvestmentSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Remote Id of the security, ie Plaid or Snaptrade security id | 
**Symbol** | **string** | The trading symbol for publicly traded securities, or a short identifier if available. | 
**Isin** | Pointer to **string** | The International Securities Identification Number (ISIN) uniquely identifies the security. | [optional] 
**Sedol** | Pointer to **string** | The Stock Exchange Daily Official List (SEDOL) code uniquely identifies the security, primarily used in the United Kingdom and Ireland. | [optional] 
**Cusip** | Pointer to **string** | The Committee on Uniform Securities Identification Procedures (CUSIP) number uniquely identifies the security, primarily used in the United States and Canada. | [optional] 
**ClosePrice** | Pointer to **float32** | The closing price of the security, in cents, at the end of the most recent trading day. The format of this value is a double. | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**Name** | Pointer to **string** | A descriptive name for the security, suitable for display. | [optional] 
**Type** | Pointer to [**FinancialConnectionsInvestmentSecurityType**](FinancialConnectionsInvestmentSecurityType.md) |  | [optional] 
**Exchange** | Pointer to [**FinancialConnectionsInvestmentSecurityExchange**](FinancialConnectionsInvestmentSecurityExchange.md) |  | [optional] 

## Methods

### NewFinancialConnectionsInvestmentSecurity

`func NewFinancialConnectionsInvestmentSecurity(remoteId string, symbol string, currency Currency, ) *FinancialConnectionsInvestmentSecurity`

NewFinancialConnectionsInvestmentSecurity instantiates a new FinancialConnectionsInvestmentSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsInvestmentSecurityWithDefaults

`func NewFinancialConnectionsInvestmentSecurityWithDefaults() *FinancialConnectionsInvestmentSecurity`

NewFinancialConnectionsInvestmentSecurityWithDefaults instantiates a new FinancialConnectionsInvestmentSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *FinancialConnectionsInvestmentSecurity) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FinancialConnectionsInvestmentSecurity) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FinancialConnectionsInvestmentSecurity) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetSymbol

`func (o *FinancialConnectionsInvestmentSecurity) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FinancialConnectionsInvestmentSecurity) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FinancialConnectionsInvestmentSecurity) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsin

`func (o *FinancialConnectionsInvestmentSecurity) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *FinancialConnectionsInvestmentSecurity) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *FinancialConnectionsInvestmentSecurity) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *FinancialConnectionsInvestmentSecurity) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetSedol

`func (o *FinancialConnectionsInvestmentSecurity) GetSedol() string`

GetSedol returns the Sedol field if non-nil, zero value otherwise.

### GetSedolOk

`func (o *FinancialConnectionsInvestmentSecurity) GetSedolOk() (*string, bool)`

GetSedolOk returns a tuple with the Sedol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedol

`func (o *FinancialConnectionsInvestmentSecurity) SetSedol(v string)`

SetSedol sets Sedol field to given value.

### HasSedol

`func (o *FinancialConnectionsInvestmentSecurity) HasSedol() bool`

HasSedol returns a boolean if a field has been set.

### GetCusip

`func (o *FinancialConnectionsInvestmentSecurity) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *FinancialConnectionsInvestmentSecurity) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *FinancialConnectionsInvestmentSecurity) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *FinancialConnectionsInvestmentSecurity) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetClosePrice

`func (o *FinancialConnectionsInvestmentSecurity) GetClosePrice() float32`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *FinancialConnectionsInvestmentSecurity) GetClosePriceOk() (*float32, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *FinancialConnectionsInvestmentSecurity) SetClosePrice(v float32)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *FinancialConnectionsInvestmentSecurity) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetCurrency

`func (o *FinancialConnectionsInvestmentSecurity) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialConnectionsInvestmentSecurity) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialConnectionsInvestmentSecurity) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetName

`func (o *FinancialConnectionsInvestmentSecurity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialConnectionsInvestmentSecurity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialConnectionsInvestmentSecurity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FinancialConnectionsInvestmentSecurity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FinancialConnectionsInvestmentSecurity) GetType() FinancialConnectionsInvestmentSecurityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialConnectionsInvestmentSecurity) GetTypeOk() (*FinancialConnectionsInvestmentSecurityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialConnectionsInvestmentSecurity) SetType(v FinancialConnectionsInvestmentSecurityType)`

SetType sets Type field to given value.

### HasType

`func (o *FinancialConnectionsInvestmentSecurity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExchange

`func (o *FinancialConnectionsInvestmentSecurity) GetExchange() FinancialConnectionsInvestmentSecurityExchange`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *FinancialConnectionsInvestmentSecurity) GetExchangeOk() (*FinancialConnectionsInvestmentSecurityExchange, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *FinancialConnectionsInvestmentSecurity) SetExchange(v FinancialConnectionsInvestmentSecurityExchange)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *FinancialConnectionsInvestmentSecurity) HasExchange() bool`

HasExchange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


