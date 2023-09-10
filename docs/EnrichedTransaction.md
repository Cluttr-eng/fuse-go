# EnrichedTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A original ID for the transaction that to help you tie data back to your systems. | 
**MerchantId** | Pointer to **string** | A Fuse defined, unique ID for the merchant associated with this transaction. | [optional] 
**MerchantName** | Pointer to **string** | The name of the merchant. | [optional] 
**Logo** | Pointer to [**MerchantLogo**](MerchantLogo.md) |  | [optional] 
**Categories** | Pointer to **[]string** | Hierarchical transaction categories: Each element narrows down from parent to nested sub-categories. Example: [&#39;personnel&#39;, &#39;employee&#39;, &#39;payroll&#39;]. | [optional] 

## Methods

### NewEnrichedTransaction

`func NewEnrichedTransaction(id string, ) *EnrichedTransaction`

NewEnrichedTransaction instantiates a new EnrichedTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichedTransactionWithDefaults

`func NewEnrichedTransactionWithDefaults() *EnrichedTransaction`

NewEnrichedTransactionWithDefaults instantiates a new EnrichedTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrichedTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrichedTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrichedTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *EnrichedTransaction) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *EnrichedTransaction) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *EnrichedTransaction) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *EnrichedTransaction) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantName

`func (o *EnrichedTransaction) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *EnrichedTransaction) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *EnrichedTransaction) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *EnrichedTransaction) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetLogo

`func (o *EnrichedTransaction) GetLogo() MerchantLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *EnrichedTransaction) GetLogoOk() (*MerchantLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *EnrichedTransaction) SetLogo(v MerchantLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *EnrichedTransaction) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCategories

`func (o *EnrichedTransaction) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *EnrichedTransaction) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *EnrichedTransaction) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *EnrichedTransaction) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


