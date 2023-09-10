# MigrateFinancialConnectionsTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionData** | [**MigrateFinancialConnectionsAggregatorConnectionData**](MigrateFinancialConnectionsAggregatorConnectionData.md) |  | 
**FuseAccessToken** | **string** | Fuse access token for the fuse connection | 
**FuseFinancialConnectionId** | **string** | Financial connection id for the fuse connection | 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewMigrateFinancialConnectionsTokenResponse

`func NewMigrateFinancialConnectionsTokenResponse(connectionData MigrateFinancialConnectionsAggregatorConnectionData, fuseAccessToken string, fuseFinancialConnectionId string, ) *MigrateFinancialConnectionsTokenResponse`

NewMigrateFinancialConnectionsTokenResponse instantiates a new MigrateFinancialConnectionsTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateFinancialConnectionsTokenResponseWithDefaults

`func NewMigrateFinancialConnectionsTokenResponseWithDefaults() *MigrateFinancialConnectionsTokenResponse`

NewMigrateFinancialConnectionsTokenResponseWithDefaults instantiates a new MigrateFinancialConnectionsTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionData

`func (o *MigrateFinancialConnectionsTokenResponse) GetConnectionData() MigrateFinancialConnectionsAggregatorConnectionData`

GetConnectionData returns the ConnectionData field if non-nil, zero value otherwise.

### GetConnectionDataOk

`func (o *MigrateFinancialConnectionsTokenResponse) GetConnectionDataOk() (*MigrateFinancialConnectionsAggregatorConnectionData, bool)`

GetConnectionDataOk returns a tuple with the ConnectionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionData

`func (o *MigrateFinancialConnectionsTokenResponse) SetConnectionData(v MigrateFinancialConnectionsAggregatorConnectionData)`

SetConnectionData sets ConnectionData field to given value.


### GetFuseAccessToken

`func (o *MigrateFinancialConnectionsTokenResponse) GetFuseAccessToken() string`

GetFuseAccessToken returns the FuseAccessToken field if non-nil, zero value otherwise.

### GetFuseAccessTokenOk

`func (o *MigrateFinancialConnectionsTokenResponse) GetFuseAccessTokenOk() (*string, bool)`

GetFuseAccessTokenOk returns a tuple with the FuseAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuseAccessToken

`func (o *MigrateFinancialConnectionsTokenResponse) SetFuseAccessToken(v string)`

SetFuseAccessToken sets FuseAccessToken field to given value.


### GetFuseFinancialConnectionId

`func (o *MigrateFinancialConnectionsTokenResponse) GetFuseFinancialConnectionId() string`

GetFuseFinancialConnectionId returns the FuseFinancialConnectionId field if non-nil, zero value otherwise.

### GetFuseFinancialConnectionIdOk

`func (o *MigrateFinancialConnectionsTokenResponse) GetFuseFinancialConnectionIdOk() (*string, bool)`

GetFuseFinancialConnectionIdOk returns a tuple with the FuseFinancialConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuseFinancialConnectionId

`func (o *MigrateFinancialConnectionsTokenResponse) SetFuseFinancialConnectionId(v string)`

SetFuseFinancialConnectionId sets FuseFinancialConnectionId field to given value.


### GetRequestId

`func (o *MigrateFinancialConnectionsTokenResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MigrateFinancialConnectionsTokenResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MigrateFinancialConnectionsTokenResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *MigrateFinancialConnectionsTokenResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


