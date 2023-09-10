# SyncTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to [**[]Transaction**](Transaction.md) | Transactions that have been added to the item since &#x60;cursor&#x60; ordered by ascending last modified time. | [optional] 
**Modified** | Pointer to [**[]Transaction**](Transaction.md) | Transactions that have been modified on the item since &#x60;cursor&#x60; ordered by ascending last modified time. | [optional] 
**Removed** | Pointer to [**[]SyncTransactionsResponseRemovedInner**](SyncTransactionsResponseRemovedInner.md) | Transactions that have been removed from the item since &#x60;cursor&#x60; ordered by ascending last modified time. | [optional] 
**NextCursor** | Pointer to **string** | Cursor used for fetching any future updates after the latest update provided in this response. The cursor obtained after all pages have been pulled (indicated by &#x60;has_next&#x60; being &#x60;false&#x60;) will be valid for at least 1 year. This cursor should be persisted for later calls. | [optional] 
**HasNext** | Pointer to **bool** | Represents if more than requested count of transaction updates exist. If true, the additional updates can be fetched by making an additional request with &#x60;cursor&#x60; set to &#x60;next_cursor&#x60;. If &#x60;has_next&#x60; is true, it&#39;s important to pull all available pages, to make it less likely for underlying data changes to conflict with pagination. | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewSyncTransactionsResponse

`func NewSyncTransactionsResponse() *SyncTransactionsResponse`

NewSyncTransactionsResponse instantiates a new SyncTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncTransactionsResponseWithDefaults

`func NewSyncTransactionsResponseWithDefaults() *SyncTransactionsResponse`

NewSyncTransactionsResponseWithDefaults instantiates a new SyncTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *SyncTransactionsResponse) GetAdded() []Transaction`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *SyncTransactionsResponse) GetAddedOk() (*[]Transaction, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *SyncTransactionsResponse) SetAdded(v []Transaction)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *SyncTransactionsResponse) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetModified

`func (o *SyncTransactionsResponse) GetModified() []Transaction`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SyncTransactionsResponse) GetModifiedOk() (*[]Transaction, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SyncTransactionsResponse) SetModified(v []Transaction)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SyncTransactionsResponse) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetRemoved

`func (o *SyncTransactionsResponse) GetRemoved() []SyncTransactionsResponseRemovedInner`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *SyncTransactionsResponse) GetRemovedOk() (*[]SyncTransactionsResponseRemovedInner, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *SyncTransactionsResponse) SetRemoved(v []SyncTransactionsResponseRemovedInner)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *SyncTransactionsResponse) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetNextCursor

`func (o *SyncTransactionsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SyncTransactionsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SyncTransactionsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *SyncTransactionsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetHasNext

`func (o *SyncTransactionsResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *SyncTransactionsResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *SyncTransactionsResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *SyncTransactionsResponse) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetRequestId

`func (o *SyncTransactionsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SyncTransactionsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SyncTransactionsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SyncTransactionsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


