# UpdatedBalanceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**IsoCurrencyCode** | **string** | The ISO-4217 currency code. | 
**Timestamp** | **string** | Datetime that the balance is accurate for In ISO-8601 format | 
**Available** | Pointer to **NullableFloat32** | The current balance of the account factoring in pending transactions. | [optional] 
**Current** | Pointer to **NullableFloat32** | The current balance of the account without factoring in pending transactions. | [optional] 

## Methods

### NewUpdatedBalanceEvent

`func NewUpdatedBalanceEvent(eventType string, isoCurrencyCode string, timestamp string, ) *UpdatedBalanceEvent`

NewUpdatedBalanceEvent instantiates a new UpdatedBalanceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedBalanceEventWithDefaults

`func NewUpdatedBalanceEventWithDefaults() *UpdatedBalanceEvent`

NewUpdatedBalanceEventWithDefaults instantiates a new UpdatedBalanceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *UpdatedBalanceEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *UpdatedBalanceEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *UpdatedBalanceEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetIsoCurrencyCode

`func (o *UpdatedBalanceEvent) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *UpdatedBalanceEvent) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *UpdatedBalanceEvent) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetTimestamp

`func (o *UpdatedBalanceEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UpdatedBalanceEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UpdatedBalanceEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetAvailable

`func (o *UpdatedBalanceEvent) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UpdatedBalanceEvent) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UpdatedBalanceEvent) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UpdatedBalanceEvent) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *UpdatedBalanceEvent) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *UpdatedBalanceEvent) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *UpdatedBalanceEvent) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *UpdatedBalanceEvent) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *UpdatedBalanceEvent) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *UpdatedBalanceEvent) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *UpdatedBalanceEvent) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *UpdatedBalanceEvent) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


