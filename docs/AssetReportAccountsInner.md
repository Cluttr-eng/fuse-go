# AssetReportAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **string** | The remote account ID of the account. | [optional] 
**Balance** | Pointer to [**AssetReportAccountsInnerBalance**](AssetReportAccountsInnerBalance.md) |  | [optional] 
**HistoricalBalances** | Pointer to [**[]AssetReportAccountsInnerHistoricalBalancesInner**](AssetReportAccountsInnerHistoricalBalancesInner.md) | An array of historical balances for the account. | [optional] 
**Transactions** | Pointer to [**[]AssetReportTransaction**](AssetReportTransaction.md) | An array of historical transactions for the account. | [optional] 

## Methods

### NewAssetReportAccountsInner

`func NewAssetReportAccountsInner() *AssetReportAccountsInner`

NewAssetReportAccountsInner instantiates a new AssetReportAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportAccountsInnerWithDefaults

`func NewAssetReportAccountsInnerWithDefaults() *AssetReportAccountsInner`

NewAssetReportAccountsInnerWithDefaults instantiates a new AssetReportAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *AssetReportAccountsInner) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *AssetReportAccountsInner) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *AssetReportAccountsInner) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *AssetReportAccountsInner) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetBalance

`func (o *AssetReportAccountsInner) GetBalance() AssetReportAccountsInnerBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AssetReportAccountsInner) GetBalanceOk() (*AssetReportAccountsInnerBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AssetReportAccountsInner) SetBalance(v AssetReportAccountsInnerBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AssetReportAccountsInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetHistoricalBalances

`func (o *AssetReportAccountsInner) GetHistoricalBalances() []AssetReportAccountsInnerHistoricalBalancesInner`

GetHistoricalBalances returns the HistoricalBalances field if non-nil, zero value otherwise.

### GetHistoricalBalancesOk

`func (o *AssetReportAccountsInner) GetHistoricalBalancesOk() (*[]AssetReportAccountsInnerHistoricalBalancesInner, bool)`

GetHistoricalBalancesOk returns a tuple with the HistoricalBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalBalances

`func (o *AssetReportAccountsInner) SetHistoricalBalances(v []AssetReportAccountsInnerHistoricalBalancesInner)`

SetHistoricalBalances sets HistoricalBalances field to given value.

### HasHistoricalBalances

`func (o *AssetReportAccountsInner) HasHistoricalBalances() bool`

HasHistoricalBalances returns a boolean if a field has been set.

### GetTransactions

`func (o *AssetReportAccountsInner) GetTransactions() []AssetReportTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AssetReportAccountsInner) GetTransactionsOk() (*[]AssetReportTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AssetReportAccountsInner) SetTransactions(v []AssetReportTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *AssetReportAccountsInner) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


