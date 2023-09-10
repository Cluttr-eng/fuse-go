# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**WebhookType**](WebhookType.md) |  | 
**FinancialConnectionId** | **string** | Financial connection id associated with the webhook | 
**Environment** | **string** |  | 
**Source** | [**WebhookSource**](WebhookSource.md) |  | 
**VerificationToken** | Pointer to **string** | Aggregator verification data needed to verify the webhook | [optional] 
**AssetReportId** | Pointer to **string** | Exists for assets.report_ready webhooks | [optional] 
**HistoricalTransactionsAvailable** | Pointer to **bool** | Exists for transactions.updates webhooks. Indicates if historical transaction information (up to 24 months) is ready to be queried. | [optional] 
**RemovedTransactionIds** | Pointer to **[]string** | Exists for transactions.updates webhooks. Currently only supported by Plaid. | [optional] 
**RemoteData** | **interface{}** |  | 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(type_ WebhookType, financialConnectionId string, environment string, source WebhookSource, remoteData interface{}, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookEvent) GetType() WebhookType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookEvent) GetTypeOk() (*WebhookType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookEvent) SetType(v WebhookType)`

SetType sets Type field to given value.


### GetFinancialConnectionId

`func (o *WebhookEvent) GetFinancialConnectionId() string`

GetFinancialConnectionId returns the FinancialConnectionId field if non-nil, zero value otherwise.

### GetFinancialConnectionIdOk

`func (o *WebhookEvent) GetFinancialConnectionIdOk() (*string, bool)`

GetFinancialConnectionIdOk returns a tuple with the FinancialConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnectionId

`func (o *WebhookEvent) SetFinancialConnectionId(v string)`

SetFinancialConnectionId sets FinancialConnectionId field to given value.


### GetEnvironment

`func (o *WebhookEvent) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WebhookEvent) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WebhookEvent) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetSource

`func (o *WebhookEvent) GetSource() WebhookSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebhookEvent) GetSourceOk() (*WebhookSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebhookEvent) SetSource(v WebhookSource)`

SetSource sets Source field to given value.


### GetVerificationToken

`func (o *WebhookEvent) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *WebhookEvent) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *WebhookEvent) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.

### HasVerificationToken

`func (o *WebhookEvent) HasVerificationToken() bool`

HasVerificationToken returns a boolean if a field has been set.

### GetAssetReportId

`func (o *WebhookEvent) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *WebhookEvent) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *WebhookEvent) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.

### HasAssetReportId

`func (o *WebhookEvent) HasAssetReportId() bool`

HasAssetReportId returns a boolean if a field has been set.

### GetHistoricalTransactionsAvailable

`func (o *WebhookEvent) GetHistoricalTransactionsAvailable() bool`

GetHistoricalTransactionsAvailable returns the HistoricalTransactionsAvailable field if non-nil, zero value otherwise.

### GetHistoricalTransactionsAvailableOk

`func (o *WebhookEvent) GetHistoricalTransactionsAvailableOk() (*bool, bool)`

GetHistoricalTransactionsAvailableOk returns a tuple with the HistoricalTransactionsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalTransactionsAvailable

`func (o *WebhookEvent) SetHistoricalTransactionsAvailable(v bool)`

SetHistoricalTransactionsAvailable sets HistoricalTransactionsAvailable field to given value.

### HasHistoricalTransactionsAvailable

`func (o *WebhookEvent) HasHistoricalTransactionsAvailable() bool`

HasHistoricalTransactionsAvailable returns a boolean if a field has been set.

### GetRemovedTransactionIds

`func (o *WebhookEvent) GetRemovedTransactionIds() []string`

GetRemovedTransactionIds returns the RemovedTransactionIds field if non-nil, zero value otherwise.

### GetRemovedTransactionIdsOk

`func (o *WebhookEvent) GetRemovedTransactionIdsOk() (*[]string, bool)`

GetRemovedTransactionIdsOk returns a tuple with the RemovedTransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedTransactionIds

`func (o *WebhookEvent) SetRemovedTransactionIds(v []string)`

SetRemovedTransactionIds sets RemovedTransactionIds field to given value.

### HasRemovedTransactionIds

`func (o *WebhookEvent) HasRemovedTransactionIds() bool`

HasRemovedTransactionIds returns a boolean if a field has been set.

### GetRemoteData

`func (o *WebhookEvent) GetRemoteData() interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *WebhookEvent) GetRemoteDataOk() (*interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *WebhookEvent) SetRemoteData(v interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *WebhookEvent) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *WebhookEvent) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


