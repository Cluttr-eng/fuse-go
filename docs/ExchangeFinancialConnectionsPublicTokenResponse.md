# ExchangeFinancialConnectionsPublicTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Token used for querying data on the user, ie account details, balances etc. This does NOT expire and should be stored securely. | 
**FinancialConnectionId** | **string** | The id of the new financial connection. Every webhook will be sent with this id. Use this id when calling the GET /financial_connection/${financial_connection_id} endpoint.  | 
**Institution** | Pointer to [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**Aggregator** | [**Aggregator**](Aggregator.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewExchangeFinancialConnectionsPublicTokenResponse

`func NewExchangeFinancialConnectionsPublicTokenResponse(accessToken string, financialConnectionId string, aggregator Aggregator, requestId string, ) *ExchangeFinancialConnectionsPublicTokenResponse`

NewExchangeFinancialConnectionsPublicTokenResponse instantiates a new ExchangeFinancialConnectionsPublicTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeFinancialConnectionsPublicTokenResponseWithDefaults

`func NewExchangeFinancialConnectionsPublicTokenResponseWithDefaults() *ExchangeFinancialConnectionsPublicTokenResponse`

NewExchangeFinancialConnectionsPublicTokenResponseWithDefaults instantiates a new ExchangeFinancialConnectionsPublicTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetFinancialConnectionId

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetFinancialConnectionId() string`

GetFinancialConnectionId returns the FinancialConnectionId field if non-nil, zero value otherwise.

### GetFinancialConnectionIdOk

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetFinancialConnectionIdOk() (*string, bool)`

GetFinancialConnectionIdOk returns a tuple with the FinancialConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnectionId

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) SetFinancialConnectionId(v string)`

SetFinancialConnectionId sets FinancialConnectionId field to given value.


### GetInstitution

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetInstitution() FinancialInstitution`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetInstitutionOk() (*FinancialInstitution, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) SetInstitution(v FinancialInstitution)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetAggregator

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetAggregator() Aggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetAggregatorOk() (*Aggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) SetAggregator(v Aggregator)`

SetAggregator sets Aggregator field to given value.


### GetRequestId

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ExchangeFinancialConnectionsPublicTokenResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


