# MigrateFinancialConnectionsTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionData** | [**MigrateFinancialConnectionsAggregatorConnectionData**](MigrateFinancialConnectionsAggregatorConnectionData.md) |  | 
**Aggregator** | **string** | The aggregator being migrated (either &#39;plaid&#39;, &#39;mx&#39; or &#39;teller). | 
**Entity** | [**MigrateFinancialConnectionsTokenRequestEntity**](MigrateFinancialConnectionsTokenRequestEntity.md) |  | 
**FuseProducts** | [**[]Product**](Product.md) | A list of Fuse products that the migrated connection will have access to. | 

## Methods

### NewMigrateFinancialConnectionsTokenRequest

`func NewMigrateFinancialConnectionsTokenRequest(connectionData MigrateFinancialConnectionsAggregatorConnectionData, aggregator string, entity MigrateFinancialConnectionsTokenRequestEntity, fuseProducts []Product, ) *MigrateFinancialConnectionsTokenRequest`

NewMigrateFinancialConnectionsTokenRequest instantiates a new MigrateFinancialConnectionsTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateFinancialConnectionsTokenRequestWithDefaults

`func NewMigrateFinancialConnectionsTokenRequestWithDefaults() *MigrateFinancialConnectionsTokenRequest`

NewMigrateFinancialConnectionsTokenRequestWithDefaults instantiates a new MigrateFinancialConnectionsTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionData

`func (o *MigrateFinancialConnectionsTokenRequest) GetConnectionData() MigrateFinancialConnectionsAggregatorConnectionData`

GetConnectionData returns the ConnectionData field if non-nil, zero value otherwise.

### GetConnectionDataOk

`func (o *MigrateFinancialConnectionsTokenRequest) GetConnectionDataOk() (*MigrateFinancialConnectionsAggregatorConnectionData, bool)`

GetConnectionDataOk returns a tuple with the ConnectionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionData

`func (o *MigrateFinancialConnectionsTokenRequest) SetConnectionData(v MigrateFinancialConnectionsAggregatorConnectionData)`

SetConnectionData sets ConnectionData field to given value.


### GetAggregator

`func (o *MigrateFinancialConnectionsTokenRequest) GetAggregator() string`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *MigrateFinancialConnectionsTokenRequest) GetAggregatorOk() (*string, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *MigrateFinancialConnectionsTokenRequest) SetAggregator(v string)`

SetAggregator sets Aggregator field to given value.


### GetEntity

`func (o *MigrateFinancialConnectionsTokenRequest) GetEntity() MigrateFinancialConnectionsTokenRequestEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *MigrateFinancialConnectionsTokenRequest) GetEntityOk() (*MigrateFinancialConnectionsTokenRequestEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *MigrateFinancialConnectionsTokenRequest) SetEntity(v MigrateFinancialConnectionsTokenRequestEntity)`

SetEntity sets Entity field to given value.


### GetFuseProducts

`func (o *MigrateFinancialConnectionsTokenRequest) GetFuseProducts() []Product`

GetFuseProducts returns the FuseProducts field if non-nil, zero value otherwise.

### GetFuseProductsOk

`func (o *MigrateFinancialConnectionsTokenRequest) GetFuseProductsOk() (*[]Product, bool)`

GetFuseProductsOk returns a tuple with the FuseProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuseProducts

`func (o *MigrateFinancialConnectionsTokenRequest) SetFuseProducts(v []Product)`

SetFuseProducts sets FuseProducts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


