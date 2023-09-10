# FinancialConnectionsAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Remote Id of the account, ie Plaid or Teller account id | 
**Ach** | Pointer to [**FinancialConnectionsAccountDetailsAch**](FinancialConnectionsAccountDetailsAch.md) |  | [optional] 
**AccountNumber** | Pointer to [**FinancialConnectionsAccountDetailsAccountNumber**](FinancialConnectionsAccountDetailsAccountNumber.md) |  | [optional] 
**RemoteData** | **interface{}** |  | 

## Methods

### NewFinancialConnectionsAccountDetails

`func NewFinancialConnectionsAccountDetails(remoteId string, remoteData interface{}, ) *FinancialConnectionsAccountDetails`

NewFinancialConnectionsAccountDetails instantiates a new FinancialConnectionsAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsAccountDetailsWithDefaults

`func NewFinancialConnectionsAccountDetailsWithDefaults() *FinancialConnectionsAccountDetails`

NewFinancialConnectionsAccountDetailsWithDefaults instantiates a new FinancialConnectionsAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *FinancialConnectionsAccountDetails) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FinancialConnectionsAccountDetails) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FinancialConnectionsAccountDetails) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetAch

`func (o *FinancialConnectionsAccountDetails) GetAch() FinancialConnectionsAccountDetailsAch`

GetAch returns the Ach field if non-nil, zero value otherwise.

### GetAchOk

`func (o *FinancialConnectionsAccountDetails) GetAchOk() (*FinancialConnectionsAccountDetailsAch, bool)`

GetAchOk returns a tuple with the Ach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAch

`func (o *FinancialConnectionsAccountDetails) SetAch(v FinancialConnectionsAccountDetailsAch)`

SetAch sets Ach field to given value.

### HasAch

`func (o *FinancialConnectionsAccountDetails) HasAch() bool`

HasAch returns a boolean if a field has been set.

### GetAccountNumber

`func (o *FinancialConnectionsAccountDetails) GetAccountNumber() FinancialConnectionsAccountDetailsAccountNumber`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FinancialConnectionsAccountDetails) GetAccountNumberOk() (*FinancialConnectionsAccountDetailsAccountNumber, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FinancialConnectionsAccountDetails) SetAccountNumber(v FinancialConnectionsAccountDetailsAccountNumber)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *FinancialConnectionsAccountDetails) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRemoteData

`func (o *FinancialConnectionsAccountDetails) GetRemoteData() interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *FinancialConnectionsAccountDetails) GetRemoteDataOk() (*interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *FinancialConnectionsAccountDetails) SetRemoteData(v interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *FinancialConnectionsAccountDetails) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *FinancialConnectionsAccountDetails) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


