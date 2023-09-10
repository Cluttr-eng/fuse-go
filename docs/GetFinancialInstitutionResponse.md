# GetFinancialInstitutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinancialInstitution** | [**FinancialInstitution**](FinancialInstitution.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialInstitutionResponse

`func NewGetFinancialInstitutionResponse(financialInstitution FinancialInstitution, requestId string, ) *GetFinancialInstitutionResponse`

NewGetFinancialInstitutionResponse instantiates a new GetFinancialInstitutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialInstitutionResponseWithDefaults

`func NewGetFinancialInstitutionResponseWithDefaults() *GetFinancialInstitutionResponse`

NewGetFinancialInstitutionResponseWithDefaults instantiates a new GetFinancialInstitutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinancialInstitution

`func (o *GetFinancialInstitutionResponse) GetFinancialInstitution() FinancialInstitution`

GetFinancialInstitution returns the FinancialInstitution field if non-nil, zero value otherwise.

### GetFinancialInstitutionOk

`func (o *GetFinancialInstitutionResponse) GetFinancialInstitutionOk() (*FinancialInstitution, bool)`

GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialInstitution

`func (o *GetFinancialInstitutionResponse) SetFinancialInstitution(v FinancialInstitution)`

SetFinancialInstitution sets FinancialInstitution field to given value.


### GetRequestId

`func (o *GetFinancialInstitutionResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialInstitutionResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialInstitutionResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


