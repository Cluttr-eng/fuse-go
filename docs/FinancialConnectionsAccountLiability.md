# FinancialConnectionsAccountLiability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Remote Id of the account, ie Plaid or Teller account id | 
**Currency** | **string** | The ISO-4217 currency code of the account. | 
**Fingerprint** | **string** | Uniquely identifies this account across all accounts for a single financial connection. Used for reconnection deduplication. See more information here: https://letsfuse.readme.io/docs/duplicate-accounts | 
**Institution** | Pointer to [**FinancialConnectionsAccountInstitution**](FinancialConnectionsAccountInstitution.md) |  | [optional] 
**Mask** | Pointer to **string** | The partial account number. | [optional] 
**Name** | **string** | The account&#39;s name, ie &#39;My Checking&#39; | 
**Type** | [**AccountType**](AccountType.md) |  | 
**Subtype** | Pointer to [**NullableAccountSubtype**](AccountSubtype.md) |  | [optional] 
**Balance** | [**FinancialConnectionsAccountCachedBalance**](FinancialConnectionsAccountCachedBalance.md) |  | 
**RemoteData** | **interface{}** |  | 
**Aprs** | Pointer to [**[]FinancialConnectionsAccountLiabilityAllOfAprs**](FinancialConnectionsAccountLiabilityAllOfAprs.md) | The various interest rates that apply to the account. If APR data is not available, this array will be empty. | [optional] 
**InterestRatePercentage** | Pointer to **float32** | The interest rate on the loan as a percentage. | [optional] 
**OriginationPrincipalAmount** | Pointer to **float32** | The original principal balance of the loan. | [optional] 
**NextPaymentDueDate** | Pointer to **string** | The due date for the next payment. The due date is null if a payment is not expected. | [optional] 
**LastPaymentDate** | Pointer to **string** | The date of the last payment. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | [optional] 
**LastPaymentAmount** | Pointer to **float32** | The amount of the last payment. | [optional] 
**MinimumPaymentAmount** | Pointer to **float32** | The minimum payment required for an account. This can apply to any debt account. | [optional] 

## Methods

### NewFinancialConnectionsAccountLiability

`func NewFinancialConnectionsAccountLiability(remoteId string, currency string, fingerprint string, name string, type_ AccountType, balance FinancialConnectionsAccountCachedBalance, remoteData interface{}, ) *FinancialConnectionsAccountLiability`

NewFinancialConnectionsAccountLiability instantiates a new FinancialConnectionsAccountLiability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsAccountLiabilityWithDefaults

`func NewFinancialConnectionsAccountLiabilityWithDefaults() *FinancialConnectionsAccountLiability`

NewFinancialConnectionsAccountLiabilityWithDefaults instantiates a new FinancialConnectionsAccountLiability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *FinancialConnectionsAccountLiability) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FinancialConnectionsAccountLiability) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FinancialConnectionsAccountLiability) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetCurrency

`func (o *FinancialConnectionsAccountLiability) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialConnectionsAccountLiability) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialConnectionsAccountLiability) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFingerprint

`func (o *FinancialConnectionsAccountLiability) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *FinancialConnectionsAccountLiability) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *FinancialConnectionsAccountLiability) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetInstitution

`func (o *FinancialConnectionsAccountLiability) GetInstitution() FinancialConnectionsAccountInstitution`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *FinancialConnectionsAccountLiability) GetInstitutionOk() (*FinancialConnectionsAccountInstitution, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *FinancialConnectionsAccountLiability) SetInstitution(v FinancialConnectionsAccountInstitution)`

SetInstitution sets Institution field to given value.

### HasInstitution

`func (o *FinancialConnectionsAccountLiability) HasInstitution() bool`

HasInstitution returns a boolean if a field has been set.

### GetMask

`func (o *FinancialConnectionsAccountLiability) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *FinancialConnectionsAccountLiability) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *FinancialConnectionsAccountLiability) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *FinancialConnectionsAccountLiability) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetName

`func (o *FinancialConnectionsAccountLiability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialConnectionsAccountLiability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialConnectionsAccountLiability) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FinancialConnectionsAccountLiability) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialConnectionsAccountLiability) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialConnectionsAccountLiability) SetType(v AccountType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *FinancialConnectionsAccountLiability) GetSubtype() AccountSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FinancialConnectionsAccountLiability) GetSubtypeOk() (*AccountSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FinancialConnectionsAccountLiability) SetSubtype(v AccountSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *FinancialConnectionsAccountLiability) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *FinancialConnectionsAccountLiability) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *FinancialConnectionsAccountLiability) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetBalance

`func (o *FinancialConnectionsAccountLiability) GetBalance() FinancialConnectionsAccountCachedBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FinancialConnectionsAccountLiability) GetBalanceOk() (*FinancialConnectionsAccountCachedBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FinancialConnectionsAccountLiability) SetBalance(v FinancialConnectionsAccountCachedBalance)`

SetBalance sets Balance field to given value.


### GetRemoteData

`func (o *FinancialConnectionsAccountLiability) GetRemoteData() interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *FinancialConnectionsAccountLiability) GetRemoteDataOk() (*interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *FinancialConnectionsAccountLiability) SetRemoteData(v interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *FinancialConnectionsAccountLiability) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *FinancialConnectionsAccountLiability) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetAprs

