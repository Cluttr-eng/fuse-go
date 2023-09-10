# FinancialConnectionsAccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteAccountId** | **string** | Remote Account Id of the transaction, ie Plaid Account Id | 
**Available** | Pointer to **NullableFloat32** | Amount in cents after factoring in pending balances. The format of this value is a double. For accounts with credit features, the available funds generally equal the credit limit. Some institutions may not provide an available balance calculation. If this is the case, Fuse will return a null value for the available balance. To ensure you have the most accurate information, we recommend obtaining the current balance by using &#39;balance.available || balance.current&#39;. | [optional] 
**Current** | Pointer to **NullableFloat32** | Amount in cents without factoring in pending balances. The format of this value is a double. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the balance. | [optional] 
**LastUpdatedDate** | Pointer to **NullableString** | The last time the account balance was updated, represented as an ISO 8601 timestamp (YYYY-MM-DDTHH:mm:ssZ). This value may not be available for some accounts. | [optional] 

## Methods

### NewFinancialConnectionsAccountBalance

`func NewFinancialConnectionsAccountBalance(remoteAccountId string, ) *FinancialConnectionsAccountBalance`

NewFinancialConnectionsAccountBalance instantiates a new FinancialConnectionsAccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsAccountBalanceWithDefaults

`func NewFinancialConnectionsAccountBalanceWithDefaults() *FinancialConnectionsAccountBalance`

NewFinancialConnectionsAccountBalanceWithDefaults instantiates a new FinancialConnectionsAccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteAccountId

`func (o *FinancialConnectionsAccountBalance) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *FinancialConnectionsAccountBalance) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *FinancialConnectionsAccountBalance) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetAvailable

`func (o *FinancialConnectionsAccountBalance) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *FinancialConnectionsAccountBalance) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *FinancialConnectionsAccountBalance) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *FinancialConnectionsAccountBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *FinancialConnectionsAccountBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *FinancialConnectionsAccountBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *FinancialConnectionsAccountBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *FinancialConnectionsAccountBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *FinancialConnectionsAccountBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *FinancialConnectionsAccountBalance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *FinancialConnectionsAccountBalance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *FinancialConnectionsAccountBalance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetIsoCurrencyCode

`func (o *FinancialConnectionsAccountBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *FinancialConnectionsAccountBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *FinancialConnectionsAccountBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *FinancialConnectionsAccountBalance) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *FinancialConnectionsAccountBalance) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *FinancialConnectionsAccountBalance) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetLastUpdatedDate

`func (o *FinancialConnectionsAccountBalance) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *FinancialConnectionsAccountBalance) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *FinancialConnectionsAccountBalance) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *FinancialConnectionsAccountBalance) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### SetLastUpdatedDateNil

`func (o *FinancialConnectionsAccountBalance) SetLastUpdatedDateNil(b bool)`

 SetLastUpdatedDateNil sets the value for LastUpdatedDate to be an explicit nil

### UnsetLastUpdatedDate
`func (o *FinancialConnectionsAccountBalance) UnsetLastUpdatedDate()`

UnsetLastUpdatedDate ensures that no value is present for LastUpdatedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


