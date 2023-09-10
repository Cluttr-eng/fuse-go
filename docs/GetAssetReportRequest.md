# GetAssetReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access fuse token corresponding to the financial account to be refresh the Asset Report for. | 
**AssetReportToken** | **string** | The token associated with the Asset Report to retrieve. | 

## Methods

### NewGetAssetReportRequest

`func NewGetAssetReportRequest(accessToken string, assetReportToken string, ) *GetAssetReportRequest`

NewGetAssetReportRequest instantiates a new GetAssetReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetReportRequestWithDefaults

`func NewGetAssetReportRequestWithDefaults() *GetAssetReportRequest`

NewGetAssetReportRequestWithDefaults instantiates a new GetAssetReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetAssetReportRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetAssetReportRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetAssetReportRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAssetReportToken

`func (o *GetAssetReportRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *GetAssetReportRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *GetAssetReportRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


