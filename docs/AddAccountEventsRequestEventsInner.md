# AddAccountEventsRequestEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the transaction | 
**EventType** | **string** |  | 
**Status** | [**InAppTransactionEventStatus**](InAppTransactionEventStatus.md) |  | 
**Amount** | **float32** |  | 
**CountryCode** | Pointer to **string** |  | [optional] [default to "US"]
**IsoCurrencyCode** | **string** | The ISO-4217 currency code. | 
**TransactionType** | Pointer to [**TransactionEventType**](TransactionEventType.md) |  | [optional] 
**TransactionDescription** | Pointer to **string** |  | [optional] 
**TransactionOwnerType** | Pointer to **string** |  | [optional] [default to "consumer"]
**MerchantName** | **string** |  | 
**Timestamp** | **string** | Datetime that the balance is accurate for In ISO-8601 format | 
**Balance** | Pointer to **float32** | The running balance of the account after the transaction has occurred, in cents. | [optional] 
**Available** | Pointer to **NullableFloat32** | The current balance of the account factoring in pending transactions. | [optional] 
**Current** | Pointer to **NullableFloat32** | The current balance of the account without factoring in pending transactions. | [optional] 

## Methods

### NewAddAccountEventsRequestEventsInner

`func NewAddAccountEventsRequestEventsInner(id string, eventType string, status InAppTransactionEventStatus, amount float32, isoCurrencyCode string, merchantName string, timestamp string, ) *AddAccountEventsRequestEventsInner`

NewAddAccountEventsRequestEventsInner instantiates a new AddAccountEventsRequestEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountEventsRequestEventsInnerWithDefaults

`func NewAddAccountEventsRequestEventsInnerWithDefaults() *AddAccountEventsRequestEventsInner`

NewAddAccountEventsRequestEventsInnerWithDefaults instantiates a new AddAccountEventsRequestEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddAccountEventsRequestEventsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddAccountEventsRequestEventsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddAccountEventsRequestEventsInner) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *AddAccountEventsRequestEventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AddAccountEventsRequestEventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AddAccountEventsRequestEventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetStatus

`func (o *AddAccountEventsRequestEventsInner) GetStatus() InAppTransactionEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddAccountEventsRequestEventsInner) GetStatusOk() (*InAppTransactionEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddAccountEventsRequestEventsInner) SetStatus(v InAppTransactionEventStatus)`

SetStatus sets Status field to given value.


### GetAmount

`func (o *AddAccountEventsRequestEventsInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddAccountEventsRequestEventsInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddAccountEventsRequestEventsInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCountryCode

`func (o *AddAccountEventsRequestEventsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddAccountEventsRequestEventsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddAccountEventsRequestEventsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AddAccountEventsRequestEventsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIsoCurrencyCode

`func (o *AddAccountEventsRequestEventsInner) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *AddAccountEventsRequestEventsInner) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *AddAccountEventsRequestEventsInner) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetTransactionType

`func (o *AddAccountEventsRequestEventsInner) GetTransactionType() TransactionEventType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *AddAccountEventsRequestEventsInner) GetTransactionTypeOk() (*TransactionEventType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *AddAccountEventsRequestEventsInner) SetTransactionType(v TransactionEventType)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *AddAccountEventsRequestEventsInner) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *AddAccountEventsRequestEventsInner) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *AddAccountEventsRequestEventsInner) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *AddAccountEventsRequestEventsInner) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *AddAccountEventsRequestEventsInner) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetTransactionOwnerType

`func (o *AddAccountEventsRequestEventsInner) GetTransactionOwnerType() string`

GetTransactionOwnerType returns the TransactionOwnerType field if non-nil, zero value otherwise.

### GetTransactionOwnerTypeOk

`func (o *AddAccountEventsRequestEventsInner) GetTransactionOwnerTypeOk() (*string, bool)`

GetTransactionOwnerTypeOk returns a tuple with the TransactionOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOwnerType

`func (o *AddAccountEventsRequestEventsInner) SetTransactionOwnerType(v string)`

SetTransactionOwnerType sets TransactionOwnerType field to given value.

### HasTransactionOwnerType

`func (o *AddAccountEventsRequestEventsInner) HasTransactionOwnerType() bool`

HasTransactionOwnerType returns a boolean if a field has been set.

### GetMerchantName

`func (o *AddAccountEventsRequestEventsInner) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *AddAccountEventsRequestEventsInner) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *AddAccountEventsRequestEventsInner) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.


### GetTimestamp

`func (o *AddAccountEventsRequestEventsInner) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AddAccountEventsRequestEventsInner) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AddAccountEventsRequestEventsInner) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetBalance

`func (o *AddAccountEventsRequestEventsInner) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AddAccountEventsRequestEventsInner) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AddAccountEventsRequestEventsInner) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AddAccountEventsRequestEventsInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetAvailable

`func (o *AddAccountEventsRequestEventsInner) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AddAccountEventsRequestEventsInner) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AddAccountEventsRequestEventsInner) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AddAccountEventsRequestEventsInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *AddAccountEventsRequestEventsInner) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *AddAccountEventsRequestEventsInner) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *AddAccountEventsRequestEventsInner) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AddAccountEventsRequestEventsInner) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AddAccountEventsRequestEventsInner) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AddAccountEventsRequestEventsInner) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *AddAccountEventsRequestEventsInner) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *AddAccountEventsRequestEventsInner) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


