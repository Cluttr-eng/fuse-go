# GetFinancialConnectionsTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access token for authentication. | 
**StartDate** | **string** | The earliest date for which data should be returned. Dates should be formatted as YYYY-MM-DD. | 
**EndDate** | **string** | The latest date for which data should be returned. Dates should be formatted as YYYY-MM-DD. | 
**Page** | **int32** | Specify current page. | 
**RecordsPerPage** | **int32** | Number of items per page. | [default to 25]

## Methods

### NewGetFinancialConnectionsTransactionsRequest

`func NewGetFinancialConnectionsTransactionsRequest(accessToken string, startDate string, endDate string, page int32, recordsPerPage int32, ) *GetFinancialConnectionsTransactionsRequest`

NewGetFinancialConnectionsTransactionsRequest instantiates a new GetFinancialConnectionsTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsTransactionsRequestWithDefaults

`func NewGetFinancialConnectionsTransactionsRequestWithDefaults() *GetFinancialConnectionsTransactionsRequest`

NewGetFinancialConnectionsTransactionsRequestWithDefaults instantiates a new GetFinancialConnectionsTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetFinancialConnectionsTransactionsRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetFinancialConnectionsTransactionsRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetFinancialConnectionsTransactionsRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetStartDate

`func (o *GetFinancialConnectionsTransactionsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetFinancialConnectionsTransactionsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetFinancialConnectionsTransactionsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GetFinancialConnectionsTransactionsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetFinancialConnectionsTransactionsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetFinancialConnectionsTransactionsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPage

`func (o *GetFinancialConnectionsTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetFinancialConnectionsTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetFinancialConnectionsTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.


### GetRecordsPerPage

`func (o *GetFinancialConnectionsTransactionsRequest) GetRecordsPerPage() int32`

GetRecordsPerPage returns the RecordsPerPage field if non-nil, zero value otherwise.

### GetRecordsPerPageOk

`func (o *GetFinancialConnectionsTransactionsRequest) GetRecordsPerPageOk() (*int32, bool)`

GetRecordsPerPageOk returns a tuple with the RecordsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsPerPage

`func (o *GetFinancialConnectionsTransactionsRequest) SetRecordsPerPage(v int32)`

SetRecordsPerPage sets RecordsPerPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


