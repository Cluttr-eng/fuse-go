# FuseApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** |  | 
**Title** | **string** |  | 
**Details** | **string** |  | 
**Code** | [**FuseApiErrorCode**](FuseApiErrorCode.md) |  | 
**Type** | [**FuseApiErrorType**](FuseApiErrorType.md) |  | 
**Source** | **string** |  | 
**Data** | Pointer to [**FuseApiErrorData**](FuseApiErrorData.md) |  | [optional] 

## Methods

### NewFuseApiError

`func NewFuseApiError(requestId string, title string, details string, code FuseApiErrorCode, type_ FuseApiErrorType, source string, ) *FuseApiError`

NewFuseApiError instantiates a new FuseApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuseApiErrorWithDefaults

`func NewFuseApiErrorWithDefaults() *FuseApiError`

NewFuseApiErrorWithDefaults instantiates a new FuseApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *FuseApiError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *FuseApiError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *FuseApiError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTitle

`func (o *FuseApiError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FuseApiError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FuseApiError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetails

`func (o *FuseApiError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FuseApiError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FuseApiError) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetCode

`func (o *FuseApiError) GetCode() FuseApiErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FuseApiError) GetCodeOk() (*FuseApiErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FuseApiError) SetCode(v FuseApiErrorCode)`

SetCode sets Code field to given value.


### GetType

`func (o *FuseApiError) GetType() FuseApiErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FuseApiError) GetTypeOk() (*FuseApiErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FuseApiError) SetType(v FuseApiErrorType)`

SetType sets Type field to given value.


### GetSource

`func (o *FuseApiError) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FuseApiError) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FuseApiError) SetSource(v string)`

SetSource sets Source field to given value.


### GetData

`func (o *FuseApiError) GetData() FuseApiErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FuseApiError) GetDataOk() (*FuseApiErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FuseApiError) SetData(v FuseApiErrorData)`

SetData sets Data field to given value.

### HasData

`func (o *FuseApiError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


