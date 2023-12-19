# SelectFinancialInstitutionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinancialInstitution** | [**FinancialInstitution**](FinancialInstitution.md) |  | 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewSelectFinancialInstitutionsResponse

`func NewSelectFinancialInstitutionsResponse(financialInstitution FinancialInstitution, ) *SelectFinancialInstitutionsResponse`

NewSelectFinancialInstitutionsResponse instantiates a new SelectFinancialInstitutionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectFinancialInstitutionsResponseWithDefaults

`func NewSelectFinancialInstitutionsResponseWithDefaults() *SelectFinancialInstitutionsResponse`

NewSelectFinancialInstitutionsResponseWithDefaults instantiates a new SelectFinancialInstitutionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinancialInstitution

`func (o *SelectFinancialInstitutionsResponse) GetFinancialInstitution() FinancialInstitution`

GetFinancialInstitution returns the FinancialInstitution field if non-nil, zero value otherwise.

### GetFinancialInstitutionOk

`func (o *SelectFinancialInstitutionsResponse) GetFinancialInstitutionOk() (*FinancialInstitution, bool)`

GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialInstitution

`func (o *SelectFinancialInstitutionsResponse) SetFinancialInstitution(v FinancialInstitution)`

SetFinancialInstitution sets FinancialInstitution field to given value.


### GetRequestId

`func (o *SelectFinancialInstitutionsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SelectFinancialInstitutionsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SelectFinancialInstitutionsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SelectFinancialInstitutionsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


