# TransactionToEnrich

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for the transaction that to help you tie data back to your systems. | 
**Description** | **string** | The name of the merchant if available or a description of the transaction. | 
**Mcc** | Pointer to **string** | The merchant category code. | [optional] 
**Amount** | **float32** | The amount of the transaction in cents, in the currency of the account. Must be a positive value. Use the type field to indicate the direction. | 
**Direction** | **string** | The direction of the transaction. | 
**CountryCode** | Pointer to **string** |  | [optional] [default to "US"]
**IsoCurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**Date** | Pointer to **string** | The date the transaction was posted. | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] [default to "consumer"]

## Methods

### NewTransactionToEnrich

`func NewTransactionToEnrich(id string, description string, amount float32, direction string, ) *TransactionToEnrich`

NewTransactionToEnrich instantiates a new TransactionToEnrich object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionToEnrichWithDefaults

`func NewTransactionToEnrichWithDefaults() *TransactionToEnrich`

NewTransactionToEnrichWithDefaults instantiates a new TransactionToEnrich object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionToEnrich) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionToEnrich) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionToEnrich) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *TransactionToEnrich) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionToEnrich) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionToEnrich) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMcc

`func (o *TransactionToEnrich) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *TransactionToEnrich) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *TransactionToEnrich) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *TransactionToEnrich) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionToEnrich) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionToEnrich) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionToEnrich) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDirection

`func (o *TransactionToEnrich) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TransactionToEnrich) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TransactionToEnrich) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetCountryCode

`func (o *TransactionToEnrich) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *TransactionToEnrich) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *TransactionToEnrich) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *TransactionToEnrich) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIsoCurrencyCode

`func (o *TransactionToEnrich) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransactionToEnrich) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransactionToEnrich) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *TransactionToEnrich) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### GetDate

`func (o *TransactionToEnrich) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionToEnrich) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionToEnrich) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionToEnrich) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetOwnerType

`func (o *TransactionToEnrich) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *TransactionToEnrich) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *TransactionToEnrich) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *TransactionToEnrich) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


