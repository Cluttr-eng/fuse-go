# AssetReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | Pointer to [**AssetReport**](AssetReport.md) |  | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewAssetReportResponse

`func NewAssetReportResponse() *AssetReportResponse`

NewAssetReportResponse instantiates a new AssetReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportResponseWithDefaults

`func NewAssetReportResponseWithDefaults() *AssetReportResponse`

NewAssetReportResponseWithDefaults instantiates a new AssetReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *AssetReportResponse) GetReport() AssetReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *AssetReportResponse) GetReportOk() (*AssetReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *AssetReportResponse) SetReport(v AssetReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *AssetReportResponse) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetRequestId

`func (o *AssetReportResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AssetReportResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AssetReportResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AssetReportResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


