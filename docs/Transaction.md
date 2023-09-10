# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Remote Id of the transaction, ie Plaid or Teller Id | 
**Fingerprint** | Pointer to **string** | Uniquely identifies this transaction across all transactions for a single financial connection. Used for reconnection deduplication. This field may or may not be present depending on if a transaction can be uniquely identified. See more information here: https://letsfuse.readme.io/docs/duplicate-accounts | [optional] 
**RemoteAccountId** | **string** | Remote Account Id of the transaction, ie Plaid Account Id | 
**Amount** | **float32** | Amount in cents associated with the transaction. The format of this value is a double.  Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative. | 
**Date** | **string** | Date of the transaction (YYYY-MM-DD) | 
**Description** | **string** | Description of the transaction | 
**Category** | **[]string** | Categories of the transaction, i.e., Computers and Electronics. You can download the categories from [here](https://fuse-public-bucket.s3.amazonaws.com/transaction-categories.csv) | 
**Merchant** | [**TransactionMerchant**](TransactionMerchant.md) |  | 
**Status** | **string** | The status of the transaction. This will be either posted or pending. | 
**Type** | **string** | Type of the transaction, ie adjustment. &#39;-&#39; means we were not able to map the upstream type. | 
**IsoCurrencyCode** | Pointer to **string** | The ISO-4217 currency code of the transaction | [optional] 
**RemoteData** | **interface{}** |  | 

## Methods

### NewTransaction

`func NewTransaction(remoteId string, remoteAccountId string, amount float32, date string, description string, category []string, merchant TransactionMerchant, status string, type_ string, remoteData interface{}, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *Transaction) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Transaction) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Transaction) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetFingerprint

`func (o *Transaction) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Transaction) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Transaction) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Transaction) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetRemoteAccountId

`func (o *Transaction) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *Transaction) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *Transaction) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetAmount

`func (o *Transaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *Transaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Transaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Transaction) SetDate(v string)`

SetDate sets Date field to given value.


### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *Transaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetMerchant

`func (o *Transaction) GetMerchant() TransactionMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *Transaction) GetMerchantOk() (*TransactionMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *Transaction) SetMerchant(v TransactionMerchant)`

SetMerchant sets Merchant field to given value.


### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.


### GetIsoCurrencyCode

`func (o *Transaction) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *Transaction) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *Transaction) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *Transaction) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### GetRemoteData

`func (o *Transaction) GetRemoteData() interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Transaction) GetRemoteDataOk() (*interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Transaction) SetRemoteData(v interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *Transaction) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Transaction) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


