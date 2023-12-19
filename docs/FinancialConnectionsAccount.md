# FinancialConnectionsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Remote Id of the account, ie Plaid or Teller account id | 
**Currency** | **string** | The ISO-4217 currency code of the account. | 
**Fingerprint** | **string** | Uniquely identifies this account across all accounts for a single financial connection. Used for reconnection deduplication. See more information here: https://letsfuse.readme.io/docs/duplicate-accounts | 
**Institution** | Pointer to [**FinancialConnectionsAccountInstitution**](FinancialConnectionsAccountInstitution.md) |  | [optional] 
**Mask** | Pointer to **NullableString** | The partial account number. | [optional] 
**Name** | **string** | The account&#39;s name, ie &#39;My Checking&#39; | 
**Type** | [**AccountType**](AccountType.md) |  | 
**Subtype** | Pointer to [**NullableAccountSubtype**](AccountSubtype.md) |  | [optional] 
**Balance** | [**FinancialConnectionsAccountCachedBalance**](FinancialConnectionsAccountCachedBalance.md) |  | 
**AdditionalBalances** | Pointer to [**[]FinancialConnectionsAccountCachedBalance**](FinancialConnectionsAccountCachedBalance.md) | An array of additional balances. This may be used for investment type accounts where the user can have multiple balances across different currencies. | [optional] 
**RemoteData** | **interface{}** |  | 

## Methods

### NewFinancialConnectionsAccount

`func NewFinancialConnectionsAccount(remoteId string, currency string, fingerprint string, name string, type_ AccountType, balance FinancialConnectionsAccountCachedBalance, remoteData interface{}, ) *FinancialConnectionsAccount`

NewFinancialConnectionsAccount instantiates a new FinancialConnectionsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsAccountWithDefaults

`func NewFinancialConnectionsAccountWithDefaults() *FinancialConnectionsAccount`

NewFinancialConnectionsAccountWithDefaults instantiates a new FinancialConnectionsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *FinancialConnectionsAccount) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FinancialConnectionsAccount) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FinancialConnectionsAccount) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetCurrency

`func (o *FinancialConnectionsAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialConnectionsAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialConnectionsAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFingerprint

`func (o *FinancialConnectionsAccount) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *FinancialConnectionsAccount) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *FinancialConnectionsAccount) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetInstitution

`func (o *FinancialConnectionsAccount) GetInstitution() FinancialConnectionsAccountInstitution`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *FinancialConnectionsAccount) GetInstitutionOk() (*FinancialConnectionsAccountInstitution, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *FinancialConnectionsAccount) SetInstitution(v FinancialConnectionsAccountInstitution)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *FinancialConnectionsAccount) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetMask

`func (o *FinancialConnectionsAccount) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *FinancialConnectionsAccount) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *FinancialConnectionsAccount) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *FinancialConnectionsAccount) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *FinancialConnectionsAccount) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *FinancialConnectionsAccount) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetName

`func (o *FinancialConnectionsAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialConnectionsAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialConnectionsAccount) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FinancialConnectionsAccount) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialConnectionsAccount) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialConnectionsAccount) SetType(v AccountType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *FinancialConnectionsAccount) GetSubtype() AccountSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FinancialConnectionsAccount) GetSubtypeOk() (*AccountSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FinancialConnectionsAccount) SetSubtype(v AccountSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *FinancialConnectionsAccount) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *FinancialConnectionsAccount) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *FinancialConnectionsAccount) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetBalance

`func (o *FinancialConnectionsAccount) GetBalance() FinancialConnectionsAccountCachedBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FinancialConnectionsAccount) GetBalanceOk() (*FinancialConnectionsAccountCachedBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FinancialConnectionsAccount) SetBalance(v FinancialConnectionsAccountCachedBalance)`

SetBalance sets Balance field to given value.


### GetAdditionalBalances

`func (o *FinancialConnectionsAccount) GetAdditionalBalances() []FinancialConnectionsAccountCachedBalance`

GetAdditionalBalances returns the AdditionalBalances field if non-nil, zero value otherwise.

### GetAdditionalBalancesOk

`func (o *FinancialConnectionsAccount) GetAdditionalBalancesOk() (*[]FinancialConnectionsAccountCachedBalance, bool)`

GetAdditionalBalancesOk returns a tuple with the AdditionalBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBalances

`func (o *FinancialConnectionsAccount) SetAdditionalBalances(v []FinancialConnectionsAccountCachedBalance)`

SetAdditionalBalances sets AdditionalBalances field to given value.

### HasAdditionalBalances

`func (o *FinancialConnectionsAccount) HasAdditionalBalances() bool`

HasAdditionalBalances returns a boolean if a field has been set.

### GetRemoteData

`func (o *FinancialConnectionsAccount) GetRemoteData() interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *FinancialConnectionsAccount) GetRemoteDataOk() (*interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *FinancialConnectionsAccount) SetRemoteData(v interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *FinancialConnectionsAccount) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *FinancialConnectionsAccount) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


