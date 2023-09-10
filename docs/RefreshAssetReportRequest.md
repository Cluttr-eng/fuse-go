# RefreshAssetReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access fuse token corresponding to the financial account to be refresh the Asset Report for. | 
**AssetReportToken** | **string** | The asset_report_token returned by the original call to /asset_report/create | 
**DaysRequested** | Pointer to **float32** | The maximum integer number of days of history to include in the Asset Report | [optional] 

## Methods

### NewRefreshAssetReportRequest

`func NewRefreshAssetReportRequest(accessToken string, assetReportToken string, ) *RefreshAssetReportRequest`

NewRefreshAssetReportRequest instantiates a new RefreshAssetReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshAssetReportRequestWithDefaults

`func NewRefreshAssetReportRequestWithDefaults() *RefreshAssetReportRequest`

NewRefreshAssetReportRequestWithDefaults instantiates a new RefreshAssetReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RefreshAssetReportRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RefreshAssetReportRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RefreshAssetReportRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAssetReportToken

`func (o *RefreshAssetReportRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *RefreshAssetReportRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *RefreshAssetReportRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.


### GetDaysRequested

`func (o *RefreshAssetReportRequest) GetDaysRequested() float32`

GetDaysRequested returns the DaysRequested field if non-nil, zero value otherwise.

### GetDaysRequestedOk

`func (o *RefreshAssetReportRequest) GetDaysRequestedOk() (*float32, bool)`

GetDaysRequestedOk returns a tuple with the DaysRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRequested

`func (o *RefreshAssetReportRequest) SetDaysRequested(v float32)`

SetDaysRequested sets DaysRequested field to given value.

### HasDaysRequested

`func (o *RefreshAssetReportRequest) HasDaysRequested() bool`

HasDaysRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


