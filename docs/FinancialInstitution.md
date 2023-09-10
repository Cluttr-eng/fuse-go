# FinancialInstitution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the financial institution id. | 
**Name** | **string** | Name for the financial institution. | 
**Logo** | Pointer to [**FinancialInstitutionLogo**](FinancialInstitutionLogo.md) |  | [optional] 
**Website** | Pointer to **string** | Website of the financial institution. | [optional] 
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | List of country codes supported by this institution | 

## Methods

### NewFinancialInstitution

`func NewFinancialInstitution(id string, name string, countryCodes []CountryCode, ) *FinancialInstitution`

NewFinancialInstitution instantiates a new FinancialInstitution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialInstitutionWithDefaults

`func NewFinancialInstitutionWithDefaults() *FinancialInstitution`

NewFinancialInstitutionWithDefaults instantiates a new FinancialInstitution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinancialInstitution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialInstitution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialInstitution) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FinancialInstitution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialInstitution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialInstitution) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *FinancialInstitution) GetLogo() FinancialInstitutionLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *FinancialInstitution) GetLogoOk() (*FinancialInstitutionLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *FinancialInstitution) SetLogo(v FinancialInstitutionLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *FinancialInstitution) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetWebsite

`func (o *FinancialInstitution) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *FinancialInstitution) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *FinancialInstitution) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *FinancialInstitution) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetCountryCodes

`func (o *FinancialInstitution) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *FinancialInstitution) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *FinancialInstitution) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


