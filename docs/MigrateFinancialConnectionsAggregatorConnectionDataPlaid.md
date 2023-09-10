# MigrateFinancialConnectionsAggregatorConnectionDataPlaid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The Plaid access token associated with the user&#39;s financial accounts. | 
**UseItemWebhook** | Pointer to **bool** | If true, any webhooks received for this new financial connection will be sent to the webhook url used when the item was created. If false, the webhook url set at the organization sandbox/production environment level will be used instead. | [optional] 

## Methods

### NewMigrateFinancialConnectionsAggregatorConnectionDataPlaid

`func NewMigrateFinancialConnectionsAggregatorConnectionDataPlaid(accessToken string, ) *MigrateFinancialConnectionsAggregatorConnectionDataPlaid`

NewMigrateFinancialConnectionsAggregatorConnectionDataPlaid instantiates a new MigrateFinancialConnectionsAggregatorConnectionDataPlaid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateFinancialConnectionsAggregatorConnectionDataPlaidWithDefaults

`func NewMigrateFinancialConnectionsAggregatorConnectionDataPlaidWithDefaults() *MigrateFinancialConnectionsAggregatorConnectionDataPlaid`

NewMigrateFinancialConnectionsAggregatorConnectionDataPlaidWithDefaults instantiates a new MigrateFinancialConnectionsAggregatorConnectionDataPlaid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetUseItemWebhook

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetUseItemWebhook() bool`

GetUseItemWebhook returns the UseItemWebhook field if non-nil, zero value otherwise.

### GetUseItemWebhookOk

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetUseItemWebhookOk() (*bool, bool)`

GetUseItemWebhookOk returns a tuple with the UseItemWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseItemWebhook

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) SetUseItemWebhook(v bool)`

SetUseItemWebhook sets UseItemWebhook field to given value.

### HasUseItemWebhook

`func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) HasUseItemWebhook() bool`

HasUseItemWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


