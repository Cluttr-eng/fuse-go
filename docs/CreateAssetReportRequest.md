# CreateAssetReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access fuse token corresponding to the financial account to be create the Asset Report for. | 
**DaysRequested** | **float32** | The maximum integer number of days of history to include in the Asset Report | 

## Methods

### NewCreateAssetReportRequest

`func NewCreateAssetReportRequest(accessToken string, daysRequested float32, ) *CreateAssetReportRequest`

NewCreateAssetReportRequest instantiates a new CreateAssetReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetReportRequestWithDefaults

`func NewCreateAssetReportRequestWithDefaults() *CreateAssetReportRequest`

NewCreateAssetReportRequestWithDefaults instantiates a new CreateAssetReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *CreateAssetReportRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CreateAssetReportRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CreateAssetReportRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetDaysRequested

`func (o *CreateAssetReportRequest) GetDaysRequested() float32`

GetDaysRequested returns the DaysRequested field if non-nil, zero value otherwise.

### GetDaysRequestedOk

`func (o *CreateAssetReportRequest) GetDaysRequestedOk() (*float32, bool)`

GetDaysRequestedOk returns a tuple with the DaysRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRequested

`func (o *CreateAssetReportRequest) SetDaysRequested(v float32)`

SetDaysRequested sets DaysRequested field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


