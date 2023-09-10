# CreateLinkTokenRequestTellerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectAccount** | Pointer to **string** | The mode of account selection: - &#39;disabled&#39;: automatically connect all the supported financial accounts associated with this user&#39;s account at the institution (default). - &#39;single&#39;: the user will see a list of supported financial accounts and will need to select one to connect - &#39;multiple&#39;: the user will see a list of supported financial accounts and will need to select one or more to connect | [optional] 
**AccountFilter** | Pointer to [**CreateLinkTokenRequestTellerConfigAccountFilter**](CreateLinkTokenRequestTellerConfigAccountFilter.md) |  | [optional] 

## Methods

### NewCreateLinkTokenRequestTellerConfig

`func NewCreateLinkTokenRequestTellerConfig() *CreateLinkTokenRequestTellerConfig`

NewCreateLinkTokenRequestTellerConfig instantiates a new CreateLinkTokenRequestTellerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkTokenRequestTellerConfigWithDefaults

`func NewCreateLinkTokenRequestTellerConfigWithDefaults() *CreateLinkTokenRequestTellerConfig`

NewCreateLinkTokenRequestTellerConfigWithDefaults instantiates a new CreateLinkTokenRequestTellerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectAccount

`func (o *CreateLinkTokenRequestTellerConfig) GetSelectAccount() string`

GetSelectAccount returns the SelectAccount field if non-nil, zero value otherwise.

### GetSelectAccountOk

`func (o *CreateLinkTokenRequestTellerConfig) GetSelectAccountOk() (*string, bool)`

GetSelectAccountOk returns a tuple with the SelectAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAccount

`func (o *CreateLinkTokenRequestTellerConfig) SetSelectAccount(v string)`

SetSelectAccount sets SelectAccount field to given value.

### HasSelectAccount

`func (o *CreateLinkTokenRequestTellerConfig) HasSelectAccount() bool`

HasSelectAccount returns a boolean if a field has been set.

### GetAccountFilter

`func (o *CreateLinkTokenRequestTellerConfig) GetAccountFilter() CreateLinkTokenRequestTellerConfigAccountFilter`

GetAccountFilter returns the AccountFilter field if non-nil, zero value otherwise.

### GetAccountFilterOk

`func (o *CreateLinkTokenRequestTellerConfig) GetAccountFilterOk() (*CreateLinkTokenRequestTellerConfigAccountFilter, bool)`

GetAccountFilterOk returns a tuple with the AccountFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilter

`func (o *CreateLinkTokenRequestTellerConfig) SetAccountFilter(v CreateLinkTokenRequestTellerConfigAccountFilter)`

SetAccountFilter sets AccountFilter field to given value.

### HasAccountFilter

`func (o *CreateLinkTokenRequestTellerConfig) HasAccountFilter() bool`

HasAccountFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


