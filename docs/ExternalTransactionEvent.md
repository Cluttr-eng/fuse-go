# ExternalTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the transaction | 
**EventType** | **string** |  | 
**Status** | [**ExternalTransactionEventStatus**](ExternalTransactionEventStatus.md) |  | 
**Amount** | **float32** | Amount in cents associated with the transaction. Use positive values to represent money going out and negative to represent money going in. | 
**CountryCode** | Pointer to **string** |  | [optional] [default to "US"]
**IsoCurrencyCode** | **string** | The ISO-4217 currency code. | 
**TransactionType** | Pointer to [**TransactionEventType**](TransactionEventType.md) |  | [optional] 
**TransactionDescription** | Pointer to **string** |  | [optional] 
**TransactionOwnerType** | Pointer to **string** |  | [optional] [default to "consumer"]
**MerchantName** | **string** |  | 
**Timestamp** | **string** | Datetime of the transaction In ISO-8601 format | 
**Balance** | Pointer to **float32** | The running balance of the account after the transaction has occurred, in cents | [optional] 

## Methods

### NewExternalTransactionEvent

`func NewExternalTransactionEvent(id string, eventType string, status ExternalTransactionEventStatus, amount float32, isoCurrencyCode string, merchantName string, timestamp string, ) *ExternalTransactionEvent`

NewExternalTransactionEvent instantiates a new ExternalTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalTransactionEventWithDefaults

`func NewExternalTransactionEventWithDefaults() *ExternalTransactionEvent`

NewExternalTransactionEventWithDefaults instantiates a new ExternalTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalTransactionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalTransactionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalTransactionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *ExternalTransactionEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ExternalTransactionEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ExternalTransactionEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetStatus

`func (o *ExternalTransactionEvent) GetStatus() ExternalTransactionEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalTransactionEvent) GetStatusOk() (*ExternalTransactionEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalTransactionEvent) SetStatus(v ExternalTransactionEventStatus)`

SetStatus sets Status field to given value.


### GetAmount

`func (o *ExternalTransactionEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExternalTransactionEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExternalTransactionEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCountryCode

`func (o *ExternalTransactionEvent) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ExternalTransactionEvent) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ExternalTransactionEvent) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ExternalTransactionEvent) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIsoCurrencyCode

`func (o *ExternalTransactionEvent) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *ExternalTransactionEvent) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *ExternalTransactionEvent) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetTransactionType

`func (o *ExternalTransactionEvent) GetTransactionType() TransactionEventType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ExternalTransactionEvent) GetTransactionTypeOk() (*TransactionEventType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ExternalTransactionEvent) SetTransactionType(v TransactionEventType)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ExternalTransactionEvent) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *ExternalTransactionEvent) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *ExternalTransactionEvent) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *ExternalTransactionEvent) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *ExternalTransactionEvent) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetTransactionOwnerType

`func (o *ExternalTransactionEvent) GetTransactionOwnerType() string`

GetTransactionOwnerType returns the TransactionOwnerType field if non-nil, zero value otherwise.

### GetTransactionOwnerTypeOk

`func (o *ExternalTransactionEvent) GetTransactionOwnerTypeOk() (*string, bool)`

GetTransactionOwnerTypeOk returns a tuple with the TransactionOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOwnerType

`func (o *ExternalTransactionEvent) SetTransactionOwnerType(v string)`

SetTransactionOwnerType sets TransactionOwnerType field to given value.

### HasTransactionOwnerType

`func (o *ExternalTransactionEvent) HasTransactionOwnerType() bool`

HasTransactionOwnerType returns a boolean if a field has been set.

### GetMerchantName

`func (o *ExternalTransactionEvent) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *ExternalTransactionEvent) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *ExternalTransactionEvent) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.


### GetTimestamp

`func (o *ExternalTransactionEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExternalTransactionEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExternalTransactionEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetBalance

`func (o *ExternalTransactionEvent) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ExternalTransactionEvent) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ExternalTransactionEvent) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *ExternalTransactionEvent) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


