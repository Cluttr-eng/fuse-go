# CreateLinkTokenRequestPlaid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Follows the same schema as Plaid&#39;s Link Token Create Schema(https://plaid.com/docs/api/tokens/#linktokencreate). &#39;products&#39;, &#39;client_id&#39;, &#39;secret&#39;, &#39;client_user_id&#39;, &#39;client_name&#39;, &#39;webhook&#39;, &#39;institution_data&#39; and &#39;country_codes&#39; (only US and Canada is supported right now) will be set by Fuse and override any values you set. | [optional] 

## Methods

### NewCreateLinkTokenRequestPlaid

`func NewCreateLinkTokenRequestPlaid() *CreateLinkTokenRequestPlaid`

NewCreateLinkTokenRequestPlaid instantiates a new CreateLinkTokenRequestPlaid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkTokenRequestPlaidWithDefaults

`func NewCreateLinkTokenRequestPlaidWithDefaults() *CreateLinkTokenRequestPlaid`

NewCreateLinkTokenRequestPlaidWithDefaults instantiates a new CreateLinkTokenRequestPlaid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateLinkTokenRequestPlaid) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLinkTokenRequestPlaid) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLinkTokenRequestPlaid) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLinkTokenRequestPlaid) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


