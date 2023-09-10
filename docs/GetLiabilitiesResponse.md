# GetLiabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Liabilities** | Pointer to [**[]FinancialConnectionsAccountLiability**](FinancialConnectionsAccountLiability.md) |  | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewGetLiabilitiesResponse

`func NewGetLiabilitiesResponse() *GetLiabilitiesResponse`

NewGetLiabilitiesResponse instantiates a new GetLiabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLiabilitiesResponseWithDefaults

`func NewGetLiabilitiesResponseWithDefaults() *GetLiabilitiesResponse`

NewGetLiabilitiesResponseWithDefaults instantiates a new GetLiabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiabilities

`func (o *GetLiabilitiesResponse) GetLiabilities() []FinancialConnectionsAccountLiability`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *GetLiabilitiesResponse) GetLiabilitiesOk() (*[]FinancialConnectionsAccountLiability, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *GetLiabilitiesResponse) SetLiabilities(v []FinancialConnectionsAccountLiability)`

SetLiabilities sets Liabilities field to given value.

### HasLiabilities

`func (o *GetLiabilitiesResponse) HasLiabilities() bool`

HasLiabilities returns a boolean if a field has been set.

### GetRequestId

`func (o *GetLiabilitiesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetLiabilitiesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetLiabilitiesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetLiabilitiesResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


