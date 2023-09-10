# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFinancialInstitutionAggregators** | [**[]Aggregator**](Aggregator.md) |  | 
**Products** | [**[]Product**](Product.md) | List of products that you would like the institutions to support | 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) | List of country codes that you would like the institutions to support | [optional] 
**Entity** | [**Entity**](Entity.md) |  | 
**AccessToken** | Pointer to **string** | The fuse access token for an existing financial connection. This will perform the process to reconnect an existing disconnected account. | [optional] 
**IsWebView** | Pointer to **bool** | Set to false for web SDKs (e.g., React), and true for mobile SDKs (e.g., React Native, Flutter, iOS, Android). | [optional] 

## Methods

### NewCreateSessionRequest

`func NewCreateSessionRequest(supportedFinancialInstitutionAggregators []Aggregator, products []Product, entity Entity, ) *CreateSessionRequest`

NewCreateSessionRequest instantiates a new CreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionRequestWithDefaults

`func NewCreateSessionRequestWithDefaults() *CreateSessionRequest`

NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFinancialInstitutionAggregators

`func (o *CreateSessionRequest) GetSupportedFinancialInstitutionAggregators() []Aggregator`

GetSupportedFinancialInstitutionAggregators returns the SupportedFinancialInstitutionAggregators field if non-nil, zero value otherwise.

### GetSupportedFinancialInstitutionAggregatorsOk

`func (o *CreateSessionRequest) GetSupportedFinancialInstitutionAggregatorsOk() (*[]Aggregator, bool)`

GetSupportedFinancialInstitutionAggregatorsOk returns a tuple with the SupportedFinancialInstitutionAggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFinancialInstitutionAggregators

`func (o *CreateSessionRequest) SetSupportedFinancialInstitutionAggregators(v []Aggregator)`

SetSupportedFinancialInstitutionAggregators sets SupportedFinancialInstitutionAggregators field to given value.


### GetProducts

`func (o *CreateSessionRequest) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateSessionRequest) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateSessionRequest) SetProducts(v []Product)`

SetProducts sets Products field to given value.


### GetCountryCodes

`func (o *CreateSessionRequest) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *CreateSessionRequest) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *CreateSessionRequest) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *CreateSessionRequest) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetEntity

`func (o *CreateSessionRequest) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CreateSessionRequest) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CreateSessionRequest) SetEntity(v Entity)`

SetEntity sets Entity field to given value.


### GetAccessToken

`func (o *CreateSessionRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CreateSessionRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CreateSessionRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *CreateSessionRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetIsWebView

`func (o *CreateSessionRequest) GetIsWebView() bool`

GetIsWebView returns the IsWebView field if non-nil, zero value otherwise.

### GetIsWebViewOk

`func (o *CreateSessionRequest) GetIsWebViewOk() (*bool, bool)`

GetIsWebViewOk returns a tuple with the IsWebView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebView

`func (o *CreateSessionRequest) SetIsWebView(v bool)`

SetIsWebView sets IsWebView field to given value.

### HasIsWebView

`func (o *CreateSessionRequest) HasIsWebView() bool`

HasIsWebView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


