# AssetReportAccountsInnerHistoricalBalancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date of the calculated historical balance, in an ISO 8601 format (YYYY-MM-DD) | [optional] 
**Current** | Pointer to **float32** | The total amount of funds in the account, calculated from the current balance in the balance object by subtracting inflows and adding back outflows according to the posted date of each transaction. | [optional] 
**IsoCurrencyCode** | Pointer to **string** | The ISO-4217 currency code of the balance. | [optional] 

## Methods

### NewAssetReportAccountsInnerHistoricalBalancesInner

`func NewAssetReportAccountsInnerHistoricalBalancesInner() *AssetReportAccountsInnerHistoricalBalancesInner`

NewAssetReportAccountsInnerHistoricalBalancesInner instantiates a new AssetReportAccountsInnerHistoricalBalancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportAccountsInnerHistoricalBalancesInnerWithDefaults

`func NewAssetReportAccountsInnerHistoricalBalancesInnerWithDefaults() *AssetReportAccountsInnerHistoricalBalancesInner`

NewAssetReportAccountsInnerHistoricalBalancesInnerWithDefaults instantiates a new AssetReportAccountsInnerHistoricalBalancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCurrent

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetIsoCurrencyCode

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *AssetReportAccountsInnerHistoricalBalancesInner) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


