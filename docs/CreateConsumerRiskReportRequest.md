# CreateConsumerRiskReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | A unique ID representing the bank account that this risk report is calculated for. Typically this will be a bank connection account ID from your application. It is currently used as a means of connecting events to the consumer risk report. | 
**IsoCurrencyCode** | **string** | The ISO-4217 currency code of the transaction | 
**EndUserName** | Pointer to **string** | The name of the end user associated with this consumer risk report. | [optional] 
**CustomizationId** | **string** | This is used to determine the timeframe and other metadata for the consumer risk report. | 

## Methods

### NewCreateConsumerRiskReportRequest

`func NewCreateConsumerRiskReportRequest(accountId string, isoCurrencyCode string, customizationId string, ) *CreateConsumerRiskReportRequest`

NewCreateConsumerRiskReportRequest instantiates a new CreateConsumerRiskReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConsumerRiskReportRequestWithDefaults

`func NewCreateConsumerRiskReportRequestWithDefaults() *CreateConsumerRiskReportRequest`

NewCreateConsumerRiskReportRequestWithDefaults instantiates a new CreateConsumerRiskReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateConsumerRiskReportRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateConsumerRiskReportRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateConsumerRiskReportRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetIsoCurrencyCode

`func (o *CreateConsumerRiskReportRequest) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *CreateConsumerRiskReportRequest) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *CreateConsumerRiskReportRequest) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetEndUserName

`func (o *CreateConsumerRiskReportRequest) GetEndUserName() string`

GetEndUserName returns the EndUserName field if non-nil, zero value otherwise.

### GetEndUserNameOk

`func (o *CreateConsumerRiskReportRequest) GetEndUserNameOk() (*string, bool)`

GetEndUserNameOk returns a tuple with the EndUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserName

`func (o *CreateConsumerRiskReportRequest) SetEndUserName(v string)`

SetEndUserName sets EndUserName field to given value.

### HasEndUserName

`func (o *CreateConsumerRiskReportRequest) HasEndUserName() bool`

HasEndUserName returns a boolean if a field has been set.

### GetCustomizationId

`func (o *CreateConsumerRiskReportRequest) GetCustomizationId() string`

GetCustomizationId returns the CustomizationId field if non-nil, zero value otherwise.

### GetCustomizationIdOk

`func (o *CreateConsumerRiskReportRequest) GetCustomizationIdOk() (*string, bool)`

GetCustomizationIdOk returns a tuple with the CustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationId

`func (o *CreateConsumerRiskReportRequest) SetCustomizationId(v string)`

SetCustomizationId sets CustomizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


