# GetInvestmentHoldingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token of the financial institution connection | 
**Options** | Pointer to [**GetInvestmentHoldingsRequestOptions**](GetInvestmentHoldingsRequestOptions.md) |  | [optional] 

## Methods

### NewGetInvestmentHoldingsRequest

`func NewGetInvestmentHoldingsRequest(accessToken string, ) *GetInvestmentHoldingsRequest`

NewGetInvestmentHoldingsRequest instantiates a new GetInvestmentHoldingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvestmentHoldingsRequestWithDefaults

`func NewGetInvestmentHoldingsRequestWithDefaults() *GetInvestmentHoldingsRequest`

NewGetInvestmentHoldingsRequestWithDefaults instantiates a new GetInvestmentHoldingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetInvestmentHoldingsRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetInvestmentHoldingsRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetInvestmentHoldingsRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetOptions

`func (o *GetInvestmentHoldingsRequest) GetOptions() GetInvestmentHoldingsRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetInvestmentHoldingsRequest) GetOptionsOk() (*GetInvestmentHoldingsRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetInvestmentHoldingsRequest) SetOptions(v GetInvestmentHoldingsRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetInvestmentHoldingsRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


