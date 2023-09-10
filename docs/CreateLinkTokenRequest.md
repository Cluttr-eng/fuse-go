# CreateLinkTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionId** | **string** | An id that is unique for an institution. | 
**Entity** | [**Entity**](Entity.md) |  | 
**ClientName** | **string** | The name of your application. This is what will be displayed to users. | 
**SessionClientSecret** | **string** | The session client secret created from the &#39;Create session client secret&#39; endpoint | 
**WebhookUrl** | Pointer to **string** | This field allows you to set a unique webhook URL for each individual entity. By specifying an entity-specific webhook URL, you can receive and process data events for each entity separately. If this field is left empty, the organization-wide webhook URL set in the sandbox/production environment will be used as the default for all entities. | [optional] 
**Mx** | Pointer to [**CreateLinkTokenRequestMx**](CreateLinkTokenRequestMx.md) |  | [optional] 
**Plaid** | Pointer to [**CreateLinkTokenRequestPlaid**](CreateLinkTokenRequestPlaid.md) |  | [optional] 
**Teller** | Pointer to [**CreateLinkTokenRequestTeller**](CreateLinkTokenRequestTeller.md) |  | [optional] 
**Snaptrade** | Pointer to [**CreateLinkTokenRequestSnaptrade**](CreateLinkTokenRequestSnaptrade.md) |  | [optional] 

## Methods

### NewCreateLinkTokenRequest

`func NewCreateLinkTokenRequest(institutionId string, entity Entity, clientName string, sessionClientSecret string, ) *CreateLinkTokenRequest`

NewCreateLinkTokenRequest instantiates a new CreateLinkTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkTokenRequestWithDefaults

`func NewCreateLinkTokenRequestWithDefaults() *CreateLinkTokenRequest`

NewCreateLinkTokenRequestWithDefaults instantiates a new CreateLinkTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionId

`func (o *CreateLinkTokenRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *CreateLinkTokenRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *CreateLinkTokenRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetEntity

`func (o *CreateLinkTokenRequest) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CreateLinkTokenRequest) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CreateLinkTokenRequest) SetEntity(v Entity)`

SetEntity sets Entity field to given value.


### GetClientName

`func (o *CreateLinkTokenRequest) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CreateLinkTokenRequest) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CreateLinkTokenRequest) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetSessionClientSecret

`func (o *CreateLinkTokenRequest) GetSessionClientSecret() string`

GetSessionClientSecret returns the SessionClientSecret field if non-nil, zero value otherwise.

### GetSessionClientSecretOk

`func (o *CreateLinkTokenRequest) GetSessionClientSecretOk() (*string, bool)`

GetSessionClientSecretOk returns a tuple with the SessionClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientSecret

`func (o *CreateLinkTokenRequest) SetSessionClientSecret(v string)`

SetSessionClientSecret sets SessionClientSecret field to given value.


### GetWebhookUrl

`func (o *CreateLinkTokenRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateLinkTokenRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateLinkTokenRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateLinkTokenRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetMx

`func (o *CreateLinkTokenRequest) GetMx() CreateLinkTokenRequestMx`

GetMx returns the Mx field if non-nil, zero value otherwise.

### GetMxOk

`func (o *CreateLinkTokenRequest) GetMxOk() (*CreateLinkTokenRequestMx, bool)`

GetMxOk returns a tuple with the Mx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMx

`func (o *CreateLinkTokenRequest) SetMx(v CreateLinkTokenRequestMx)`

SetMx sets Mx field to given value.

### HasMx

`func (o *CreateLinkTokenRequest) HasMx() bool`

HasMx returns a boolean if a field has been set.

### GetPlaid

`func (o *CreateLinkTokenRequest) GetPlaid() CreateLinkTokenRequestPlaid`

GetPlaid returns the Plaid field if non-nil, zero value otherwise.

### GetPlaidOk

`func (o *CreateLinkTokenRequest) GetPlaidOk() (*CreateLinkTokenRequestPlaid, bool)`

GetPlaidOk returns a tuple with the Plaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaid

`func (o *CreateLinkTokenRequest) SetPlaid(v CreateLinkTokenRequestPlaid)`

SetPlaid sets Plaid field to given value.

### HasPlaid

`func (o *CreateLinkTokenRequest) HasPlaid() bool`

HasPlaid returns a boolean if a field has been set.

### GetTeller

`func (o *CreateLinkTokenRequest) GetTeller() CreateLinkTokenRequestTeller`

GetTeller returns the Teller field if non-nil, zero value otherwise.

### GetTellerOk

`func (o *CreateLinkTokenRequest) GetTellerOk() (*CreateLinkTokenRequestTeller, bool)`

GetTellerOk returns a tuple with the Teller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeller

`func (o *CreateLinkTokenRequest) SetTeller(v CreateLinkTokenRequestTeller)`

SetTeller sets Teller field to given value.

### HasTeller

`func (o *CreateLinkTokenRequest) HasTeller() bool`

HasTeller returns a boolean if a field has been set.

### GetSnaptrade

`func (o *CreateLinkTokenRequest) GetSnaptrade() CreateLinkTokenRequestSnaptrade`

GetSnaptrade returns the Snaptrade field if non-nil, zero value otherwise.

### GetSnaptradeOk

`func (o *CreateLinkTokenRequest) GetSnaptradeOk() (*CreateLinkTokenRequestSnaptrade, bool)`

GetSnaptradeOk returns a tuple with the Snaptrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnaptrade

`func (o *CreateLinkTokenRequest) SetSnaptrade(v CreateLinkTokenRequestSnaptrade)`

SetSnaptrade sets Snaptrade field to given value.

### HasSnaptrade

`func (o *CreateLinkTokenRequest) HasSnaptrade() bool`

HasSnaptrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


