# CreateLinkTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkToken** | **string** | Token needed by the frontend sdk to initiate the connection | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewCreateLinkTokenResponse

`func NewCreateLinkTokenResponse(linkToken string, requestId string, ) *CreateLinkTokenResponse`

NewCreateLinkTokenResponse instantiates a new CreateLinkTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkTokenResponseWithDefaults

`func NewCreateLinkTokenResponseWithDefaults() *CreateLinkTokenResponse`

NewCreateLinkTokenResponseWithDefaults instantiates a new CreateLinkTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkToken

`func (o *CreateLinkTokenResponse) GetLinkToken() string`

GetLinkToken returns the LinkToken field if non-nil, zero value otherwise.

### GetLinkTokenOk

`func (o *CreateLinkTokenResponse) GetLinkTokenOk() (*string, bool)`

GetLinkTokenOk returns a tuple with the LinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToken

`func (o *CreateLinkTokenResponse) SetLinkToken(v string)`

SetLinkToken sets LinkToken field to given value.


### GetRequestId

`func (o *CreateLinkTokenResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateLinkTokenResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateLinkTokenResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


