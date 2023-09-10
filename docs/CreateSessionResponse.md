# CreateSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | **string** | 4 hours from the point of creation | 
**ClientSecret** | **string** | Token needed by the frontend sdk to start the process | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewCreateSessionResponse

`func NewCreateSessionResponse(expiration string, clientSecret string, requestId string, ) *CreateSessionResponse`

NewCreateSessionResponse instantiates a new CreateSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionResponseWithDefaults

`func NewCreateSessionResponseWithDefaults() *CreateSessionResponse`

NewCreateSessionResponseWithDefaults instantiates a new CreateSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *CreateSessionResponse) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CreateSessionResponse) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CreateSessionResponse) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetClientSecret

`func (o *CreateSessionResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateSessionResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateSessionResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetRequestId

`func (o *CreateSessionResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateSessionResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateSessionResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


