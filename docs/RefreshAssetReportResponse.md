# RefreshAssetReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetReportToken** | Pointer to **string** | A token that can be provided to endpoints such as /asset_report to fetch an asset report. | [optional] 
**AssetReportId** | Pointer to **string** | A unique ID identifying an Asset Report. | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewRefreshAssetReportResponse

`func NewRefreshAssetReportResponse() *RefreshAssetReportResponse`

NewRefreshAssetReportResponse instantiates a new RefreshAssetReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshAssetReportResponseWithDefaults

`func NewRefreshAssetReportResponseWithDefaults() *RefreshAssetReportResponse`

NewRefreshAssetReportResponseWithDefaults instantiates a new RefreshAssetReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetReportToken

`func (o *RefreshAssetReportResponse) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *RefreshAssetReportResponse) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *RefreshAssetReportResponse) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.

### HasAssetReportToken

`func (o *RefreshAssetReportResponse) HasAssetReportToken() bool`

HasAssetReportToken returns a boolean if a field has been set.

### GetAssetReportId

`func (o *RefreshAssetReportResponse) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *RefreshAssetReportResponse) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *RefreshAssetReportResponse) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.

### HasAssetReportId

`func (o *RefreshAssetReportResponse) HasAssetReportId() bool`

HasAssetReportId returns a boolean if a field has been set.

### GetRequestId

`func (o *RefreshAssetReportResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RefreshAssetReportResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RefreshAssetReportResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RefreshAssetReportResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


