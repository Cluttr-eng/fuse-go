# SearchFinancialInstitutionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinancialInstitutions** | [**[]FinancialInstitution**](FinancialInstitution.md) |  | 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewSearchFinancialInstitutionsResponse

`func NewSearchFinancialInstitutionsResponse(financialInstitutions []FinancialInstitution, ) *SearchFinancialInstitutionsResponse`

NewSearchFinancialInstitutionsResponse instantiates a new SearchFinancialInstitutionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFinancialInstitutionsResponseWithDefaults

`func NewSearchFinancialInstitutionsResponseWithDefaults() *SearchFinancialInstitutionsResponse`

NewSearchFinancialInstitutionsResponseWithDefaults instantiates a new SearchFinancialInstitutionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinancialInstitutions

`func (o *SearchFinancialInstitutionsResponse) GetFinancialInstitutions() []FinancialInstitution`

GetFinancialInstitutions returns the FinancialInstitutions field if non-nil, zero value otherwise.

### GetFinancialInstitutionsOk

`func (o *SearchFinancialInstitutionsResponse) GetFinancialInstitutionsOk() (*[]FinancialInstitution, bool)`

GetFinancialInstitutionsOk returns a tuple with the FinancialInstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialInstitutions

`func (o *SearchFinancialInstitutionsResponse) SetFinancialInstitutions(v []FinancialInstitution)`

SetFinancialInstitutions sets FinancialInstitutions field to given value.


### GetRequestId

`func (o *SearchFinancialInstitutionsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SearchFinancialInstitutionsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SearchFinancialInstitutionsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SearchFinancialInstitutionsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


