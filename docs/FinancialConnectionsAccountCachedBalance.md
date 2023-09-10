# FinancialConnectionsAccountCachedBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **NullableFloat32** | The amount of funds available, in cents, to be withdrawn from the account, as determined by the financial institution. The format of this value is a double. Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by /financial_connections/balances. For accounts with credit features, the available funds generally equal the credit limit. Some institutions may not provide an available balance calculation. If this is the case, Fuse will return a null value for the available balance. To ensure you have the most accurate information, we recommend obtaining the current balance by using &#39;balance.available || balance.current&#39;. | [optional] 
**Current** | Pointer to **NullableFloat32** | Amount without factoring in pending balances, in cents. The format of this value is a double. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the balance. | [optional] 
**LastUpdatedDate** | Pointer to **NullableString** | The date of the last update to the balance. | [optional] 

## Methods

### NewFinancialConnectionsAccountCachedBalance

`func NewFinancialConnectionsAccountCachedBalance() *FinancialConnectionsAccountCachedBalance`

NewFinancialConnectionsAccountCachedBalance instantiates a new FinancialConnectionsAccountCachedBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsAccountCachedBalanceWithDefaults

`func NewFinancialConnectionsAccountCachedBalanceWithDefaults() *FinancialConnectionsAccountCachedBalance`

NewFinancialConnectionsAccountCachedBalanceWithDefaults instantiates a new FinancialConnectionsAccountCachedBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *FinancialConnectionsAccountCachedBalance) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *FinancialConnectionsAccountCachedBalance) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *FinancialConnectionsAccountCachedBalance) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *FinancialConnectionsAccountCachedBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *FinancialConnectionsAccountCachedBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *FinancialConnectionsAccountCachedBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *FinancialConnectionsAccountCachedBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *FinancialConnectionsAccountCachedBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *FinancialConnectionsAccountCachedBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *FinancialConnectionsAccountCachedBalance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *FinancialConnectionsAccountCachedBalance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *FinancialConnectionsAccountCachedBalance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetIsoCurrencyCode

`func (o *FinancialConnectionsAccountCachedBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *FinancialConnectionsAccountCachedBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *FinancialConnectionsAccountCachedBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *FinancialConnectionsAccountCachedBalance) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *FinancialConnectionsAccountCachedBalance) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *FinancialConnectionsAccountCachedBalance) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetLastUpdatedDate

`func (o *FinancialConnectionsAccountCachedBalance) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *FinancialConnectionsAccountCachedBalance) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *FinancialConnectionsAccountCachedBalance) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *FinancialConnectionsAccountCachedBalance) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### SetLastUpdatedDateNil

`func (o *FinancialConnectionsAccountCachedBalance) SetLastUpdatedDateNil(b bool)`

 SetLastUpdatedDateNil sets the value for LastUpdatedDate to be an explicit nil

### UnsetLastUpdatedDate
`func (o *FinancialConnectionsAccountCachedBalance) UnsetLastUpdatedDate()`

UnsetLastUpdatedDate ensures that no value is present for LastUpdatedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


