# InAppTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the transaction | 
**EventType** | **string** |  | 
**Status** | [**InAppTransactionEventStatus**](InAppTransactionEventStatus.md) |  | 
**Amount** | **float32** |  | 
**IsoCurrencyCode** | **string** | The ISO-4217 currency code. | 
**TransactionType** | Pointer to [**TransactionEventType**](TransactionEventType.md) |  | [optional] 
**MerchantName** | **string** |  | 
**Timestamp** | **string** | Datetime of the transaction In ISO-8601 format | 
**Balance** | Pointer to **float32** | The running balance of the account after the transaction has occurred, in cents. | [optional] 

## Methods

### NewInAppTransactionEvent

`func NewInAppTransactionEvent(id string, eventType string, status InAppTransactionEventStatus, amount float32, isoCurrencyCode string, merchantName string, timestamp string, ) *InAppTransactionEvent`

NewInAppTransactionEvent instantiates a new InAppTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInAppTransactionEventWithDefaults

`func NewInAppTransactionEventWithDefaults() *InAppTransactionEvent`

NewInAppTransactionEventWithDefaults instantiates a new InAppTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InAppTransactionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InAppTransactionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InAppTransactionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *InAppTransactionEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *InAppTransactionEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *InAppTransactionEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetStatus

`func (o *InAppTransactionEvent) GetStatus() InAppTransactionEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InAppTransactionEvent) GetStatusOk() (*InAppTransactionEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InAppTransactionEvent) SetStatus(v InAppTransactionEventStatus)`

SetStatus sets Status field to given value.


### GetAmount

`func (o *InAppTransactionEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InAppTransactionEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InAppTransactionEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *InAppTransactionEvent) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *InAppTransactionEvent) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *InAppTransactionEvent) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetTransactionType

`func (o *InAppTransactionEvent) GetTransactionType() TransactionEventType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *InAppTransactionEvent) GetTransactionTypeOk() (*TransactionEventType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *InAppTransactionEvent) SetTransactionType(v TransactionEventType)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *InAppTransactionEvent) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetMerchantName

`func (o *InAppTransactionEvent) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *InAppTransactionEvent) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *InAppTransactionEvent) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.


### GetTimestamp

`func (o *InAppTransactionEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InAppTransactionEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InAppTransactionEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetBalance

`func (o *InAppTransactionEvent) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *InAppTransactionEvent) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *InAppTransactionEvent) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *InAppTransactionEvent) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


