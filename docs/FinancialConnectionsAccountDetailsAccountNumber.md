# FinancialConnectionsAccountDetailsAccountNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Unique identifier representing the account, typically referred to as the account number. | 
**SortCode** | Pointer to **string** | A six-digit number used by banks in the United Kingdom and Ireland to identify the branch where the account is held. | [optional] 
**Iban** | Pointer to **string** | International Bank Account Number (IBAN) is an internationally agreed system of identifying bank accounts across national borders to facilitate the communication and processing of cross border transactions. | [optional] 
**SwiftBic** | Pointer to **string** | SWIFT/BIC code is an international identifier used to distinctively recognize a particular bank during financial transactions, such as money transfers. | [optional] 
**Bsb** | Pointer to **string** | Bank-State-Branch (BSB) code is a six-digit numerical code used to identify an individual branch of a financial institution in Australia. | [optional] 
**Bic** | Pointer to **string** | Bank Identifier Code (BIC) is an international standard identifier used by banks and financial institutions globally to carry out transactions. | [optional] 

## Methods

### NewFinancialConnectionsAccountDetailsAccountNumber

`func NewFinancialConnectionsAccountDetailsAccountNumber(number string, ) *FinancialConnectionsAccountDetailsAccountNumber`

NewFinancialConnectionsAccountDetailsAccountNumber instantiates a new FinancialConnectionsAccountDetailsAccountNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsAccountDetailsAccountNumberWithDefaults

`func NewFinancialConnectionsAccountDetailsAccountNumberWithDefaults() *FinancialConnectionsAccountDetailsAccountNumber`

NewFinancialConnectionsAccountDetailsAccountNumberWithDefaults instantiates a new FinancialConnectionsAccountDetailsAccountNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FinancialConnectionsAccountDetailsAccountNumber) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetSortCode

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *FinancialConnectionsAccountDetailsAccountNumber) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.

### HasSortCode

`func (o *FinancialConnectionsAccountDetailsAccountNumber) HasSortCode() bool`

HasSortCode returns a boolean if a field has been set.

### GetIban

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *FinancialConnectionsAccountDetailsAccountNumber) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *FinancialConnectionsAccountDetailsAccountNumber) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwiftBic

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSwiftBic() string`

GetSwiftBic returns the SwiftBic field if non-nil, zero value otherwise.

### GetSwiftBicOk

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetSwiftBicOk() (*string, bool)`

GetSwiftBicOk returns a tuple with the SwiftBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftBic

`func (o *FinancialConnectionsAccountDetailsAccountNumber) SetSwiftBic(v string)`

SetSwiftBic sets SwiftBic field to given value.

### HasSwiftBic

`func (o *FinancialConnectionsAccountDetailsAccountNumber) HasSwiftBic() bool`

HasSwiftBic returns a boolean if a field has been set.

### GetBsb

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBsb() string`

GetBsb returns the Bsb field if non-nil, zero value otherwise.

### GetBsbOk

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBsbOk() (*string, bool)`

GetBsbOk returns a tuple with the Bsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsb

`func (o *FinancialConnectionsAccountDetailsAccountNumber) SetBsb(v string)`

SetBsb sets Bsb field to given value.

### HasBsb

`func (o *FinancialConnectionsAccountDetailsAccountNumber) HasBsb() bool`

HasBsb returns a boolean if a field has been set.

### GetBic

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *FinancialConnectionsAccountDetailsAccountNumber) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *FinancialConnectionsAccountDetailsAccountNumber) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *FinancialConnectionsAccountDetailsAccountNumber) HasBic() bool`

HasBic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


