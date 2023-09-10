# AssetReportTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Remote Id of the transaction, ie Plaid or Teller Id | 
**RemoteAccountId** | **string** | Remote Account Id of the transaction, ie Plaid Account Id | 
**Amount** | **float32** | Amount in cents associated with the transaction. The format of this value is a double. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative. | 
**Date** | **string** | Date of the transaction (YYYY-MM-DD) | 
**Description** | **string** | Description of the transaction | 
**Category** | **[]string** | Categories of the transaction, ie Computers and Electronics. &#39;-&#39; means we could not map the upstream category. | 
**Merchant** | [**TransactionMerchant**](TransactionMerchant.md) |  | 
**Status** | **string** | The status of the transaction. This will be either posted or pending. | 
**IsoCurrencyCode** | Pointer to **string** | The ISO-4217 currency code of the transaction | [optional] 
**RemoteData** | **interface{}** |  | 

## Methods

### NewAssetReportTransaction

`func NewAssetReportTransaction(remoteId string, remoteAccountId string, amount float32, date string, description string, category []string, merchant TransactionMerchant, status string, remoteData interface{}, ) *AssetReportTransaction`

NewAssetReportTransaction instantiates a new AssetReportTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportTransactionWithDefaults

`func NewAssetReportTransactionWithDefaults() *AssetReportTransaction`

NewAssetReportTransactionWithDefaults instantiates a new AssetReportTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *AssetReportTransaction) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *AssetReportTransaction) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *AssetReportTransaction) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetRemoteAccountId

`func (o *AssetReportTransaction) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *AssetReportTransaction) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *AssetReportTransaction) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetAmount

`func (o *AssetReportTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AssetReportTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AssetReportTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *AssetReportTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AssetReportTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AssetReportTransaction) SetDate(v string)`

SetDate sets Date field to given value.


### GetDescription

`func (o *AssetReportTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetReportTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetReportTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *AssetReportTransaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AssetReportTransaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AssetReportTransaction) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetMerchant

`func (o *AssetReportTransaction) GetMerchant() TransactionMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *AssetReportTransaction) GetMerchantOk() (*TransactionMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *AssetReportTransaction) SetMerchant(v TransactionMerchant)`

SetMerchant sets Merchant field to given value.


### GetStatus

`func (o *AssetReportTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetReportTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetReportTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsoCurrencyCode

`func (o *AssetReportTransaction) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *AssetReportTransaction) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *AssetReportTransaction) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *AssetReportTransaction) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### GetRemoteData

`func (o *AssetReportTransaction) GetRemoteData() interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *AssetReportTransaction) GetRemoteDataOk() (*interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *AssetReportTransaction) SetRemoteData(v interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *AssetReportTransaction) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *AssetReportTransaction) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


