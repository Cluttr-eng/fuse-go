# SyncTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token of the financial institution connection | 
**Cursor** | Pointer to **string** | The cursor value represents the last update requested. Providing it will cause the response to only return changes after this update. If omitted, the entire history of updates will be returned, starting with the first-added transactions on the item. | [optional] 
**Count** | Pointer to **int32** | The number of transaction updates to fetch. | [optional] 

## Methods

### NewSyncTransactionsRequest

`func NewSyncTransactionsRequest(accessToken string, ) *SyncTransactionsRequest`

NewSyncTransactionsRequest instantiates a new SyncTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncTransactionsRequestWithDefaults

`func NewSyncTransactionsRequestWithDefaults() *SyncTransactionsRequest`

NewSyncTransactionsRequestWithDefaults instantiates a new SyncTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *SyncTransactionsRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SyncTransactionsRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SyncTransactionsRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetCursor

`func (o *SyncTransactionsRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SyncTransactionsRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SyncTransactionsRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SyncTransactionsRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetCount

`func (o *SyncTransactionsRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SyncTransactionsRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SyncTransactionsRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SyncTransactionsRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


