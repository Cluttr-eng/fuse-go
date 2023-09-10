# SyncFinancialConnectionsDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Response message | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewSyncFinancialConnectionsDataResponse

`func NewSyncFinancialConnectionsDataResponse() *SyncFinancialConnectionsDataResponse`

NewSyncFinancialConnectionsDataResponse instantiates a new SyncFinancialConnectionsDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncFinancialConnectionsDataResponseWithDefaults

`func NewSyncFinancialConnectionsDataResponseWithDefaults() *SyncFinancialConnectionsDataResponse`

NewSyncFinancialConnectionsDataResponseWithDefaults instantiates a new SyncFinancialConnectionsDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SyncFinancialConnectionsDataResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyncFinancialConnectionsDataResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyncFinancialConnectionsDataResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyncFinancialConnectionsDataResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRequestId

`func (o *SyncFinancialConnectionsDataResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SyncFinancialConnectionsDataResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SyncFinancialConnectionsDataResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SyncFinancialConnectionsDataResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