`func (o *FinancialConnectionsAccountLiability) GetAprs() []FinancialConnectionsAccountLiabilityAllOfAprs`

GetAprs returns the Aprs field if non-nil, zero value otherwise.

### GetAprsOk

`func (o *FinancialConnectionsAccountLiability) GetAprsOk() (*[]FinancialConnectionsAccountLiabilityAllOfAprs, bool)`

GetAprsOk returns a tuple with the Aprs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprs

`func (o *FinancialConnectionsAccountLiability) SetAprs(v []FinancialConnectionsAccountLiabilityAllOfAprs)`

SetAprs sets Aprs field to given value.

### HasAprs

`func (o *FinancialConnectionsAccountLiability) HasAprs() bool`

HasAprs returns a boolean if a field has been set.

### GetInterestRatePercentage

`func (o *FinancialConnectionsAccountLiability) GetInterestRatePercentage() float32`

GetInterestRatePercentage returns the InterestRatePercentage field if non-nil, zero value otherwise.

### GetInterestRatePercentageOk

`func (o *FinancialConnectionsAccountLiability) GetInterestRatePercentageOk() (*float32, bool)`

GetInterestRatePercentageOk returns a tuple with the InterestRatePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRatePercentage

`func (o *FinancialConnectionsAccountLiability) SetInterestRatePercentage(v float32)`

SetInterestRatePercentage sets InterestRatePercentage field to given value.

### HasInterestRatePercentage

`func (o *FinancialConnectionsAccountLiability) HasInterestRatePercentage() bool`

HasInterestRatePercentage returns a boolean if a field has been set.

### GetOriginationPrincipalAmount

`func (o *FinancialConnectionsAccountLiability) GetOriginationPrincipalAmount() float32`

GetOriginationPrincipalAmount returns the OriginationPrincipalAmount field if non-nil, zero value otherwise.

### GetOriginationPrincipalAmountOk

`func (o *FinancialConnectionsAccountLiability) GetOriginationPrincipalAmountOk() (*float32, bool)`

GetOriginationPrincipalAmountOk returns a tuple with the OriginationPrincipalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationPrincipalAmount

`func (o *FinancialConnectionsAccountLiability) SetOriginationPrincipalAmount(v float32)`

SetOriginationPrincipalAmount sets OriginationPrincipalAmount field to given value.

### HasOriginationPrincipalAmount

`func (o *FinancialConnectionsAccountLiability) HasOriginationPrincipalAmount() bool`

HasOriginationPrincipalAmount returns a boolean if a field has been set.

### GetNextPaymentDueDate

`func (o *FinancialConnectionsAccountLiability) GetNextPaymentDueDate() string`

GetNextPaymentDueDate returns the NextPaymentDueDate field if non-nil, zero value otherwise.

### GetNextPaymentDueDateOk

`func (o *FinancialConnectionsAccountLiability) GetNextPaymentDueDateOk() (*string, bool)`

GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentDueDate

`func (o *FinancialConnectionsAccountLiability) SetNextPaymentDueDate(v string)`

SetNextPaymentDueDate sets NextPaymentDueDate field to given value.

### HasNextPaymentDueDate

`func (o *FinancialConnectionsAccountLiability) HasNextPaymentDueDate() bool`

HasNextPaymentDueDate returns a boolean if a field has been set.

### GetLastPaymentDate

`func (o *FinancialConnectionsAccountLiability) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *FinancialConnectionsAccountLiability) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *FinancialConnectionsAccountLiability) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.

### HasLastPaymentDate

`func (o *FinancialConnectionsAccountLiability) HasLastPaymentDate() bool`

HasLastPaymentDate returns a boolean if a field has been set.

### GetLastPaymentAmount

`func (o *FinancialConnectionsAccountLiability) GetLastPaymentAmount() float32`

GetLastPaymentAmount returns the LastPaymentAmount field if non-nil, zero value otherwise.

### GetLastPaymentAmountOk

`func (o *FinancialConnectionsAccountLiability) GetLastPaymentAmountOk() (*float32, bool)`

GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAmount

`func (o *FinancialConnectionsAccountLiability) SetLastPaymentAmount(v float32)`

SetLastPaymentAmount sets LastPaymentAmount field to given value.

### HasLastPaymentAmount

`func (o *FinancialConnectionsAccountLiability) HasLastPaymentAmount() bool`

HasLastPaymentAmount returns a boolean if a field has been set.

### GetMinimumPaymentAmount

`func (o *FinancialConnectionsAccountLiability) GetMinimumPaymentAmount() float32`

GetMinimumPaymentAmount returns the MinimumPaymentAmount field if non-nil, zero value otherwise.

### GetMinimumPaymentAmountOk

`func (o *FinancialConnectionsAccountLiability) GetMinimumPaymentAmountOk() (*float32, bool)`

GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPaymentAmount

`func (o *FinancialConnectionsAccountLiability) SetMinimumPaymentAmount(v float32)`

SetMinimumPaymentAmount sets MinimumPaymentAmount field to given value.

### HasMinimumPaymentAmount

`func (o *FinancialConnectionsAccountLiability) HasMinimumPaymentAmount() bool`

HasMinimumPaymentAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


