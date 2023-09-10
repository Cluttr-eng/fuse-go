# GetFinancialConnectionsBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access token for authentication | 
**Options** | Pointer to [**GetFinancialConnectionsBalanceRequestOptions**](GetFinancialConnectionsBalanceRequestOptions.md) |  | [optional] 

## Methods

### NewGetFinancialConnectionsBalanceRequest

`func NewGetFinancialConnectionsBalanceRequest(accessToken string, ) *GetFinancialConnectionsBalanceRequest`

NewGetFinancialConnectionsBalanceRequest instantiates a new GetFinancialConnectionsBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsBalanceRequestWithDefaults

`func NewGetFinancialConnectionsBalanceRequestWithDefaults() *GetFinancialConnectionsBalanceRequest`

NewGetFinancialConnectionsBalanceRequestWithDefaults instantiates a new GetFinancialConnectionsBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetFinancialConnectionsBalanceRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetFinancialConnectionsBalanceRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetFinancialConnectionsBalanceRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetOptions

`func (o *GetFinancialConnectionsBalanceRequest) GetOptions() GetFinancialConnectionsBalanceRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetFinancialConnectionsBalanceRequest) GetOptionsOk() (*GetFinancialConnectionsBalanceRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetFinancialConnectionsBalanceRequest) SetOptions(v GetFinancialConnectionsBalanceRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetFinancialConnectionsBalanceRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


