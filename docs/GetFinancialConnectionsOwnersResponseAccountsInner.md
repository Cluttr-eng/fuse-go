# GetFinancialConnectionsOwnersResponseAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteAccountId** | **string** | The remote account id of the account | 
**Owners** | [**[]FinancialConnectionsOwner**](FinancialConnectionsOwner.md) |  | 

## Methods

### NewGetFinancialConnectionsOwnersResponseAccountsInner

`func NewGetFinancialConnectionsOwnersResponseAccountsInner(remoteAccountId string, owners []FinancialConnectionsOwner, ) *GetFinancialConnectionsOwnersResponseAccountsInner`

NewGetFinancialConnectionsOwnersResponseAccountsInner instantiates a new GetFinancialConnectionsOwnersResponseAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsOwnersResponseAccountsInnerWithDefaults

`func NewGetFinancialConnectionsOwnersResponseAccountsInnerWithDefaults() *GetFinancialConnectionsOwnersResponseAccountsInner`

NewGetFinancialConnectionsOwnersResponseAccountsInnerWithDefaults instantiates a new GetFinancialConnectionsOwnersResponseAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteAccountId

`func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *GetFinancialConnectionsOwnersResponseAccountsInner) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetOwners

`func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetOwners() []FinancialConnectionsOwner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *GetFinancialConnectionsOwnersResponseAccountsInner) GetOwnersOk() (*[]FinancialConnectionsOwner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *GetFinancialConnectionsOwnersResponseAccountsInner) SetOwners(v []FinancialConnectionsOwner)`

SetOwners sets Owners field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


