# AssetReportAccountsInnerBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **NullableFloat32** | Amount after factoring in pending balances | [optional] 
**Current** | Pointer to **NullableFloat32** | Amount without factoring in pending balances | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the balance. | [optional] 

## Methods

### NewAssetReportAccountsInnerBalance

`func NewAssetReportAccountsInnerBalance() *AssetReportAccountsInnerBalance`

NewAssetReportAccountsInnerBalance instantiates a new AssetReportAccountsInnerBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportAccountsInnerBalanceWithDefaults

`func NewAssetReportAccountsInnerBalanceWithDefaults() *AssetReportAccountsInnerBalance`

NewAssetReportAccountsInnerBalanceWithDefaults instantiates a new AssetReportAccountsInnerBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AssetReportAccountsInnerBalance) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AssetReportAccountsInnerBalance) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AssetReportAccountsInnerBalance) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AssetReportAccountsInnerBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *AssetReportAccountsInnerBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *AssetReportAccountsInnerBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *AssetReportAccountsInnerBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AssetReportAccountsInnerBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AssetReportAccountsInnerBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AssetReportAccountsInnerBalance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *AssetReportAccountsInnerBalance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *AssetReportAccountsInnerBalance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetIsoCurrencyCode

`func (o *AssetReportAccountsInnerBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *AssetReportAccountsInnerBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *AssetReportAccountsInnerBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *AssetReportAccountsInnerBalance) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *AssetReportAccountsInnerBalance) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *AssetReportAccountsInnerBalance) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


