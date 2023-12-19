# SelectFinancialInstitutionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionClientSecret** | **string** | The session client secret created. | 
**FinancialInstitutionId** | **string** | The selected financial institution id | 

## Methods

### NewSelectFinancialInstitutionsRequest

`func NewSelectFinancialInstitutionsRequest(sessionClientSecret string, financialInstitutionId string, ) *SelectFinancialInstitutionsRequest`

NewSelectFinancialInstitutionsRequest instantiates a new SelectFinancialInstitutionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectFinancialInstitutionsRequestWithDefaults

`func NewSelectFinancialInstitutionsRequestWithDefaults() *SelectFinancialInstitutionsRequest`

NewSelectFinancialInstitutionsRequestWithDefaults instantiates a new SelectFinancialInstitutionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionClientSecret

`func (o *SelectFinancialInstitutionsRequest) GetSessionClientSecret() string`

GetSessionClientSecret returns the SessionClientSecret field if non-nil, zero value otherwise.

### GetSessionClientSecretOk

`func (o *SelectFinancialInstitutionsRequest) GetSessionClientSecretOk() (*string, bool)`

GetSessionClientSecretOk returns a tuple with the SessionClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionClientSecret

`func (o *SelectFinancialInstitutionsRequest) SetSessionClientSecret(v string)`

SetSessionClientSecret sets SessionClientSecret field to given value.


### GetFinancialInstitutionId

`func (o *SelectFinancialInstitutionsRequest) GetFinancialInstitutionId() string`

GetFinancialInstitutionId returns the FinancialInstitutionId field if non-nil, zero value otherwise.

### GetFinancialInstitutionIdOk

`func (o *SelectFinancialInstitutionsRequest) GetFinancialInstitutionIdOk() (*string, bool)`

GetFinancialInstitutionIdOk returns a tuple with the FinancialInstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialInstitutionId

`func (o *SelectFinancialInstitutionsRequest) SetFinancialInstitutionId(v string)`

SetFinancialInstitutionId sets FinancialInstitutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


