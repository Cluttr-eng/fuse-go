# AssetReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetReportId** | Pointer to **string** | A unique ID identifying an Asset Report. | [optional] 
**AssetReportToken** | Pointer to **string** | A token that can be provided to endpoints such as /asset_report/get or /asset_report/pdf/get to fetch or update an Asset Report. | [optional] 
**DateGenerated** | Pointer to **string** | The date and time when the Asset Report was created, in ISO 8601 format | [optional] 
**DaysRequested** | Pointer to **int32** | The duration of transaction history you requested | [optional] 
**Accounts** | Pointer to [**[]AssetReportAccountsInner**](AssetReportAccountsInner.md) | An array of Asset Reports, one for each account in the Asset Report. | [optional] 

## Methods

### NewAssetReport

`func NewAssetReport() *AssetReport`

NewAssetReport instantiates a new AssetReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportWithDefaults

`func NewAssetReportWithDefaults() *AssetReport`

NewAssetReportWithDefaults instantiates a new AssetReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetReportId

`func (o *AssetReport) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *AssetReport) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *AssetReport) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.

### HasAssetReportId

`func (o *AssetReport) HasAssetReportId() bool`

HasAssetReportId returns a boolean if a field has been set.

### GetAssetReportToken

`func (o *AssetReport) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReport) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReport) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.

### HasAssetReportToken

`func (o *AssetReport) HasAssetReportToken() bool`

HasAssetReportToken returns a boolean if a field has been set.

### GetDateGenerated

`func (o *AssetReport) GetDateGenerated() string`

GetDateGenerated returns the DateGenerated field if non-nil, zero value otherwise.

### GetDateGeneratedOk

`func (o *AssetReport) GetDateGeneratedOk() (*string, bool)`

GetDateGeneratedOk returns a tuple with the DateGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateGenerated

`func (o *AssetReport) SetDateGenerated(v string)`

SetDateGenerated sets DateGenerated field to given value.

### HasDateGenerated

`func (o *AssetReport) HasDateGenerated() bool`

HasDateGenerated returns a boolean if a field has been set.

### GetDaysRequested

`func (o *AssetReport) GetDaysRequested() int32`

GetDaysRequested returns the DaysRequested field if non-nil, zero value otherwise.

### GetDaysRequestedOk

`func (o *AssetReport) GetDaysRequestedOk() (*int32, bool)`

GetDaysRequestedOk returns a tuple with the DaysRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRequested

`func (o *AssetReport) SetDaysRequested(v int32)`

SetDaysRequested sets DaysRequested field to given value.

### HasDaysRequested

`func (o *AssetReport) HasDaysRequested() bool`

HasDaysRequested returns a boolean if a field has been set.

### GetAccounts

`func (o *AssetReport) GetAccounts() []AssetReportAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AssetReport) GetAccountsOk() (*[]AssetReportAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AssetReport) SetAccounts(v []AssetReportAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AssetReport) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


